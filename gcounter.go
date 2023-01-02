package crdt

type GCounter struct {
	value int
}

func NewGCounter() *GCounter {
	counter := new(GCounter)
	counter.value = 0
	return counter
}

func (counter *GCounter) Value() int {
	return counter.value
}

func (counter *GCounter) Increment() {
	counter.value++
}

func (this *GCounter) Merge(that *GCounter) {
	value := this.value
	if that.value > value {
		value = that.value
	}
	this.value = value
}
