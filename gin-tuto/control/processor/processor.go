package processor

import (
	"log"
	"net/http"

	"webzen.com/notifyhandler/control/notifysender"
	"webzen.com/notifyhandler/model"
)

type EMaillConvertable interface {
	ToEMail() (*model.NotifyEMail, error)
}
type TeamsConvertable interface {
	ToTeams() (*model.NotifyTeams, error)
}

func EMail(s EMaillConvertable) map[string]interface{} {
	var resultSet = make(map[string]interface{})

	jsonEMail, err := s.ToEMail()
	if err != nil {
		log.Println("raise error", err)
		resultSet["responseCode"] = http.StatusBadRequest
		resultSet["errorTitle"] = "request struct error"
		return resultSet
	}

	resp, err := notifysender.SendEMail(jsonEMail)
	if err != nil {
		log.Println(err)
	}
	if resp != nil {
		log.Println("resp", resp)
		resultSet["responseCode"] = resp.StatusCode()
	}
	return resultSet
}
func Teams(s TeamsConvertable) map[string]interface{} {
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
