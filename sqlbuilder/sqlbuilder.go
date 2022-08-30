package sqlbuilder

import (
	"errors"
	"github.com/Hiddennn/go-druid/builder/query"
	"strings"
)

var (
	// ErrTableEmpty table not set
	ErrTableEmpty = errors.New("table empty")
)

// SQLBuilder sql builder
type SQLBuilder struct {
	_select       string
	_table        string
	_join         string
	_where        string
	_groupBy      string
	_having       string
	_orderBy      string
	_limit        string
	_whereParams  []query.SQLParameter
	_havingParams []query.SQLParameter
	_limitParams  []query.SQLParameter
	_joinParams   []query.SQLParameter
}

// NewSQLBuilder init sql builder
func NewSQLBuilder() *SQLBuilder {
	return &SQLBuilder{}
}

// GetQuerySQL get sql
func (sb *SQLBuilder) GetQuerySQL() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}
	var buf strings.Builder

	buf.WriteString("SELECT ")
	if sb._select != "" {
		buf.WriteString(sb._select)
	} else {
		buf.WriteString("*")
	}
	buf.WriteString(" FROM ")
	buf.WriteString(sb._table)
	if sb._join != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._join)
	}
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}
	if sb._groupBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._groupBy)
	}
	if sb._having != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._having)
	}
	if sb._orderBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._orderBy)
	}
	if sb._limit != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._limit)
	}

	return buf.String(), nil
}

// GetQueryParams get params
func (sb *SQLBuilder) GetQueryParams() []query.SQLParameter {
	params := make([]query.SQLParameter, 0)
	params = append(params, sb._joinParams...)
	params = append(params, sb._whereParams...)
	params = append(params, sb._havingParams...)
	params = append(params, sb._limitParams...)
	return params
}

// Table set table
func (sb *SQLBuilder) Table(table string) *SQLBuilder {
	sb._table = table

	return sb
}

// Select set select cols
func (sb *SQLBuilder) Select(cols ...string) *SQLBuilder {
	var buf strings.Builder

	for k, col := range cols {
		buf.WriteString(col)

		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}

	sb._select = buf.String()
	return sb
}

// GenPlaceholders generate placeholders
func GenPlaceholders(n int) string {
	var buf strings.Builder

	for i := 0; i < n-1; i++ {
		buf.WriteString("?,")
	}

	if n > 0 {
		buf.WriteString("?")
	}

	return buf.String()
}

func ConvertValueToSQLParameter(value interface{}) query.SQLParameter {
	var sqlType string

	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		sqlType = "BIGINT"
	case float32:
		sqlType = "FLOAT"
	case float64:
		sqlType = "DOUBLE"
	case string:
		sqlType = "VARCHAR"
	default:
		return query.SQLParameter{}
	}

	return query.SQLParameter{
		Type:  sqlType,
		Value: value,
	}
}
