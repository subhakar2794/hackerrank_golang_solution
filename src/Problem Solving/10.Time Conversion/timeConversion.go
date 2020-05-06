package main
import "fmt"
func main(){
    var h,m,s int
    var suff string
    fmt.Scanf("%d:%d:%d%s",&h,&m,&s,&suff)
    if suff == "AM" && h == 12{
           h=0
        }
    if suff=="PM" && h != 12 {
        h+=12
    }
    fmt.Printf("%02d:%02d:%02d",h,m,s)
}
