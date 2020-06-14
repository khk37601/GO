package main

import "fmt"
func main(){

    var array []int
    r := append(array, 1,2,3)
    fmt.Println(r)
    fmt.Println(double())
    test(1,2,3,4,5,6)
    test(r...)
    x := 10

    change(&x)
    fmt.Println(x)
}

func change(x * int){
   *x = 20
}


func test(args ...int){
    for _, value := range(args){
        fmt.Println(value)
    }
}

func double() (int, int){
    return  1, 2
}

func average(num []float64) float64{
    total := 0.0
    for _, value := range(num){
        total += value
    }
    return total/ float64(len(num))

}