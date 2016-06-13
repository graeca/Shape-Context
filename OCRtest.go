package main

import (
	"code.google.com/p/graphics-go/graphics"
	//"code.google.com/p/graphics-go/graphics/edge"
	mPkg "./mPkg"
	"github.com/Comdex/imgo"
	//"errors"
	//"fmt"
	//"image"
	//"image/draw"
	//"math"
	//"os"
	//"errors"
	"fmt"
	"image"
	"image/color"
	//"image/png"
	"log"
	//"os"
	"runtime"
	"time"
)

type TBackData struct {
	Value             string
	Name              string
	Cmp               string
	N                 float64
	CostAvg           float64
	EuclideanDistance float64
	PPlen             int
}

var (
	Bl1     float64 = 0
	Bl2     float64 = 0
	srcStep int     = 0
	bei     float64 = 0.2
	outPng  bool    = false
	ccc     int     = 0
)

func test4(dir, namee2 string, Tmd mPkg.TmodData, BackData chan TBackData) {
	// yuansrc1, decodeErr1 := imgo.DecodeImage("img/En/" + namee1 + ".png")
	// if decodeErr1 != nil {
	// 	panic(decodeErr1)
	// }

	//startTime := time.Now() //>>>>>>>>>>>>>>>>>>
	yuansrc2, decodeErr2 := imgo.DecodeImage(dir + namee2 + ".png")
	if decodeErr2 != nil {
		panic(decodeErr2)
	}
	// ysW1 := float64(yuansrc1.Bounds().Dx())
	// ysH1 := float64(yuansrc1.Bounds().Dy())
	ysW2 := float64(yuansrc2.Bounds().Dx())
	ysH2 := float64(yuansrc2.Bounds().Dy())

	// ysW1 = ysW1 * Bl1 / ysH1
	// ysH1 = Bl1
	if Bl2 != 0 {
		ysW2 = ysW2 * Bl2 / ysH2
		ysH2 = Bl2
	}

	//src1 := image.NewRGBA(image.Rect(0, 0, int(ysW1), int(ysH1)))
	src2 := image.NewRGBA(image.Rect(0, 0, int(ysW2), int(ysH2)))
	//graphics.Scale(src1, yuansrc1)
	graphics.Scale(src2, yuansrc2)

	//dst1 := mPkg.Edge_Canny(src1)

	dst2 := mPkg.Edge_Canny(src2)
	//dst2 := mPkg.Edge_LaplacianOfGaussian(src2)

	//piximg1 := mPkg.Img2Matrix(dst1)

	//piximg2 := mPkg.Img2Matrix(dst2)
	//newImg1 := mPkg.Edge_Canny(srcImg)
	piximg := mPkg.Img2Matrix2(dst2)
	//fmt.Print("cccc")
	piximg2 := mPkg.FindEdges(piximg)

	TT1 := Tmd.TT
	//TT1 := mPkg.GetPoints(piximg1)
	TT2 := mPkg.GetPoints(piximg2, srcStep)
	//TT2 := mPkg.PurificationMod(TT2tmp, bei)

	//mPkg.SaveImgAsPNG(mPkg.Matrix2Img(piximg1), "imgout/En_"+namee1+".png")
	//mPkg.SaveImgAsPNG(mPkg.Matrix2Img(piximg2), "imgout/LD_"+namee2+".png")

	TT1, TT2 = mPkg.ResetCenterMass(TT1, TT2)

	pp, pCostAvg := mPkg.CompareTPCC(TT1, TT2)

	//CC, _ := mPkg.CosineSum2(TT1, TT2, pp)
	CC := mPkg.CosineSum(TT1, TT2, pp)

	//TT1, TT3 = mPkg.ResetCenterMass(TT1, TT3)

	if outPng {
		ourBei := 3
		outImg := mPkg.ImgPlusImg(piximg2, piximg2)
		out := mPkg.Matrix2Img(outImg)

		wh1 := out.Bounds().Dx()
		if wh1 < out.Bounds().Dy() {
			wh1 = out.Bounds().Dy()
		}

		dst := image.NewRGBA(image.Rect(0, 0, wh1*ourBei, wh1*ourBei)) //图片放大
		graphics.Scale(dst, out)
		//fmt.Println(namee1+":"+namee2+">>>>>>", len(pp))

		for i := 0; i < len(TT1); i++ {
			//fmt.Println(TT1[i])
			if (TT1[i].P.Y*ourBei < dst.Bounds().Dy()) && (TT1[i].P.X*ourBei < dst.Bounds().Dx()) {
				_, G, _, A := dst.At(TT1[i].P.X*ourBei, TT1[i].P.Y*ourBei).RGBA()
				dst.Set(TT1[i].P.X*ourBei, TT1[i].P.Y*ourBei, color.RGBA{255, uint8(G), 100, uint8(A)}) //[0] = 250
				//outImg[TT1[i].P.Y][TT1[i].P.X][2] = 200
			}
		}
		//fmt.Println(len(TT2))
		for i := 0; i < len(TT2); i++ {
			//fmt.Println(TT2[i])
			if (TT2[i].P.Y*ourBei < dst.Bounds().Dy()) && (TT2[i].P.X*ourBei < dst.Bounds().Dx()) {
				R, _, _, A := dst.At(TT2[i].P.X*ourBei, TT2[i].P.Y*ourBei).RGBA()
				dst.Set(TT2[i].P.X*ourBei, TT2[i].P.Y*ourBei, color.RGBA{uint8(R), 250, 100, uint8(A)}) //[1] = 250
				//outImg[TT2[i].P.Y][TT2[i].P.X][2] = 200
			}
		}

		wh := dst.Bounds().Dx()
		if wh < dst.Bounds().Dy() {
			wh = dst.Bounds().Dy()
		}

		dstO := image.NewRGBA(image.Rect(0, 0, wh*ourBei, wh*ourBei)) //图片放大
		graphics.Scale(dstO, dst)

		for k, v := range pp {
			if v != 0 {
				x1 := TT1[k].P.X * ourBei * ourBei
				y1 := TT1[k].P.Y * ourBei * ourBei
				x2 := TT2[v].P.X * ourBei * ourBei
				y2 := TT2[v].P.Y * ourBei * ourBei
				//if (x1 < wh*ourBei) && (y1 < wh*ourBei) && (x2 < wh*ourBei) && (y2 < wh*ourBei) && (x1 >= 0) && (y1 >= 0) && (x2 >= 0) && (y2 >= 0) {
				mPkg.Drawline(x1, y1, x2, y2, func(x, y int) {
					dstO.Set(x, y, color.RGBA{100, 100, 100, 255})
					//dstO.Set(x, y, color.White)
				})
				//}
			}
		}

		mPkg.SaveImgAsPNG(dstO, "imgout/"+Tmd.V+"_"+namee2+"-"+Tmd.Name+fmt.Sprintf("_%d", ccc)+".png")
		ccc++
	}
	//fmt.Print("----", namee1)
	//return mPkg.CosineSum(TT1, TT2, pp)
	//fmt.Println("["+namee2+"]-["+Tmd.V+"]", mPkg.CosineSum(TT1, TT2, pp), pCostAvg)
	//panic("停止测试")
	BackData <- TBackData{Tmd.V, Tmd.Name, namee2, CC, pCostAvg, mPkg.EuclideanDistanceSum(TT1, TT2, pp), len(pp)}
}

