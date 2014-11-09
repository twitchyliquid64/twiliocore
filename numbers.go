package twiliocore

import (
	"encoding/json"
	"net/url"
	"errors"
)

type AvailableNumber struct {
	FriendlyName		string `json:"friendly_name"`
	PhoneNumber			string `json:"phone_number"`
	IsoCountry			string `json:"iso_country"`
	Capabilities		map[string]bool `json:"capabilities"`
}

type AvailableNumbers struct {
	Uri					string `json:"uri"`
	Available_Numbers	[]AvailableNumber `json:"available_phone_numbers"`
}



func (inst *TwilioClient)AvailableNumbersByCountryLocal(ISOcountryCode string, needVoice bool, needSMS bool)(*AvailableNumbers,error){
	params := url.Values{}
	if needSMS{
		params.Set("SmsEnabled", "true")
	}
	if needVoice{
		params.Set("VoiceEnabled", "true")
	}
	
	var availableNumbers *AvailableNumbers
	
	res, err := inst.get(params, inst.RootUrl()+"/AvailablePhoneNumbers/" + ISOcountryCode + "/Local.json")

	if err != nil {
		return availableNumbers, err
	}

	availableNumbers = new(AvailableNumbers)
	err = json.Unmarshal(res, availableNumbers)


	return availableNumbers, err
	
}


func (inst *TwilioClient)AvailableNumbersByCountryMobile(ISOcountryCode string, needVoice bool, needSMS bool)(*AvailableNumbers,error){
	params := url.Values{}
	if needSMS{
		params.Set("SmsEnabled", "true")
	}
	if needVoice{
		params.Set("VoiceEnabled", "true")
	}
	
	var availableNumbers *AvailableNumbers
	
	res, err := inst.get(params, inst.RootUrl()+"/AvailablePhoneNumbers/" + ISOcountryCode + "/Mobile.json")

	if err != nil {
		return availableNumbers, err
	}

	availableNumbers = new(AvailableNumbers)
	err = json.Unmarshal(res, availableNumbers)

	if len(availableNumbers.Available_Numbers) > 0{
		return availableNumbers, errors.New("Error: No numbers given")
	}

	return availableNumbers, err
	
}
