package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mvording/pwned/model"
)

// Breaches : Retrieve count of password usages
func Breaches(accounts []string) []model.Breaches {
	response := make([]model.Breaches, len(accounts))

	var results int

	for i := 0; i < len(accounts); i++ {
		acct := strings.Trim(accounts[i], "\n\r\t ")
		if len(acct) > 0 {

			b, err := retrieveBreaches(acct)

			response[results].Account = acct
			response[results].Breaches = b

			if err != nil {
				response[results].Error = err.Error()
			} else {
				// retrieve pastes if the account is an email address
				if strings.Contains(acct, "@") {
					p, err2 := retrievePastes(acct)

					response[results].Pastes = p
					if err != nil {
						response[results].Error = err2.Error()
					}
				}
			}

			results++
		}
	}

	return response[:results]
}

// retrieveBreaches: call https://haveibeenpwned.com API
func retrieveBreaches(account string) ([]model.Breach, error) {
	var response []model.Breach

	if len(account) < 1 {
		return response, errors.New("Account not provided")
	}

	param := url.QueryEscape(account)

	url := "https://haveibeenpwned.com/api/v2/breachedaccount/" + param
	res, err := http.NewRequest("GET", url, nil)

	timeout := time.Duration(15 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(res)

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	} else {

		if err == nil && resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				err = json.Unmarshal(body, &response)
			}
		}
	}
	return response, err
}

// retrievePastes: call https://haveibeenpwned.com API
func retrievePastes(account string) ([]model.Paste, error) {
	var response []model.Paste

	if len(account) < 1 {
		return response, errors.New("Account not provided")
	}

	param := url.QueryEscape(account)

	url := "https://haveibeenpwned.com/api/v2/pasteaccount/" + param
	res, err := http.NewRequest("GET", url, nil)

	timeout := time.Duration(15 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(res)

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	} else {

		if err == nil && resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				err = json.Unmarshal(body, &response)
			}
		}
	}
	return response, err
}
