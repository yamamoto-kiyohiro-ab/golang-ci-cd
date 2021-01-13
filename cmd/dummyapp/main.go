package main

import "fmt"

func main() {
	fmt.Println(getHoge() + getFuga() + getPuyo())
}

func getHoge() string {
	return "hoge"
}

func getFuga() string {
	return "fuga"
}

func getPuyo() string {
	return "puyo"
}

func getPiyo() string {
	return "puyo"
}
