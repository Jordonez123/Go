package controls

import "fmt"


func PrintControls() {
	if 6%2 == 0 {
		fmt.Println("6 is even")
	}else{
		fmt.Println("6 is odd")
	}

	if 10%4 == 0 {
		fmt.Println("10 is divisible by 4")
	}

	if 6%2 == 0 || 5%2 == 0 {
		fmt.Println("either 6 or 5 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10{
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}