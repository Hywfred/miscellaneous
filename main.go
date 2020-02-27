package main

import (
	`fmt`

	`github.com/Hywfred/miscellaneous/util`
)

func main() {
	dir, err := util.FindDir("hello")
	util.CheckErr(err)
	fmt.Println(dir)
}
