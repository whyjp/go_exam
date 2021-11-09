package notifysender

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"webzen.com/notifyhandler/model"
)

var defaultBaseURL = "http://10.105.33.38"
var mailPath = "/alert/api/v2/mail"
var teamsPath = "/alert/api/v2/teams"

var client *resty.Client

type authSuccess struct {
}
type authError struct {
}

type Error struct {
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

func init() {
	client = resty.New()
	client.SetDebug(false)
	client.SetTimeout(1 * time.Minute)
	// Try getting Accounts API base URL from env var
	apiURL := os.Getenv("API_ADDR")
	if apiURL == "" {
		apiURL = defaultBaseURL
	}
	client.SetHostURL(apiURL)
	// Setting global error struct that maps to Form3's error response
	client.SetError(&Error{})
}
func SetConfig(config *viper.Viper) {
	defaultBaseURL = config.GetString("notifyserver.baseuri")
	mailPath = config.GetString("notifyserver.mail")
	teamsPath = config.GetString("notifyserver.teams")
}

func getAPIError(resp *resty.Response) error {
	apiError := resp.Error().(*Error)
	return fmt.Errorf("request failed [%s]: %s", apiError.Code, apiError.Message)
}
func printResponse(err error, resp *resty.Response) {
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
}
func SendTeams(jsonTeams *model.StNotifyTeams) (*resty.Response, error) {
	resp, err := client.R().
		SetBody(jsonTeams).
		SetResult(authSuccess{}).
		SetError(&authError{}).
		Post(teamsPath)

	if err != nil {
		return nil, fmt.Errorf("send teams to notify server failed: %s", err)
	}
	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}
	printResponse(err, resp)
	return resp, nil
}

func SendMail(jsonMail *model.StNotifyMail) (*resty.Response, error) {
	resp, err := client.R().
		SetBody(jsonMail).
		SetResult(authSuccess{}).
		SetError(&authError{}).
		Post(mailPath)

	if err != nil {
		return nil, fmt.Errorf("send mail to notify server failed: %s", err)
	}
	if resp.Error() != nil {
		return nil, getAPIError(resp)
	}

	printResponse(err, resp)
	return resp, nil
}
