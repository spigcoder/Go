package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string{
	n := len(s)
	if n <= 3{
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] 
}

func intsToString(valuse []int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range valuse{
		if i > 0{
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main(){
	var f float32 = 16777216
	fmt.Println(f == f+1)
}