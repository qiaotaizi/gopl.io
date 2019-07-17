package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

//生成李萨如图像
//运行命令: go run ./ch1/lissajous.go > out.gif
//
//func main() {
//
//	lissajous(os.Stdout)
//
//}

//调色板(一个颜色数组)
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	//书上的代码有点问题
	//需要设置一个随机种子
	//否则每次生成的随机数都是相同的序列
	rand.Seed(time.Now().Unix())
	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount:nframes}//gif动画对象
	phase:=0.0
	for i:=0;i<nframes;i++{
		rect:=image.Rect(0,0,2*size+1,2*size+1)
		img:=image.NewPaletted(rect,palette)
		for t:=0.0;t<cycles*2*math.Pi;t+=res{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(size+int(x*size+0.5),size+int(y*size+0.5),blackIndex)
		}

		phase+=0.1
		anim.Delay=append(anim.Delay,delay)
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(out,&anim)
}
