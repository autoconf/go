package main

import "fmt"

// ** değişken tanımlanırken ilklendirme yapılabilir **
var i, j int = 1, 2

func main(){
/* ** ilklendirme yapıldığında değişken türü belirtilmeyebilir.
      bu durumda değişken türü ilklendirme değeri ile aynı türde olur.
*/
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}
