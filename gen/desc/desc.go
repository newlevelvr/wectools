package desc

import (
	"fmt"
	"gen/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func printColoredTool(tool, color string, parser *cases.Caser) {
	fmt.Printf("## %s %s\nThis item does nothing\n", config.GetHumanColor(color, parser), config.GetHumanTool(tool, parser))
}

func printTool(tool string, parser *cases.Caser) {
	fmt.Printf("\n# More %ss!!!\n", strings.Replace(tool, "_", " ", -1))
	for _, color := range config.Colors {
		printColoredTool(tool, color, parser)
	}
}

func Generate() {
	fmt.Println("**This mod adds these UNIQUE items:**")
	parser := cases.Title(language.Und)
	for _, tool := range config.Tools {
		printTool(tool, &parser)
	}
	fmt.Println("\n# A Wand!!!\nThis item does nothing")
}
