package processor

import (
	"log"
	"net/http"

	"webzen.com/notifyhandler/control/notifysender"
	"webzen.com/notifyhandler/model"
)

type MaillCenter interface {
	ToMail() (*model.StNotifyMail, error)
}
type TeamsCenter interface {
	ToTeams() (*model.StNotifyTeams, error)
}

func Mail(s MaillCenter) map[string]interface{} {
	var resultSet = make(map[string]interface{})

	jsonMail, err := s.ToMail()
	if err != nil {
		log.Println("raise error", err)
		resultSet["responseCode"] = http.StatusBadRequest
		resultSet["errorTitle"] = "request struct error"
		return resultSet
	}

	resp, err := notifysender.SendMail(jsonMail)
	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		log.Println("resp", resp)
		resultSet["responseCode"] = resp.StatusCode()
	}
	return resultSet
}
func Teams(s TeamsCenter) map[string]interface{} {
	var resultSet = make(map[string]interface{})

	jsonTeams, err := s.ToTeams()
	if err != nil {
		log.Println("raise error", err)
		resultSet["responseCode"] = http.StatusBadRequest
		resultSet["errorTitle"] = "request struct error"
		return resultSet
	}

	resp, err := notifysender.SendTeams(jsonTeams)
	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		log.Println("resp", resp)
		resultSet["responseCode"] = resp.StatusCode()
	}
	return resultSet
}
