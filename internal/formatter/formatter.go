package formatter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v3"
)

type Formatter interface {
	Format(data interface{}) error
}

func New(format string) Formatter {
	switch strings.ToLower(format) {
	case "yaml", "yml":
		return &YAMLFormatter{}
	case "table":
		return &TableFormatter{}
	case "json":
		fallthrough
	default:
		return &JSONFormatter{}
	}
}

type JSONFormatter struct{}

func (f *JSONFormatter) Format(data interface{}) error {
	// If data is already parsed, use it directly
	var v interface{}
	switch d := data.(type) {
	case []byte:
		if err := json.Unmarshal(d, &v); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}
	case string:
		if err := json.Unmarshal([]byte(d), &v); err != nil {
			// If it's not JSON, just print the string
			fmt.Println(d)
			return nil
		}
	default:
		v = data
	}

	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	if isTerminal() {
		fmt.Println(colorizeJSON(string(pretty)))
	} else {
		fmt.Println(string(pretty))
	}

	return nil
}

type YAMLFormatter struct{}

func (f *YAMLFormatter) Format(data interface{}) error {
	// If data is already parsed, use it directly
	var v interface{}
	switch d := data.(type) {
	case []byte:
		if err := json.Unmarshal(d, &v); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}
	case string:
		if err := json.Unmarshal([]byte(d), &v); err != nil {
			// If it's not JSON, just print the string
			fmt.Println(d)
			return nil
		}
	default:
		v = data
	}

	yamlData, err := yaml.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to convert to YAML: %w", err)
	}

	fmt.Println(string(yamlData))
	return nil
}

type TableFormatter struct{}

func (f *TableFormatter) Format(data interface{}) error {
	// If data is already parsed, use it directly
	var v interface{}
	switch d := data.(type) {
	case []byte:
		if err := json.Unmarshal(d, &v); err != nil {
			return fmt.Errorf("failed to parse JSON: %w", err)
		}
	case string:
		if err := json.Unmarshal([]byte(d), &v); err != nil {
			// If it's not JSON, just print the string
			fmt.Println(d)
			return nil
		}
	default:
		v = data
	}

	switch val := v.(type) {
	case []interface{}:
		return f.formatArray(val)
	case map[string]interface{}:
		return f.formatObject(val)
	default:
		fmt.Println(v)
		return nil
	}
}

func (f *TableFormatter) formatArray(arr []interface{}) error {
	if len(arr) == 0 {
		fmt.Println("No data")
		return nil
	}

	first, ok := arr[0].(map[string]interface{})
	if !ok {
		return fmt.Errorf("cannot format as table: not an array of objects")
	}

	headers := make([]string, 0, len(first))
	for k := range first {
		headers = append(headers, k)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, item := range arr {
		obj, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = formatValue(obj[h])
		}
		table.Append(row)
	}

	table.Render()
	return nil
}

func (f *TableFormatter) formatObject(obj map[string]interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Field", "Value"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for k, v := range obj {
		table.Append([]string{k, formatValue(v)})
	}

	table.Render()
	return nil
}

func formatValue(v interface{}) string {
	switch val := v.(type) {
	case nil:
		return ""
	case string:
		return val
	case float64:
		if val == float64(int(val)) {
			return fmt.Sprintf("%d", int(val))
		}
		return fmt.Sprintf("%f", val)
	case bool:
		return fmt.Sprintf("%t", val)
	default:
		data, _ := json.Marshal(val)
		return string(data)
	}
}

func isTerminal() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func colorizeJSON(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "\"") && strings.Contains(trimmed, ":") {
			parts := strings.SplitN(trimmed, ":", 2)
			if len(parts) == 2 {
				key := color.CyanString(parts[0])
				lines[i] = strings.Replace(line, parts[0], key, 1)
			}
		}

		if strings.Contains(line, ": \"") {
			start := strings.Index(line, ": \"")
			end := strings.LastIndex(line, "\"")
			if start != -1 && end > start+2 {
				value := line[start+2 : end+1]
				coloredValue := color.GreenString(value)
				lines[i] = line[:start+1] + " " + coloredValue + line[end+1:]
			}
		}

		if strings.Contains(line, ": true") || strings.Contains(line, ": false") {
			lines[i] = strings.Replace(line, "true", color.YellowString("true"), -1)
			lines[i] = strings.Replace(lines[i], "false", color.YellowString("false"), -1)
		}

		if strings.Contains(line, ": null") {
			lines[i] = strings.Replace(line, "null", color.MagentaString("null"), -1)
		}
	}

	return strings.Join(lines, "\n")
}
