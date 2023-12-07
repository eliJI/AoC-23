package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const ZERO_INT = 48

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getNums(bytes []byte) []byte {
	fmt.Println(string(bytes))
	nums := make([]byte, 0)

	if len(bytes) < 1 {
		return []byte{byte('0'), byte('0')}
	}

	for _, v := range bytes {
		if unicode.IsNumber(rune(v)) {
			nums = append(nums, v)
		}
	}

	if len(bytes) > 1 {
		return []byte{nums[0], nums[len(nums)-1]}
	}

	return []byte{nums[0], nums[0]}
}

func replaceWordNums(line []byte) []byte {
	for k, v := range numberMap {
		line = []byte(strings.ReplaceAll(string(line), k, strconv.Itoa(v)))
	}
	return line
}

func addone(intptr *int) {
	*intptr += 1
}

func main() {
	sum := 0

	f, err := os.Open("testInput.txt")
	checkErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := replaceWordNums([]byte(scanner.Text()))
		fmt.Println(string(line))
	}
	b := 1
	addone(&b)
	fmt.Println(b)
	fmt.Println(sum)
}
