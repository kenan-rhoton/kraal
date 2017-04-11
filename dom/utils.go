package dom

import "strings"

func indent(in string) string {
	lines := strings.Split(in, "\n")
	for i, line := range lines {
		if line != "" {
			lines[i] = "  " + line
		}
	}
	return strings.Join(lines, "\n")
}
