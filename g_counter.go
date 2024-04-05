package crdt

type GCounter struct {
	node   string
	values map[string]int
}

func NewGCounter(node string, nodes []string) *GCounter {
	counter := new(GCounter)
	counter.node = node
	counter.values = make(map[string]int)
	for _, n := range nodes {
		counter.values[n] = 0
	}
	return counter
}

func (counter *GCounter) Value() int {
	sum := 0
	for _, val := range counter.values {
		sum += val
	}
	return sum
}

func (counter *GCounter) Increment() {
	counter.values[counter.node]++
}

func (counter *GCounter) Merge(that *GCounter) {
	for key, value := range counter.values {
		if value < that.values[key] {
			counter.values[key] = that.values[key]
		}
	}
}
