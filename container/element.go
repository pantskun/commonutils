package container

type Element interface{}

// type Element interface {
// 	Equal(Element) bool
// }

// func Equal(a, b Element) bool {
// 	if IsBasicType(a) && !IsBasicType(b) {
// 		return false
// 	}

// 	if !IsBasicType(a) && IsBasicType(b) {
// 		return false
// 	}

// 	if IsBasicType(a) && IsBasicType(b) {
// 		if GetBasicTypeName(a) != GetBasicTypeName(b) {
// 			return false
// 		}

// 		return a == b
// 	}

// 	return a.Equal(b)
// }

// func IsBasicType(a interface{}) bool {
// 	switch a.(type) {
// 	case bool:
// 		return true
// 	case string:
// 		return true
// 	case int:
// 		return true
// 	case int8:
// 		return true
// 	case int16:
// 		return true
// 	case int32:
// 		return true
// 	case int64:
// 		return true
// 	case uint:
// 		return true
// 	case uint8:
// 		return true
// 	case uint16:
// 		return true
// 	case uint32:
// 		return true
// 	case uint64:
// 		return true
// 	case uintptr:
// 		return true
// 	case float32:
// 		return true
// 	case float64:
// 		return true
// 	case complex64:
// 		return true
// 	case complex128:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func GetBasicTypeName(a interface{}) string {
// 	switch a.(type) {
// 	case bool:
// 		return "bool"
// 	case string:
// 		return "string"
// 	case int:
// 		return "int"
// 	case int8:
// 		return "int8"
// 	case int16:
// 		return "int16"
// 	case int32:
// 		return "int32"
// 	case int64:
// 		return "int64"
// 	case uint:
// 		return "uint"
// 	case uint8:
// 		return "uint8"
// 	case uint16:
// 		return "uint16"
// 	case uint32:
// 		return "uint32"
// 	case uint64:
// 		return "uint64"
// 	case uintptr:
// 		return "uintptr"
// 	case float32:
// 		return "float32"
// 	case float64:
// 		return "float64"
// 	case complex64:
// 		return "complex64"
// 	case complex128:
// 		return "complex128"
// 	default:
// 		return "err"
// 	}
// }
