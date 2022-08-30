package sqlbuilder

import (
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

	for _, value := range values {
		sb._havingParams = append(sb._havingParams, ConvertValueToSQLParameter(value))
	}

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
