package twiliocore

import (
	"encoding/json"
	"net/url"
	"errors"
)

type Message struct {
	Sid         string `json:"sid"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	DateSent    string `json:"date_sent"`
	AccountSid  string `json:"account_sid"`
	From        string `json:"from"`
	To          string `json:"to"`
	Body        string `json:"body"`
	NumSegments string `json:"num_segments"`
	Status      string `json:"status"`
	Direction   string `json:"direction"`
	Price       string `json:"price"`
	PriceUnit   string `json:"price_unit"`
	ApiVersion  string `json:"api_version"`
	Uri         string `json:"uri"`
}



func (client *TwilioClient)NewMessage(from string, to string, body string, mediaurl string) (*Message, error) {
	var message *Message

	params := url.Values{}
	params.Set("From", from)
	params.Set("To", to)

	if body == "" && mediaurl == ""{
		return message, errors.New("Must have either a body or a mediaURL section.")
	}
	
	if body != "" && mediaurl != ""{
		return message, errors.New("Must have EITHER a body or a mediaURL section. Not both.")
	}
	
	if body != ""{
		params.Set("Body", body)
	}
	if mediaurl != ""{
		params.Set("MediaUrl", mediaurl)
	}

	res, err := client.post(params, client.RootUrl()+"/Messages.json")

	if err != nil {
		return message, err
	}

	message = new(Message)
	err = json.Unmarshal(res, message)

	return message, err
}
