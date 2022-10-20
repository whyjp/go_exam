package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"fastdataconsumer/config"
	kafka "fastdataconsumer/consumer"
	"fastdataconsumer/types"
	"fastdataconsumer/util"

	"github.com/Shopify/sarama"
	flag "github.com/spf13/pflag"
)

var (
	brokers                  = ""
	version                  = ""
	group                    = ""
	topics                   = ""
	assignor                 = ""
	oldest                   = true
	verbose                  = false
	autocommit               = true
	logfilename              = "log.log"
	stdout                   = false
	consumeouttofile         = false
	checktypeseqno           = false
	checklosslist            = false
	checkduplicatelist       = false
	typeresult_onlysensitive = false
	checklist_asstring       = false

	filelock      sync.Mutex
	writefilename = "_out_"
)

func init() {
	today := time.Now().Format("20060102_150405")
	writefilename = today + writefilename
	flag.StringVarP(&brokers, "kafka.brokers", "b", "", "Kafka bootstrap brokers to connect to, as a comma separated list")
	flag.StringVarP(&group, "kafka.group", "g", "", "Kafka consumer group definition")
	flag.StringVarP(&version, "kafka.version", "v", "2.1.1", "Kafka cluster version")
	flag.StringVarP(&topics, "kafka.topics", "t", "", "Kafka topics to be consumed, as a comma separated list")
	flag.StringVarP(&assignor, "kafka.assignor", "s", "range", "Consumer group partition assignment strategy (range, roundrobin, sticky)")
	flag.BoolVarP(&oldest, "kafka.oldest", "o", true, "Kafka consumer consume initial offset from oldest")
	flag.BoolVarP(&autocommit, "kafka.autocommit", "a", true, "auto commit")
	flag.Parse()

	config.Init("fastdataconsumer")

	config_ := config.GetConfig()

	x := config_.AllSettings()
	util.LogingStructure(x)

	brokers = config_.GetString("kafka.brokers")
	topics = config_.GetString("kafka.topics")
	group = config_.GetString("kafka.group")
	autocommit = config_.GetBool("kafka.autocommit")
	logfilename = config_.GetString("app.logfile")
	stdout = config_.GetBool("app.stdout")
	consumeouttofile = config_.GetBool("app.consumeouttofile")
	checktypeseqno = config_.GetBool("app.checktypeseqno")
	checklosslist = config_.GetBool("app.checklosslist")
	checkduplicatelist = config_.GetBool("app.checklosslist")
	typeresult_onlysensitive = config_.GetBool("app.typeresult_onlysensitive")
	checklist_asstring = config_.GetBool("app.checklist_asstring")

	if len(brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

func main() {
	file, err := os.OpenFile(logfilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	var multiWriter io.Writer
	if stdout {
		multiWriter = io.MultiWriter(os.Stdout, file)
	} else {
		multiWriter = io.MultiWriter(file)
	}
	log.SetOutput(multiWriter)

	keepRunning := true
	log.Printf("Starting a fastdata consumer pid:[%d]", os.Getpid())

	if verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	switch assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategySticky}
	case "roundrobin":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRoundRobin}
	case "range":
		config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.BalanceStrategyRange}
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", assignor)
	}

	if oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
		log.Println("offset to oldest")
	}
	if !autocommit {
		config.Consumer.Offsets.AutoCommit.Enable = false
	}
	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := kafka.NewConsumer()

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(brokers, ","), group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	//consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(topics, ","), consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.Ready = make(chan bool)
		}
	}()

	<-consumer.Ready // Await till the consumer has been set up
	log.Println("fastdataconsumer up and running!...")
	log.Printf("connected at broker [%s]", brokers)

	//sigusr1 := make(chan os.Signal, 1)
	//signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	var checker_ = types.Checker{
		Seqnos:           make(map[int]int),
		FdSeqnos:         make(map[int]int),
		Type_seqnos:      make(map[string]map[int]int),
		TransferredBytes: 0,
	}
	curIdx := 0
	curMax := 4 * 3
	curMap := [...]string{"|", "|", "|", "/", "/", "/", "-", "-", "-", "\\", "\\", "\\"}
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		//case <-sigusr1:
		//	toggleConsumptionFlow(client, &consumptionIsPaused)
		case curType := <-consumer.Type_cur:
			//log.Printf("type: %s, seq: %d, typeseq: %d", curType.Type, int(curType.Sequenceno), int(curType.Typeseqno))
			if curType.Sequenceno != -1 {
				_, exist := checker_.Seqnos[curType.Sequenceno]
				if exist {
					checker_.Seqnos[curType.Sequenceno]++
				} else {
					checker_.Seqnos[curType.Sequenceno] = 1
				}
			}
			if curType.Fdseqno != -1 {
				_, exist := checker_.FdSeqnos[curType.Fdseqno]
				if exist {
					checker_.FdSeqnos[curType.Fdseqno]++
				} else {
					checker_.FdSeqnos[curType.Fdseqno] = 1
				}
			}
			if curType.Typeseqno != -1 {
				_, exist_type := checker_.Type_seqnos[curType.Type]
				if !exist_type {
					checker_.Type_seqnos[curType.Type] = map[int]int{}
				}
				_, exist_t := checker_.Type_seqnos[curType.Type][curType.Typeseqno]
				if exist_t {
					checker_.Type_seqnos[curType.Type][curType.Typeseqno]++
				} else {
					checker_.Type_seqnos[curType.Type][curType.Typeseqno] = 1
				}
			}
			checker_.TransferredBytes += curType.TransferredByte
		case jsonfail_ := <-consumer.Jsonfail:
			if jsonfail_ {
				checker_.FailedtoJson++
			}
		case summ_ := <-consumer.Summ:
			checker_.Summarys = append(checker_.Summarys, summ_)
		case recvmsg_ := <-consumer.Recvmsg:
			if consumeouttofile {
				writedata(recvmsg_)
			}
			checker_.RecvCnt++
			cur := curMap[curIdx]
			curIdx++
			if curIdx > curMax-1 {
				curIdx = 0
			}
			//fmt.Printf("status: %s\r", cur)
			fmt.Printf("consuming......%s\r", cur)
		case beforeEOF_ := <-consumer.BeforeEOF:
			consumer.EOFstatLock.Lock()

			if beforeEOF_.EOFstat {
				log.Printf("current Topic:\"%s\" partition:#%d consumed to end", beforeEOF_.Topic, beforeEOF_.Partition)

				consumer.PartitionsEOFStat[beforeEOF_.Topic][beforeEOF_.Partition] = true
				bEnd := true

				for _, partition := range consumer.PartitionsEOFStat {
					for _, stat := range partition {
						if !stat {
							bEnd = false
						}
					}
				}
				if bEnd {
					log.Printf("all topics partitions offset at the end [%d/%d]", len(consumer.PartitionsEOFStat), len(consumer.PartitionsEOFStat))
					//checkResult_ := types.NewCheckResult()
					//checkResult_.FinalizeConsumer(checker_)
				}
			} else {
				consumer.PartitionsEOFStat[beforeEOF_.Topic][beforeEOF_.Partition] = false
			}
			consumer.EOFstatLock.Unlock()
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	} else {
		log.Printf("end of consumer")
		checkResult_ := types.NewCheckResult(checktypeseqno, checklosslist, checkduplicatelist, typeresult_onlysensitive, checklist_asstring)
		checkResult_.FinalizeConsumer(checker_)
	}
}

func writedata(str string) {
	filelock.Lock()
	defer filelock.Unlock()

	writefile, err := os.OpenFile(writefilename+strconv.Itoa(os.Getpid())+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Panic(err)
	}
	defer writefile.Close()

	_, err_ := writefile.Seek(0, 2)
	if err_ != nil {
		log.Panicln("unable to seek to the end of")
	}
	writefile.Write([]byte(str + "\n"))
}

/*
func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}
*/
