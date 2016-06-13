package mPkg

import (
	//"code.google.com/p/graphics-go/graphics/edge"
	"fmt"
	"image"
	"image/color"
	"image/png"
	//"github.com/Comdex/imgo"
	"encoding/gob"
	"errors"
	//"image/draw"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"code.google.com/p/graphics-go/graphics"
)

type Putpixel func(x, y int)

type TmodData struct {
	V    string
	Name string
	TT   []TPointCC
}

var (
	ModTTs      []TmodData = []TmodData{}
	pathSplit   string         //目录分割方式
	modlistfile []string       //获取文件列表
	modStep     int        = 2 //模板采样矩阵
	modZoom     float64    = 1 //模板放大倍数
	modbaei     float64    = 1 //模板采集点精简倍数
	modename    string         //模板名字
)

func Listfunc(path string, f os.FileInfo, err error) error {
	strRet, _ := os.Getwd()
	strRet += pathSplit

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	strRet += path //+ "\r\n"

	//用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
	ok := strings.HasSuffix(strRet, ".png")
	if ok {
		modlistfile = append(modlistfile, path) //将目录push到listfile []string中
	}
	//fmt.Println(ostype) // print ostype
	//fmt.Println(path) //list the file

	return nil
}

func getFileList(path string) string { //获取模板目录文件列表
	//var strRet string
	err := filepath.Walk(path, Listfunc) //

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	return " "
}

func getVn(path string) (string, string) { //截取对应真实字符和相关信息
	pp := strings.Split(path, pathSplit)
	if len(pp) < 1 {
		panic("no mod file")
	}
	//fmt.Println(pp, ">>", pathSplit)
	fn := pp[len(pp)-1]
	ret := strings.Split(fn, "_")
	if len(ret) < 2 {
		return "", ""
	}
	return ret[0], ret[1]
}

func PurificationMod(src []TPointCC, xBei float64) []TPointCC { //模板简化
	ret := []TPointCC{}
	//xBei := float64(0.5)
	for i := 0; i < len(src); i++ {
		src[i].P.X = int(float64(src[i].P.X) * xBei)
		src[i].P.Y = int(float64(src[i].P.Y) * xBei)
	}
	flaMap := make(map[string]bool)
	for i := 0; i < len(src); i++ {
		if _, ok := flaMap[strconv.Itoa(src[i].P.X)+"_"+strconv.Itoa(src[i].P.Y)]; !ok {
			flaMap[strconv.Itoa(src[i].P.X)+"_"+strconv.Itoa(src[i].P.Y)] = true
			ret = append(ret, src[i])
		}
	}

	for i := 0; i < len(ret); i++ {
		ret[i].P.X = int(float64(ret[i].P.X) / xBei)
		ret[i].P.Y = int(float64(ret[i].P.Y) / xBei)
	}
	return ret
}

func doLoadMods(p []string) {
	for _, value := range p {
		yuansrc, decodeErr := DecodeImage(value)
		if decodeErr != nil {
			panic(decodeErr)
		}
		//yuansrc = ImgZoom(yuansrc, 2)
		//newImg := Edge_Prewitt2(yuansrc)
		yuansrc = ImgZoom(yuansrc, modZoom)
		//newImg := Edge_Canny(yuansrc)
		piximg0 := Img2Matrix(yuansrc)
		piximg := FindEdges(piximg0)

		//piximg := Img2Matrix(yuansrc)

		TT := GetPoints(piximg, modStep)

		//TT := PurificationMod(TTtmp, modbaei)

		tv, tn := getVn(value)
		if tv != "" {
			ModTTs = append(ModTTs, TmodData{tv, tn, TT})
		}
		//fmt.Println(value, tv, tn)
		fmt.Print(".")
	}
}

func MakeMods(listpath string, modfile string) {
	modename = modfile
	ostype := runtime.GOOS
	if ostype == "windows" {
		pathSplit = "\\"
	} else {
		pathSplit = "/"
	}

	//fmt.Println(runtime.GOOS)
	getFileList(listpath)
	//fmt.Println(modlistfile)
	doLoadMods(modlistfile)
	file, err := os.Create(modename)
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := gob.NewEncoder(file)
	err2 := enc.Encode(ModTTs)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func LoadMods(modfile string) {
	modename = modfile
	file, err := os.Open(modename)
	if err != nil {
		fmt.Println(err)
		return
	}
	dec := gob.NewDecoder(file)
	err2 := dec.Decode(&ModTTs)

	if err2 != nil {
		fmt.Println(err2)
		return
	}
}

func FileDecodeImage(path string) (image.Image, error) {
	img, err := DecodeImage(path)
	return img, err
}

func Img2Matrix(img image.Image) (imgMatrix [][][]uint8) {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	src := ConvertToNRGBA(img)
	imgMatrix = NewRGBAMatrix(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := src.At(j, i)
			r, g, b, a := c.RGBA()
			imgMatrix[i][j][0] = uint8(r)
			imgMatrix[i][j][1] = uint8(g)
			imgMatrix[i][j][2] = uint8(b)
			imgMatrix[i][j][3] = uint8(a)
		}
	}
	return
}

