package twiliocore

import (
	"encoding/json"
	"net/url"
)

type Account struct {
	Sid         		string `json:"sid"`
	Friendly_name		string `json:"friendly_name"`
	Type				string `json:"type"`
	Status				string `json:"status"`
	Date_created		string `json:"date_created"`
	Date_updated		string `json:"date_updated"`
	Auth_token			string `json:"auth_token"`
	Uri					string `json:"uri"`
	Subresource_uris	map[string]string `json:"subresource_uris"`
}

func (inst *TwilioClient)Account()(*Account,error){
	params := url.Values{}
	var account *Account
	
	res, err := inst.get(params, inst.RootUrl()+".json")

	if err != nil {
		return account, err
	}

	account = new(Account)
	err = json.Unmarshal(res, account)

	return account, err
	
}
