package mPkg

import (
	//"fmt"
	"image"
	"image/color"

	"code.google.com/p/graphics-go/graphics/edge"
	"github.com/coraldane/resize"
)

func Edge_GaussianSmooth(srcImg image.Image, sigma float64, radius int) image.Image { //高斯模糊
	return resize.GaussianSmooth(srcImg, sigma, radius)
}

func Edge_LaplacianOfGaussian(src image.Image) (dst *image.Gray) { //log算法
	dst = image.NewGray(src.Bounds())
	edge.LaplacianOfGaussian(dst, src)
	return
}

func Edge_DifferenceOfGaussians(src image.Image, sd0, sd1 float64) (dst *image.Gray) { //Dog算法
	dst = image.NewGray(src.Bounds())
	edge.DifferenceOfGaussians(dst, src, sd0, sd1)
	return
}

func Edge_Canny(src image.Image) (dst *image.Gray) { //Canny算法
	dst = image.NewGray(src.Bounds())
	if err := edge.Canny(dst, src); err != nil {
		panic(err)
	}
	return
}

func Edge_Sobel(src image.Image) (dir *image.Gray) { //Sobel(索贝尔)算法,白底黑字
	mag := image.NewGray(src.Bounds())
	dir = image.NewGray(src.Bounds())
	if err := edge.Sobel(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 0) && (uint8(G) != 0) && (uint8(B) != 0) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Scharr(src image.Image) (dir *image.Gray) { //类似于Sobel(索贝尔)算法,角误差更少,白底黑字
	mag := image.NewGray(src.Bounds())
	dir = image.NewGray(src.Bounds())
	if err := edge.Scharr(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 0) && (uint8(G) != 0) && (uint8(B) != 0) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Prewitt(src image.Image) (mag *image.Gray) { //Prewitt(普瑞维特)算法,白底黑字
	mag = image.NewGray(src.Bounds())
	dir := image.NewGray(src.Bounds())
	if err := edge.Prewitt(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 0) && (uint8(G) != 0) && (uint8(B) != 0) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Canny2(src image.Image) (dst *image.Gray) { //Canny算法,高亮区
	dst = image.NewGray(src.Bounds())
	if err := edge.Canny(dst, src); err != nil {
		panic(err)
	}
	bounds := dst.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dst.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 255) && (uint8(G) != 255) && (uint8(B) != 255) {
				dst.Set(i, j, color.Black)
			} else {
				dst.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Sobel2(src image.Image) (dir *image.Gray) { //Sobel(索贝尔)算法,黑底白字
	mag := image.NewGray(src.Bounds())
	dir = image.NewGray(src.Bounds())
	if err := edge.Sobel(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 45) && (uint8(G) != 45) && (uint8(B) != 45) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Scharr2(src image.Image) (dir *image.Gray) { //类似于Sobel(索贝尔)算法,角误差更少,黑底白字
	mag := image.NewGray(src.Bounds())
	dir = image.NewGray(src.Bounds())
	if err := edge.Scharr(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 45) && (uint8(G) != 45) && (uint8(B) != 45) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}

func Edge_Prewitt2(src image.Image) (dir *image.Gray) { //Prewitt(普瑞维特)算法,黑底白字
	mag := image.NewGray(src.Bounds())
	dir = image.NewGray(src.Bounds())
	if err := edge.Prewitt(mag, dir, src); err != nil {
		panic(err)
	}
	bounds := dir.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			R, G, B, _ := dir.At(i, j).RGBA()
			//fmt.Println(uint8(R), uint8(G), uint8(B), uint8(A))
			if (uint8(R) != 45) && (uint8(G) != 45) && (uint8(B) != 45) {
				dir.Set(i, j, color.Black)
			} else {
				dir.Set(i, j, color.White)
			}
		}
	}
	return
}
