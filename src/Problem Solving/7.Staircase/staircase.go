package main

import "fmt"

func main() {
    n:=0
    fmt.Scan(&n)
    for i:=0;i<n;i++{
        for j:=0;j<n;j++{
             if i+j<=(n-2){
                    fmt.Print(" ")
            }else{
                fmt.Print("#")
                 if j==n-1{
                      fmt.Println()
                 }
            }
            
        }
    }
}
