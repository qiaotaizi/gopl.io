package memo3

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

//这版优化,缩小了临界区的范围
//仅当操作cache时才上锁
//而耗时较长的网络io操作是不处于临界区的
//但是这样做的问题是
//同样的请求由于http访问一直没有结束
//后续的请求并没有利用起缓存来
//也就是发生了所谓的缓存穿透
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
