package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png" //将png包匿名导入,如果没有这行导入语句,程序将无法识别png格式的图片,匿名导入这个包,是为了可以初始化包级变量Decoder
	"io"
	"os"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input fomart=", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func main() {
	if err := toJPEG(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