func test5(src image.Image, key string, Tmd mPkg.TmodData, BackData chan TBackData) {
	ysW2 := float64(src.Bounds().Dx())
	ysH2 := float64(src.Bounds().Dy())

	if Bl2 != 0 {
		ysW2 = ysW2 * Bl2 / ysH2
		ysH2 = Bl2
	}

	src2 := image.NewRGBA(image.Rect(0, 0, int(ysW2), int(ysH2)))
	graphics.Scale(src2, src)

	dst2 := mPkg.Edge_Canny(src2)
	piximg := mPkg.Img2Matrix2(dst2)
	piximg2 := mPkg.FindEdges(piximg)

	TT1 := Tmd.TT
	TT2 := mPkg.GetPoints(piximg2, srcStep)
	TT1, TT2 = mPkg.ResetCenterMass(TT1, TT2)
	pp, pCostAvg := mPkg.CompareTPCC(TT1, TT2)
	CC := mPkg.CosineSum(TT1, TT2, pp)
	if outPng {
		ourBei := 3
		outImg := mPkg.ImgPlusImg(piximg2, piximg2)
		out := mPkg.Matrix2Img(outImg)

		wh1 := out.Bounds().Dx()
		if wh1 < out.Bounds().Dy() {
			wh1 = out.Bounds().Dy()
		}

		dst := image.NewRGBA(image.Rect(0, 0, wh1*ourBei, wh1*ourBei)) //图片放大
		graphics.Scale(dst, out)
		//fmt.Println(namee1+":"+namee2+">>>>>>", len(pp))

		for i := 0; i < len(TT1); i++ {
			//fmt.Println(TT1[i])
			if (TT1[i].P.Y*ourBei < dst.Bounds().Dy()) && (TT1[i].P.X*ourBei < dst.Bounds().Dx()) {
				_, G, _, A := dst.At(TT1[i].P.X*ourBei, TT1[i].P.Y*ourBei).RGBA()
				dst.Set(TT1[i].P.X*ourBei, TT1[i].P.Y*ourBei, color.RGBA{255, uint8(G), 100, uint8(A)}) //[0] = 250
				//outImg[TT1[i].P.Y][TT1[i].P.X][2] = 200
			}
		}
		//fmt.Println(len(TT2))
		for i := 0; i < len(TT2); i++ {
			//fmt.Println(TT2[i])
			if (TT2[i].P.Y*ourBei < dst.Bounds().Dy()) && (TT2[i].P.X*ourBei < dst.Bounds().Dx()) {
				R, _, _, A := dst.At(TT2[i].P.X*ourBei, TT2[i].P.Y*ourBei).RGBA()
				dst.Set(TT2[i].P.X*ourBei, TT2[i].P.Y*ourBei, color.RGBA{uint8(R), 250, 100, uint8(A)}) //[1] = 250
				//outImg[TT2[i].P.Y][TT2[i].P.X][2] = 200
			}
		}

		wh := dst.Bounds().Dx()
		if wh < dst.Bounds().Dy() {
			wh = dst.Bounds().Dy()
		}

		dstO := image.NewRGBA(image.Rect(0, 0, wh*ourBei, wh*ourBei)) //图片放大
		graphics.Scale(dstO, dst)

		for k, v := range pp {
			if v != 0 {
				x1 := TT1[k].P.X * ourBei * ourBei
				y1 := TT1[k].P.Y * ourBei * ourBei
				x2 := TT2[v].P.X * ourBei * ourBei
				y2 := TT2[v].P.Y * ourBei * ourBei
				//if (x1 < wh*ourBei) && (y1 < wh*ourBei) && (x2 < wh*ourBei) && (y2 < wh*ourBei) && (x1 >= 0) && (y1 >= 0) && (x2 >= 0) && (y2 >= 0) {
				mPkg.Drawline(x1, y1, x2, y2, func(x, y int) {
					dstO.Set(x, y, color.RGBA{100, 100, 100, 255})
					//dstO.Set(x, y, color.White)
				})
				//}
			}
		}

		mPkg.SaveImgAsPNG(dstO, "imgout/"+Tmd.V+"_"+key+"-"+Tmd.Name+fmt.Sprintf("_%d", ccc)+".png")
		ccc++
	}
	BackData <- TBackData{Tmd.V, Tmd.Name, key, CC, pCostAvg, mPkg.EuclideanDistanceSum(TT1, TT2, pp), len(pp)}
}

