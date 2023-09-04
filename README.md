#Public Suffix

Simple Go command that prints domain information using the [Public Suffix List](https://publicsuffix.org/).

## Installation

Assuming you have Go installed and your `$GOPATH` set:

1. Clone the repository
2. Run `go install`

## Usage

Standard input is used to read the domain name. The output is a JSON object with the following fields:

- `input`: The input
- `tld`: The top level domain
- `sld`: The second level domain
- `trd`: The third level domain
- `base`: The base domain
- `error`: Any error that occurred

Input can be a full URL or just a domain name. If a URL is provided, the scheme and path are ignored.

```
$ echo 'www.example.com' | publicsuffix
{"input":"www.example.com","tld":"com","sld":"example","trd":"www","base":"example.com","error":null}

$ echo 'https://www.example.com' | publicsuffix
{"input":"www.example.com","tld":"com","sld":"example","trd":"www","base":"example.com","error":null}


$ cat domains.txt | publicsuffix > domains.json
```
