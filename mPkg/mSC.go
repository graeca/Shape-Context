package mPkg

import (
	//"github.com/Comdex/imgo"
	//"image"
	"fmt"
	//"image"
	"math"
	//"strconv"
)

type TPoint struct {
	X int
	Y int
}

type TPointCC struct { //TPointCharacteristicCollection
	P     TPoint
	CList [][]int
}

type TPAD struct {
	P   TPoint
	Pad uint64
}

var (
	splitR int = 3 //5
	splitS int = 8 //12 必须大于4, 尽可能4的倍数
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//func FindEdges(src [][][]uint8) [][][]uint8 { //查找边缘
//	height := len(src)
//	width := len(src[0])
//	for i := 0; i < height; i++ { //二值化
//		for j := 0; j < width; j++ {
//			if (src[i][j][0] + src[i][j][1] + src[i][j][2]) > 0 {
//				src[i][j][0] = 255 //二值化数据
//				if (i == 0) || (i == height-1) || (j == 0) || (j == width-1) {
//					src[i][j][1] = 255 //边缘标记
//				} else {
//					src[i][j][1] = 0 //边缘标记
//				}
//				src[i][j][2] = 0 //平滑后边缘数据
//				src[i][j][3] = 255
//			} else {
//				src[i][j][0] = 0
//				src[i][j][1] = 0
//				src[i][j][2] = 0
//				src[i][j][3] = 255
//			}
//		}
//	}

//	for n := 0; n < 3; n++ { //3次补值
//		for i := 0; i < height-2; i++ { //补值
//			for j := 0; j < width-2; j++ {
//				if src[i][j][0] > 0 && src[i+1][j][0] == 0 && src[i+2][j][0] > 0 { //i方向补值
//					src[i+1][j][0] = 255
//				}
//				if src[i][j][0] > 0 && src[i][j+1][0] == 0 && src[i][j+2][0] > 0 { //j方向补值
//					src[i][j+1][0] = 255
//				}

//				if src[i][j][0] > 0 && src[i+1][j+1][0] == 0 && src[i+2][j+2][0] > 0 { //ij方向补值
//					src[i+1][j+1][0] = 255
//				}
//			}
//		}
//	}

//	Check8Direction := func(i, j int) bool { //八个方向检测, 非全包围点返回 true
//		if src[i][j][0] == 0 { //无值点跳过
//			return false
//		}
//		a := 0 //有值计数
//		if src[i+1][j][0] != 0 {
//			a++
//		}
//		if src[i][j+1][0] != 0 {
//			a++
//		}
//		if src[i+1][j+1][0] != 0 {
//			a++
//		}
//		if src[i-1][j][0] != 0 {
//			a++
//		}
//		if src[i][j-1][0] != 0 {
//			a++
//		}
//		if src[i-1][j-1][0] != 0 {
//			a++
//		}
//		if src[i+1][j+1][0] != 0 {
//			a++
//		}
//		if src[i-1][j+1][0] != 0 {
//			a++
//		}
//		if (a != 0) && (a != 8) {
//			return true
//		} else {
//			return false
//		}
//	}
//	for i := 1; i < height-1; i++ { //查找边缘
//		for j := 1; j < width-1; j++ {
//			if Check8Direction(i, j) {
//				src[i][j][1] = 255
//			}
//		}
//	}

//	CheckDepression := func(i, j int) bool { //凹陷点检测
//		if src[i][j][1] != 0 { //有值点跳过
//			return false
//		}
//		a := 0                   //下
//		b := 0                   //右
//		c := 0                   //右下
//		d := 0                   //上
//		e := 0                   //左
//		f := 0                   //左上
//		g := 0                   //右上
//		h := 0                   //左下
//		if src[i+1][j][0] != 0 { //下
//			a = 1
//		}
//		if src[i][j+1][0] != 0 { //右
//			b = 0
//		}
//		if src[i+1][j+1][0] != 0 { //右下
//			c = 1
//		}
//		if src[i-1][j][0] != 0 { //上
//			d = 1
//		}
//		if src[i][j-1][0] != 0 { //左
//			e = 1
//		}
//		if src[i-1][j-1][0] != 0 { //左上
//			f = 1
//		}

//		if src[i+1][j+1][0] != 0 { //右上
//			g = 1
//		}
//		if src[i-1][j+1][0] != 0 { //左下
//			h = 1
//		}
//		if ((f+e+h == 0) && (d+g+b+c+a == 5)) || ((f+d+g == 0) && (b+c+a+h+e == 5)) || ((g+b+c == 0) && (a+h+e+f+d == 5)) || ((c+a+h == 0) && (e+f+d+g+b == 5)) {
//			return true
//		} else {
//			return false
//		}
//	}
//	for i := 1; i < height-1; i++ { //平滑凹陷
//		for j := 1; j < width-1; j++ {
//			if CheckDepression(i, j) {
//				src[i][j][1] = 255
//			}
//		}
//	}

//	CheckExtrude := func(i, j int) { //处理突出点
//		if src[i][j][1] == 0 { //无值点跳过
//			return
//		}
//		a := 0                   //下
//		b := 0                   //右
//		c := 0                   //右下
//		d := 0                   //上
//		e := 0                   //左
//		f := 0                   //左上
//		g := 0                   //右上
//		h := 0                   //左下
//		if src[i+1][j][0] != 0 { //下
//			a = 1
//		}
//		if src[i][j+1][0] != 0 { //右
//			b = 0
//		}
//		if src[i+1][j+1][0] != 0 { //右下
//			c = 1
//		}
//		if src[i-1][j][0] != 0 { //上
//			d = 1
//		}
//		if src[i][j-1][0] != 0 { //左
//			e = 1
//		}
//		if src[i-1][j-1][0] != 0 { //左上
//			f = 1
//		}

//		if src[i+1][j+1][0] != 0 { //右上
//			g = 1
//		}
//		if src[i-1][j+1][0] != 0 { //左下
//			h = 1
//		}

//		//a 下
//		//b 右
//		//c 右下
//		//d 上
//		//e 左
//		//f 左上

//		//		if a+b+c+d+e+f == 2 { //线状突出, 填充缺失后删除

//		//		}
//		if ((f+e+h == 3) && (d+g+b+c+a == 0)) || ((f+d+g == 3) && (b+c+a+h+e == 0)) || ((g+b+c == 3) && (a+h+e+f+d == 0)) || ((c+a+h == 3) && (e+f+d+g+b == 0)) { //点状突出, 直接删除
//			src[i][j][1] = 0
//		}
//	}

//	for n := 0; n < 2; n++ { //二次平滑突出
//		for i := 1; i < height-1; i++ { //平滑突出
//			for j := 1; j < width-1; j++ {
//				CheckExtrude(i, j)
//			}
//		}
//	}
//	return src
//}

func FindEdges(src [][][]uint8) [][][]uint8 { //查找边缘
	height := len(src)
	width := len(src[0])
	for i := 0; i < height; i++ { //二值化
		for j := 0; j < width; j++ {
			if (i == 0) || (j == 0) || (i == height-1) || (j == width-1) {
				src[i][j][0] = 0
				src[i][j][1] = 0
				src[i][j][2] = 0
				src[i][j][3] = 255
			} else {
				if (src[i][j][0] + src[i][j][1] + src[i][j][2]) > 0 {
					src[i][j][0] = 255 //二值化数据
					if (i == 0) || (i == height-1) || (j == 0) || (j == width-1) {
						src[i][j][1] = 255 //边缘标记
					} else {
						src[i][j][1] = 0 //边缘标记
					}
					src[i][j][2] = 0 //平滑后边缘数据
					src[i][j][3] = 255
				} else {
					src[i][j][0] = 0
					src[i][j][1] = 0
					src[i][j][2] = 0
					src[i][j][3] = 255
				}
			}
		}
	}

	for n := 0; n < 3; n++ { //3次补值
		for i := 0; i < height-2; i++ { //补值
			for j := 0; j < width-2; j++ {
				if src[i][j][0] > 0 && src[i+1][j][0] == 0 && src[i+2][j][0] > 0 { //i方向补值
					src[i+1][j][0] = 255
				}
				if src[i][j][0] > 0 && src[i][j+1][0] == 0 && src[i][j+2][0] > 0 { //j方向补值
					src[i][j+1][0] = 255
				}

				if src[i][j][0] > 0 && src[i+1][j+1][0] == 0 && src[i+2][j+2][0] > 0 { //ij方向补值
					src[i+1][j+1][0] = 255
				}
			}
		}
	}

	Check8Direction := func(i, j int) bool { //八个方向检测, 非全包围点返回 true
		if src[i][j][0] == 0 { //无值点跳过
			return false
		}
		a := 0 //有值计数
		if src[i+1][j][0] != 0 {
			a++
		}
		if src[i][j+1][0] != 0 {
			a++
		}
		if src[i+1][j+1][0] != 0 {
			a++
		}
		if src[i-1][j][0] != 0 {
			a++
		}
		if src[i][j-1][0] != 0 {
			a++
		}
		if src[i-1][j-1][0] != 0 {
			a++
		}
		if src[i+1][j+1][0] != 0 {
			a++
		}
		if src[i-1][j+1][0] != 0 {
			a++
		}
		if (a != 0) && (a != 8) {
			return true
		} else {
			return false
		}
	}
	for i := 1; i < height-1; i++ { //查找边缘
		for j := 1; j < width-1; j++ {
			if Check8Direction(i, j) {
				src[i][j][1] = 255
			}
		}
	}

	CheckExtrude := func(i, j int) { //处理突出点
		if src[i][j][1] == 0 { //无值点跳过
			return
		}
		a := 0 //左上
		b := 0 //上
		c := 0 //右上
		d := 0 //左
		e := 0 //右
		f := 0 //左下
		g := 0 //下
		h := 0 //右下
		if src[i-1][j-1][1] != 0 {
			a = 1
		}
		if src[i-1][j][1] != 0 {
			b = 1
		}
		if src[i+1][j+1][1] != 0 {
			c = 1
		}
		if src[i][j-1][1] != 0 {
			d = 1
		}
		if src[i][j+1][1] != 0 {
			e = 1
		}
		if src[i+1][j-1][1] != 0 {
			f = 1
		}
		if src[i+1][j][1] != 0 {
			g = 1
		}
		if src[i+1][j+1][1] != 0 {
			h = 1
		}

		//线状突出, 填充缺失后删除
		if (a+c == 2) && (b+d+e+f+g+h == 0) {
			src[i][j][1] = 0
			src[i-1][j][1] = 255
		}
		if (c+h == 2) && (a+b+d+e+f+g == 0) {
			src[i][j][1] = 0
			src[i][j+1][1] = 255
		}
		if (f+h == 2) && (a+b+c+d+e+g == 0) {
			src[i][j][1] = 0
			src[i+1][j][1] = 255
		}
		if (a+f == 2) && (b+c+d+e+g+h == 0) {
			src[i][j][1] = 0
			src[i][j-1][1] = 255
		}

		if (a+c+d+e == 4) && (b+f+g+h == 0) {
			src[i][j][1] = 0
			src[i][j-1][1] = 0
			src[i][j+1][1] = 0
			src[i-1][j][1] = 255
		}
		//		if (c+h+b+g == 4) && (a+d+e+f == 0) {
		//			src[i][j][1] = 0
		//			src[i-1][j][1] = 0
		//			src[i+1][j][1] = 0
		//			src[i][j+1][1] = 255
		//		}
		//		if (f+h+d+e == 4) && (a+b+c+g == 0) {
		//			src[i][j][1] = 0
		//			src[i][j-1][1] = 0
		//			src[i][j+1][1] = 0
		//			src[i+1][j][1] = 255
		//		}
		//		if (a+f+b+g == 4) && (c+d+e+h == 0) {
		//			src[i][j][1] = 0
		//			src[i-1][j][1] = 0
		//			src[i+1][j][1] = 0
		//			src[i][j-1][1] = 255
		//		}
		//点状突出, 直接删除
		if ((a+d+f == 3) && (b+c+e+h+g == 0)) || ((a+b+c == 3) && (d+e+f+g+h == 0)) || ((c+e+h == 3) && (a+b+d+f+g == 0)) || ((f+g+h == 3) && (a+b+c+d+e == 0)) {
			src[i][j][1] = 0
		}
	}

	for n := 0; n < 3; n++ { //二次平滑突出
		for i := 1; i < height-1; i++ { //平滑突出
			for j := 1; j < width-1; j++ {
				CheckExtrude(i, j)
			}
		}
	}

	//	CheckDepression := func(i, j int) bool { //凹陷点检测
	//		if src[i][j][1] != 0 { //有值点跳过
	//			return false
	//		}
	//		a := 0 //左上
	//		b := 0 //上
	//		c := 0 //右上
	//		d := 0 //左
	//		e := 0 //右
	//		f := 0 //左下
	//		g := 0 //下
	//		h := 0 //右下
	//		if src[i-1][j-1][1] != 0 {
	//			a = 1
	//		}
	//		if src[i-1][j][1] != 0 {
	//			b = 1
	//		}
	//		if src[i+1][j+1][1] != 0 {
	//			c = 1
	//		}
	//		if src[i][j-1][1] != 0 {
	//			d = 1
	//		}
	//		if src[i][j+1][1] != 0 {
	//			e = 1
	//		}
	//		if src[i+1][j-1][1] != 0 {
	//			f = 1
	//		}
	//		if src[i+1][j][1] != 0 {
	//			g = 1
	//		}
	//		if src[i+1][j+1][1] != 0 {
	//			h = 1
	//		}
	//		if ((a+d+f == 0) && (b+c+e+h+g == 5)) || ((a+b+c == 0) && (d+e+f+g+h == 5)) || ((c+e+h == 0) && (a+b+d+f+g == 5)) || ((f+g+h == 0) && (a+b+c+d+e == 5)) {
	//			return true
	//		} else {
	//			return false
	//		}
	//	}
	//	for i := 1; i < height-1; i++ { //平滑凹陷
	//		for j := 1; j < width-1; j++ {
	//			if CheckDepression(i, j) {
	//				src[i][j][1] = 255
	//			}
	//			src[i][j][0] = 0
	//		}
	//	}

	for i := 0; i < height; i++ { //恢复图像
		for j := 0; j < width; j++ {
			src[i][j][0] = src[i][j][1]
			src[i][j][2] = src[i][j][1]
		}
	}

	return src
}

func GetPoints(src [][][]uint8, step int) []TPointCC { //生成 5*12 形状直方图
	height := len(src)
	width := len(src[0])

	//step := 2 //采样矩阵 step*step

	ret := []TPointCC{}

	if (height <= step) || (width <= step) {
		return []TPointCC{}
	}

	for i := step + 1; i < height+step; i += step {
		step_i := step
		tmpi := i
		if i > height-1 {
			tmpi = height - 1
			step_i = step - (i - height)
			//break
		}
		for j := step; j < width+step; j += step {
			pixV := 0
			step_j := step
			tmpj := j
			if j > width-1 {
				tmpj = width - 1
				step_j = step - (j - width)
				//break
			}
			xsumTmp := 0
			ysumTmp := 0
			xyCountTmp := 0

			for iT := tmpi; iT >= i-step_i; iT-- { //遍历 step * step 窗口里的值, 判断存在有效像素
				for jT := tmpj; jT >= j-step_j; jT-- {
					if (src[iT][jT][0] + src[iT][jT][1] + src[iT][jT][2]) > 0 {
						pixV++
						ysumTmp += iT
						xsumTmp += jT
						xyCountTmp++
					}
				}
			}
			if step-1 < pixV { //窗口内, 有效像素采样
				TmpCList := make([][]int, splitR, splitR) //初始化上下文矩阵
				for Ti := 0; Ti < splitR; Ti++ {
					s2 := make([]int, splitS, splitS)
					TmpCList[Ti] = s2
				}

				var TmpTPointCC TPointCC

				TmpTPointCC.P.X = xsumTmp / xyCountTmp
				TmpTPointCC.P.Y = ysumTmp / xyCountTmp

				TmpTPointCC.CList = TmpCList

				ret = append(ret, TmpTPointCC)
			}
		}
	}

	minX := width
	maxX := 0
	minY := height
	maxY := 0

	for i := 0; i < len(ret); i++ {
		if maxX < ret[i].P.X {
			maxX = ret[i].P.X
		}
		if minX > ret[i].P.X {
			minX = ret[i].P.X
		}
		if maxY < ret[i].P.Y {
			maxY = ret[i].P.Y
		}
		if minY > ret[i].P.Y {
			minY = ret[i].P.Y
		}

	}

	heightLine := maxY - minY
	widthLine := maxX - minX

	getRJ := func(A, B float64) (float64, float64) { //获得向量长度与角度
		if (A == 0) || (B == 0) {
			if A == 0 {
				return B, 0
			} else {
				return A, 0
			}
		} else {
			R := math.Sqrt(A*A + B*B) //斜边R

			JJ := (math.Acos(B/R) * 180) / math.Pi
			return R, JJ
		}
	}

	R := float64(min(heightLine, widthLine) * 1 / 2) //math.Sqrt(float64(heightLine*heightLine+widthLine*widthLine)) / 2
	Rlimt := R / float64(splitR)

	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret); j++ {
			if ret[i].P != ret[j].P {
				//math.Acos(x)
				pX1 := float64(ret[j].P.X)
				pX0 := float64(ret[i].P.X)
				pY1 := float64(ret[j].P.Y)
				pY0 := float64(ret[i].P.Y)
				pA := float64(0)
				pB := float64(0)
				pR := float64(0)
				pJ := float64(0)
				index0 := 0 //角度分区
				index1 := 0 //半径分区
				ok := false

				if pX1 >= pX0 { //右半圆
					pA = pX1 - pX0
					if pY1 >= pY0 { //右下半圆
						//fmt.Println("右下半圆")
						pB = pY1 - pY0 // 3 4 5
						//TODO
						pR, pJ = getRJ(pB, pA)
						if pR < R { //在有效采集范围内 (360/float64(splitS)=30
							index0 = int(pJ/(360/float64(splitS)) + float64(splitS/4)) //2 3 - 2 id=1  (float64(splitS)*1)
							index1 = int(pR / Rlimt)
							ok = true
						}
					} else { //右上半圆
						//fmt.Println("右上半圆")
						pB = pY0 - pY1 //0 1 2
						//TODO
						pR, pJ = getRJ(pA, pB)
						if pR < R { //在有效采集范围内
							index0 = int(pJ/(360/float64(splitS)) + 0) //0 0 - 1 id=0  (float64(splitS)*0)
							index1 = int(pR / Rlimt)
							ok = true
						}
					}
				} else { //左半圆
					pA = pX0 - pX1
					if pY1 >= pY0 { //左下半圆
						//fmt.Println("左下半圆")
						pB = pY1 - pY0 //6 7 8
						//TODO
						pR, pJ = getRJ(pA, pB)
						if pR < R { //在有效采集范围内
							index0 = int(pJ/(360/float64(splitS)) + (float64(splitS/4) * 2)) //4 6 - 3 id=2 (float64(splitS)*2)
							index1 = int(pR / Rlimt)
							ok = true
						}
					} else { //左上半圆
						//fmt.Println("左上半圆")
						pB = pY0 - pY1 //9 10 11
						//TODO
						pR, pJ = getRJ(pB, pA)
						if pR < R { //在有效采集范围内
							index0 = int(pJ/(360/float64(splitS)) + (float64(splitS/4) * 3)) //6 9 - 4 id=3 (float64(splitS)*3)
							index1 = int(pR / Rlimt)
							ok = true
						}
					}
				}
				if ok {
					//fmt.Println("------>", index1, index0)
					ret[i].CList[index1][index0]++
				}
			}
		}
	}
	return ret
}

