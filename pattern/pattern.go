package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {

	defer executionTime(time.Now().UnixNano())

	searchIn := "Joho: 2578.34 William: 4567.32 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	if matched, _ := regexp.Match(pat, []byte(searchIn)); matched {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	v, _ := strconv.ParseFloat(searchIn, 32)
	fmt.Println(v)
}

func executionTime(start int64) {
	end := time.Now().UnixNano()
	fmt.Printf("Execution time is %v ns\n", (end - start))
}
