package module5

import (
	"fmt"
	"net/http"
)

// GetExampleDotCom uses the "net/http" package to send a GET request to example.com
func GetExampleDotCom() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("something went wrong")
	}

	defer resp.Body.Close()
}
