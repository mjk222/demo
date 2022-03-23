package notes

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

func RegexNotes() {
	str := "abcd"
	regex := regexp2.MustCompile(`a(?=b)`, 0)
	matches, err := regex.MatchString(str)
	if err != nil {
		fmt.Printf("err is %v", err)
	}
	fmt.Println(matches)
}
