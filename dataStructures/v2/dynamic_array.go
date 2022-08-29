package v2

//type dynamicArray struct {
//	store []any
//	count int
//}
//
//func (da *dynamicArray) append(element any) {
//	if da.count >= len(da.store) {
//		const newSize = 2 * len(da.store)
//		newArr := [newSize]any{}
//
//		for index, value := range da.store {
//			newArr[index] = value
//		}
//		da.store = newArr
//	}
//
//	// append the new element to the new array.
//	da.store[da.count+1] = element
//	da.count += 1
//}
