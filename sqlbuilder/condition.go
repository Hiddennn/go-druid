package sqlbuilder

import "strings"

func ConditionForRangeOr(s string, num int) string {
	if num == 0 {
		return ""
	}

	var buf strings.Builder
	buf.WriteString("(")

	for num > 0 {
		buf.WriteString(s)

		if num > 1 {
			buf.WriteString(" OR ")
		}
		num--
	}

	buf.WriteString(")")
	return buf.String()
}
