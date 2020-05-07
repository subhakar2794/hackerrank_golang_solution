package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the migratoryBirds function below.

//func migratoryBirds(arr []int) int {
//	mod := -1
//	maxcount := -1
//	for i := 0; i < len(arr); i++ {
//		count := 0
//		for j := 0; j < len(arr); j++ {
//			if arr[i] == arr[j] {
//				count++
//			}
//		}
//		if count > maxcount {
//			mod = arr[i]
//			maxcount = count
//		}
//	}
//	return mod

//}
func migratoryBirds(arr []int) int {
	var i, j, k, l, m, mod int
	for _, v := range arr {
		switch {
		case (v == 1):
			i++
		case (v == 2):
			j++
		case (v == 3):
			k++
		case (v == 4):
			l++
		case (v == 5):
			m++
		}

	}
	count := []int{i, j, k, l, m}
	max := 0
	for _, v := range count {
		if max < v {
			max = v
		}
	}
	if max == i {
		mod = 1
	} else if max == j {
		mod = 2
	} else if max == k {
		mod = 3
	} else if max == l {
		mod = 4
	} else if max == m {
		mod = 5
	}

	return mod

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
