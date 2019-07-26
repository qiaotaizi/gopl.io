package visit

import (
	"golang.org/x/net/html"
)

//递归获取页面中的a标签
func Visit(links []string, n *html.Node) []string {
	//n是dom节点且为a标签
	if n.Type == html.ElementNode && n.Data == "a" {
		//遍历标签的属性,查找href属性
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
				break
			}
		}
	}

	//声明c变量,赋值为n节点的第一个子节点
	//循环每进行一次,将c赋值为c的下一个兄弟节点
	//直到没有下一个兄弟节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}
