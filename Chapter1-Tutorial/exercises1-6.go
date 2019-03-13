// Fetchall fetches URLs in parallel and reports their times and sizes

// 11. Try fetchall with longer argument lists, such as samples from the top
// million web sites available at alexa.com. How does the program behave if a
// web site just doesn't respond.

// NOTE: INCOMPLETE. COULD NOT GET ALEXA API TO WORK.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // recieve from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// adding headers
// client := &http.Client{
// 	CheckRedirect: redirectPolicyFunc,
// }

// resp, err := client.Get("http://example.com")

func fetch(url string, ch chan<- string) {
	start := time.Now()
	// add http:// if not there in url
	if strings.HasPrefix(url, "http://") == false {
		url = "http://"+url
	}
	
	// load secret keys
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessKey := os.Getenv("S3_BUCKET")
	client := &http.Client {
		Authorization: "AWS4-HMAC-SHA256 Credential=<%s>/20171114/us-west-1/AlexaTopSites/aws4_request, SignedHeaders=host;x-amz-date, Signature=<GENERATED_AUTH_V4_SIGNATURE>"%accessKey,
		Content-Type: "application/xml",
		X-Amz-Date: "20171114T190104Z",
		Accept: "application/xml",
	}
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// or to print data too use this instead: 
	// nbytes, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}