package internal

type poolElement interface {
	isAlive() bool
}

type pool interface {
	create(element poolElement)
	purge()
	list() []poolElement
}

func newPool() pool {
	return &poolType{
		pool: make([]poolElement, 0),
	}
}

type poolType struct {
	pool []poolElement
}

func (p *poolType) create(e poolElement) {
	p.pool = append(p.pool, e)
}

func (p *poolType) list() []poolElement {
	return p.pool
}

func (p *poolType) purge() {
	var newPool []poolElement
	for _, element := range p.pool {
		if element.isAlive() {
			newPool = append(newPool, element)
		}
	}
	p.pool = newPool
}
