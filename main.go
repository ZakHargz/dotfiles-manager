package main

import "github.com/ZakHargz/dotfiles_manager/cli"

func main() {
	app := cli.NewDotfilesManager()
	app.Run()
}
