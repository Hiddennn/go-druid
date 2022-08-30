package sqlbuilder

import (
	"fmt"
	"testing"
)

func TestSQLBuilder(t *testing.T) {
	builder := NewSQLBuilder()
	builder.Select("aa", "bb").Table("table").OrderBy(OrderPolicy{
		Field:     "aa",
		Direction: "asc",
	}, OrderPolicy{
		Field:     "bb",
		Direction: "desc",
	}).GroupBy("aa", "bb").Limit(10, 10).
		Where("aa", "=", "aa")

	sql, err := builder.GetQuerySQL()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql)

	params := builder.GetQueryParams()
	fmt.Println(params)
}
