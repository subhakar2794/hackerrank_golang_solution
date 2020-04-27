package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the birthday function below.
func birthday(s []int, d int, m int) int {
    var sum,counter int
    for i:=0;i<=(len(s)-m);i++ {
        for j:=0;j<m;j++ {
            sum+=s[i+j]
        }
        if sum == d{
            counter++
        }
        sum=0
    }
    return counter

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int(nTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int

    for i := 0; i < int(n); i++ {
        sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        sItem := int(sItemTemp)
        s = append(s, sItem)
    }

    dm := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    dTemp, err := strconv.ParseInt(dm[0], 10, 64)
    checkError(err)
    d := int(dTemp)

    mTemp, err := strconv.ParseInt(dm[1], 10, 64)
    checkError(err)
    m := int(mTemp)

    result := birthday(s, d, m)

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
