// Fibonacci seriex
package main
import "fmt"
import "sort"

func main() {
	var fib3 int
	no:=10
	fib1:=0
	fib2:=1
	fmt.Print(fib1, " ", fib2, " ")
	for i:=2; i<no; i++{
		fib3=fib1+fib2
		fmt.Print(fib3, " ")
		fib1=fib2
		fib2=fib3
	}
	fmt.Println()
	fmt.Println(is_Prime((10)))
	fmt.Println(factorial(5))
	is_Palindrome("malayalam")
	arr1:=[...]int{1,42,6,3,45,23,67}
	do_sort(arr1[:])
}

func is_Prime(a int)string {
	if a<1{
		return "Not a prime"
	}else{
		for i:=2; i<a; i++{
			if a%i==0{
				return "Not a prime number"
			}else{
				return "Prime"
			}
		}
	}
	return "None"
}

func factorial(no int)int {
	var factorial1 int=1
	for i:=no; i>1; i--{
		factorial1=factorial1*i
	}
	return factorial1
}


func is_Palindrome(str1 string) {
	SI:=0
	EI:=len(str1)-1
	var flag bool=true
	for SI<EI{
		if string(str1[SI])!=string(str1[EI]){
			flag=false
			break
		}
		SI++
		EI--
	}
	if flag{
			fmt.Println("Palindrome")
		}else{
			fmt.Println("Not a palindrome")
		}
}

func do_sort(arr1[] int){
	sort.Ints(arr1)
	fmt.Println(arr1)
}
