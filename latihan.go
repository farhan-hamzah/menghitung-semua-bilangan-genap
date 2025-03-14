package main 
import "fmt"

func hitung(n int)int{
    var hasil int
    if n <= 0{
        return 0
    }
    if n%2 == 0{
        hasil += n
    }
    hasil += hitung(n-1)
    return hasil
}
func main(){
    var masukan int
    fmt.Scan(&masukan)
    fmt.Println(hitung(masukan))
}