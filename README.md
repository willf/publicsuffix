#Public Suffix

Simple Go command that prints domain information using the [Public Suffix List](https://publicsuffix.org/).

## Installation

Assuming you have Go installed and your `$GOPATH` set:

1. Clone the repository
2. Run `go install`

## Usage

```
$ echo 'www.example.com' | publicsuffix
{"name":"www.example.com","tld":"com","sld":"example","trd":"www","base":"example.com","error":null}

$ cat domains.txt | publicsuffix > domains.json
```
