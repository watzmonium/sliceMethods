package sliceMethods

import (
	"fmt"
	"strings"
	"testing"
)

func TestFilterWithInt(t *testing.T) {
	intSlice := []int{1, 2, 3, 4, 5, 6}
	isEven := func(i int) bool {
		if i&1 == 1 {
			return false
		}
		return true
	}
	newSlice := Filter(&intSlice, isEven)

	if len(newSlice) != 3 {
		t.Fatalf("did not correctly filter ints")
	}
}

func TestFilterWithFloat(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	lessFour := func(f float64) bool {
		if f > 4 {
			return false
		}
		return true
	}
	newSlice := Filter(&floatSlice, lessFour)

	if len(newSlice) != 3 {
		t.Fatalf("did not correctly filter floats")
	}
}

func TestFilterWithString(t *testing.T) {
	strSlice := []string{"abc", "def", "ghi", "jkl", "m"}
	isVowel := func(s string) bool {
		for _, r := range s {
			if strings.ContainsRune("aeiou", r) {
				return false
			}
		}
		return true
	}
	newSlice := Filter(&strSlice, isVowel)

	if len(newSlice) != 2 {
		t.Fatalf("did not correctly filter strings")
	}
}

func TestMapWithIntReturnInt(t *testing.T) {
	intSlice := []int{1, 2, 3}
	timesTwo := func(i int) any {
		return i * 2
	}

	newSlice := Map(&intSlice, timesTwo)

	if newSlice[0] != 2 || newSlice[1] != 4 || newSlice[2] != 6 {
		t.Fatalf("did not map ints to ints")
	}
}

func TestMapWithIntReturnFloat(t *testing.T) {
	intSlice := []int{1, 2, 3}
	toFloat := func(i int) any {
		return float64(i)
	}

	newSlice := Map(&intSlice, toFloat)
	var first interface{} = newSlice[0]
	_, isFloat := first.(float64)
	if !isFloat {
		t.Fatalf("did not map ints to floats")
	}
}

func TestMapWithIntReturnString(t *testing.T) {
	intSlice := []int{1, 2, 3}
	toStr := func(i int) any {
		return fmt.Sprint(i)
	}

	newSlice := Map(&intSlice, toStr)
	var first interface{} = newSlice[0]
	_, isStr := first.(string)
	if !isStr {
		t.Fatalf("did not map ints to floats")
	}
}

func TestMapWithFloatReturnFloat(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3}
	timesTwo := func(f float64) any {
		return 2 * f
	}

	newSlice := Map(&floatSlice, timesTwo)

	if newSlice[0] != 2.2 || newSlice[1] != 4.4 || newSlice[2] != 6.6 {
		t.Fatalf("did not map floats to floats")
	}
}

func TestMapWithFloatReturnInt(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3}
	toInt := func(f float64) any {
		return int(f)
	}

	newSlice := Map(&floatSlice, toInt)

	if newSlice[0] != 1 || newSlice[1] != 2 || newSlice[2] != 3 {
		t.Fatalf("did not map floats to ints")
	}
}

func TestMapWithFloatReturnString(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3}
	toStr := func(f float64) any {
		return fmt.Sprint(f)
	}

	newSlice := Map(&floatSlice, toStr)

	if newSlice[0] != "1.1" {
		t.Fatalf("did not map floats to strings")
	}
}

func TestMapWithStringReturnString(t *testing.T) {
	strSlice := []string{"abc", "d", "efg"}
	toStr := func(s string) any {
		return s + "a"
	}

	newSlice := Map(&strSlice, toStr)

	if newSlice[0] != "abca" || newSlice[1] != "da" || newSlice[2] != "efga" {
		t.Fatalf("did not map strings to strings")
	}
}

func TestReduceWithInt(t *testing.T) {
	intSlice := []int{1, 2, 3}
	add := func(i int, acc int) int {
		return i + acc
	}

	sum := Reduce(&intSlice, add, 0)

	if sum != 6 {
		t.Fatalf("reduce did not add all int numbers")
	}
}

func TestReduceWithFloat(t *testing.T) {
	floatSlice := []float64{1.1, 2.2, 3.3}
	add := func(i float64, acc float64) float64 {
		return i + acc
	}

	sum := Reduce(&floatSlice, add, 0)

	if sum != 6.6 {
		t.Fatalf("reduce did not add all float numbers")
	}
}

func TestReduceWithString(t *testing.T) {
	strSlice := []string{"a", "bc", "def"}
	add := func(i string, acc string) string {
		return acc + i
	}

	sum := Reduce(&strSlice, add, "")
	if sum != "abcdef" {
		t.Fatalf("reduce did not add all strings")
	}
}