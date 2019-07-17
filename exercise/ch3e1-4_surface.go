package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

//练习 3.1： 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素
//（虽然许多SVG渲染器会妥善处理这类问题） 。修改程序跳过无效的多边形。
//练习 3.2： 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或asaddle图案?
//练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色(#ff0000)，谷部将是蓝色
//(#0000ff)。
//练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返
//回SVG数据给客户端。服务器必须设置Content-Type头部：
//w.Header().Set("Content-Type", "image/svg+xml")
//（这一步在Lissajous例子中不是必须的，因为服务器使用标准的PNG图像格式，可以根据前
//面的512个字节自动输出对应的头部。） 允许客户端通过HTTP请求参数设置高度、宽度和颜
//色等参数。

func main() {
	http.HandleFunc("/surface", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "image/svg+xml")
		surface(writer)
	})
	err:=http.ListenAndServe("localhost:8888",nil)
	if err!=nil {
		log.Fatalf("surface: %v\n",err)
	}
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func surface(out io.Writer) {
	fmt.Fprintf(out,"<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay,z := corner(i+1, j)
			bx, by,_ := corner(i, j)
			cx, cy,_ := corner(i, j+1)
			dx, dy,_ := corner(i+1, j+1)
			clr:=color_(z);
			fmt.Fprintf(out,"<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy,clr)
		}
	}
	fmt.Fprintln(out,"</svg>")
}

func corner(i, j int) (float64, float64,float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy,z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//输入z轴的值,映射颜色
func color_(z float64) string{
	clr:=int(k*z+b)
	//整数转16进制字符串输出
	return fmt.Sprintf("#%x",clr)
}

const b  = 3855/2
const k  = 3840-b
