package test

import (
	"fmt"
	"testing"

	"github.com/elgris/sqrl"
)

func TestSqlr(t *testing.T) {
	sqlStr := sqrl.Select("*").From("m.table").
		Where("map_id = ?", "mapID").
		Where("vin = ?", "vin").
		Where("start_time <= ?", "time").
		Where("status NOT IN (?,?)", "util.TaskStatusCancelled", "util.TaskStatusWaitCancel").
		OrderBy("end_time DESC").
		Limit(1)

	query, args, err := sqlStr.ToSql()

	fmt.Printf("sql: %+v, args: %+v, err: %+v", query, args, err)
}
