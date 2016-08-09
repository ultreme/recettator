package recettator

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"strings"

	"github.com/moul/advanced-ssh-config/pkg/templates"
)

type Recettator struct {
	// components
	title       string
	people      uint64
	steps       Steps
	ingredients Ingredients

	// internal
	seed  int64
	ready bool
	rnd   *rand.Rand // global random, used to add ingredients and steps
}

func New(seed int64) Recettator {
	return Recettator{
		seed:        seed,
		steps:       make(Steps, 0),
		ingredients: make(Ingredients, 0),
		rnd:         rand.New(rand.NewSource(seed)),
	}
}

func (r *Recettator) AddRandomIngredient() error { r.ready = false; return nil }
func (r *Recettator) AddRandomStep() error       { r.ready = false; return nil }

func (r *Recettator) prepare() {
	if r.ready {
		return
	}

	// dedicated random for prepare
	rnd := rand.New(rand.NewSource(r.seed))

	r.title = "some random words"
	r.people = uint64(rnd.Intn(4) + 1)

	r.ready = true
}

func (r *Recettator) Seed() int64              { r.prepare(); return r.seed }
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

func (r *Recettator) JSON() string {
	export := make(map[string]interface{}, 0)
	export["seed"] = r.Seed()
	export["title"] = r.Title()
	export["steps"] = r.Steps()
	export["people"] = r.People()
	export["ingredients"] = r.Ingredients()

	output, _ := json.MarshalIndent(export, "", "  ")
	return string(output)
}
