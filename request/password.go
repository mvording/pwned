package request

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mvording/pwned/model"
)

// PasswordCount : Retrieve count of password usages
func PasswordCount(passwords []string) []model.Password {
	if len(passwords) < 1 {
		return []model.Password{}
	}
	response := make([]model.Password, len(passwords))

	var results int

	for i := 0; i < len(passwords); i++ {
		pass := strings.Trim(passwords[i], "\n\r\t ")
		if len(pass) > 0 {

			hash := calcHash(pass)
			count, err := retrieveByHash(hash)

			response[results].Password = pass
			response[results].Count = count
			if err != nil {
				response[results].Error = err.Error()
			}

			results++
		}
	}

	return response[:results]
}

// calcHash : return MD-1 hash as hex of password
func calcHash(password string) string {
	h := sha1.New()
	io.WriteString(h, password)
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

// retrieveByHash : call https://haveibeenpwned.com API
func retrieveByHash(hash string) (int, error) {
	if len(hash) != 40 {
		return -1, errors.New("Invalid hash length")
	}

	url := "https://api.pwnedpasswords.com/range/" + hash[:5]
	resp, err := http.Get(url)

	count := 0
	if err == nil {

		body, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			lines := strings.Split(string(body), "\r\n")

			for i := 0; i < len(lines); i++ {
				line := strings.Split(lines[i], ":")
				if len(line) > 1 {
					if line[0] == hash[5:] {
						count, err = strconv.Atoi(line[1])
						break
					}
				}
			}
		}
	}

	return count, err
}
