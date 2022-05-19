package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	urls := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// this runs in series which take a lot of time
	// for _, link := range urls {
	// 	checkLink(link) // this is a "blocking call"
	// }

	// this way you start a go routine to "paralelize" the code execution
	// by adding the "go" command before the function a "routine" is called
	// a routine starts and allows more routines to run while the the routine itself if "waiting"
	// this is used to handle "blocking" code

	c := make(chan string) // a channel is a type of variable that allows the go "routines" to comunicate
	for _, link := range urls {
		go checkLink(link, c)
		// fmt.Println(<-c) // if this statement is witin the loop the code takes ~5 sec to run and is NOT parelalized
	}

	// loop with a certain amunt of loops (in this case the lenght of "urls")
	// start index; number of loops; (equivalent to "for range urls")
	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	for l := range c {
		// if using only "func()", the var "l" will alwyas contain the same value
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c) // This way
			fmt.Println("code ran in ", time.Since(start))
		}(l)
	}

	fmt.Println("code ran in ", time.Since(start))
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding")
		c <- "Might be down, I think..."
		return
	}

	fmt.Println(link, "is working!")
	c <- link
}
