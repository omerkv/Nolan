package main

import "flag"

func main() {
	declare_flags()
	println("Welcome to Nolan!\nA golang CLI for downloading/searching films")
	println("use nolan -h for help :)")
}

func declare_flags() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCmd.String("title", "", "Film Title")
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchCmd.String("title", "", "Film Title")
	flag.Parse()
}
