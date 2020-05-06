package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []float64

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := float64(arItemTemp)
		ar = append(ar, arItem)

	}
	m := mean(ar)
	md := median(ar)
	mod := mode(ar)
	fmt.Fprintf(writer, "%f\n", m)
	fmt.Fprintf(writer, "%f\n", md)
	fmt.Fprintf(writer, "%f\n", mod)
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

func mean(a []float64) float64 {
	var sum float64
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum / float64(len(a)))
	return (sum / float64(len(a)))
}
func median(a []float64) float64 {
	sort.Float64s(a)
	var md float64
	if len(a)%2 == 0 {

		md = (a[int(len(a)/2)] + a[(len(a)-1)/2]) / 2

	} else {
		md = a[len(a)/2]
	}
	fmt.Println(md)
	return md
}

func mode(a []float64) float64 {
	mod := -1.0
	maxcount := -1
	for i := 0; i < len(a); i++ {
		count := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count > maxcount {
			mod = a[i]
			maxcount = count
		}
	}
	return mod

}
