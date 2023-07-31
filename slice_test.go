package exercise

import (
	"fmt"
	"testing"
)

var less256IntTests = []struct {
	s1   []int
	i    int
	want bool
}{
	{
		[]int{1, 2, 3, 4, 5},
		2,
		true,
	},
	{
		[]int{1, 2, 3, 4, 5},
		-1,
		false,
	},
}

var less256StringTests = []struct {
	s1   []string
	i    int
	want bool
}{
	{
		[]string{"a", "b", "c", "d", "e"},
		1,
		true,
	},
	{
		[]string{"a", "b", "c", "d", "e"},
		6,
		false,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range less256IntTests {
		if _, err := DeleteElement(test.s1, test.i); (err != nil) == test.want {
			t.Errorf("DeleteElement(%v, %v), got %t", test.s1, test.i, test.want)
		}
	}
	for _, test := range less256StringTests {
		if _, err := DeleteElement(test.s1, test.i); (err != nil) == test.want {
			t.Errorf("DeleteElement(%v, %v), got %t", test.s1, test.i, test.want)
		}
	}

}

func TestLenMore256(t *testing.T) {
	myArray := [...]int{256: 256}
	mySlice := myArray[:]
	fmt.Println("before append cap: ", cap(mySlice))
	for i := 0; i < 258; i++ {
		mySlice = append(mySlice, i)
	}
	fmt.Println("after append, cap: ", cap(mySlice))
	var newSlice []int = mySlice
	var err error
	for i := 0; i < 256; i++ {
		newSlice, err = DeleteElement(newSlice, i)
	}
	if err == nil {
		fmt.Println("after: ", newSlice)
	} else {
		t.Errorf("DeleteElement(%v), resize err occur, got %v", myArray, err)
	}
}
