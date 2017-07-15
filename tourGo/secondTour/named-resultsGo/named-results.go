package main

import "fmt"

/*
** sonuç değerleri değişken gibi adlandırılabilir böyle yapıldığında değişken
   fonksiyonun başında tanımlanmış gibi muamele görür **
** argümansız return, adlandırılmış sonuç değerlerini döndürür.
   "çıplak(naked)" return olarak bilinir **


!!! Kodun okunurluğu için sadece kısa fonksiyonlarda kullanılmalıdır !!!
*/
func split(sum int) (x, y int)  {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main(){
  fmt.Println(split(17))
}
