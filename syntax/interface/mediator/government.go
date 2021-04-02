package main

type Government struct {
	maxUniv     int
	maxStudents int
	maxTuition  int
	univQueue   []University
}

func NewGovern() *Government {
	return &Government{maxUniv: 3, maxStudents: 50, maxTuition: 1000}
}

func (g *Government) MakePublic(u University) bool {
	if len(g.univQueue) > g.maxUniv {
		return false
	}
	univQueue = append(univQueue, u)
	return true
}

func (g *Government) isFull(u University, students int) {
	if g.maxStudents > u.students+students {
		return true
	}
	return false

}

func (g *Government) isExpensive(u University, fee int) {
	if g.maxTuition > u.fee+fee {
		return true
	}
	return false
}
