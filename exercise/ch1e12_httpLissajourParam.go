package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/lissa", func(writer http.ResponseWriter, request *http.Request) {
		if err := request.ParseForm(); err != nil {
			log.Println(err)
		}
		//参数处理
		control := new(lissajousControl)
		control.cycles = cyclesDef
		control.res = resDef
		control.size = sizeDef
		control.nframes = nframesDef
		control.delay = delayDef
		//可以用反射改造一下
		dealWithParam("cycles", request, control)
		dealWithParam("res", request, control)
		dealWithParam("size", request, control)
		dealWithParam("nframes", request, control)
		dealWithParam("delay", request, control)
		//添加其他参数的处理...
		lissajousParams(writer, control)
	})
	log.Println("服务启动")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func dealWithParam(paramName string, request *http.Request, control *lissajousControl) {
	//这个方法以后可以用反射改造一下
	paramValueStr := request.Form.Get(paramName)
	if paramValueStr==""{
		log.Println(paramName+"参数为空")
		return
	}
	switch paramName {
	case "cycles":
		paramValue, err := strconv.Atoi(paramValueStr)
		if err != nil {
			log.Println(paramName + "参数转化失败")
			return
		}
		control.cycles = paramValue
	case "res":
		paramValue, err := strconv.ParseFloat(paramValueStr,64)
		if err != nil {
			log.Println(paramName + "参数转化失败")
			return
		}
		control.res = paramValue
	case "size":
		paramValue, err := strconv.Atoi(paramValueStr)
		if err != nil {
			log.Println(paramName + "参数转化失败")
			return
		}
		control.size = paramValue
	case "nframes":
		paramValue, err := strconv.Atoi(paramValueStr)
		if err != nil {
			log.Println(paramName + "参数转化失败")
			return
		}
		control.size = paramValue
	case "delay":
		paramValue, err := strconv.Atoi(paramValueStr)
		if err != nil {
			log.Println(paramName + "参数转化失败")
			return
		}
		control.delay = paramValue
	}
}

const (
	cyclesDef  = 5
	resDef     = 0.001
	sizeDef    = 100
	nframesDef = 64
	delayDef   = 8
)

type lissajousControl struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

//调色板(一个颜色数组)
var myPalette = []color.Color{color.White, color.RGBA{G: 0xff, B: 0x3e, A: 0xff}}

func lissajousParams(writer http.ResponseWriter, params *lissajousControl) {
	log.Println(params.cycles)
	log.Println(params.res)
	log.Println(params.size)
	log.Println(params.nframes)
	log.Println(params.delay)
	rand.Seed(time.Now().Unix())
	freq:=rand.Float64()*3.0
	anim:=gif.GIF{LoopCount:params.nframes}//gif动画对象
	phase:=0.0
	for i:=0;i<params.nframes;i++{
		rect:=image.Rect(0,0,2*params.size+1,2*params.size+1)
		img:=image.NewPaletted(rect,myPalette)
		for t:=0.0;t<float64(params.cycles)*2*math.Pi;t+=params.res{
			x:=math.Sin(t)
			y:=math.Sin(t*freq+phase)
			img.SetColorIndex(params.size+int(x*float64(params.size)+0.5),
				params.size+int(y*float64(params.size)+0.5),1)
		}

		phase+=0.1
		anim.Delay=append(anim.Delay,params.delay)
		anim.Image=append(anim.Image,img)
	}
	gif.EncodeAll(writer,&anim)
	log.Println("图像输出")
}
