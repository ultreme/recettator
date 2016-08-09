package recettator

import (
	"bytes"
	"math/rand"
	"strings"

	"github.com/moul/advanced-ssh-config/pkg/templates"
)

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

	r.title = "some random words"
	r.people = uint64(rand.Intn(4) + 1)

	r.ready = true
}

func (r *Recettator) Seed() uint64             { r.prepare(); return r.seed }
func (r *Recettator) Title() string            { r.prepare(); return r.title }
func (r *Recettator) People() uint64           { r.prepare(); return r.people }
func (r *Recettator) Ingredients() Ingredients { r.prepare(); return r.ingredients }
func (r *Recettator) Steps() Steps             { r.prepare(); return r.steps }

func (r *Recettator) Markdown() (string, error) {
	var buff bytes.Buffer
	tmpl, err := templates.New(strings.TrimSpace(`
# {{ .Title }}

Pour {{ .People }} personnes.

## Ingrédients

{{ .Ingredients }}

## Etapes

{{ .Steps }}
`))
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&buff, r); err != nil {
		return "", err
	}

	return buff.String(), nil
}
