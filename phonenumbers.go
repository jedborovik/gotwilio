package gotwilio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AvailablePhoneNumbersResponse struct {
	Numbers []AvailablePhoneNumberResponse `json:"available_phone_numbers"`
	Url     string                         `json:"uri"`
}

type AvailablePhoneNumberResponse struct {
	FriendlyName string `json:"friendly_name"`
	PhoneNumber  string `json:"phone_number"`
	Region       string `json:"string"`
	// Lata int `json:"lata"`
	// RateCenter string `json:"rate_center"`
	// Latitude float32 `json:"latitude"`
}

type IncomingPhoneNumbersResponse struct {
	Numbers []IncomingPhoneNumberResponse `json:"incoming_phone_numbers"`
	Url     string                        `json:"uri"`
}

type IncomingPhoneNumberResponse struct {
	PhoneNumber string `json:"phone_number"`
}

func (twilio *Twilio) AvailablePhoneNumbers(iso string) (availablePhoneNumbersResponse *AvailablePhoneNumbersResponse, exception *Exception, err error) {

	twilioUrl := twilio.BaseUrl + "/Accounts/" + twilio.AccountSid +
		"/AvailablePhoneNumbers/" + iso + "/Local.json"

	formValues := url.Values{}
	formValues.Set("ExcludeAllAddressRequired", "false")
	formValues.Set("ExcludeLocalAddressRequired", "false")
	formValues.Set("ExcludeForeignAddressRequired", "false")

	res, err := twilio.get(formValues, twilioUrl)
	if err != nil {
		// Actually add the stuff here
		return
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusCreated {

	}

	availablePhoneNumbersResponse = new(AvailablePhoneNumbersResponse)
	err = json.Unmarshal(responseBody, availablePhoneNumbersResponse)

	return
}

func (twilio *Twilio) IncomingPhoneNumbers() (incomingPhoneNumbersResponse *IncomingPhoneNumbersResponse, exception *Exception, err error) {
	twilioUrl := twilio.BaseUrl + "/Accounts/" + twilio.AccountSid + "/IncomingPhoneNumbers.json"

	res, err := twilio.get(url.Values{}, twilioUrl)
	if err != nil {
		return
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	incomingPhoneNumbersResponse = new(IncomingPhoneNumbersResponse)
	err = json.Unmarshal(responseBody, incomingPhoneNumbersResponse)
	return
}
