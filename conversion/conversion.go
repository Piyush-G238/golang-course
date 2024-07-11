package conversion

import "strconv"

func StringToFloat(val string) (float64, error) {
	return strconv.ParseFloat(val, 64)
}
