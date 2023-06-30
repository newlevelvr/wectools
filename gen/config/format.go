package config

import "fmt"

const ToolIDFormat = "%s_%s"

func ToID(tool, color string) string {
	return fmt.Sprintf(ToolIDFormat, tool, color)
}
