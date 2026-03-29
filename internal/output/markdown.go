package output

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/glamour"
)

// PrintMarkdown extracts a field from result and renders it as terminal markdown.
// If field is ".", treats the entire result as markdown when it's a string.
// Returns false if the field wasn't a string (caller should fall back to Print).
func PrintMarkdown(result interface{}, field string) bool {
	var md string
	if field == "." {
		s, ok := result.(string)
		if !ok {
			return false
		}
		md = s
	} else {
		b, err := json.Marshal(result)
		if err != nil {
			return false
		}
		var m map[string]interface{}
		if err := json.Unmarshal(b, &m); err != nil {
			return false
		}
		s, ok := m[field].(string)
		if !ok {
			return false
		}
		md = s
	}
	rendered, err := glamour.Render(md, "dark")
	if err != nil {
		return false
	}
	fmt.Print(rendered)
	return true
}
