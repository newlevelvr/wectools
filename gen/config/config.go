package config

import "fmt"

const ModID = "wectools"

var Colors = [...]string{
	"gray", "lightgray", "white", "candle", "brown", "copper", "orange", "yellow", "lime",
	"green", "cyan", "lightblue", "blue", "black", "purple", "violet", "pink", "red",
}

var CustomColorNames = map[string]string{
	"lightblue": "Light Blue",
	"lightgray": "Light Gray",
}

var Tools = [...]string{
	"wooden_pickaxe", "wooden_axe", "wooden_shovel", "wooden_hoe", "stick",
}

var OutputFolderBase = fmt.Sprintf("src/main/resources/assets/%s", ModID)
