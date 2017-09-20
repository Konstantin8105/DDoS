package ddos_test

import (
	"testing"

	ddos "github.com/Konstantin8105/DDoS"
)

func TestNewDDoS(t *testing.T) {
	d := ddos.NewDDoS("127.0.0.1")
	if d == nil {
		t.Error("Cannot create a new ddos structure")
	}
}
