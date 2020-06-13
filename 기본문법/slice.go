package main
/*
 go언어 슬라이스

*/
import "fmt"

func main(){
  array := []int{1,2,3,4}
  var array2 = []int{1,3,4,5,6}
  fmt.Println(array2)

  array = append(array, 5, 6, 7, 8)
  fmt.Println(array)

  a := []float64{10.0, 34.6, 434.5}
  b := []float64{1.0,2.4,6.0,3.1,7.8,5.3}
  // 배열의 타입이 서로 같아야 합니다.
  a = append(a, b...)
  fmt.Println(a)
  // 슬라이스는 레퍼런스 복사 입니다.

  fmt.Println(len(array), cap(array))
  array = append(array, 0,3)
  fmt.Println(len(array), cap(array))

  // make를 통해서 슬라이스를 생성 할 수 있습니다.
  //make([]타입, 길이, cap)
  test := make([]int64, 5, 10)
  fmt.Println(len(test), cap(test))
  test = append(test,1,2,3,4)
  fmt.Println(test[4:])
  fmt.Println(a[:4])

}