func doCompareTPCC(mod, src []TPointCC) (map[int]int, float64) { //穷举法匹配

	modLen := len(mod)
	srcLen := len(src)

	retTmp := make(map[int]int)
	//retTmpOK := make(map[int]int)

	inRet := func(n int) bool {
		if _, ok := retTmp[n]; ok {
			return true
		} else {
			return false
		}
	}

	//	getMinLimit := func(ret []TPointCC) int {
	//		minX := 99999
	//		maxX := 0
	//		minY := 99999
	//		maxY := 0

	//		for i := 0; i < len(ret); i++ {
	//			if maxX < ret[i].P.X {
	//				maxX = ret[i].P.X
	//			}
	//			if minX > ret[i].P.X {
	//				minX = ret[i].P.X
	//			}
	//			if maxY < ret[i].P.Y {
	//				maxY = ret[i].P.Y
	//			}
	//			if minY > ret[i].P.Y {
	//				minY = ret[i].P.Y
	//			}
	//		}

	//		return min(maxY-minY, maxX-minX)
	//	}

	//	getMaxLimit := func(ret []TPointCC) int {
	//		minX := 99999
	//		maxX := 0
	//		minY := 99999
	//		maxY := 0

	//		for i := 0; i < len(ret); i++ {
	//			if maxX < ret[i].P.X {
	//				maxX = ret[i].P.X
	//			}
	//			if minX > ret[i].P.X {
	//				minX = ret[i].P.X
	//			}
	//			if maxY < ret[i].P.Y {
	//				maxY = ret[i].P.Y
	//			}
	//			if minY > ret[i].P.Y {
	//				minY = ret[i].P.Y
	//			}
	//		}

	//		return max(maxY-minY, maxX-minX)
	//	}

	CostAvg := float64(0)
	//maxLimit := float64(max(getMinLimit(mod), getMinLimit(src)))

	for i := 0; i < modLen; i++ {
		minNum := float64(0)
		srcIndex := 0
		fFlag := true

		for j := 0; j < srcLen; j++ {
			if !inRet(j) {
				t := GetCost(mod[i].CList, src[j].CList)
				//fmt.Println(t)
				//if t < 30 {
				if fFlag {
					minNum = t
					fFlag = false
					srcIndex = j
				} else {
					if minNum > t {
						srcIndex = j
						minNum = t
					}
				}
				//}
			}
		}
		//if srcIndex != 0 {
		//fmt.Println("--->", minNum, srcIndex, src[srcIndex].P)

		//		A := float64(mod[i].P.X - src[srcIndex].P.X)
		//		B := float64(mod[i].P.Y - src[srcIndex].P.Y)
		//		tR := math.Sqrt(A*A + B*B)
		//if (maxLimit/2 > tR) && (minNum < 30) {
		retTmp[srcIndex] = i
		CostAvg += minNum
		//fmt.Println("[", minNum, "]")
		//}
		//}
	}

	ret := make(map[int]int)
	for k, v := range retTmp {
		if _, ok := ret[v]; !ok {
			ret[v] = k
		}
	}

	CostAvg = math.Sqrt(CostAvg / float64(len(retTmp)))
	return ret, CostAvg
}

