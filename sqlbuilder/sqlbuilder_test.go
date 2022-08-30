package sqlbuilder

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/proto"
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
		Where("aa", "=", proto.String("aa"))

	sql, err := builder.GetQuerySQL()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sql)

	params := builder.GetQueryParams()
	spew.Dump(params)
}
