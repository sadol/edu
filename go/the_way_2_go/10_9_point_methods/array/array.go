package array

//the last element of the output tuple is NUMBER OF ARGS, in case of putting
//array of floats or slice of floats use `argName...' construct
func MinMaxArrFloat64(elements ...float64) (min float64, max float64, noOfElements uint) {
	for _, value := range elements {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
		noOfElements++
	}
	return
}
