package vcodeHMAC

import (
	"fmt"
	"testing"
)

func Test_getHostNoPort(t *testing.T) {
	givenHost := "lolcathost"
	urlString := fmt.Sprintf("http://%s/meow", givenHost)

	host, err := getHost(urlString)

	if err != nil {
		t.Errorf("Unexpected error in getHost(): %v", err)
	}
	if givenHost != host {
		t.Errorf("expection does not match result \"%s\" != \"%s\"", givenHost, host)
	}

}

func Test_getHostWithPort(t *testing.T) {
	givenHost := "lolcathost"
	givenPort := "3000"
	urlString := fmt.Sprintf("http://%s:%s/meow", givenHost, givenPort)

	host, err := getHost(urlString)

	if err != nil {
		t.Errorf("Unexpected error in getHost(): %v", err)
	}
	if givenHost != host {
		t.Errorf("expection does not match result \"%s\" != \"%s\"", givenHost, host)
	}
}