func Img2Matrix2(img image.Image) (imgMatrix [][][]uint8) { //屏蔽Canny算法产生的边缘
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	src := ConvertToNRGBA(img)
	imgMatrix = NewRGBAMatrix(height, width)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := src.At(j, i)
			r, g, b, a := c.RGBA()
			if (i == 0) || (i == height-1) || (j == 0) || (j == width-1) {
				imgMatrix[i][j][0] = 0
				imgMatrix[i][j][1] = 0
				imgMatrix[i][j][2] = 0
				imgMatrix[i][j][3] = uint8(a)
			} else {
				imgMatrix[i][j][0] = uint8(r)
				imgMatrix[i][j][1] = uint8(g)
				imgMatrix[i][j][2] = uint8(b)
				imgMatrix[i][j][3] = uint8(a)
			}
		}
	}
	return
}

func Matrix2Img(imgMatrix [][][]uint8) image.Image {
	height := len(imgMatrix)
	width := len(imgMatrix[0])

	if height == 0 || width == 0 {
		panic(errors.New("The input of matrix is illegal!"))
	}

	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{imgMatrix[i][j][0], imgMatrix[i][j][1], imgMatrix[i][j][2], imgMatrix[i][j][3]})
		}
	}
	return nrgba
}

func SaveImgAsPNG(src image.Image, path string) {
	dirf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dirf.Close()

	png.Encode(dirf, src)
}

func RGB2BW(src [][][]uint8, lim uint8) [][][]uint8 {

	height := len(src)
	width := len(src[0])

	imgMatrix := NewRGBAMatrix(height, width)
	copy(imgMatrix, src)

	Rsum := uint(0)
	Gsum := uint(0)
	Bsum := uint(0)
	Ravg := uint8(0)
	Gavg := uint8(0)
	Bavg := uint8(0)
	//RGBavg := uint8(0)
	pixCount := uint(height * width)

	//fmt.Println("pixCount:", pixCount)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			Rsum += uint(imgMatrix[i][j][0])
			Gsum += uint(imgMatrix[i][j][1])
			Bsum += uint(imgMatrix[i][j][2])
		}
	}
	Ravg = uint8(Rsum / pixCount)
	Gavg = uint8(Gsum / pixCount)
	Bavg = uint8(Bsum / pixCount)

	getNewPix := func(RGBValue uint8, RGBavg uint8) uint8 {
		if RGBavg > lim {
			RGBavg = lim
		}
		if RGBValue < RGBavg {
			return 0 //RGBValue
		} else {
			return 255
		}
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			Tr := getNewPix(imgMatrix[i][j][0], Ravg)
			Tg := getNewPix(imgMatrix[i][j][1], Gavg)
			Tb := getNewPix(imgMatrix[i][j][2], Bavg)
			Trgb := uint8(255)
			if (Tr == 0) || (Tg == 0) || (Tb == 0) {
				Trgb = 0
			}

			imgMatrix[i][j][0] = Trgb
			imgMatrix[i][j][1] = Trgb
			imgMatrix[i][j][2] = Trgb
		}
	}

	return imgMatrix
}

func ImgPlusImg(img1, img2 [][][]uint8) [][][]uint8 {
	height1 := len(img1)
	width1 := len(img1[0])
	height2 := len(img2)
	width2 := len(img2[0])
	getmax := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	height := getmax(height1, height2)
	width := getmax(width1, width2)

	imgMatrix := NewRGBAMatrix(height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			imgMatrix[i][j][0] = 0
			imgMatrix[i][j][1] = 0
			imgMatrix[i][j][2] = 0
			imgMatrix[i][j][3] = 255
		}
	}
	return imgMatrix
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Bresenham's algorithm, http://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// TODO: handle int overflow etc.
func Drawline(x0, y0, x1, y1 int, brush Putpixel) {
	dx := Abs(x1 - x0)
	dy := Abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		brush(x0, y0)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func ImgZoom(src image.Image, multiple float64) image.Image {
	bounds := src.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	ret := image.NewRGBA(image.Rect(0, 0, int(float64(width)*multiple), int(float64(height)*multiple)))
	if err := graphics.Scale(ret, src); err != nil {
		panic(err)
	}
	return ret
}
