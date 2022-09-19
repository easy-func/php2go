package in_array

import (
	"log"
	"strings"
)

type Type any

type CaseTo int8

const (
	CaseUpper = CaseTo(1)
	CaseLower = CaseTo(2)
)

//Do
/**
	from	https://www.php.net/manual/zh/function.array-change-key-case.php
    php		array_change_key_case(array $array, int $case = CASE_LOWER): array
	注释		array_change_key_case() 将 array 数组中的所有键名改为全小写或大写。本函数不改变数字索引。
*/
func Do(mapArray map[string]any, caseTo CaseTo) map[string]any {

	var funcChange = func(s string) (s1 string) { return }
	if caseTo == CaseUpper {
		funcChange = strings.ToUpper
	} else if caseTo == CaseLower {
		funcChange = strings.ToLower
	} else {
		log.Println("caseTo 代码传入值错误")
		funcChange = strings.ToUpper
	}

	var ret = map[string]any{}

	for key := range mapArray {
		ret[funcChange(key)] = mapArray[key]
	}

	return ret
}
