package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

type Result struct {
	Input       string `json:"input"`
	BasicDomain string `json:"base"`
	TopLevel    string `json:"tld"`
	SecondLevel string `json:"sld"`
	ThirdLevel  string `json:"trd"`
	Error       error  `json:"error"`
}

func main() {
	if len(os.Args) >= 2 {
		fmt.Println("Returns parsed domains using the Public Suffix List.\n\nUsage:\n\tcat domains.txt | publicsuffix")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		domainName := input

		// Parse the input as a URL
		u, err := url.Parse(domainName)
		if err == nil && u.Host != "" {
			domainName = u.Host
		}

		parsed, err := publicsuffix.Parse(domainName)
		if err != nil {
			result := Result{Input: domainName, Error: err}
			jsonData, _ := json.Marshal(result)
			fmt.Println(string(jsonData))
			continue
		}
		tld := parsed.TLD
		sld := parsed.SLD
		trd := parsed.TRD
		sldPlusTld := sld + "." + tld

		result := Result{Input: input, BasicDomain: sldPlusTld, TopLevel: tld, SecondLevel: sld, ThirdLevel: trd, Error: err}
		jsonData, _ := json.Marshal(result)
		fmt.Println(string(jsonData))
	}
}
