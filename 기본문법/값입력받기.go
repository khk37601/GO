package main

// c 언어에서 stdio.h 와 같다고 생각하면 됩니다.
import "fmt"

func main(){
    // 값 입력받기
    // fmt.Scanf C언어와 형식이 같습니다.
    var num, num1 int
    fmt.Scanf("%d",&num)
    fmt.Println(num)

    // 새 줄로 구분해서 입력을 받습니다.
    i, j := fmt.Scan(&num, &num1)
    fmt.Println(num, num1)
    // i 입렫 받은 갯수. nil == null == None
    fmt.Println(i, j)

    // 공백 구분해서 입력을 받습니다.
    fmt.Scanln(&num, &num1)
    fmt.Println(num, num1)

}