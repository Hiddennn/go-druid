package sqlbuilder

import "strings"

// GroupBy set group by fields
func (sb *SQLBuilder) GroupBy(fields ...string) *SQLBuilder {
	if fields == nil || len(fields) == 0 {
		return sb
	}

	var buf strings.Builder
	buf.WriteString("GROUP BY ")

	for k, field := range fields {
		buf.WriteString(field)

		if k != len(fields)-1 {
			buf.WriteString(",")
		}
	}

	sb._groupBy = buf.String()
	return sb
}
