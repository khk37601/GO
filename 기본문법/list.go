package main

// c 언어에서 stdio.h 와 같다고 생각하면 됩니다.
import "fmt"

func main(){
    //배열

    // 배열 선언 배열의 크기는 10으로 다 0으로 채워져 있게 됩니다.
    var array [10]int
    fmt.Println(array)
    array[0] = 10
    //[10 0 0 0 0 0 0 0 0 0]
    fmt.Println(array)

    array2 := [5]int{1,2,3,4,5}
    fmt.Println(array2)

    for i:=0; i<len(array2); i++{
        fmt.Println(array2[i])
    }

    for index, value := range array2{
        fmt.Println(index, value)
    }


}