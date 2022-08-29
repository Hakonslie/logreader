package parser

import (
	"strings"
)

func ParseLine(line string) []string {
	lines := strings.Split(line, " ")
	if lines[0] == "[RECV]" {
		lines = lines[2:]
		trimleft := strings.Replace(lines[0], "[", "", 1)
		trimmed := strings.Replace(trimleft, "]", "", 1)
		return append(make([]string, 0, 0), trimmed, strings.Join(lines[1:], " "))
	} else if strings.Contains(line, "NEW GAME STARTED") {
		return []string{"NewGameStarted", ""}
	}
	return []string{"", ""}
}
