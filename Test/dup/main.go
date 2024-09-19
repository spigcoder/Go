package main

import (
	"bufio"
	"fmt"

	// "io/ioutil"
	"os"
	// "strings"
)

// dup2 一行一行的从file中进行读取
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Println(f.Name())
			break
		}
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//read from stdin
		countLines(os.Stdin, counts)
	} else {
		//read from files
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// for line, n := range counts{
	// 	if n > 1{
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
}

//dup3 现将file的内容全部读取过来再进行分片操作
// func main(){
// 	counts := make(map[string]int)

// 	for _, filename := range os.Args[1:]{
// 		data, err := ioutil.ReadFile(filename)
// 		if err != nil{
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 		}
// 		for _, line := range strings.Split(string(data), "\n"){
// 			counts[line]++
// 		}
// 	}

// 	for line, n := range counts{
// 			fmt.Printf("%d\t%s\n", n, line)
// 	}
// }

