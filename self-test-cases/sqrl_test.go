package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/elgris/sqrl"
)

func TestSqrl(t *testing.T) {
	// sqlStr := sqrl.Select("*").From("task").
	// 	Where(sqrl.GtOrEq{"start_time": 11111111111}).
	// 	Where(sqrl.LtOrEq{"end_time": 2222222222})

	sqlStr := sqrl.Select("*").From("task").
		Where(`
			(
				start_time >= ? AND
				end_time   <= ?
			)
			OR
			(
				start_time < ? AND
				end_time   > ?
			)
		`, 11, 22, 11, 22)
	fmt.Printf("sqlStr: %#v\n", sqlStr)
	query, args, err := sqlStr.ToSql()
	if err != nil {
		fmt.Printf("query task sql str error: %+v", err)
	}
	fmt.Printf("query: %v\n", query)
	fmt.Printf("args: %v\n", args)

	taskIDs := []string{"t1", "t2", "t3"}
	pred := map[string]interface{}{
		"task_id": taskIDs,
	}
	sqlStr2 := sqrl.Select("*").From("task").
		Where("map_id = ?", "map1").
		Where(pred)
	query2, args2, err := sqlStr2.ToSql()
	fmt.Printf("sqlStr2: %#v\n", sqlStr2)
	fmt.Printf("query2: %v\n", query2)
	fmt.Printf("args2: %v\n", args2)
	fmt.Printf("taskIDs: %v, err: %+v\n", reflect.ValueOf(taskIDs).Kind(), err)
}
