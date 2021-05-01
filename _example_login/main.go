package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/nu50218/go-nagoyau"
	"github.com/nu50218/nucredit"
	"golang.org/x/crypto/ssh/terminal"
)

func getInput(key string) (string, error) {
	fmt.Printf("%s: ", key)
	b, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(b), err
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	username, err := getInput("username")
	checkError(err)

	password, err := getInput("password")
	checkError(err)

	client, err := nagoyau.NewClient(username, password, nagoyau.Portal)
	checkError(err)

	resp, err := client.Get("https://app3.nagoya-u.ac.jp/rishu/gkca0470")
	checkError(err)

	f, _ := os.Create("a.html")
	io.Copy(f, resp.Body)

	subjects, err := nucredit.FromResponse(resp)
	checkError(err)

	for _, s := range subjects {
		fmt.Println(s.Name)
	}
}
