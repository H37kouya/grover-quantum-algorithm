package num

import "math"

// Int64ToInt Int64をIntに変換
func Int64ToInt(i int64) int {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	} else {
		return int(i)
	}
}
