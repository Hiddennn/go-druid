package sqlbuilder

import (
	"strings"
)

// JoinRaw join with raw sql
func (sb *SQLBuilder) JoinRaw(join string, values ...interface{}) *SQLBuilder {
	var buf strings.Builder

	buf.WriteString(sb._join)
	if buf.Len() != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(join)

	sb._join = buf.String()

	for _, value := range values {
		sb._joinParams = append(sb._joinParams, ConvertValueToSQLParameter(value))
	}

	return sb
}
