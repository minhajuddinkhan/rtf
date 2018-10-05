package main

import (
	"os"
	"github.com/urfave/cli"
	"github.com/minhajuddinkhan/rtf/commands"
)


func main() {
	path := "/usr/lib/libreoffice/program"
	libreCLIApp := cli.NewApp()
	libreCLIApp.Name = "Libre Office converter - Converts to your desired data format through LibreOffice 5.2"

	libreCLIApp.Commands = []cli.Command{
		*commands.Convert(path),
	}

	err := libreCLIApp.Run(os.Args)
	if err != nil {
		panic(err)
	}

}