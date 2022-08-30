package sqlbuilder

import (
	"strings"
)

// WhereRaw set where raw string
func (sb *SQLBuilder) WhereRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.whereRaw("AND", s, values)
}

// OrWhereRaw set where raw string
func (sb *SQLBuilder) OrWhereRaw(s string, values ...interface{}) *SQLBuilder {
	return sb.whereRaw("OR", s, values)
}

func (sb *SQLBuilder) whereRaw(operator string, s string, values []interface{}) *SQLBuilder {
	if values == nil || len(values) == 0 {
		return sb
	}

	var buf strings.Builder
	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._where = buf.String()

	for _, value := range values {
		sb._whereParams = append(sb._whereParams, ConvertValueToSQLParameter(value))
	}

	return sb
}

// Where set where cond
func (sb *SQLBuilder) Where(field string, condition string, value interface{}) *SQLBuilder {
	return sb.where("AND", condition, field, value)
}

// OrWhere set or where cond
func (sb *SQLBuilder) OrWhere(field string, condition string, value interface{}) *SQLBuilder {
	return sb.where("OR", condition, field, value)
}

func (sb *SQLBuilder) where(operator string, condition string, field string, value interface{}) *SQLBuilder {
	if value == nil {
		return sb
	}

	var buf strings.Builder
	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
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

	sb._where = buf.String()

	sb._whereParams = append(sb._whereParams, ConvertValueToSQLParameter(value))

	return sb
}

// WhereIn set where in cond
func (sb *SQLBuilder) WhereIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("AND", "IN", field, values)
}

// OrWhereIn set or where in cond
func (sb *SQLBuilder) OrWhereIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("OR", "IN", field, values)
}

// WhereNotIn set where not in cond
func (sb *SQLBuilder) WhereNotIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("AND", "NOT IN", field, values)
}

// OrWhereNotIn set or where not in cond
func (sb *SQLBuilder) OrWhereNotIn(field string, values ...interface{}) *SQLBuilder {
	return sb.whereIn("OR", "NOT IN", field, values)
}

func (sb *SQLBuilder) whereIn(operator string, condition string, field string, values []interface{}) *SQLBuilder {
	if values == nil || len(values) == 0 {
		return sb
	}

	var buf strings.Builder
	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	plhs := GenPlaceholders(len(values))
	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("(")
	buf.WriteString(plhs)
	buf.WriteString(")")

	sb._where = buf.String()

	for _, value := range values {
		sb._whereParams = append(sb._whereParams, ConvertValueToSQLParameter(value))
	}

	return sb
}
