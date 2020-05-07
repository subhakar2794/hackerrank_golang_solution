package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */
func getTotalX(a []int, b []int) int {
    // Write your code here
    var c int
    sort.Ints(a)
    sort.Ints(b)
    for i:=a[0];i<=b[len(b)-1];i++{
        factor:=true
        for _,v:=range a{
            if i % v != 0{
                factor = false
            }
        }
        for _,v:=range b{
            if v%i!=0{
                factor=false
            }
        }
        if factor == true{
            c++
        }
    }
    return c
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int(nTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    m := int(mTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int(arrItemTemp)
        arr = append(arr, arrItem)
    }

    brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var brr []int

    for i := 0; i < int(m); i++ {
        brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
        checkError(err)
        brrItem := int(brrItemTemp)
        brr = append(brr, brrItem)
    }

    total := getTotalX(arr, brr)

    fmt.Fprintf(writer, "%d\n", total)

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
