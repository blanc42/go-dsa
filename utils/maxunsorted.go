package utils

import "errors"

func MaxUnsorted(arr []int) (int, error) {
	if len(arr) == 0 {
		err := errors.New("list is empty")
		return 0, err
	} else {
		max := arr[1]
		for _, a := range arr {
			if a > max {
				max = a
			}
		}
		return max, nil
	}
}
