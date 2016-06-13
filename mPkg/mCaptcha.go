package mPkg

import (
	"fmt"
	"image"
	"image/color"
)

//	"code.google.com/p/graphics-go/graphics/edge"
//	"github.com/coraldane/resize"

type TPointVectorProperties struct { //八方向量, 颜色向量
	C_toR int16
	C_toG int16
	C_toB int16
	V_上   int16
	V_下   int16
	V_左   int16
	V_右   int16
	V_左上  int16
	V_右上  int16
	V_左下  int16
	V_右下  int16
	Index uint16
}

//func Captcha_FindEdges(src image.Image) image.Image {
//	for i := 0; i < src.Bounds().Dx(); i++ {
//		for j := 0; j < src.Bounds().Dy(); j++ {
//		}
//	}
//}

func Captcha_split(src image.Image, name string) {
	W := src.Bounds().Dx()
	H := src.Bounds().Dy()
	pvpR := make([][]TPointVectorProperties, W, W)
	for i := 0; i < W; i++ {
		pvpR[i] = make([]TPointVectorProperties, H, H)
	}

	pvpG := make([][]TPointVectorProperties, W, W)
	for i := 0; i < W; i++ {
		pvpG[i] = make([]TPointVectorProperties, H, H)
	}

	pvpB := make([][]TPointVectorProperties, W, W)
	for i := 0; i < W; i++ {
		pvpB[i] = make([]TPointVectorProperties, H, H)
	}

	src2 := src //Edge_LaplacianOfGaussian(src)
	index := uint16(0)
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, G, B, _ := src2.At(i, j).RGBA()
			if (i == 0) || (j == 0) || (i == W-1) || (j == H-1) {
			} else {
				pvpR[i][j].C_toB = int16(uint8(R)) - int16(uint8(B))
				pvpG[i][j].C_toB = int16(uint8(G)) - int16(uint8(B))
				pvpB[i][j].C_toB = 0

				pvpR[i][j].C_toG = int16(uint8(R)) - int16(uint8(G))
				pvpG[i][j].C_toG = 0
				pvpB[i][j].C_toG = int16(uint8(B)) - int16(uint8(G))

				pvpR[i][j].C_toR = 0
				pvpG[i][j].C_toR = int16(uint8(G)) - int16(uint8(R))
				pvpB[i][j].C_toR = int16(uint8(B)) - int16(uint8(R))

				Rt, Gt, Bt, _ := src2.At(i, j-1).RGBA()
				pvpR[i][j].V_上 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_上 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_上 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i, j+1).RGBA()
				pvpR[i][j].V_下 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_下 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_下 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i+1, j).RGBA()
				pvpR[i][j].V_右 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_右 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_右 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i+1, j-1).RGBA()
				pvpR[i][j].V_右上 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_右上 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_右上 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i+1, j+1).RGBA()
				pvpR[i][j].V_右下 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_右下 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_右下 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i-1, j).RGBA()
				pvpR[i][j].V_左 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_左 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_左 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i-1, j-1).RGBA()
				pvpR[i][j].V_左上 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_左上 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_左上 = int16(uint8(B)) - int16(uint8(Bt))

				Rt, Gt, Bt, _ = src2.At(i-1, j+1).RGBA()
				pvpR[i][j].V_左下 = int16(uint8(R)) - int16(uint8(Rt))
				pvpG[i][j].V_左下 = int16(uint8(G)) - int16(uint8(Gt))
				pvpB[i][j].V_左下 = int16(uint8(B)) - int16(uint8(Bt))
				pvpR[i][j].Index = index
			}
			index++
		}
	}
	//fmt.Println(pvpR)
	dst := image.NewRGBA(src.Bounds())
	thisAbs := func(a int16) int16 {
		if a < 0 {
			return 0 - a
		} else {
			return a
		}
	}

	limit := int16(50)

	yesLimit := func(pvp [][]TPointVectorProperties, i, j int) bool {
		return thisAbs(pvp[i][j].V_上) > limit || (thisAbs(pvp[i][j].V_下) > limit) || (thisAbs(pvp[i][j].V_右) > limit) || (thisAbs(pvp[i][j].V_右上) > limit) || (thisAbs(pvp[i][j].V_右下) > limit) || (thisAbs(pvp[i][j].V_左) > limit) || (thisAbs(pvp[i][j].V_左上) > limit) || (thisAbs(pvp[i][j].V_左下) > limit)
	}

	for i := 0; i < W; i++ {
		for j := 1; j < H; j++ {
			//R, _, _, A := src2.At(i, j).RGBA()
			//if thisAbs(pvpR[i][j].V_上) < limit || (thisAbs(pvpR[i][j].V_下) < limit) || (thisAbs(pvpR[i][j].V_右) < limit) || (thisAbs(pvpR[i][j].V_右上) < limit) || (thisAbs(pvpR[i][j].V_右下) < limit) || (thisAbs(pvpR[i][j].V_左) < limit) || (thisAbs(pvpR[i][j].V_左上) < limit) || (thisAbs(pvpR[i][j].V_左下) < limit) {
			yesR := yesLimit(pvpR, i, j)
			yesG := yesLimit(pvpG, i, j)
			yesB := yesLimit(pvpB, i, j)
			if (yesR && yesG) || (yesR && yesB) || (yesB && yesG) || (yesR && yesB && yesG) {
				dst.Set(i, j, color.White)
			} else {
				//if (uint8(R) > 100) && (uint8(R) < 200) {
				//} else {
				dst.Set(i, j, color.Black)
				//	dst.Set(i, j, color.White)
				//}
			}
		}
	}
	fmt.Println(name)

	xTJ := make([]int, W, W)

	sum := 0

	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, _, _, _ := dst.At(i, j).RGBA()
			sum += int(uint8(R))
		}
	}
	avg := sum / (W * H)
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, _, _, _ := dst.At(i, j).RGBA()
			if int(uint8(R)) > avg {
				xTJ[i]++
			}
		}
	}

	srcZFT := image.NewRGBA(src.Bounds())
	for i := 0; i < W; i++ {
		for j := H - 1; j >= 0; j-- {
			if xTJ[i] > (H - j) {
				srcZFT.Set(i, j, color.Black)
			} else {
				srcZFT.Set(i, j, color.White)
			}
		}
	}

	SaveImgAsPNG(dst, "CaptchaOut/xxl_"+name+".png")
	SaveImgAsPNG(srcZFT, "CaptchaOut/srcZFT_"+name+".png")
}

