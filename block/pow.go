package block

type Pow struct {
	Block *Block
}

func NewPOW(block *Block) *Pow {
	return &Pow{block}
}

func (p *Pow) Run() ([]byte, int64) {
	return nil, 0
}
