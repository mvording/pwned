package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/mvording/pwned/request"
)

func parseFileRequests(fileName string) ([]string, error) {
	requests := []string{}

	if len(fileName) > 0 {
		b, err := ioutil.ReadFile(fileName)
		if err != nil {
			return requests, err
		}
		body := strings.Replace(string(b), "\r", "", -1)
		requests = strings.Split(body, "\n")
	}

	return requests, nil
}

func parsePipedRequests() []string {
	requests := []string{}

	// see if data is being piped into process
	info, _ := os.Stdin.Stat()

	if info.Size() > 0 || info.Mode()&os.ModeNamedPipe > 0 {
		reader := bufio.NewReader(os.Stdin)

		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if len(text) > 0 {
				requests = append(requests, text)
			}
		}
	}
	return requests
}

func main() {

	var account, password, mode, file, format string
	flag.StringVar(&account, "account", "", "typically an email address")
	flag.StringVar(&password, "password", "", "")
	flag.StringVar(&mode, "mode", "", "'breach' or 'password' - specify when using a pipe or file to send multiple requests")
	flag.StringVar(&file, "file", "", "filename to process")
	flag.StringVar(&format, "format", "text", "output format - text or json")
	flag.Parse()

	var requests []string
	var err error

	if len(mode) > 0 {

		if len(file) > 0 {
			requests, err = parseFileRequests(file)
			if err != nil {
				println("error: " + err.Error())
				return
			}
		} else {
			requests = parsePipedRequests()
		}
	} else {
		if len(account) > 0 {
			requests = []string{account}
			mode = "breach"
		}

		if len(password) > 0 {
			requests = []string{password}
			mode = "password"
		}
	}

	switch mode {
	case "breach":
		breaches := request.Breaches(requests)
		if format == "json" {
			out, _ := json.MarshalIndent(breaches, "", "  ")
			println(string(out))
		} else {
			for i := 0; i < len(breaches); i++ {
				if err != nil {
					println("error: " + breaches[i].Error)
				} else {
					count := strconv.Itoa(len(breaches[i].Breaches) + len(breaches[i].Pastes))
					print(breaches[i].Account + "=" + count + " // ")
					for j := 0; j < len(breaches[i].Breaches); j++ {
						print(breaches[i].Breaches[j].Name + " ")
					}
					for j := 0; j < len(breaches[i].Pastes); j++ {
						print(breaches[i].Pastes[j].Source + "(" + breaches[i].Pastes[j].ID + ") ")
					}
					println()
				}
			}
		}
	case "password":
		passwords := request.PasswordCount(requests)
		if format == "json" {
			out, _ := json.MarshalIndent(passwords, "", "  ")
			println(string(out))
		} else {
			for i := 0; i < len(passwords); i++ {
				if err != nil {
					println("error: " + passwords[i].Error)
				} else {
					println(passwords[i].Password + "=" + strconv.Itoa(passwords[i].Count))
				}
			}
		}
	}

}
