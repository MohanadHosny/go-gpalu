package gpalu

import (
	"io"
	"log"
	"net/http"
)

func ReadBody(resp *http.Response, err error) []byte {
	if err != nil { log.Fatal(err) }
	defer resp.Body.Close()

	body, bErr := io.ReadAll(resp.Body)
	if bErr != nil { log.Fatal(bErr) }

	return body
}
