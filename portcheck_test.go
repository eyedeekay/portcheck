package pc

import (
	"log"
	"net"
	"testing"
)

func Test_HostIsOccupied(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	b, e := CheckLocal(8090)
	if e != nil {
		t.Fatal(e)
	}
	if !b {
		t.Fatal("Negative test failed, port should be occupied")
	}
	log.Println(b)
}

func Test_HostListIsOccupied(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	b, e := CheckLocal(8090, "127.0.0.1", "0.0.0.0")
	if e != nil {
		t.Fatal(e)
	}
	if !b {
		t.Fatal("Negative test failed, port should be occupied")
	}
	log.Println(b)
}

func Test_the_other_thing(t *testing.T) {
	b, e := CheckLocal(8090, "127.0.0.1")
	if e != nil {
		t.Fatal(e)
	}
	if b {
		t.Fatal("Test failed, port should not be occupied")
	}
	log.Println(b)
}

func Test_the_last_thing(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	i := FindLocal(8090)
	if i == 0 {
		t.Fatal("Port not found: i == ", i)
	}
	log.Println(i)
}
