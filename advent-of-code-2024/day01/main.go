package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	file, ferr := os.Open("input-list")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	var slice1 []string
	var slice2 []string

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")
		slice1 = append(slice1, items[0])
		slice2 = append(slice2, items[1])

		// fmt.Printf("First list:%s Second List:%s\n", items[0], items[1])
	}

	// fmt.Println(reflect.TypeOf(slice))

	slice3 := []int{}
	slice4 := []int{}

	for i := 0; i < len(slice1); i++ {
		num, err1 := strconv.Atoi(slice1[i])
		num2, err2 := strconv.Atoi(slice2[i])
		if err1 != nil || err2 != nil {
			// fmt.Println("Error:", err1)
		} else {
			// fmt.Println("Number:", num) // Output: Number: 42
		}

		slice3 = append(slice3, num)
		slice4 = append(slice4, num2)

		fmt.Println(slice3[i] + slice4[i])
	}

	sort.Ints(slice3)
	sort.Ints(slice4)

	slice5 := []int{}
	var mynum int

	for i := 0; i < len(slice3); i++ {
		slice5 = append(slice5, (absInt(slice3[i] - slice4[i])))
		fmt.Print(i)
		fmt.Println(" Line number")
		mynum = mynum + slice5[i]
	}
	fmt.Println(mynum)

}
