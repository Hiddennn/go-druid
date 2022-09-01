package sqlbuilder

import (
	"github.com/Hiddennn/go-druid/builder/query"
	"strings"
)

// JoinRaw join with raw sql
func (sb *SQLBuilder) JoinRaw(join string, values ...interface{}) *SQLBuilder {
	if values == nil || len(values) == 0 {
		return sb
	}

	tempParams := make([]query.SQLParameter, len(values))
	for i, value := range values {
		if value == nil {
			return sb
		}
		tempParams[i] = ConvertValueToSQLParameter(value)
	}
	sb._joinParams = append(sb._joinParams, tempParams...)

	var buf strings.Builder

	buf.WriteString(sb._join)
	if buf.Len() != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(join)

	sb._join = buf.String()

	return sb
}
