package main

import (
	"fmt"
	"os"
	// /"io"
	"bufio"
	"path/filepath"
)

func main() {
	// data, _ := os.ReadFile("testfile")
	// fmt.Println(string(data))

	//b := make([]byte, 5)
	f, _ := os.OpenFile("testfile", os.O_APPEND|os.O_RDWR, 777)
	//f.Read(b)
	//fmt.Println(b)
	// _, err := f.WriteString("dqdqdwq11111111")
	// fmt.Println(err)
	// f.Sync()

	//d1, _ := io.ReadAll(f)

	//fmt.Println(string(d1))


	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	writer := bufio.NewWriter(f)

	writer.WriteString("ddwqdwqdqdwqdwqdqwf2222222")

	writer.Flush()


	fmt.Println(filepath.Base("/dqd/dqwdq/fwqfq/filename"))
	fmt.Println(filepath.Dir("/dqd/dqwdq/fwqfq/filename"))
	fmt.Println(filepath.Join("dqd", "dqwdq",  "fwqfq/filename"))


	

}