package pc

import (
	"log"
	"net"
	"testing"
)

func Test_thing(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	b, e := CheckLocal(8090)
	if e != nil {
		t.Fatal(e.Error())
	}
	log.Println(b)
}

func Test_other_thing(t *testing.T) {
	n, e := net.Listen("tcp", ":8090")
	if e != nil {
		t.Fatal(e)
	}
	defer n.Close()
	b, e := CheckLocal(8090, "127.0.0.1", "128.0.0.1")
	if e != nil {
		t.Fatal(e.Error())
	}
	log.Println(b)
}
