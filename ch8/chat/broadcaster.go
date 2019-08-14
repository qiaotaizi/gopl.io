package main

type client chan<- string //底层是一个只写的channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) //管理所有客户端连接
	for {
		select {
		case msg := <-messages:
			//有消息发来时,
			//遍历所有的客户端,向其channel发送消息
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			//当有新的客户端接入时,在map中管理这个客户端连接
			clients[cli] = true;

		case cli := <-leaving:
			//当有客户端退出时,删除map中相应的连接对象,并关闭其channel
			delete(clients, cli)
			close(cli)
		}
	}
}
