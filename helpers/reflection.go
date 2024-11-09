package helpers

import (
	"fmt"
	"reflect"
)

// Expected type is passed as a generic and this method will use reflection to compare the underlying type agains T.
// Returns the casted type or panics if it fails. TODO: Return error instead of panic
// ExpectType 函数用于检查传入的参数 r 是否为期望的类型 T。
// 如果 r 的类型与期望的类型 T 匹配，则返回 r 的类型转换结果。
// 如果类型不匹配，则会触发一个 panic 并输出详细的错误信息。
//
// 参数:
//   - r: 任意类型的参数，函数将检查其类型是否为期望的类型 T。
//
// 返回值:
//   - 返回类型为 T 的值，如果 r 的类型与期望的类型 T 匹配。
//
// 泛型:
//   - T: 期望的类型。
//
// 示例:
//
//	var result int = ExpectType[int](42) // 正确，42 是 int 类型
//	var result string = ExpectType[string](42) // 错误，42 不是 string 类型，将触发 panic
//
// 注意:
//   - 如果类型不匹配，函数将触发 panic 并输出错误信息，格式为 "Expected %T but instead received %T inside ExpectType[T](r)"。
func ExpectType[T any](r any) T {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	recievedType := reflect.TypeOf(r)

	if expectedType == recievedType {
		return r.(T)
	}

	panic(fmt.Sprintf("Expected %T but instead recived %T inside ExpectType[T](r)\n", expectedType, recievedType))
}
