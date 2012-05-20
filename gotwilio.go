// Package gotwilio is a library for interacting with http://www.twilio.com/ API.
// It's very incomplete at the moment.
package gotwilio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Twilio struct {
	AccountSid string
	AuthToken  string
	BaseUrl    string
}

func NewTwilioClient(accountSid, authToken string) *Twilio {
	twilioUrl := "https://api.twilio.com/2010-04-01" // Should this be moved into a constant?
	return &Twilio{accountSid, authToken, twilioUrl}
}

// SendTextMessage uses Twilio to send a text message.
// See http://www.twilio.com/docs/api/rest/sending-sms for more information.
func (twilio *Twilio) SendTextMessage(from, to, body string) (string, error) {
	twilioUrl := twilio.BaseUrl + "/Accounts/" + twilio.AccountSid + "/SMS/Messages.json" // needs a better variable name

	formValues := url.Values{}
	formValues.Set("From", from)
	formValues.Set("To", to)
	formValues.Set("Body", body)

	return twilio.post(formValues, twilioUrl)
}

func (twilio *Twilio) post(formValues url.Values, twilioUrl string) (string, error) {
	var respError error
	req, _ := http.NewRequest("POST", twilioUrl, strings.NewReader(formValues.Encode())) // TODO: check error
	req.SetBasicAuth(twilio.AccountSid, twilio.AuthToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	resBody, _ := ioutil.ReadAll(resp.Body) // TODO: check error
	results := string(resBody)

	if resp.StatusCode != 200 {
		respError = errors.New(resp.Status)
	}
	return results, respError
}