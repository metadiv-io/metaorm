package query

import "strings"

func SafeField(field string) string {
	var prefix string
	if strings.HasPrefix(field, "*") {
		prefix = "*"
		field = strings.TrimPrefix(field, "*")
	}
	field = strings.ReplaceAll(field, "`", "")
	field = strings.ReplaceAll(field, ".", "`.`")
	return prefix + "`" + field + "`"
}
