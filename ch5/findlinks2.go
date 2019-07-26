package main

import (
	"fmt"
	"golang.org/x/net/html"
	"gopl.io/ch5/visit"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		links, err := findlinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2 :%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findlinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	//获取页面中的a标签
	return visit.Visit(nil, doc), nil
}