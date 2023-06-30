package lang

import (
	"encoding/json"
	"fmt"
	"gen/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
)

var outputFolder = config.OutputFolderBase + "/lang"

type Lang map[string]string

func (l Lang) PutTranslation(key, id, value string) {
	l[fmt.Sprintf("%s.%s.%s", key, config.ModID, id)] = value
}

func Generate() {
	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	f, err := os.Create(outputFolder + "/en_us.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	lang := make(Lang)
	titler := cases.Title(language.Und)
	for _, tool := range config.Tools {
		for _, color := range config.Colors {
			id := config.ToID(tool, color)
			humanColor := config.GetHumanColor(color, &titler)
			humanTool := config.GetHumanTool(tool, &titler)
			lang.PutTranslation("item", id, fmt.Sprintf("%s %s", humanColor, humanTool))
		}
	}
	lang.PutTranslation("item", "wand", "Wand")
	lang.PutTranslation("itemGroup", "tools", "WorldEdit Colored Tools")
	res, err := json.MarshalIndent(lang, "", "\t")
	if err != nil {
		panic(err)
	}
	f.Write(res)
}
