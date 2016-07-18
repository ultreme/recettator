package recettator

type Recettator struct {
	seed        uint64
	title       string
	people      uint64
	steps       Steps
	ingredients Ingredients

	ready bool
}

func New(seed uint64) Recettator {
	return Recettator{
		seed:        seed,
		steps:       make(Steps, 0),
		ingredients: make(Ingredients, 0),
	}
}

func (r *Recettator) AddRandomIngredient() error { return nil }
func (r *Recettator) AddRandomStep() error       { return nil }

func (r *Recettator) prepare() {
	if r.ready {
		return
	}
}

func (r *Recettator) Seed() uint64             { r.prepare(); return r.seed }
func (r *Recettator) Title() string            { r.prepare(); return r.title }
func (r *Recettator) People() uint64           { r.prepare(); return r.people }
func (r *Recettator) Ingredients() Ingredients { r.prepare(); return r.ingredients }
func (r *Recettator) Steps() Steps             { r.prepare(); return r.steps }
