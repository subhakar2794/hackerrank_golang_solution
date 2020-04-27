package main
import "fmt"
func main(){
    var l,p,n,z,t int
    fmt.Scan(&l)
    for i:=0;i<l;i++{
        fmt.Scan(&t)
        if t>0{
            p++
        }else if t<0{
            n++
        }else{
            z++
        }
   
    }
     pr:=float64(p)/float64(l)
    nr:=float64(n)/float64(l)
    zr:=float64(z)/float64(l)
    fmt.Printf("%.6f\n",pr)
    fmt.Printf("%.6f\n",nr)
    fmt.Printf("%.6f",zr)
}
