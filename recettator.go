package recettator // import "ultre.me/recettator"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"text/template"

	"ultre.me/recettator/pkg/ingredients"
)

type Settings struct {
	MainIngredients      uint64
	SecondaryIngredients uint64
	Steps                uint64
	// vegan
}

type Recettator struct {
	// components
	title  string
	people uint64
	steps  []string
	pool   *ingredients.IngredientsPool

	// settings
	settings Settings

	// internal
	seed  int64
	ready bool
	rnd   *rand.Rand // global random, used to add ingredients and steps
}

func New(seed int64) Recettator {
	rnd := rand.New(rand.NewSource(seed))
	return Recettator{
		seed:  seed,
		steps: make([]string, 0),
		pool:  ingredients.NewPool(rnd),
		rnd:   rnd,
	}
}

func (r *Recettator) applyDefaults() {
	if r.settings.MainIngredients == 0 {
		r.settings.MainIngredients = uint64(r.rnd.Intn(3) + 1)
	}
	if r.settings.SecondaryIngredients == 0 {
		r.settings.SecondaryIngredients = uint64(r.rnd.Intn(3) + 1)
	}
	if r.settings.Steps == 0 {
		r.settings.Steps = uint64(r.rnd.Intn(3) + 1)
	}
}

func (r *Recettator) pickItems() {
	for i := uint64(0); i < r.settings.MainIngredients; i++ {
		r.pool.MainIngredients.Pick()
	}
	for i := uint64(0); i < r.settings.SecondaryIngredients; i++ {
		r.pool.SecondaryIngredients.Pick()
	}

	for _, ingredient := range r.pool.MainIngredients.Picked {
		if r.rnd.Intn(20) < 2 {
			continue
		}
		if method := r.pool.IngredientMethods.Pick(); method != nil {
			ingredient.SetMethod(method)
		}
	}
}

func (r *Recettator) isValid() error {
	if r.settings.MainIngredients+r.settings.SecondaryIngredients < 1 {
		return fmt.Errorf("not enough ingredients.")
	}
	return nil
}

func (r *Recettator) prepare() {
	if r.ready {
		return
	}

	// pick items
	r.applyDefaults()
	if err := r.isValid(); err != nil {
		panic(err)
	}
	r.pickItems()

	// compute fields
	titleParts := []string{}
	var left ingredients.Ingredient

	pickedIngredients := append(
		r.pool.MainIngredients.Picked,
		r.pool.SecondaryIngredients.Picked...,
	)

	for _, ingredient := range pickedIngredients {
		titleParts = append(titleParts, ingredient.TitlePart(left))
		left = ingredient
	}
	r.title = strings.Join(titleParts, " ")
	r.people = uint64(r.rnd.Intn(4) + 1)

	steps := append(
		r.pool.MainIngredients.GetSteps(),
		r.pool.SecondaryIngredients.GetSteps()...,
	)
	r.steps = steps.List(r.rnd)

	r.ready = true
}

func (r *Recettator) Seed() int64                        { return r.seed }
func (r *Recettator) Settings() Settings                 { return r.settings }
func (r *Recettator) Title() string                      { r.prepare(); return r.title }
func (r *Recettator) People() uint64                     { r.prepare(); return r.people }
func (r *Recettator) Pool() *ingredients.IngredientsPool { r.prepare(); return r.pool }
func (r *Recettator) Steps() []string                    { r.prepare(); return r.steps }

func (r *Recettator) SetSettings(settings Settings) {
	r.settings = settings
}

func (r *Recettator) Markdown() (string, error) {
	var buff bytes.Buffer
	tmpl := template.Must(template.New("markdown").Parse(strings.TrimSpace(`
# {{ .Title }}

Pour {{ .People }} {{ if eq .People 1 }}personne{{ else }}personnes{{ end }}.

## Ingrédients

{{ range .Pool.MainIngredients.Picked }}* {{ .NameAndQuantity }}
{{ end }}{{ range .Pool.SecondaryIngredients.Picked }}* {{ .NameAndQuantity }}
{{ end }}
## Etapes

{{ range .Steps }}* {{.}}
{{end}} `)))

	if err := tmpl.Execute(&buff, r); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func (r *Recettator) ToMap() map[string]interface{} {
	export := make(map[string]interface{}, 0)
	r.prepare()
	export["seed"] = r.seed
	export["title"] = r.title
	export["steps"] = r.steps
	export["people"] = r.people
	export["settings"] = r.settings
	export["pool"] = map[string][]ingredients.IngredientMap{
		"main-ingredients":      r.pool.MainIngredients.Picked.ToMap(),
		"secondary-ingredients": r.pool.SecondaryIngredients.Picked.ToMap(),
	}
	return export
}

func (r *Recettator) JSON() string {
	export := r.ToMap()
	output, _ := json.MarshalIndent(export, "", "  ")
	return string(output)
}
