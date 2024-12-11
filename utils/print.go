package utils

import (
	"fmt"
	"strings"
)

// FormatString python的format，{}表示占位符
func FormatString(template string, values ...interface{}) string {
	for _, value := range values {
		template = strings.Replace(template, "{}", fmt.Sprintf("%v", value), 1)
	}
	return template
}

// Print 输出python的format
func Print(template string, values ...interface{}) {
	result := FormatString(template, values...)
	fmt.Println(result)
}

// ColorPrint 带颜色的打印
func ColorPrint(content string, color string) {
	colorCodes := map[string]string{
		"black":         "30",
		"red":           "31",
		"green":         "32",
		"yellow":        "33",
		"blue":          "34",
		"magenta":       "35",
		"cyan":          "36",
		"white":         "37",
		"gray":          "90",
		"light_red":     "91",
		"light_green":   "92",
		"light_yellow":  "93",
		"light_blue":    "94",
		"light_magenta": "95",
		"light_cyan":    "96",
		"light_white":   "97",
	}
	colorCode, exists := colorCodes[color]
	if !exists {
		fmt.Println(content)
	} else {
		fmt.Printf("\033[%sm%s\033[0m\n", colorCode, content)
	}
}
