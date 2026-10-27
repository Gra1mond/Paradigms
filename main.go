// package main

// import (
// 	"fmt"
// 	"math"
// 	"net/http"
// )

// func main(){
// 	var a, b, c float64
// 	fmt.Println("Введите a, b, c для ax⁴ + bx² + c = 0:")
// 	if _, err := fmt.Scan(&a, &b, &c); err != nil {
// 		fmt.Println("Ошибка ввода")
// 		return
// 	}
// 	ts := FindSqrt(a,b,c)
// 	found:=false
// 	for _,t:=range ts{
// 		if t>0{
// 			x:=math.Sqrt(t)
// 			fmt.Printf("x = %g и x = %g\n",x,-x)
// 			found = true
// 		}else if t==0{
// 			fmt.Printf("x = 0\n")
// 			found=true
// 		}
// 	}
// 	if !found{
// 		fmt.Println("Действительных корней нет")
// 	}
// }


// func FindSqrt(a,b,c float64)[]float64{
// 	if a == 0{
// 		fmt.Println("а не должно быть равно 0")
// 		return nil
// 	}
// 	d:=b*b-4*a*c
// 	if d<0{
// 		fmt.Println("действительных корней нет")
// 		return nil
// 	}

// 	ts := []float64{(-b + math.Sqrt(d)) / (2 * a)}
// 	if d > 0 {
// 		ts = append(ts, (-b-math.Sqrt(d))/(2*a))
// 	}
// 	return ts
// }

// func FindHandler(w http.ResponseWriter, r *http.Request){

// }
