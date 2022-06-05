/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/6/3
*/

package util

import "strconv"

func Str2Int64(slice []string) []int64 {
	ans := make([]int64, len(slice))
	for i, item := range slice {
		ans[i], _ = strconv.ParseInt(item, 10, 64)
	}
	return ans
}
