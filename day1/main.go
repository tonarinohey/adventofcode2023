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
		line, isPrefix, err := reader.ReadLine() // size を超えるとisPrefixがfalse
		println(string(line))
		strNum := getNumFromString(string(line))
		sumNumForLine := sumNumbers(strNum)
		println(sumNumForLine)
		sumNumForInput = sumNumForInput + sumNumForLine

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
	}

	println(sumNumForInput)
}

// 文字列から数字だけを抽出する(ex. 47hcbcnhjgxhdfn7rrbonehvlmddbone -> 477)
func getNumFromString(strVal string) string {
	rex := regexp.MustCompile(`\d+`)
	var addNums string
	nums := rex.FindAllString(strVal, -1)

	for i := range nums {
		addNums = addNums + nums[i]
	}

	println(addNums)

	return addNums
}

// 数字だけの文字列（ex. "477"）をsplitし、先頭と末端を繋いだ数字(77)を返す
func sumNumbers(input string) int64 {
	slice := strings.Split(input, "")
	len := len(slice)
	// println(len)

	// 0文字の場合は0を返す
	if len == 0 {
		return 0
	}

	// 1文字の場合はその数字で２桁の数字を作る（ex. 7 なら 77）
	if len == 1 {
		val, _ := strconv.ParseInt(slice[0]+slice[0], 10, 64)

		return val
	}

	val, _ := strconv.ParseInt(slice[0]+slice[len-1], 10, 64)

	return val
}
