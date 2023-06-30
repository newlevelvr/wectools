package config

import (
	"golang.org/x/text/cases"
	"strings"
)

func GetHumanColor(color string, parser *cases.Caser) string {
	custom, exists := CustomColorNames[color]
	if exists {
		return custom
	}
	return parser.String(color)
}

func GetHumanTool(tool string, parser *cases.Caser) string {
	return parser.String(strings.Replace(tool, "_", " ", -1))
}