/*func doCompareTPCC(mod, src []TPointCC) (map[int]int, float64) { //穷举法匹配

	modLen := len(mod)
	srcLen := len(src)

	retTmp := make(map[int]int)
	//retTmpOK := make(map[int]int)

	inRet := func(n int) bool {
		//tret := false
		// for _, v := range retTmp {
		// 	if v == n {
		// 		tret = true
		// 	}
		// }
		if _, ok := retTmp[n]; ok {
			return true
		} else {
			return false
		}
	}

	CostAvg := float64(0)

	for i := 0; i < modLen; i++ {
		minNum := float64(0)
		srcIndex := 0
		fFlag := true

		srcIndexTmp := make(map[int]int)
		minNumTmp := make(map[int]float64)
		index := 0

		for j := 0; j < srcLen; j++ {
			if !inRet(j) {
				t := GetCost(mod[i].CList, src[j].CList)
				if fFlag {
					minNum = t
					fFlag = false
					srcIndexTmp[index] = j
					minNumTmp[index] = t
				} else {
					if minNum > t {
						srcIndexTmp = make(map[int]int)
						minNumTmp = make(map[int]float64)
						index = 0
						srcIndex = j
						minNum = t
						srcIndexTmp[index] = j
						minNumTmp[index] = t
					} else {
						if minNum == t {
							index++
							srcIndexTmp[index] = j
							minNumTmp[index] = t
						}
					}
				}
			}
		}
		if srcIndex != 0 {
			if index > 0 {
				//fmt.Println(index)
				minR := float64(9999999)
				minRi := 0
				//minRsrcIndex := 0
				for ii, j := range srcIndexTmp {
					srcX := src[j].P.X
					srcY := src[j].P.Y
					modX := mod[i].P.X
					modY := mod[i].P.Y
					tmpR := math.Sqrt(float64((srcX-modX)*(srcX-modX) + (srcY-modY)*(srcY-modY)))
					if minR > tmpR {
						minR = tmpR
						minRi = ii
					}
				}
				retTmp[srcIndexTmp[minRi]] = i
				CostAvg += minNumTmp[minRi]
			} else {
				retTmp[srcIndex] = i
				//fmt.Println(minNum)
				CostAvg += minNum
			}
		}
	}

	ret := make(map[int]int)
	for k, v := range retTmp {
		if _, ok := ret[v]; !ok {
			ret[v] = k
		}
	}

	CostAvg = CostAvg / float64(len(retTmp))
	return ret, CostAvg
}*/

