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
		Aliases: []string{"convert"},
		Usage: "does file conversion",
		Action: func(c *cli.Context) error {
			
			office, err := libreofficekit.NewOffice(path)
			
			document, err := office.LoadDocument("abc.html")
			if err != nil {
				return errors.New(DOC_ERR + err.Error())
			}
			err = document.SaveAs("yolo.rtf", fileType, format)
			if err != nil {
				return errors.New(SAVE_ERR + err.Error())
			}
			document.Close()
			office.Close()
			return err
		},
	}
}