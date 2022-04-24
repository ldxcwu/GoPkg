package regexp_test

import (
	"fmt"
	"regexp"
	"testing"
)

func ExampleReplaceAll() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb"), []byte("T")))
	// Output:
	// -T-T
}

func TestMatchString(t *testing.T) {
	re := regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	if re.MatchString(" adam[23]") ||
		re.MatchString("eve[7] ") ||
		re.MatchString("Job[30]") ||
		re.MatchString("snakey") {
		t.Error("Need false")
	}
}

func TestMathch(t *testing.T) {
	matched, err := regexp.Match(`foo`, []byte("seafood"))
	fmt.Println(matched, err)
}

func TestXx(t *testing.T) {
}
