package pc

import (
	"log"
	"net"
	"testing"
)

func Test_HostPortIsOccupied(t *testing.T) {
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

func Test_HostPortListIsOccupied(t *testing.T) {
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

func Test_HostPortIsAvailable(t *testing.T) {
	b, e := CheckLocal(8090, "127.0.0.1")
	if e != nil {
		t.Fatal(e)
	}
	if b {
		t.Fatal("Test failed, port should not be occupied")
	}
	log.Println(b)
}

func Test_FindNewHostPort(t *testing.T) {
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

func Test_FindNewHostPortShort(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	i := FindLocal(8091)
	if i == 0 {
		t.Fatal("Port not found: i == ", i)
	}
	log.Println(i)
}
