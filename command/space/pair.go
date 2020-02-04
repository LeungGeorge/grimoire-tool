package space

// Pair ...
type Pair struct {
	Key   string
	Value uint64
}

// PairList ...
type PairList []Pair

// Swap ...
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Len ...
func (p PairList) Len() int { return len(p) }

// Less ...
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
