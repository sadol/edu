package miner

type Miner interface {
	Less(i, j int) bool
	Len() int
}

// linear complexity
func Min(miner Miner) (minIndex int) {
	if miner.Len() > 0 {
		minIndex = 0
		for i := 1; i < miner.Len(); i++ {
			if miner.Less(i, minIndex) {
				minIndex = i
			}
		}
	} else {
		minIndex = -1
	}
	return
}
