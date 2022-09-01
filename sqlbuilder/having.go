package sqlbuilder

import (
	"github.com/Hiddennn/go-druid/builder/query"
	"strings"
)

// HavingRaw set having raw string
func (sb *SQLBuilder) HavingRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.havingRaw("AND", s, values)
}

// OrHavingRaw set having raw string
func (sb *SQLBuilder) OrHavingRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.havingRaw("OR", s, values)
}

func (sb *SQLBuilder) havingRaw(operator string, s string, values []interface{}) *SQLBuilder {
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
	sb._havingParams = append(sb._havingParams, tempParams...)

	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._having = buf.String()

	return sb
}

// Having set having cond
func (sb *SQLBuilder) Having(field string, condition string, value interface{}) *SQLBuilder {
	return sb.having("AND", condition, field, value)
}

// OrHaving set or having cond
func (sb *SQLBuilder) OrHaving(field string, condition string, value interface{}) *SQLBuilder {
	return sb.having("OR", condition, field, value)
}

func (sb *SQLBuilder) having(operator string, condition string, field string, value interface{}) *SQLBuilder {
	if value == nil {
		return sb
	}

	if sb._groupBy == "" { // group by not set
		return sb
	}

	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("?")

	sb._having = buf.String()

	sb._havingParams = append(sb._havingParams, ConvertValueToSQLParameter(value))

	return sb
}
