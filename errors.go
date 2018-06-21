package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}


func Sqrt(x float64) (float64, error) {
	if x >=0 {
		z := 1.0
		count := 1
		temp := 0.0 
		diff := 1.0
		for math.Abs(diff) > math.Pow(10,-15) {
			temp = z
			z -= (z*z - x) / (2*z)
			diff = temp - z
			count++
		}
		//fmt.Println("",count)
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/*
Question: a call to fmt.Sprint(e) inside the Error method will send the program into 
an infinite loop. It can be avoided by converting e first: fmt.Sprint(float64(e)). Why?

Answer: fmt.Sprintf(e) will call e.Error() to convert e to a string. 
In the Error() methods, it calls fmt.Sprintf() again so there is a recursion. 
How to fix? —> Change the type of e to a non e-type 
(From ErrNegativeSqrt to float64 in this problem, then it will call float64.Error() )
*/