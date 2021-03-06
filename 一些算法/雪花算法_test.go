package algorithm

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetSnowflakeId(t *testing.T) {
	//var ids = []int64{}
	var ids = make([]int64, 0)

	//设置一个机器标识，如IP编码,防止分布式机器生成重复码
	SetMachineId(192168100101)

	fmt.Println("start", time.Now().Format("2006-01-02 15:04"))
	for i := 0; i < 10000000; i++ {
		id := GetSnowflakeId()
		ids = append(ids, id)
	}
	fmt.Println("end  ", time.Now().Format("2006/1/2 15:04:05"))

	result := Duplicate(ids)
	fmt.Println("去重后数量：", len(result))
	fmt.Println(result[10], result[11], result[12], result[13], result[14])
	fmt.Println(result[9990], result[9991], result[9992], result[9993], result[9994])
}

//去重
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}
