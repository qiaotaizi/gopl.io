package main

import (
	"fmt"
	"github.com/qiaotaizi/gopl.io/ch5/links"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

//defer关键字的使用

func main() {
	for _, url := range os.Args[1:] {
		err := title(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//打印html文档中title节点下首个子节点的内容
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	links.ForEachNode(doc, visitNode, nil)
	return nil
}
