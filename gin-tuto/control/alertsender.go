package control

import (
	"log"

	"github.com/go-resty/resty/v2"
	"webzen.com/notifyhandler/model"
)

type StAlertSender struct {
}

type AuthSuccess struct {
}
type AuthError struct {
}

func SendTeams(jsonTeams *model.StNotifyTeams) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(jsonTeams).
		SetResult(AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).   // or SetError(AuthError{}).
		Post("http://10.105.33.38/alert/api/v2/teams")

	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	log.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	log.Println("  DNSLookup     :", ti.DNSLookup)
	log.Println("  ConnTime      :", ti.ConnTime)
	log.Println("  TCPConnTime   :", ti.TCPConnTime)
	log.Println("  TLSHandshake  :", ti.TLSHandshake)
	log.Println("  ServerTime    :", ti.ServerTime)
	log.Println("  ResponseTime  :", ti.ResponseTime)
	log.Println("  TotalTime     :", ti.TotalTime)
	log.Println("  IsConnReused  :", ti.IsConnReused)
	log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	log.Println("  RequestAttempt:", ti.RequestAttempt)
	//log.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	return resp, nil
}

func SendMail(jsonMail *model.StNotifyMail) (*resty.Response, error) {
	client := resty.New()
	resp, err := client.R().
		SetBody(jsonMail).
		SetResult(AuthSuccess{}). // or SetResult(AuthSuccess{}).
		SetError(&AuthError{}).   // or SetError(AuthError{}).
		Post("http://10.105.33.38/alert/api/v2/email")

	// Explore response object
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	log.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	log.Println("  DNSLookup     :", ti.DNSLookup)
	log.Println("  ConnTime      :", ti.ConnTime)
	log.Println("  TCPConnTime   :", ti.TCPConnTime)
	log.Println("  TLSHandshake  :", ti.TLSHandshake)
	log.Println("  ServerTime    :", ti.ServerTime)
	log.Println("  ResponseTime  :", ti.ResponseTime)
	log.Println("  TotalTime     :", ti.TotalTime)
	log.Println("  IsConnReused  :", ti.IsConnReused)
	log.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	log.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	log.Println("  RequestAttempt:", ti.RequestAttempt)
	//log.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	return resp, nil
}
