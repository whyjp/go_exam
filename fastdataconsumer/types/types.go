package types

import (
	"fastdataconsumer/util"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

type Checker struct {
	RecvCnt          int
	TransferredBytes int
	FailedtoJson     int
	Seqnos           map[int]int
	FdSeqnos         map[int]int
	Type_seqnos      map[string]map[int]int
	Summarys         []NXLog_SdkSummary_M
}
type NoCheckResult struct {
	Cnt           int
	Max           int
	CntLost       int
	LostList      interface{}
	CntDuplicate  int
	DuplicateList interface{}
}
type CheckResult struct {
	RecvCnt          int
	TransferredBytes int
	FailedtoJson     int

	Seqno_CheckResult   interface{}
	Fdno_CheckResult    NoCheckResult
	TypeCnt             int
	TypeList            []string
	Typeno_CheckResults map[string]NoCheckResult
}

type EOF_status struct {
	EOFstat   bool
	Topic     string
	Partition int32
}

type Type_common struct {
	Type            string `json:"type"`
	Sequenceno      int    `json:"sequenceno"`
	Fdseqno         int    `json:"fdsdqno"`
	Typeseqno       int    `json:"typeseqno"`
	TransferredByte int    `json:"bytes"`
}

type NXLog_SdkSummary_M struct {
	Version         string `json:"version"`
	Regionid        string `json:"regionid"`
	Type            string `json:"type"`
	Ipaddress       string `json:"ipaddress"`
	Sequenceno      int    `json:"sequenceno"`
	Serviceid       string `json:"serviceid"`
	Typeseqno       int    `json:"typeseqno"`
	Servername      string `json:"servername"`
	Processid       int    `json:"processid"`
	Worldid         string `json:"worldid"`
	NXLogSdkSummary struct {
		SentMessages []struct {
			Count    int    `json:"count"`
			Typename string `json:"typename"`
		} `json:"sent_messages"`
		Resentmessages []struct {
			Count    int    `json:"count"`
			Typename string `json:"typename"`
		} `json:"resent_messages"`
		Failedmessages []struct {
			Count    int    `json:"count"`
			Typename string `json:"typename"`
		} `json:"failed_messages"`
		Deletedmessages []struct {
			Count    int    `json:"count"`
			Typename string `json:"typename"`
		} `json:"deleted_messages"`
		LogServerProtocol string    `json:"log_server_protocol"`
		EndDatetime       time.Time `json:"end_datetime"`
		LogServerAddress  []string  `json:"log_server_address"`
		BeginDatetime     time.Time `json:"begin_datetime"`
		RequestedMessages []struct {
			Count    int    `json:"count"`
			Typename string `json:"typename"`
		} `json:"requested_messages"`
	} `json:"NXLog_SdkSummary"`
	Mid        string    `json:"mid"`
	Createdate time.Time `json:"createdate"`
	Timesync   bool      `json:"timesync"`
	Encoding   string    `json:"encoding"`
}

func getSortedKeys(from map[int]int) []int {
	keys := make([]int, 0, len(from))
	for k, _ := range from {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

var (
	checktypeseqno           = false
	checklosslist            = false
	checkduplicatelist       = false
	typeresult_onlysensitive = false
	checklist_asstring       = false
)

func NewCheckResult(_checktypeseqno bool, _checklosslist bool, _checkduplicatelist bool, _typeresult_onlysensitive bool, _checklist_asstring bool) *CheckResult {
	checkResult_ := CheckResult{
		Fdno_CheckResult:    NoCheckResult{},
		Typeno_CheckResults: make(map[string]NoCheckResult),
	}
	checktypeseqno = _checktypeseqno
	checklosslist = _checklosslist
	checkduplicatelist = _checkduplicatelist
	typeresult_onlysensitive = _typeresult_onlysensitive
	checklist_asstring = _checklist_asstring
	return &checkResult_
}

func (checkResult_ *CheckResult) FinalizeConsumer(_checker Checker) {
	checkResult_.RecvCnt = _checker.RecvCnt
	checkResult_.TransferredBytes = _checker.TransferredBytes
	checkResult_.FailedtoJson = _checker.FailedtoJson

	type dupSet struct {
		duplicate    int
		duplicateKey []int
		loss         int
		lossKey      []int
	}

	if len(_checker.Seqnos) > 0 {
		Seqno_CheckResult_ := NoCheckResult{}
		Seqno_CheckResult_.Cnt = len(_checker.Seqnos)

		keys := getSortedKeys(_checker.Seqnos)

		k := keys[len(keys)-1]
		Seqno_CheckResult_.Max = k

		dup := dupSet{
			duplicate: 0,
			loss:      0,
		}
		for cur := 1; cur <= keys[len(keys)-1]; cur++ {
			cnt, exist := _checker.Seqnos[cur]
			if !exist {
				dup.loss++
				dup.lossKey = append(dup.lossKey, cur)
			} else {
				if cnt > 1 {
					dup.duplicate += cnt
					dup.duplicateKey = append(dup.duplicateKey, cur)
				}
			}
		}
		Seqno_CheckResult_.CntLost = dup.loss
		if checklosslist {
			if checklist_asstring {
				strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.lossKey)), "|"), "[]")
				Seqno_CheckResult_.LostList = strList
			} else {
				Seqno_CheckResult_.LostList = dup.lossKey
			}
		} else {
			Seqno_CheckResult_.LostList = "disable"
		}

		Seqno_CheckResult_.CntDuplicate = dup.duplicate
		if checkduplicatelist {
			if checklist_asstring {
				strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.duplicateKey)), "|"), "[]")
				Seqno_CheckResult_.DuplicateList = strList
			} else {
				Seqno_CheckResult_.DuplicateList = dup.duplicateKey
			}
		} else {
			Seqno_CheckResult_.DuplicateList = "disable"
		}

		checkResult_.Seqno_CheckResult = Seqno_CheckResult_
	} else {
		checkResult_.Seqno_CheckResult = "disable"
	}

	if len(_checker.FdSeqnos) > 0 {
		checkResult_.Fdno_CheckResult.Cnt = len(_checker.FdSeqnos)

		keys := getSortedKeys(_checker.FdSeqnos)

		k := keys[len(keys)-1]
		checkResult_.Fdno_CheckResult.Max = k

		dup := dupSet{
			duplicate: 0,
			loss:      0,
		}
		for cur := 1; cur <= keys[len(keys)-1]; cur++ {
			cnt, exist := _checker.FdSeqnos[cur]
			if !exist {
				dup.loss++
				dup.lossKey = append(dup.lossKey, cur)
			} else {
				if cnt > 1 {
					dup.duplicate += cnt
					dup.duplicateKey = append(dup.duplicateKey, cur)
				}
			}
		}
		checkResult_.Fdno_CheckResult.CntLost = dup.loss
		if checklosslist {
			if checklist_asstring {
				strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.lossKey)), "|"), "[]")
				checkResult_.Fdno_CheckResult.LostList = strList
			} else {
				checkResult_.Fdno_CheckResult.LostList = dup.lossKey
			}
		} else {
			checkResult_.Fdno_CheckResult.LostList = "disable"
		}

		checkResult_.Fdno_CheckResult.CntDuplicate = dup.duplicate
		if checkduplicatelist {
			if checklist_asstring {
				strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.duplicateKey)), "|"), "[]")
				checkResult_.Fdno_CheckResult.DuplicateList = strList
			} else {
				checkResult_.Fdno_CheckResult.DuplicateList = dup.duplicateKey
			}
		} else {
			checkResult_.Fdno_CheckResult.DuplicateList = "disable"
		}
	}
	if checktypeseqno {
		for typename, seqs := range _checker.Type_seqnos {
			if len(seqs) > 0 {
				keys := getSortedKeys(seqs)

				k := keys[len(keys)-1]
				cur_CheckResult := NoCheckResult{}
				cur_CheckResult.Max = k

				dup := dupSet{
					duplicate: 0,
					loss:      0,
				}
				for cur := 1; cur <= keys[len(keys)-1]; cur++ {
					cnt, exist := _checker.Type_seqnos[typename][cur]
					if !exist {
						dup.loss++
						if checklosslist {
							dup.lossKey = append(dup.lossKey, cur)
						}
					} else {
						if cnt > 1 {
							dup.duplicate += cnt
							dup.duplicateKey = append(dup.duplicateKey, cur)
						}
					}
				}

				cur_CheckResult.Cnt = len(seqs)
				cur_CheckResult.CntLost = dup.loss
				if checklosslist {
					if checklist_asstring {
						strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.lossKey)), "|"), "[]")
						cur_CheckResult.LostList = strList
					} else {
						cur_CheckResult.LostList = dup.lossKey
					}
				} else {
					cur_CheckResult.LostList = "disable"
				}

				cur_CheckResult.CntDuplicate = dup.duplicate
				if checkduplicatelist {
					if checklist_asstring {
						strList := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(dup.duplicateKey)), "|"), "[]")
						cur_CheckResult.DuplicateList = strList
					} else {
						cur_CheckResult.DuplicateList = dup.duplicateKey
					}
				} else {
					cur_CheckResult.DuplicateList = "disable"
				}
				if !typeresult_onlysensitive || dup.loss > 0 || dup.duplicate > 0 {
					_, exist := checkResult_.Typeno_CheckResults[typename]
					if !exist {
						checkResult_.Typeno_CheckResults[typename] = cur_CheckResult
					} else {
						log.Println("type check exist... overwriten... ")
						checkResult_.Typeno_CheckResults[typename] = cur_CheckResult
					}
				}
			}
			checkResult_.TypeCnt++
			checkResult_.TypeList = append(checkResult_.TypeList, typename)
		}
	}

	util.LogingStructure(checkResult_)
}