func OutPointCount(mod, src []TPointCC, pp map[int]int) float64 { //统计未覆盖点占比
	if (len(mod) == len(src)) && (len(src) == len(pp)) { //全部匹配,不存在未覆盖点
		return 0
	}
	modTj := make(map[int](map[int]int)) //[x][y]int

	markLine := func(x0, y0, x1, y1 int) {
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
			//brush(x0, y0)
			if _, ok := modTj[x0]; ok {
				if _, ok := modTj[x0][y0]; ok {
					modTj[x0][y0]++
				}
			}
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

	for i := 0; i < len(mod); i++ {
		tmp := mod[i].P
		if modTj[tmp.X] == nil {
			modTj[tmp.X] = make(map[int]int)
		}
		modTj[tmp.X][tmp.Y] = 0
	}
	for k, v := range pp {
		X0 := mod[k].P.X
		Y0 := mod[k].P.Y
		X1 := src[v].P.X
		Y1 := src[v].P.Y
		markLine(X0, Y0, X1, Y1)
	}
	//fmt.Println(modTj)
	allCount := float64(len(mod))
	zeroCount := float64(0)

	for _, v1 := range modTj {
		for _, v2 := range v1 {
			//fmt.Print(v2, " ")
			if v2 == 0 {
				zeroCount++
			}
		}
	}
	//fmt.Println(zeroCount / allCount)
	return zeroCount / allCount
}

func CompareTPCC(mod, src []TPointCC) (map[int]int, float64) {
	if len(mod) < len(src) {
		a, _ := doCompareTPCC(mod, src)
		ret := make(map[int]int)
		for k, v := range a {
			ret[v] = k
		}
		return a, OutPointCount(src, mod, ret)
	} else {
		a, _ := doCompareTPCC(src, mod)
		ret := make(map[int]int)
		for k, v := range a {
			ret[v] = k
		}
		return ret, OutPointCount(mod, src, ret)
	}
}

func GetCost(src1, src2 [][]int) float64 { //两个形状直方图的匹配代价
	syn := float64(0)
	for i := 0; i < splitR; i++ {
		for j := 0; j < splitS; j++ {
			t := float64(src1[i][j] - src2[i][j])
			t = t * t
			d := float64(src1[i][j] + src2[i][j])
			if d != 0 {
				syn += t / d
			}
		}
	}
	return syn / 2
}

func CCtest(mod, src []TPointCC) {
	modLen := len(mod)
	srcLen := len(src)
	for i := 0; i < modLen; i++ {
		for j := 0; j < srcLen; j++ {
			fmt.Println(GetCost(mod[i].CList, src[j].CList))
		}
	}
}

func CosineSum(mod, src []TPointCC, pp map[int]int) float64 { //夹角余弦距离(Cosine)
	mod, src = ResetCenterMass(mod, src) //归一化
	a := float64(0)
	b := float64(0)
	c := float64(0)
	for k, v := range pp {
		if v < len(src) {
			x1 := float64(mod[k].P.X)
			y1 := float64(mod[k].P.Y)
			//fmt.Println(" >", v, len(src))
			//fmt.Println(" >>", src[v].P.X)
			x2 := float64(src[v].P.X)
			y2 := float64(src[v].P.Y)
			a += x1*x2 + y1*y2
			b += x1*x1 + y1*y1
			c += x2*x2 + y2*y2
		}
	}
	t := (math.Sqrt(b) * math.Sqrt(c))
	ret := float64(0)
	if a == 0 {
		ret = 0
	} else {
		if t == 0 {
			ret = 1
		} else {
			ret = a / t
		}
	}
	return ret
}

//(x平移rx0,y平移ry0,角度a对应-RotaryAngle
//x'= (x - rx0)*cos(RotaryAngle) + (y - ry0)*sin(RotaryAngle) + rx0 ;
//y'=-(x - rx0)*sin(RotaryAngle) + (y - ry0)*cos(RotaryAngle) + ry0 ;
//根据新的坐标点求源坐标点的公式为：
//x=(x'- rx0)*cos(RotaryAngle) - (y'- ry0)*sin(RotaryAngle) + rx0 ;
//y=(x'- rx0)*sin(RotaryAngle) + (y'- ry0)*cos(RotaryAngle) + ry0 ;
//角度 := (弧度 * 180) / math.Pi
//弧度 := (角度 * math.Pi) / 180
func XzRotary(src []TPointCC, pp map[int]int, RotaryAngle float64) []TPointCC { //旋转变换
	//TODO
	//math.Cos()
	//math.Sin()
	//X0, Y0 := GetCenterMass(src)
	//	X0 := src[0].P.X
	//	Y0 := src[0].P.Y
	h := (RotaryAngle * math.Pi) / 180

	for _, v := range pp {
		X := float64(src[v].P.X)
		Y := float64(src[v].P.Y)
		//		src[i].P.X = int((X-float64(X0))*math.Cos(h) - (Y-float64(X0))*math.Sin(h) + float64(X0))
		//		src[i].P.Y = int(-(X-float64(X0))*math.Sin(h) + (Y-float64(Y0))*math.Cos(h) + float64(Y0))
		src[v].P.X = int(X*math.Cos(h) + Y*math.Sin(h))
		src[v].P.Y = int(Y*math.Cos(h) - X*math.Sin(h))
	}

	//	for i := 0; i < len(src); i++ {
	//		X := float64(src[i].P.X)
	//		Y := float64(src[i].P.Y)
	//		//		src[i].P.X = int((X-float64(X0))*math.Cos(h) - (Y-float64(X0))*math.Sin(h) + float64(X0))
	//		//		src[i].P.Y = int(-(X-float64(X0))*math.Sin(h) + (Y-float64(Y0))*math.Cos(h) + float64(Y0))
	//		src[i].P.X = int(X*math.Cos(h) + Y*math.Sin(h))
	//		src[i].P.Y = int(Y*math.Cos(h) - X*math.Sin(h))
	//	}
	//	minX := 99999
	//	minY := 99999
	//	for i := 0; i < len(src); i++ { //获取最小XY轴
	//		if minX > src[i].P.X {
	//			minX = src[i].P.X
	//		}
	//		if minY > src[i].P.Y {
	//			minY = src[i].P.Y
	//		}
	//	}
	//	pX := 0
	//	if minX < 0 { //x轴平移
	//		pX = 0 - minX
	//	}
	//	pY := 0
	//	if minY < 0 { //y轴平移
	//		pY = 0 - minY
	//	}
	//	for i := 0; i < len(src); i++ {
	//		src[i].P.X += pX
	//		src[i].P.Y += pY
	//	}
	return src
}

func CosineSum2(mod, srcY []TPointCC, pp map[int]int) (float64, float64) { //夹角余弦距离(Cosine), 最优旋转变换, 顺时针逆时针60度, 步进10度
	maxRotaryAngle := float64(0)
	maxC := CosineSum(mod, srcY, pp)
	//maxSrc := src
	src := XzRotary(srcY, pp, -35)

	xzRotary := func(RotaryAngle float64) {
		h := (RotaryAngle * math.Pi) / 180

		for _, v := range pp {
			X := float64(src[v].P.X)
			Y := float64(src[v].P.Y)
			//		src[i].P.X = int((X-float64(X0))*math.Cos(h) - (Y-float64(X0))*math.Sin(h) + float64(X0))
			//		src[i].P.Y = int(-(X-float64(X0))*math.Sin(h) + (Y-float64(Y0))*math.Cos(h) + float64(Y0))
			src[v].P.X = int(X*math.Cos(h) + Y*math.Sin(h))
			src[v].P.Y = int(Y*math.Cos(h) - X*math.Sin(h))
		}
	}

	for i := float64(-30); i <= 30; i += 5 {
		//fmt.Println(i)
		if i != 0 {
			//newSrc := XzRotary(src, i)
			xzRotary(5)
			newC := CosineSum(mod, src, pp)
			if maxC < newC {
				maxRotaryAngle = i
				maxC = newC
			}
		}
	}
	//fmt.Println("修正角度:", maxRotaryAngle)
	return maxC, maxRotaryAngle
}

func EuclideanDistanceSum(mod, src []TPointCC, pp map[int]int) float64 { //欧氏距离(Euclidean Distance)
	a := float64(0)
	for k, v := range pp {
		x1 := float64(mod[k].P.X)
		y1 := float64(mod[k].P.Y)
		x2 := float64(src[v].P.X)
		y2 := float64(src[v].P.Y)

		a += (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	}
	ret := math.Sqrt(a)
	//fmt.Println("...EuclideanDistanceSum:", ret)
	return ret
}

func ManhattanDistanceSum(mod, src []TPointCC, pp map[int]int) float64 { //曼哈顿距离(Manhattan Distance)
	a := float64(0)
	for k, v := range pp {
		x1 := float64(mod[k].P.X)
		y1 := float64(mod[k].P.Y)
		x2 := float64(src[v].P.X)
		y2 := float64(src[v].P.Y)

		a += math.Sqrt((x1-x2)*(x1-x2)) + math.Sqrt((y1-y2)*(y1-y2))
	}
	//fmt.Println("...ManhattanDistanceSum:", a)
	return a
}

func GetCenterMass(src []TPointCC) (int, int) { //质心的计算
	xSum := 0
	ySum := 0
	for i := 0; i < len(src); i++ {
		xSum += src[i].P.X
		ySum += src[i].P.Y
	}
	if len(src) == 0 {
		return 0, 0
	}
	xAvg := xSum / len(src)
	yAvg := ySum / len(src)
	return xAvg, yAvg
}

func getTPointCCWH(src []TPointCC) (int, int, int, int) { //获取采样集合长宽
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	minXF := true
	minYF := true
	for i := 0; i < len(src); i++ {
		if minXF {
			minX = src[i].P.X
			minXF = false
		} else {
			if minX >= src[i].P.X {
				minX = src[i].P.X
			}
		}
		if minYF {
			minY = src[i].P.Y
			minYF = false
		} else {
			if minY >= src[i].P.Y {
				minY = src[i].P.Y
			}
		}
		if maxX <= src[i].P.X {
			maxX = src[i].P.X
		}
		if maxY <= src[i].P.Y {
			maxY = src[i].P.Y
		}
	}
	return maxX - minX, maxY - minY, minX, minY
}

func ResetCenterMass(src1, src2 []TPointCC) ([]TPointCC, []TPointCC) { //同步质心, 归一化处理
	//同比长宽
	w1, h1, _, _ := getTPointCCWH(src1)
	w2, h2, _, _ := getTPointCCWH(src2)

	r1 := math.Sqrt(float64((w1 * w1) + (h1 * h1)))
	r2 := math.Sqrt(float64((w2 * w2) + (h2 * h2)))

	if r1 > r2 {
		xBei := float64(r1) / float64(r2)
		for i := 0; i < len(src2); i++ {
			src2[i].P.X = int(float64(src2[i].P.X) * xBei)
			src2[i].P.Y = int(float64(src2[i].P.Y) * xBei)
		}
	} else {
		xBei := float64(r2) / float64(r1)
		for i := 0; i < len(src1); i++ {
			src1[i].P.X = int(float64(src1[i].P.X) * xBei)
			src1[i].P.Y = int(float64(src1[i].P.Y) * xBei)
		}
	}

	//同步动质心
	x1CM, y1CM := GetCenterMass(src1)
	x2CM, y2CM := GetCenterMass(src2)
	if (x1CM == 0) && (y1CM == 0) {
		return src1, src2
	}

	if (x2CM == 0) && (y2CM == 0) {
		return src1, src2
	}

	xMove := int(math.Abs(float64(x1CM - x2CM)))
	if x1CM > x2CM {
		for i := 0; i < len(src2); i++ {
			src2[i].P.X = src2[i].P.X + xMove
		}
	} else {
		if x1CM != x2CM {
			for i := 0; i < len(src1); i++ {
				src1[i].P.X = src1[i].P.X + xMove
			}
		}
	}
	yMove := int(math.Abs(float64(y1CM - y2CM)))
	if y1CM > y2CM {
		for i := 0; i < len(src2); i++ {
			src2[i].P.Y = src2[i].P.Y + yMove
		}
	} else {
		if y1CM != y2CM {
			for i := 0; i < len(src1); i++ {
				src1[i].P.Y = src1[i].P.Y + yMove
			}
		}
	}

	//趋近极坐标
	_, _, left1, top1 := getTPointCCWH(src1)
	_, _, left2, top2 := getTPointCCWH(src2)

	minleft := min(left1, left2)
	mintop := min(top1, top2)
	//	minleft = left1
	//	mintop = top1

	//	ret1 := []TPointCC{}
	//	ret2 := []TPointCC{}

	//	flaMap1 := make(map[string]bool)
	//	flaMap2 := make(map[string]bool)

	for i := 0; i < len(src1); i++ {
		src1[i].P.Y = src1[i].P.Y - mintop
		src1[i].P.X = src1[i].P.X - minleft /*
			if _, ok := flaMap1[strconv.Itoa(src1[i].P.X)+"_"+strconv.Itoa(src1[i].P.Y)]; !ok {
				flaMap1[strconv.Itoa(src1[i].P.X)+"_"+strconv.Itoa(src1[i].P.Y)] = true
				ret1 = append(ret1, src1[i])
			}*/
	}
	for i := 0; i < len(src2); i++ {
		src2[i].P.Y = src2[i].P.Y - mintop
		src2[i].P.X = src2[i].P.X - minleft /*
			if _, ok := flaMap2[strconv.Itoa(src2[i].P.X)+"_"+strconv.Itoa(src2[i].P.Y)]; !ok {
				flaMap2[strconv.Itoa(src2[i].P.X)+"_"+strconv.Itoa(src2[i].P.Y)] = true
				ret2 = append(ret2, src2[i])
			}*/
	}

	return src1, src2
}

func ResetCenterHeightMass(src1, src2 []TPointCC) ([]TPointCC, []TPointCC) { //同步高度质心, 归一化处理
	//同比长宽
	//w1, h1, _, _ := getTPointCCWH(src1)
	//w2, h2, _, _ := getTPointCCWH(src2)

	//	x1CM, y1CM := GetCenterMass(src1)
	//	x2CM, y2CM := GetCenterMass(src2)

	//	r1 := math.Sqrt(float64((x1CM * x1CM) + (y1CM * y1CM)))
	//	r2 := math.Sqrt(float64((x2CM * x2CM) + (y2CM * y2CM)))

	//	//if r1 > r2 {
	//	xBei := float64(r1) / float64(r2)
	//	//fmt.Println(h1, w1, h2, w2)
	//	for i := 0; i < len(src2); i++ {
	//		src2[i].P.X = int(float64(src2[i].P.X) * xBei)
	//		src2[i].P.Y = int(float64(src2[i].P.Y) * xBei)
	//	}
	//	} else {
	//		xBei := float64(r2) / float64(r1)
	//		for i := 0; i < len(src1); i++ {
	//			src1[i].P.X = int(float64(src1[i].P.X) * xBei)
	//			src1[i].P.Y = int(float64(src1[i].P.Y) * xBei)
	//		}
	//	}

	//同步动质心
	x1CM, y1CM := GetCenterMass(src1)
	x2CM, y2CM := GetCenterMass(src2)
	if (x1CM == 0) && (y1CM == 0) {
		return src1, src2
	}

	if (x2CM == 0) && (y2CM == 0) {
		return src1, src2
	}

	xMove := int(math.Abs(float64(x1CM - x2CM)))
	if x1CM > x2CM {
		for i := 0; i < len(src2); i++ {
			src2[i].P.X = src2[i].P.X + xMove
		}
	} else {
		if x1CM != x2CM {
			for i := 0; i < len(src1); i++ {
				src1[i].P.X = src1[i].P.X + xMove
			}
		}
	}
	yMove := int(math.Abs(float64(y1CM - y2CM)))
	if y1CM > y2CM {
		for i := 0; i < len(src2); i++ {
			src2[i].P.Y = src2[i].P.Y + yMove
		}
	} else {
		if y1CM != y2CM {
			for i := 0; i < len(src1); i++ {
				src1[i].P.Y = src1[i].P.Y + yMove
			}
		}
	}

	//趋近极坐标
	_, _, left1, top1 := getTPointCCWH(src1)
	_, _, left2, top2 := getTPointCCWH(src2)

	minleft := min(left1, left2)
	mintop := min(top1, top2)
	//	minleft = left1
	//	mintop = top1

	//	ret1 := []TPointCC{}
	//	ret2 := []TPointCC{}

	//	flaMap1 := make(map[string]bool)
	//	flaMap2 := make(map[string]bool)

	for i := 0; i < len(src1); i++ {
		src1[i].P.Y = src1[i].P.Y - mintop
		src1[i].P.X = src1[i].P.X - minleft /*
			if _, ok := flaMap1[strconv.Itoa(src1[i].P.X)+"_"+strconv.Itoa(src1[i].P.Y)]; !ok {
				flaMap1[strconv.Itoa(src1[i].P.X)+"_"+strconv.Itoa(src1[i].P.Y)] = true
				ret1 = append(ret1, src1[i])
			}*/
	}
	for i := 0; i < len(src2); i++ {
		src2[i].P.Y = src2[i].P.Y - mintop
		src2[i].P.X = src2[i].P.X - minleft /*
			if _, ok := flaMap2[strconv.Itoa(src2[i].P.X)+"_"+strconv.Itoa(src2[i].P.Y)]; !ok {
				flaMap2[strconv.Itoa(src2[i].P.X)+"_"+strconv.Itoa(src2[i].P.Y)] = true
				ret2 = append(ret2, src2[i])
			}*/
	}

	return src1, src2
}

func GetR(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}

func Coverage(mod, src []TPointCC, pp map[int]int) { //覆盖率
	// maxTPointCC := func() (bool, []TPointCC) {
	// 	if len(mod) > len(mod) {
	// 		return true, mod
	// 	} else {
	// 		return false, src
	// 	}
	// }
	// ms, orz := maxTPointCC()
}