func run5(name string) {
	//name := "202"
	yuansrc1, decodeErr1 := imgo.DecodeImage("img/" + name + ".png")
	if decodeErr1 != nil {
		panic(decodeErr1)
	}
	mPkg.Captcha_split(yuansrc1, name)
}

func run(dir, cmp string) bool {
	startTime := time.Now()
	//mods := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "#a", "#b", "#c", "#d", "#e", "#f", "#g", "#h", "#i", "#j", "#k", "#l", "#m", "#n", "#o", "#p", "#q", "#r", "#s", "#t", "#u", "#v", "#w", "#x", "#y", "#z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	max := float64(0)
	EuclideanDistance := float64(0)
	//minCostAvg := float64(999999)
	maxStr := ""
	maxName := ""
	BackData := make(chan TBackData)
	maxpplen := 0
	//acg := float64(0)

	fmt.Print(cmp)

	treadCount := 20
	treadCount2 := treadCount
	//go func() {
	for i := 0; i < len(mPkg.ModTTs); i++ {
		if treadCount > 0 {
			go test4(dir, cmp, mPkg.ModTTs[i], BackData)
			treadCount--
		} else {
			data := <-BackData
			//fmt.Println(">"+data.Value, data.N, data.CostAvg)
			if data.PPlen > 5 {
				td := data.N //* 10 / data.CostAvg
				//CostAvg := data.CostAvg
				//			if data.EuclideanDistance > 120 {
				//				fmt.Println(data)
				//			}
				//if data.PPlen > 30 {
				if max < td {
					max = td
					maxStr = data.Value
					maxName = data.Name
					//acg = data.CostAvg
					EuclideanDistance = data.EuclideanDistance
					maxpplen = data.PPlen
				}
			}
			//}
			if i%2 == 0 {
				fmt.Print(".")
			}
			go test4(dir, cmp, mPkg.ModTTs[i], BackData)
		}
	}
	//}()

	for i := 0; i < treadCount2; i++ {
		data := <-BackData
		if data.PPlen > 5 {
			//fmt.Println(">", data.N, data.CostAvg)
			td := data.N //* 10 / data.CostAvg
			//CostAvg := data.CostAvg
			//		if data.EuclideanDistance > 120 {
			//			fmt.Println(data)
			//		}
			//if data.PPlen > 30 {
			if max < td {
				max = td
				maxStr = data.Value
				maxName = data.Name
				//acg = data.CostAvg
				EuclideanDistance = data.EuclideanDistance
				maxpplen = data.PPlen
			}
		}
		//}
		if i%2 == 0 {
			fmt.Print(".")
		}
	}

	if EuclideanDistance > 120 {
		fmt.Println("ok ["+cmp+"]-["+maxStr+"]", maxName, max, EuclideanDistance, maxpplen, "time:", (time.Now().Sub(startTime)))
		log.Println("ok ["+cmp+"]-["+maxStr+"]", maxName, max, EuclideanDistance, maxpplen, "time:", (time.Now().Sub(startTime)))
	} else {
		fmt.Println("bad ["+cmp+"]-["+maxStr+"]", maxName, max, EuclideanDistance, maxpplen, "time:", (time.Now().Sub(startTime)))
		log.Println("bad ["+cmp+"]-["+maxStr+"]", maxName, max, EuclideanDistance, maxpplen, "time:", (time.Now().Sub(startTime)))
	}
	return true
}

