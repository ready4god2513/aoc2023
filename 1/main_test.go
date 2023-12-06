package main

import "testing"

func Test_replaceWords(t *testing.T) {
	var cases = []struct {
		in, want string
	}{
		{"75twoeightfourbcgeight", "75284bcg8"},
		{"7vtcgkpgqqzcxdxrjmpbjone1pvxcjjtpdtcprkq5mqxkjpbqkd", "7vtcgkpgqqzcxdxrjmpbj11pvxcjjtpdtcprkq5mqxkjpbqkd"},
		{"one", "1"},
		{"skdjkls1813", "skdjkls1813"},
		{"xtwone3four", "x2ne34"},
		{"6oneeightnine6", "61896"},
		{"oneeight3", "183"},
		{"twoonetwone", "212ne"},
		{"4nineeightseven2", "49872"},
	}

	for _, c := range cases {
		got := replaceWords(c.in)
		if got != c.want {
			t.Errorf("replaceWords(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func Test_firstNum(t *testing.T) {
	var cases = []struct {
		in   string
		want int
	}{
		{"79ewwew9wet", 7},
		{"07vtcgkpgqqzcxdxrjmpbjone1pvxcjjtpdtcprkq5mqxkjpbqkd", 7},
		{"1", 1},
		{"skdjkls1813", 1},
		{"qone3nineqfqdcfc", 3},
		{"1", 1},
		{"12", 1},
		{"sdjgdskjgdss", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
	}

	for _, c := range cases {
		got := firstNum(c.in)
		if got != c.want {
			t.Errorf("firstNum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test_lastNum(t *testing.T) {
	var cases = []struct {
		in   string
		want int
	}{
		{"75twoeightfourbcgeight7", 7},
		{"7vtcgkpgqqzcxdxrjmpbjone1pvxcjjtpdtcprkq5mqxkjpbqkd", 5},
		{"one2", 2},
		{"skdjkls1813", 3},
		{"qone3nineqfqdcfc", 3},
		{"1", 1},
		{"12", 2},
		{"121221", 1},
		{"1222342", 2},
		{"sjdhsgskdhs", 0},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
	}

	for _, c := range cases {
		got := lastNum(c.in)
		if got != c.want {
			t.Errorf("lastNum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test_joinDigits(t *testing.T) {
	var cases = []struct {
		first, last, want int
	}{
		{1, 2, 12},
		{7, 5, 75},
		{0, 0, 0},
		{1, 0, 10},
		{0, 1, 1},
	}

	for _, c := range cases {
		got, err := joinDigits(c.first, c.last)
		if err != nil {
			t.Errorf("joinDigits(%d, %d) == %d, want %d", c.first, c.last, got, c.want)
		}
	}
}

func Test_integration(t *testing.T) {
	var cases = []struct {
		in   string
		want int
	}{
		{"75twoeightfourbcgeight7", 77},
		{"twoeightfourbcgeight", 28},
		{"7vtcgkpgqqzcxdxrjmpbjone1pvxcjjtpdtcprkq5mqxkjpbqkd", 75},
		{"one2", 12},
		{"skdjkls1813", 13},
		{"qone3nineqfqdcfc", 19},
		{"1", 11},
		{"12", 12},
		{"121221", 11},
		{"1222342", 12},
		{"6oneeightnine6", 66},
		{"oneeight3", 13},
		{"seventhree8one293", 73},
		{"ninenunnine", 99},
		{"one", 11},
		{"two", 22},
		{"three", 33},
		{"four", 44},
		{"five", 55},
		{"six", 66},
		{"seven", 77},
		{"eight", 88},
		{"nine", 99},
		{"37qdmsqzsq72clfntfxqfrhbxtmfourzcjxfmmfz", 34},
	}

	for _, c := range cases {
		got, err := calibrationValues(c.in)
		if err != nil {
			t.Errorf("calibrationValues(%q) == %d, want %d", c.in, got, c.want)
		}

		if got != c.want {
			t.Errorf("calibrationValues(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test_sumDigits(t *testing.T) {
	var cases = []struct {
		in   []string
		want int
	}{
		{
			[]string{
				"75twoeightfourbcgeight7",                             // 77
				"7vtcgkpgqqzcxdxrjmpbjone1pvxcjjtpdtcprkq5mqxkjpbqkd", // 75
			},
			152,
		},
		{
			[]string{
				"two1nine",         // 29
				"eightwothree",     // 83
				"abcone2threexyz",  // 13
				"xtwone3four",      // 24
				"4nineeightseven2", // 42
				"zoneight234",      // 84
				"7pqrstsixteen",    // 76
			},
			281,
		},
	}

	for _, c := range cases {
		got, err := sumDigits(c.in)
		if err != nil {
			t.Errorf("sumDigits(%q) == %d, want %d", c.in, got, c.want)
		}

		if got != c.want {
			t.Errorf("sumDigits(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func Test_readFile(t *testing.T) {
	lines, err := readFile(inputFile)
	if err != nil {
		t.Errorf("readFile() == %v, want nil", err)
	}

	if len(lines) == 0 {
		t.Errorf("readFile() == %d, want > 0", len(lines))
	}

	if lines[0] != "kjrqmzv9mmtxhgvsevenhvq7" {
		t.Errorf("readFile() == %q, want %q", lines[0], "kjrqmzv9mmtxhgvsevenhvq7s")
	}

	if lines[len(lines)-1] != "8ninejseven5" {
		t.Errorf("readFile() == %q, want %q", lines[len(lines)-1], "8ninejseven5")
	}
}
