package main

import (
	"fmt"
	"test/mystringutil"
	"net/http"
	"io/ioutil"
	"strings"
	"mvdan.cc/xurls"
)

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

 
func main() {
	fmt.Println("go hello")
	fmt.Println(mystringutil.Reverse("go hello"))

	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("err")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	urls := xurls.Strict().FindAllString(string(body), -1)
	for _, url := range urls {
		fmt.Println(url)
	}

}
