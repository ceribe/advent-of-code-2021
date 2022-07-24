package advent_of_code_2016

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// ReadInput reads lines from given input txt file and returns them
func ReadInput(day string, filename string) []string {
	content, err := ioutil.ReadFile(day + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return SplitIntoLines(string(content))
}

// ReadFirstLine reads given input txt file and returns it's first line
func ReadFirstLine(day string, filename string) string {
	return ReadInput(day, filename)[0]
}

// Check checks if given parameters are equal. If not throws an exception.
func Check(expected string, actual string) {
	if expected != actual {
		log.Fatal("Expected ", expected, ", got ", actual, "\n")
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Sort(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func SplitIntoLines(text string) []string {
	return strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
}

func Hash(data string) string {
	h := md5.Sum([]byte(data))
	return hex.EncodeToString(h[:])
}
