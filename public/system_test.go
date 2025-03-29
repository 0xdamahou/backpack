package public

import (
	"log"
	"testing"
)

func TestBackpackPublicClient_Status(t *testing.T) {
	status, err := bpbc.Status()
	if err != nil {
		return
	}
	log.Println(status)
}

func TestBackpackPublicClient_Ping(t *testing.T) {
	ping, err := bpbc.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(ping)
}

func TestBackpackPublicClient_GetSystemTime(t *testing.T) {
	time, err := bpbc.GetSystemTime()
	if err != nil {
		return
	}
	log.Println(time)
}
