# DDoS

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/DDoS/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/DDoS?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/DDoS.svg?branch=master)](https://travis-ci.org/Konstantin8105/DDoS)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/DDoS)](https://goreportcard.com/report/github.com/Konstantin8105/DDoS)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/DDoS/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/DDoS?status.svg)](https://godoc.org/github.com/Konstantin8105/DDoS)

DDoS attack. Creating infinite http GET requests.
If you need more information, then please see [wiki](https://en.wikipedia.org/wiki/Denial-of-service_attack#Defense_techniques).

**Library created just for education task.**

Minimal example of using:

```golang
func main() {
	workers := 100
	d, err := ddos.New("http://127.0.0.1:80", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://127.0.0.1:80")
	// Output: DDoS attack server: http://127.0.0.1:80
}
```