/*
func Captcha_split(src image.Image, name string) {
	//src2 := Edge_GaussianSmooth(src, 10, 3)
	W := src.Bounds().Dx()
	H := src.Bounds().Dy()
	src3 := image.NewRGBA(src.Bounds())
	thisAbs := func(a int) int {
		if a < 0 {
			return 0 - a
		} else {
			return a
		}
	}
	sumR := 0
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, G, B, _ := src.At(i, j).RGBA()
			sumR += thisAbs(int(uint8(R))-int(uint8(G))) + thisAbs(int(uint8(R))-int(uint8(B))) + thisAbs(int(uint8(G))-int(uint8(B)))
		}
	}

	avg := sumR / (W * H)
	fmt.Println(sumR, sumR/(W*H))

	xTJ := make([]int, W, W)
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, G, B, _ := src.At(i, j).RGBA()
			xx := thisAbs(int(uint8(R))-int(uint8(G))) + thisAbs(int(uint8(R))-int(uint8(B))) + thisAbs(int(uint8(G))-int(uint8(B)))
			if avg < xx {
				src3.Set(i, j, color.White)
				xTJ[i]++
			} else {
				src3.Set(i, j, color.Black)
			}
		}
	}
	srcZFT := image.NewRGBA(src.Bounds())
	for i := 0; i < W; i++ {
		for j := H - 1; j >= 0; j-- {
			if xTJ[i] > (H - j) {
				srcZFT.Set(i, j, color.Black)
			} else {
				srcZFT.Set(i, j, color.White)
			}
		}
	}
	SaveImgAsPNG(src3, "CaptchaOut/dst0_"+name+".png")
	SaveImgAsPNG(srcZFT, "CaptchaOut/srcZFT_"+name+".png")
}
*/
/**
func Captcha_split(src image.Image, name string) {
	W := src.Bounds().Dx()
	H := src.Bounds().Dy()
	src3 := Edge_Prewitt2(src)

	fmt.Println(name)

	xTJ := make([]int, W, W)

	sum := 0

	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, _, _, _ := src.At(i, j).RGBA()
			sum += int(uint8(R))
		}
	}
	avg := sum / (W * H)
	for i := 0; i < W; i++ {
		for j := 0; j < H; j++ {
			R, _, _, _ := src.At(i, j).RGBA()
			if int(uint8(R)) > avg {
				xTJ[i]++
			}
		}
	}

	srcZFT := image.NewRGBA(src.Bounds())
	for i := 0; i < W; i++ {
		for j := H - 1; j >= 0; j-- {
			if xTJ[i] > (H - j) {
				srcZFT.Set(i, j, color.Black)
			} else {
				srcZFT.Set(i, j, color.White)
			}
		}
	}
	SaveImgAsPNG(src3, "CaptchaOut/dst0_"+name+".png")
	SaveImgAsPNG(srcZFT, "CaptchaOut/srcZFT_"+name+".png")
}**/
