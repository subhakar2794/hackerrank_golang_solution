package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the divisibleSumPairs function below.
func solveMeFirst(a int,b int) int{
	// Hint: Type return (a+b) below
	return a+b
  }

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    n := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(n[0], 10, 64)
    checkError(err)
    a:= int(nTemp)

	k := strings.Split(readLine(reader), " ")
    kTemp, err := strconv.ParseInt(k[0], 10, 64)
    checkError(err)
    b := int(kTemp)

    

    result := solveMeFirst(a,b)

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