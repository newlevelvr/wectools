package models

import (
	"fmt"
	"gen/config"
	"os"
)

const template = `{
  "parent": "minecraft:item/handheld",
  "textures": {
    "layer0": "%s:item/%s"
  }
}
`

var outputFolder = config.OutputFolderBase + "/models/item"

func createModelFile(id string) error {
	f, err := os.Create(fmt.Sprintf("%s/%s.json", outputFolder, id))
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(f, template, config.ModID, id)
	if err != nil {
		return err
	}
	return nil
}

func Generate() {
	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for _, tool := range config.Tools {
		for _, color := range config.Colors {
			err := createModelFile(config.ToID(tool, color))
			if err != nil {
				panic(err)
			}
		}
	}
	err = createModelFile("wand")
	if err != nil {
		panic(err)
	}
}
