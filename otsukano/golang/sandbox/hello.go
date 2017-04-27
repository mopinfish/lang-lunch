package main

import (
  "fmt" 
  "reflect"
)

func main() {
  username := `otsukano`
  fmt.Printf(username + "\n")

  fmt.Println(reflect.TypeOf(username))
  fmt.Printf("こんにちはせかい\n")

  const v = iota // constキーワードが出現する度に、iotaは置き直されます。ここではv == 0です。

  const ( 
      e, f, g = iota, iota, iota //e=0,f=0,g=0 iotaの同一行は同じです
  )
  fmt.Println(e, f, g)

  arr := [3]int{1, 2, 3}
  fmt.Println(arr)
  fmt.Println(reflect.TypeOf(arr)) 

  arr2 := []int{1, 2, 3}
  fmt.Println(arr2)
  fmt.Println(reflect.TypeOf(arr2)) 

  // keyを文字列で宣言します。値はintとなるディクショナリです。この方法は使用される前にmakeで初期化される必要があります。
  // var numbers map[string]int
  // もうひとつのmapの宣言方法
  numbers := make(map[string]int)
  numbers["one"] = 1  //代入
  numbers["ten"] = 10 //代入
  numbers["three"] = 3

  fmt.Println("３つ目の数字は: ", numbers["three"]) // データの取得
  // "３つ目の数字は： 3"という風に出力されます。
  // ディクショナリを初期化します。
  rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
  // mapは２つの戻り値があります。２つ目の戻り値では、もしkeyが存在しなければ、okはfalseに、存在すればokはtrueになります。
  csharpRating, ok := rating["C#"]
  if ok {
    fmt.Println("C# is in the map and its rating is ", csharpRating)
  } else {
    fmt.Println("We have no rating associated with C# in the map")
  }

  delete(rating, "C")  // keyがCの要素を削除します。
  fmt.Println(rating)

  i := 3
  switch i {
  case 1:
    fmt.Println("i is equal to 1")
  case 2, 3, 4:
    fmt.Println("i is equal to 2, 3 or 4")
  case 10:
    fmt.Println("i is equal to 10")
  default:
    fmt.Println("All I know is that i is an integer")
  }
}
