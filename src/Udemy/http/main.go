package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// give me empty slice with 99999 elements
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(bs)

	//take reponse body and automatic logout the info into the terminal
	// io.Copy(dst Writer, src Reader) (written int64, err error)
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// this functions add the "functionality" logWriter to the the type
// "Write" which can also be used with interface "Writer"
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println((string(bs)))
	fmt.Println("Jsut wrote this bytes:", len(bs))
	return len(bs), nil
}
