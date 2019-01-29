package pc

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	//"log"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// CheckLocal determines if a port is *in use*. If a port is in use locally, it
// will return "true" if the port is used, and will return "false" if the port
// is unavailable. If it encounters an error, it will return true and the error.
// If it you pass it a list of hosts, it will return true if *any* of them is
// in use.
func CheckLocal(port int, hosts ...string) (bool, error) {
	p := ":" + strconv.Itoa(port)
	if len(hosts) < 1 {
		if temp, err := net.Listen("tcp", p); err == nil {
			temp.Close()
			return false, nil
		} else if strings.Contains(err.Error(), "in use") {
			return true, nil
		} else {
			return true, err
		}
	}
	for _, h := range hosts {
		if temp, err := net.Listen("tcp", h+p); err == nil {
			temp.Close()
			return false, nil
		} else if strings.Contains(err.Error(), "in use") {
			return true, nil
		} else {
			return true, err
		}
	}
	return false, nil
}

// ShortCheckLocal is CheckLocal but ignoring errors
func ShortCheckLocal(port int) bool {
	b, _ := CheckLocal(port)
	return b
}

// SCL is a shortened alias for ShortCheckLocal
func SCL(port int) bool {
	return ShortCheckLocal(port)
}

func FindLocal(port int) int {

	server, err := net.Listen("tcp", ":0")

	if err != nil {
		return 0
	}

	defer server.Close()

	hostString := server.Addr().String()

	_, portString, err := net.SplitHostPort(hostString)
	if err != nil {
		return 0
	}

	n, err := strconv.Atoi(portString)
	if err != nil {
		return 0
	}
	return n
}

// CheckRemote tries to connect to a remote port, and returns false if it fails.
// If it fails because of a timeout it returns no error, but if it fails because
// of any other reason it returns false and the error.
func CheckRemote(host string, port int) (bool, error) {
	conn, err := net.DialTimeout("tcp", host, time.Duration(600)*time.Second)
	defer conn.Close()

	if err, ok := err.(*net.OpError); ok && err.Timeout() {
		fmt.Printf("Timeout error: %s\n", err)
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

// CR is a shortened alias for CheckRemote, which returns true if a port on a
// remote server is open and false if appears to be closed or unavailable.
func CR(port int) (bool, error) {
	return CheckLocal(port)
}

// ShortCheckRemote is CheckRemote but ignoring errors
func ShortCheckRemote(host string, port int) bool {
	b, _ := CheckRemote(host, port)
	return b
}

// SCR is a shortened alias for ShortCheckRemote
func SCR(host string, port int) bool {
	return ShortCheckRemote(host, port)
}
