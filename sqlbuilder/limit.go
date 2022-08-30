package sqlbuilder

import (
	"github.com/Hiddennn/go-druid/builder/query"
	"strings"
)

// Limit set limit
func (sb *SQLBuilder) Limit(offset, num int) *SQLBuilder {
	if num <= 0 || offset < 0 {
		return sb
	}

	var buf strings.Builder
	buf.WriteString("LIMIT ? OFFSET ?")
	sb._limit = buf.String()

	sb._limitParams = append(sb._limitParams,
		query.SQLParameter{Type: "INTEGER", Value: num},
		query.SQLParameter{Type: "INTEGER", Value: offset},
	)

	return sb
}
