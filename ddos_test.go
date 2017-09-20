package ddos_test

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	ddos "github.com/Konstantin8105/DDoS"
	freeport "github.com/Konstantin8105/FreePort"
)

func TestNewDDoS(t *testing.T) {
	d := ddos.NewDDoS("127.0.0.1", 0)
	if d == nil {
		t.Error("Cannot create a new ddos structure")
	}
}

func TestDDoS(t *testing.T) {
	port, err := freeport.Get()
	if err != nil {
		t.Errorf("Cannot found free tcp port. Error = ", err)
	}
	createServer(port, t)

	url := "http://127.0.0.1:" + strconv.Itoa(port)
	d := ddos.NewDDoS(url, 100)
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	success, amount := d.Result()
	if success == 0 || amount == 0 {
		t.Errorf("Negative result of DDoS attack.\nSuccess requests = %v.\nAmount requests = %v", success, amount)
	}
}

// Create a simple go server
func createServer(port int, t *testing.T) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		})
		if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
			t.Fatalf("Server is down. %v", err)
		}
	}()
}
