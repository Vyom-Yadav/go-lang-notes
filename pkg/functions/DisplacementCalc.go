package functions

import "math"

func GenDispFn(acc, initVel, initDisp float64) func(time float64) float64 {
	return func(time float64) float64 {
		return initDisp + initVel*time + (0.5)*(acc*math.Pow(time, 2))
	}
}
