package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := "logfile"
	filedata1, err := readFileBufio(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filedata1)

	filedata2, err := readFileIoutil(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(filedata2)

}

func readFileBufio(path string) (string, error) {
	var fdata string

	f, err := os.Open(path)
	if err != nil {
		return fdata, err
	}
	defer f.Close()

	scn := bufio.NewScanner(f)

	for scn.Scan() {
		fdata += scn.Text() + "\n"
	}

	return fdata, nil
}

func readFileIoutil(path string) (string, error) {
	fdata, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(fdata), nil
}

func writeFileBufio(path, data string) error {
	f := os.Create()
}
