package in_array

type Type any

//Do
/**
	from	https://www.php.net/manual/zh/function.in-array.php
    php		in_array(mixed $needle, array $haystack, bool $strict = false): bool
	注释		大海捞针，在大海（haystack）中搜索针（ needle），如果没有设置 strict 则使用宽松的比较。
	备注		本次实现不考虑 strict 参数
*/

func Do(needle any, array []any) bool {

	for index, _ := range array {
		if array[index] == needle {
			return true
		}
	}

	return false
}
