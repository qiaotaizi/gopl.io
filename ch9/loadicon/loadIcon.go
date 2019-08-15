package loadicon

import (
	"image"
	"sync"
)

var (
	loadIconsOnce sync.Once
	icons         map[string]image.Image
)

func loadIcons() {
	icons = map[string]image.Image{
		"spade.png": loadIcon("spade.png"),
		"hearts.png": loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png": loadIcon("clubs.png"),
	}
}

//从map中获取已经加载好的icon
//因为函数可能被并发访问
//这里使用sync.Once锁来保证初始化函数loadIcons只会被调用一次
//之后再进行访问,loadIcons其实是不会执行的
func Icon(name string)image.Image{
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func loadIcon(name string) image.Image {
	return image.NewCMYK(image.Rect(0,0,1,1))
}
