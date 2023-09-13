package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

type NullableString string

func (s NullableString) MarshalJSON() ([]byte, error) {
	if s == "" {
		return []byte("null"), nil
	}
	return json.Marshal(string(s))
}

type Result struct {
	Input        string         `json:"input"`
	BasicDomain  string         `json:"base"`
	TopLevel     string         `json:"tld"`
	SecondLevel  string         `json:"sld"`
	ThirdLevel   string         `json:"trd"`
	ErrorMessage NullableString `json:"error"`
}

func RFC3490Check(str string) (ok bool) {
	// make sure there are no extraneous things in the tld or the sld
	// regular expression for tld is: ^[a-z0-9\-]+$ more or less ...
	// can't start or end with a dash
	re := regexp.MustCompile(`^[a-zA-Z0-9-]{1,63}$`)
	return len(str) >= 1 && re.MatchString(str) && str[:1] != "-" && str[len(str)-1:] != "-"
}

func FinalCheck(tld string, sld string, trd string) (err error) {
	if tld == "" {
		return errors.New("tld is empty")
	}
	if sld == "" {
		return errors.New("sld is empty")
	}
	// make sure there are no extraneous things in the tld or the sld
	// regular expression for tld is: ^[a-z0-9\-]+$
	if !RFC3490Check(tld) {
		return errors.New("tld contains invalid characters")
	}
	// regular expression for sld is: ^[a-z0-9\-]+$
	if !RFC3490Check(sld) {
		return errors.New("sld contains invalid characters")
	}
	// regular expression for trd is: ^[a-z0-9\-]+$
	if trd != "" && !RFC3490Check(trd) {
		return errors.New("trd contains invalid characters")
	}
	return nil
}

func ParseHost(input string, host string) (result Result) {
	parsed, err := publicsuffix.Parse(host)
	if err != nil {
		return Result{Input: input, ErrorMessage: NullableString(err.Error())}
	} else {
		tld := parsed.TLD
		sld := parsed.SLD
		trd := parsed.TRD
		sldPlusTld := sld + "." + tld
		finalCheck := FinalCheck(tld, sld, trd)
		if finalCheck != nil {
			return Result{Input: input, ErrorMessage: NullableString(finalCheck.Error())}
		}
		return Result{Input: input, BasicDomain: sldPlusTld, TopLevel: tld, SecondLevel: sld, ThirdLevel: trd}
	}
}

func Parse(input string) (result Result) {
	u, err := url.Parse(input)
	if err == nil && u.Host != "" {
		return ParseHost(input, u.Host)
	} else {
		u, err = url.Parse("https://" + input)
		if err == nil && u.Host != "" {
			return ParseHost(input, u.Host)
		}
	}
	return ParseHost(input, input)
}

func main() {
	if len(os.Args) >= 2 {
		fmt.Println("Returns parsed domains using the Public Suffix List.\n\nUsage:\n\tcat domains.txt | publicsuffix")
	} else {

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			input := scanner.Text()
			result := Parse(input)
			jsonData, _ := json.Marshal(result)
			fmt.Println(string(jsonData))
		}
	}
}
