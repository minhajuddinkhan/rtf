package commands

import (
	"github.com/urfave/cli"
	"github.com/dveselov/go-libreofficekit"
	"errors"
)


var (
	DOC_ERR = "cannot load document to libre office"
	SAVE_ERR = "cannot save as document to libre office"
)



func Convert(path string) *cli.Command{
	format := "Rich Text Document"
	fileType := "rtf"
	return &cli.Command{
		Name: "convert",
		Aliases: []string{"convert path/input.html path/output.html"},
		Usage: "does file conversion",
		Action: func(c *cli.Context) error {
			

			args := c.Args()
			pathToInput := args.Get(0)
			pathToOutput := args.Get(1)
			if len(pathToInput) == 0 {
				return errors.New("Invalid Args [Input missing] ")
			} 
			if len(pathToOutput) == 0 {
				return errors.New("Invalid Args [Output missing] ")
			}

			office, err := libreofficekit.NewOffice(path)			
			document, err := office.LoadDocument(pathToInput)
			if err != nil {
				return errors.New(DOC_ERR + err.Error())
			}
			err = document.SaveAs(pathToOutput, fileType, format)
			if err != nil {
				return errors.New(SAVE_ERR + err.Error())
			}
			document.Close()
			office.Close()
			return err
		},
	}
}