package leetcode

type MovingAverage struct {
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return *new(MovingAverage)
}

func (this *MovingAverage) Next(val int) float64 {
	return 0.1
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
