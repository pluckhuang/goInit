package exercise

import (
	"fmt"
)

func DeleteElement[T any](oldSlice []T, index int) (newSlice []T, err error) {

	oldSliceLen, oldSliceCap := len(oldSlice), cap(oldSlice)

	if index >= oldSliceLen || index < 0 {
		return oldSlice, fmt.Errorf("invalid index")
	} else {
		if oldSliceCap > 256 && oldSliceCap > ((oldSliceLen*5/4)*5/4) { // len 大于256 并且 小于2次 1.25倍 扩容大小
			newSlice := make([]T, oldSliceLen, (oldSliceCap * 4 / 5)) // 缩容为 上次 控容时的 cap大小
			fmt.Println("reSize occur, new cap:", cap(newSlice))
			copy(newSlice[:index], oldSlice[:index])
			copy(newSlice[index:], oldSlice[index+1:])
			return newSlice[:oldSliceLen-1], nil
		} else {
			copy(oldSlice[index:], oldSlice[index+1:])
			return oldSlice[:oldSliceLen-1], nil
		}
	}
}
