package notes

import (
	"fmt"
)

var (
	season int
)

func ControlStrucNotes() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		//fallthrough
	default:
		fmt.Println("default case")
	}
	//Season()
	outputTest()
}

// 练习：写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称（不用考虑月份的日期）
func Season() {
	fmt.Println("Please enter season :")
	fmt.Scanln(&season)
	switch {
	case season >= 3 && season < 6:
		fmt.Printf("%d is spring", season)
	case season >= 6 && season < 9:
		fmt.Printf("%d is summer", season)
	case season >= 9 && season < 12:
		fmt.Printf("%d is autumn", season)
	case (season >= 1 && season < 3) || season == 12:
		fmt.Printf("%d is winter", season)
	default:
		fmt.Println("season input is invalid")
	}
}

// 练习：以下函数的输出结果
func outputTest() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d \n", v)
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }
	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

LABEL:
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		for j := 0; j < 10; j++ {
			if j == 2 {
				break LABEL
			}
			fmt.Println(j)
		}
		if i == 3 {
			break
		}
	}

	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}

	count := 15
LOOP:
	if count > 0 {
		fmt.Printf("count is %d\n", count)
		count--
		goto LOOP
	}

	// for s := "G"; ; {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(s)
	// 	if len(s) < 15 {
	// 		s += "G"
	// 	}
	// }

	for i := 1; i < 101; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
