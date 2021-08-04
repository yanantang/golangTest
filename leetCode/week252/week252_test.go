package week252

import (
	"fmt"
	"testing"
)

func Test_isThree(t *testing.T) {
	s := isThree(81)
	fmt.Printf("%+v = false\n ", s)
	s = isThree(2)
	fmt.Printf("%+v = false\n ", s)
	s = isThree(4)
	fmt.Printf("%+v = true\n ", s)
}

func Test_numberOfWeeks(t *testing.T) {
	s:= minimumPerimeter(13)
	fmt.Printf("%+v = 16\n ", s)
}

func Test_numberOfRounds(t *testing.T) {
	s := numberOfRounds("00:00", "23:59")
	fmt.Printf("%+v = 95\n ", s)

	s = numberOfRounds("22:46", "23:01")
	fmt.Printf("%+v = 0\n ", s)

	s = numberOfRounds("00:01", "00:01")
	fmt.Printf("%+v = 95\n ", s)
}