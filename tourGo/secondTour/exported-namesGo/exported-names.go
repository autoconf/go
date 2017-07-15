package main

import(
  "fmt"
  "math"
)

func main(){
/*
  !!! nesnelerde public / private gibi yazarak ayrım yok !!!
  bunun yerine public olması için baş harfin büyük olması yeterli
  yani private nesne küçük harfle
       public  nesne büyük harfle başlar.

  bu örnekte math paketine ait pi nesnesi public olduğu için küçük harfle
yazıldığında hata veriyor.
*/
//fmt.Println(math.pi)
  fmt.Println(math.Pi)
}
