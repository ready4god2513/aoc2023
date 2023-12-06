package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputFile = "./input.txt"
)

var numbers = map[string]int{
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

type word struct {
	index int
	value string
}

func main() {
	lines, err := readFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := sumDigits(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func readFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	var lines []string

	for {
		str, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		lines = append(lines, strings.TrimSuffix(str, "\n"))
		if err == io.EOF {
			break
		}
	}

	return lines, nil
}

func sumDigits(lines []string) (int, error) {
	var sum int
	for _, line := range lines {
		n, err := calibrationValues(line)
		if err != nil {
			return 0, err
		}

		sum += n
	}

	return sum, nil
}

func calibrationValues(str string) (int, error) {
	replaced := replaceWords(strings.ToLower(str))

	n, err := joinDigits(firstNum(replaced), lastNum(replaced))
	if err != nil {
		return 0, err
	}

	if n < 10 {
		return 0, fmt.Errorf("calibrationValues: %d is less than 10", n)
	}

	if n > 99 {
		return 0, fmt.Errorf("calibrationValues: %d is greater than 99", n)
	}

	return n, nil
}

func joinDigits(first, last int) (int, error) {
	n, err := strconv.ParseInt(fmt.Sprintf("%d%d", first, last), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

// replaceWords replaces words with numbers.  However, there is a caveat.  For
// example, this string xtwone3four should be replaced with x2ne34.  So we need
// to find the index of all replacements and then replace left to right.
func replaceWords(str string) string {
	for {
		replaced := replaceFirstWord(str)
		if replaced == str {
			return replaced
		}

		str = replaced
	}
}

func replaceFirstWord(str string) string {
	var words []word
	for num := range numbers {
		index := strings.Index(str, num)
		if index != -1 {
			words = append(words, word{index, num})
		}
	}

	if len(words) == 0 {
		return str
	}

	// Sort words in ascending order.
	sort.Slice(words, func(i, j int) bool {
		return words[i].index < words[j].index
	})

	// Replace the first word.
	return strings.Replace(
		str,
		words[0].value,
		strconv.Itoa(numbers[words[0].value]), 1)
}

func firstNum(str string) int {
	for _, c := range str {
		if c > '0' && c <= '9' {
			char, _ := strconv.Atoi(string(c))
			return char
		}
	}

	return 0
}

func lastNum(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] > '0' && str[i] <= '9' {
			char, _ := strconv.Atoi(string(str[i]))
			return char
		}
	}

	return 0
}
