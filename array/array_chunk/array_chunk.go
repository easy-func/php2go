package in_array

import "log"

type Type any

type CaseTo int8

const (
	CaseUpper = CaseTo(1)
	CaseLower = CaseTo(2)
)

//Do
/**
	from	https://www.php.net/manual/zh/function.array-chunk.php
    php		array_chunk(array $array, int $length, bool $preserve_keys = false): array
	注释		将一个数组分割成多个数组，其中每个数组的单元数目由 length 决定。最后一个数组的单元数目可能会少于 length 个。
	备注		由于go的数组不像php是array+map，下标可以自定义，所以本函数不实现 preserveKeys 这个参数
*/
func Do(array []any, length int) [][]any {

	if length < 1 {
		log.Println("error length 必须大于0")
		length = 1
	}

	var ret [][]any
	if len(array) == 0 {
		log.Println("error length 必须大于0")
		return [][]any{}
	}

	var splitArray []any
	var i = 0
	for index, _ := range array {
		i++
		splitArray = append(splitArray, array[index])
		if i == length {
			i = 1
			ret = append(ret, splitArray)
			splitArray = []any{}
		}

	}

	return ret
}
