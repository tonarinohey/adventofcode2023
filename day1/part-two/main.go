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

	reader := bufio.NewReaderSize(fp, 64)
	var tmp []byte

	var sumNumForInput int64

	// 各行処理
	for {
		// 各行文字列取得
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		println(string(line))

		// 先頭の数字を取得する
		strFirstNum := convertStringToNum(getFirstNumFromString(string(line)))
		println(strFirstNum)
		// 最後尾の数字を取得する
		strLastNum := convertStringToNum(getLastNumFromString(string(line)))
		println(strLastNum)

		// 先頭と末尾の数字を合体し2桁の数字とする
		numInLine, _ := strconv.ParseInt(strFirstNum+strLastNum, 10, 64)
		println(numInLine)

		sumNumForInput = sumNumForInput + numInLine

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
		println(sumNumForInput)
		println()
	}

	println(sumNumForInput)
}

// 文字列から数字だけを抽出する(ex. 47hcbcnhjgxhdfn7rrbonehvlmddbone -> 477)
func getNumFromString(strVal string) string {
	// 正規表現指定
	r := `(\d+|one|two|three|four|five|six|seven|eight|nine)`
	rex := regexp.MustCompile(r)

	// 最初にヒットした数を返す
	return rex.FindString(strVal)
}

func getFirstNumFromString(strVal string) string {
	// 一行の文字列を1文字ずつ分割し配列に格納する
	splitStrVal := strings.Split(strVal, "")
	var firstString string

	// 分割した文字列先頭から1字ずつ再度結合していく
	for i := 0; i < len(splitStrVal); i++ {
		firstString = firstString + splitStrVal[i]

		if getNumFromString(firstString) != "" {
			return getNumFromString(firstString)
		}
	}

	return getNumFromString(firstString)
}

func getLastNumFromString(strVal string) string {
	// 一行の文字列を1文字ずつ分割し配列に格納する
	splitStrVal := strings.Split(strVal, "")
	var lastString string

	// 分割した文字列末端から1字ずつ再度結合していく
	for i := len(splitStrVal) - 1; i > -1; i-- {
		lastString = splitStrVal[i] + lastString

		if getNumFromString(lastString) != "" {
			return getNumFromString(lastString)
		}
	}

	return getNumFromString(lastString)
}

// // 数字だけの文字列（ex. "477"）をsplitし、先頭と末端を繋いだ数字(47)を返す
// func sumNumbers(input string) int64 {
// 	slice := strings.Split(input, "")
// 	len := len(slice)

// 	// 0文字の場合は0を返す
// 	if len == 0 {
// 		return 0
// 	}

// 	// 1文字の場合はその数字で２桁の数字を作る（ex. 7 なら 77）
// 	if len == 1 {
// 		val, _ := strconv.ParseInt(slice[0]+slice[0], 10, 64)
// 		return val
// 	}

// 	val, _ := strconv.ParseInt(slice[0]+slice[len-1], 10, 64)

// 	return val
// }

func convertStringToNum(input string) string {
	if input == "one" {
		return "1"
	} else if input == "two" {
		return "2"
	} else if input == "three" {
		return "3"
	} else if input == "four" {
		return "4"
	} else if input == "five" {
		return "5"
	} else if input == "six" {
		return "6"
	} else if input == "seven" {
		return "7"
	} else if input == "eight" {
		return "8"
	} else if input == "nine" {
		return "9"
	}

	return input
}
