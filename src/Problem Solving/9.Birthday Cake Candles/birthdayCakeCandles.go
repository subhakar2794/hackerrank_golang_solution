package main

import (
    "fmt"
    
)

func main(){
    var n,max,count int
    fmt.Scan(&n)
    a:=make([]int,n,n)
    for  i:=0;i<n;i++{
        fmt.Scan(&a[i])
        if a[i]==max{
            count++
        }else if a[i]>max{
            max=a[i]
            count=1
        }
    }
    fmt.Println(count)
}