func MakeMods() {
	fmt.Print("Make Mods ...")
	mPkg.MakeMods("ModImgs_1", "ModImgs_1.data")
	fmt.Println("ok")
}

func LoadMods() {
	fmt.Print("Load Mods ...")
	mPkg.LoadMods("ModImgs_1.data")
	fmt.Println("ok")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//	run5("101")
	//	run5("102")
	//	run5("103")
	//	run5("104")
	//	run5("105")
	//	run5("106")
	//	run5("107")
	//	run5("108")
	//	run5("109")
	//	run5("110")
	//	run5("110_0")
	//	run5("111")
	//	run5("112")
	//	run5("113")
	//	run5("114")
	//	run5("115")
	//	run5("116")
	//	run5("116-2")
	//	run5("117")
	//	run5("118")
	//	run5("119")
	//	run5("120")
	//	run5("121")
	//	run5("122")
	//	run5("201")
	//	run5("202")
	//	run5("203")
	//	run5("204")
	//	run5("205")
	//	run5("206")
	//	run5("207")
	//	run5("pic2")
	//	run5("at-col")
	//	return
	LoadMods()

	//fmt.Println(mPkg.ModTTs)
	//Bl1 = float64(60)
	Bl2 = 60
	srcStep = 3
	bei = 1
	outPng = false

	tests := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "#a", "#b", "#c", "#d", "#e", "#f", "#g", "#h", "#i", "#j", "#k", "#l", "#m", "#n", "#o", "#p", "#q", "#r", "#s", "#t", "#u", "#v", "#w", "#x", "#y", "#z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(tests); i++ {
		run("img/En/", tests[i])
	}
	return
	//  run("img/test/", "5-test")
	//	run("img/test/", "6-test")
	//	run("img/test/", "8-test")
	//	run("img/test/", "E-test")
	//	run("img/test/", "Q-test")
	//	run("img/test/", "U-test")
	//	run("img/test/", "5")
	//	run("img/test/", "7")
	//	run("img/test/", "A")
	//	run("img/test/", "a2")
	//	run("img/test/", "B")
	//	run("img/test/", "k2")
	//	run("img/test/", "k3")
	//	run("img/test/", "m")
	//	run("img/test/", "p")
	//	run("img/test/", "s")
	//	run("img/test/", "t")
	//	run("img/test/", "y")

	// run("img/", "101")
	// run("img/", "102")
	// run("img/", "103")
	// run("img/", "104")
	// run("img/", "110_0")
	//run("", "baseball-template")

	//run("img/captcha/", "A")
	//run("img/captcha/", "d")
	//run("img/captcha/", "e")
	//run("img/captcha/", "p2")
	// run("img/captcha/", "tz")

	//	run("img/", "101")
	run("img/", "102")
	//	run("img/", "103")
	run("img/", "104")
	//	run("img/", "105")
	run("img/", "106")
	//	run("img/", "107")
	//	run("img/", "108")
	//	run("img/", "109")
	run("img/", "110_0")
	run("img/", "111")
	//	run("img/", "112")
	//	run("img/", "113")
	//	run("img/", "114")
	//	run("img/", "115")
	//	run("img/", "116")
	//	run("img/", "116-2")
	//	run("img/", "117")
	//	run("img/", "118")
	//	run("img/", "119")
	run("img/", "120")
	//	run("img/", "121")
	//	run("img/", "122")
	//	run("img/", "201")
	//	run("img/", "202")
	//	run("img/", "203")
	//	run("img/", "204")
	//	run("img/", "205")
	//	run("img/", "206")
	//	run("img/", "207")
	//	run("img/", "cp")
	//	run("img/", "nb1")
	//	run("img/", "nb2")
	//	run("img/", "nb3")
	//	run("img/", "nb4")
	//	run("img/", "nb6")
	//	run("img/", "nb7")
	//	run("img/", "nb8")
	//	run("img/", "nb9")
	//	run("img/", "nb10")
	//	run("img/", "nb11")
	//	run("img/", "nb12")
	//	run("img/", "nb13")
	//	run("img/", "nb15")
	//	run("img/", "nb16")
	return

	fmt.Print("cccc")
	srcImg, decodeErr := imgo.DecodeImage("ModImgs_1/0_comic_52.png")
	if decodeErr != nil {
		panic(decodeErr)
	}
	//	mPkg.SaveImgAsPNG(newImg1, "imgoutTest1.png")
	//	//newImg2 := mPkg.Edge_Canny(newImg1)
	//newImg2 := mPkg.Edge_LaplacianOfGaussian(newImg1)
	//srcImg = mPkg.Edge_GaussianSmooth(srcImg, 6, 4)
	//newImg1 := mPkg.Edge_Canny2(srcImg)
	//newImg1 := mPkg.Edge_Prewitt2(srcImg)
	//newImg1 := mPkg.Edge_Scharr2(srcImg) //ok
	//newImg1 := mPkg.Edge_Sobel2(srcImg)
	srcImg = mPkg.ImgZoom(srcImg, 4)
	newImg1 := mPkg.Edge_Canny(srcImg)
	piximg := mPkg.Img2Matrix(newImg1)
	fmt.Print("cccc")
	piximg2 := mPkg.FindEdges(piximg)
	//piximg2 = mPkg.FindEdges(piximg2)
	outImg := mPkg.Matrix2Img(piximg2)
	fmt.Print("cccc")
	mPkg.SaveImgAsPNG(outImg, "imgoutTest1.png")
}
