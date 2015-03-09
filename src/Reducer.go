package src

type Reducer struct{}

func NewReducer() *Reducer {
	return &Reducer{}
}

func (this *Reducer) Reduce(values []float64) float64 {
	var sum float64
	var i int
	var length = len(values)
	for i = 0; i < length; i++ {
		sum += values[i]
	}
	return sum / float64(length)
}
