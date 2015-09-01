package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		panic("URLを第1引数に指定してください!")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	err2 := ioutil.WriteFile("./file.txt", []byte(body), 0666)
	if err2 != nil {
		fmt.Println(err)
	}
}
