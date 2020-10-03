package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// PraseError is a struct
type PraseError struct {
	Index int
	Word  string
	Err   error
}

func (e *PraseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
}

// Parse is a function
func Parse(input string) (numbers []int, err error) {
	defer func() {

		if r := recover(); r != nil {

			var ok bool
			err, ok = r.(error)

			if !ok {
				err = fmt.Errorf("pkg %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)

	return
}

func fields2numbers(fields []string) (numbers []int) {

	if len(fields) == 0 {
		panic("no words to parse")
	}

	for idx, field := range fields {

		num, err := strconv.Atoi(field)

		if err != nil {
			panic(&PraseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}

	return
}
