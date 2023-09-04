package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

type Result struct {
	Name        string `json:"name"`
	TopLevel    string `json:"tld"`
	SecondLevel string `json:"sld"`
	ThirdLevel  string `json:"trd"`
	BasicDomain string `json:"base"`
	Error       error  `json:"error"`
}

func main() {
	if len(os.Args) >= 2 {
		fmt.Println("Returns parsed domains using the Public Suffix List.\n\nUsage:\n\tcat domains.txt | publicsuffix")
		os.Exit(0)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		domainName := scanner.Text()
		parsed, err := publicsuffix.Parse(domainName)
		if err != nil {
			result := Result{Name: domainName, Error: err}
			jsonData, _ := json.Marshal(result)
			fmt.Println(string(jsonData))
			continue
		}
		tld := parsed.TLD
		sld := parsed.SLD
		trd := parsed.TRD
		sldPlusTld := sld + "." + tld

		result := Result{Name: domainName, BasicDomain: sldPlusTld, TopLevel: tld, SecondLevel: sld, ThirdLevel: trd, Error: err}
		jsonData, _ := json.Marshal(result)
		fmt.Println(string(jsonData))
	}
}
