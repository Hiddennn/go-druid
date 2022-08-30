package sqlbuilder

import "strings"

type OrderPolicy struct {
	Field     string
	Direction string
}

// OrderBy set order by fields
func (sb *SQLBuilder) OrderBy(orderPolicys ...OrderPolicy) *SQLBuilder {
	if orderPolicys == nil || len(orderPolicys) == 0 {
		return sb
	}

	var buf strings.Builder
	buf.WriteString("ORDER BY ")

	for k, orderPolicy := range orderPolicys {
		buf.WriteString(orderPolicy.Field)
		buf.WriteString(" ")
		buf.WriteString(orderPolicy.Direction)

		if k != len(orderPolicys)-1 {
			buf.WriteString(",")
		}
	}

	sb._orderBy = buf.String()
	return sb
}
