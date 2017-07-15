package main

import "fmt"

func main()  {
  var i, j int = 1, 2
  /* Fonksiyon içinde 'var' yerine
                      ':=' kısa atama deyimi ile değişken tanımlanabilir

  !!! Fonksiyon alanı dışında her yapı bir anahtar kelime ile başlar.
      Fonksiyon alanı dışında ':= deyimi kullanılamaz. !!!
  */
  k := 3
  c, python, java := true, false, "no!"

  fmt.Println(i, j, k, c, python, java)
}
