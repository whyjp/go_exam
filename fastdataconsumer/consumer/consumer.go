package kafka

import (
	"encoding/json"
	"fastdataconsumer/types"
	"log"
	"reflect"
	"sync"

	"github.com/Shopify/sarama"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	Ready             chan bool
	Type_cur          chan types.Type_common
	Jsonfail          chan bool
	Summ              chan types.NXLog_SdkSummary_M
	Recvmsg           chan string
	BeforeEOF         chan types.EOF_status
	PartitionsEOFStat map[string]map[int32]bool
	EOFstatLock       sync.Mutex
}

func NewConsumer() *Consumer {
	consumer_ := Consumer{
		Ready:     make(chan bool),
		Type_cur:  make(chan types.Type_common),
		Jsonfail:  make(chan bool),
		Summ:      make(chan types.NXLog_SdkSummary_M),
		Recvmsg:   make(chan string),
		BeforeEOF: make(chan types.EOF_status),
	}
	return &consumer_
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)
	consumer.PartitionsEOFStat = make(map[string]map[int32]bool, len(session.Claims()))

	for topic, claim := range session.Claims() {
		consumer.PartitionsEOFStat[topic] = make(map[int32]bool, len(claim))
		for _, partition := range claim {
			consumer.PartitionsEOFStat[topic][partition] = false
		}
	}
	for topic, partition := range consumer.PartitionsEOFStat {
		log.Printf("topic: %s partition cnt: %d", topic, len(partition))
	}
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	if claim.HighWaterMarkOffset() == 0 || claim.HighWaterMarkOffset() == getInitialOffset(claim) {
		eofstat := types.EOF_status{
			EOFstat:   true,
			Topic:     claim.Topic(),
			Partition: claim.Partition(),
		}
		consumer.BeforeEOF <- eofstat
	}
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			consumer.Recvmsg <- string(message.Value)
			//log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", strings.TrimRight(string(message.Value[:len(message.Value)-1]), "\\n"), message.Timestamp, message.Topic)
			var objmap map[string]interface{}
			if err := json.Unmarshal([]byte(message.Value), &objmap); err != nil {
				//log.Println(err)
				consumer.Jsonfail <- true
			} else {
				val, exists := objmap["type"]
				if exists {
					//log.Printf("type : %s", val.(string))
					var curType = types.Type_common{
						Sequenceno: -1,
						Fdseqno:    -1,
						Typeseqno:  -1,
					}
					curType.Type = val.(string)
					if false {
						val_seq, exists_seq := objmap["sequenceno"]
						if exists_seq {
							curType.Sequenceno = int(val_seq.(float64))
						}
					}
					val_fdseq, exists_fdseq := objmap["fdseqno"]
					if exists_fdseq {
						curType.Fdseqno = int(val_fdseq.(float64))
					}
					val_typeseq, exists_typeseq := objmap["typeseqno"]
					if exists_typeseq {
						curType.Typeseqno = int(val_typeseq.(float64))
					}
					curType.TransferredByte = len(message.Value)
					consumer.Type_cur <- curType
					if val.(string) == "NXLog_SdkSummary" {
						var sumCur types.NXLog_SdkSummary_M
						err := json.Unmarshal([]byte(message.Value), &sumCur)
						if err != nil {
							//log.Println(err)
							consumer.Jsonfail <- true
						} else {
							consumer.Summ <- sumCur
						}
					}
				}
				if claim.HighWaterMarkOffset() == message.Offset+1 {
					eofstat := types.EOF_status{
						EOFstat:   true,
						Topic:     message.Topic,
						Partition: message.Partition,
					}
					consumer.BeforeEOF <- eofstat
				} else {
					eofstat := types.EOF_status{
						EOFstat:   false,
						Topic:     message.Topic,
						Partition: message.Partition,
					}
					consumer.BeforeEOF <- eofstat
				}
			}

			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

// Sarama does not export partition initial offset so we need use reflection to get it
func getInitialOffset(claim sarama.ConsumerGroupClaim) int64 {
	cgc := reflect.ValueOf(claim).Elem()        // sarama.consumerGroupClaim
	pci := cgc.FieldByName("PartitionConsumer") // sarama.PartitionConsumer
	pc := pci.Elem().Elem()                     // sarama.partitionConsumer
	o := pc.FieldByName("offset")
	return o.Int()
}
