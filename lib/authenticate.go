package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PingAuthService(token string) {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
