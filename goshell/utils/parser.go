package utils

import "strings"

// ParseInput trims and splits user input into command and arguments
func ParseInput(input string) (string, []string) {
    input = strings.TrimSpace(input)
    parts := strings.Fields(input)
    if len(parts) == 0 {
        return "", []string{}
    }
    return parts[0], parts[1:]
}
