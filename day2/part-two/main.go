package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileName := "input.txt"
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 512) // 1行が長いので拡張する
	var tmp []byte

	var sumVal int

	id := 0

	// 各行処理
	for {
		// ID加算
		id++

		// -----------------ファイル読み取り---------------------
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		tmp = append(tmp, line...)
		if !isPrefix {
			tmp = nil
		}

		// -----------------処理---------------------
		// ": "で行を区切り前半（ID部）は捨てる
		omittedIDLine := strings.Split(string(line), ": ")

		// "; "で行を区切りsliceに格納する
		splittedLine := strings.Split(string(omittedIDLine[1]), "; ")

		sumVal = sumVal + check(splittedLine)
	}
	println(sumVal)
}

func check(splittedLine []string) int {
	// 上限
	largestRedCubeCount := 0
	largestBlueCubeCount := 0
	largestGreenCubeCount := 0

	rexBlue := regexp.MustCompile(`(.+?) blue`)
	rexGreen := regexp.MustCompile(`(.+?) green`)
	rexRed := regexp.MustCompile(`(.+?) red`)

	//
	for i := range splittedLine {
		// カンマでsplitする
		splittedLineWithComma := strings.Split(string(splittedLine[i]), ", ")

		for j := range splittedLineWithComma {
			arrBlue := rexBlue.FindStringSubmatch(splittedLineWithComma[j])
			if len(arrBlue) > 0 {
				numBlue, _ := strconv.ParseInt(arrBlue[1], 10, 64)
				// バーストしたら即return
				if int(numBlue) > largestBlueCubeCount {
					largestBlueCubeCount = int(numBlue)
				}
			}

			arrGreen := rexGreen.FindStringSubmatch(splittedLineWithComma[j])
			if len(arrGreen) > 0 {
				numGreen, _ := strconv.ParseInt(arrGreen[1], 10, 64)
				// バーストしたら即return
				if int(numGreen) > largestGreenCubeCount {
					largestGreenCubeCount = int(numGreen)
				}
			}

			arrRed := rexRed.FindStringSubmatch(splittedLineWithComma[j])
			if len(arrRed) > 0 {
				numRed, _ := strconv.ParseInt(arrRed[1], 10, 64)
				// バーストしたら即return
				if int(numRed) > largestRedCubeCount {
					largestRedCubeCount = int(numRed)
				}
			}
		}
	}

	return largestRedCubeCount * largestGreenCubeCount * largestBlueCubeCount
}
