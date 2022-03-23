package notes

import (
	"fmt"
	"strconv"
	"strings"
)

func StringsNotes() {
	//前缀和后缀,包含关系,索引,替换
	str := "Hello World , majikun !"
	s1 := "e"
	s2 := "E"
	s3 := []string{"mjk"}
	// 前缀
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?    %t\n", str, "He", strings.HasPrefix(str, "He"))
	// 后缀
	fmt.Printf("T/F? Does the string \"%s\" have Suffix %s?    %t\n", str, "He", strings.HasSuffix(str, "He"))
	// 包含关系
	fmt.Printf("T/F? Does the string \"%s\" contains    %s?    %t\n", str, "He", strings.Contains(str, "He"))
	// 索引
	fmt.Printf("The first index of %s in the string \"%s\" is %d\n", str, "l", strings.Index(str, "l"))
	fmt.Printf("The last  index of %s in the string \"%s\" is %d\n", str, "l", strings.Index(str, "l"))
	// 替换
	fmt.Printf("Replace %s with %s,the result is %s\n", s1, s2, strings.Replace(str, s1, s2, -1))
	// 计数
	fmt.Printf("Number of H's in %s is: %d,", str, strings.Count(str, "H"))
	// 重复
	fmt.Println(strings.Repeat(str, 2))
	// 修改大小写
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))
	// 修剪
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Trim(str, "l"))
	// 分割
	fmt.Println(strings.Fields(str))
	fmt.Println("Split result: ", strings.Split(str, ":")[0])
	// 拼接
	fmt.Printf("The result of string splicing is %v\n", strings.Join(s3, ":"))
	// 读取
	reader, _ := strings.NewReader(str).ReadByte()
	fmt.Println(reader)
	// 字符串和其他类型转换
	fmt.Println(strconv.Itoa(1))
	intStr, _ := strconv.Atoi(str)
	fmt.Println(intStr)
}
