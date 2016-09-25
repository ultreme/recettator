package ingredients

import (
	"fmt"
	"math/rand"
	"strings"
)

func (i *PoolCategory) append(ingredient Ingredient) {
	i.Availables = append(i.Availables, ingredient)
}

type Ingredient interface {
	Name() string
	Kind() string
	NameAndQuantity() string
}

type Ingredients []Ingredient

type IngredientsPool struct {
	rand                 *rand.Rand
	MainIngredients      PoolCategory
	SecondaryIngredients PoolCategory
}

type PoolCategory struct {
	rand       *rand.Rand
	Availables Ingredients
	Picked     Ingredients
}

func (i *Ingredients) shuffle(rnd *rand.Rand) {
	for a := range *i {
		b := rnd.Intn(len(*i))
		(*i)[a], (*i)[b] = (*i)[b], (*i)[a]
	}
}

func (i *PoolCategory) Pick() Ingredient {
	i.Availables.shuffle(i.rand)
	i.Picked = append(i.Picked, i.Availables[0])
	i.Availables = i.Availables[1:]
	return i.Availables[0]
}

type MainIngredient struct {
	name     string
	quantity string

	Gender   string
	Multiple bool
}

func NewMainIngredient(name, gender string, multiple bool, rnd *rand.Rand) MainIngredient {
	ingredient := MainIngredient{
		name:     name,
		Gender:   gender,
		Multiple: multiple,
	}

	var words []string

	switch i := rnd.Intn(3); i {
	case 0, 1:
		var value int
		var unit string
		switch i {
		case 0:
			value = (rnd.Intn(50) + 1) * 10
			if value == 1 {
				unit = "gramme"
			} else {
				unit = "grammes"
			}
			break
		case 1:
			value = rnd.Intn(6) + 2
			if value == 1 {
				unit = "tranche"
			} else {
				unit = "tranches"
			}
			break
		}

		words = append(words, fmt.Sprintf("%d", value), unit)

		if beginsWithVoyel(ingredient.name) {
			words = append(words, "d'")
		} else {
			words = append(words, "de ")
		}
		ingredient.quantity = strings.Join(words, " ")
		break
	case 2:
		options := []string{}

		if ingredient.Gender == "male" && !ingredient.Multiple {
			options = append(options, "un bon gros ")
			options = append(options, "un assez gros ")
			options = append(options, "un plutôt gros ")
			options = append(options, "un relativement gros ")
			options = append(options, "du ")
			options = append(options, "un moyen ")
		}
		if ingredient.Gender == "female" && !ingredient.Multiple {
			options = append(options, "une bonne grosse ")
			options = append(options, "une assez grosse ")
			options = append(options, "une plutôt grosse ")
			options = append(options, "une relativement grosse ")
			options = append(options, "de la ")
			options = append(options, "une moyenne ")
		}
		if ingredient.Gender == "male" && ingredient.Multiple {
			options = append(options, "plusieurs gros ")
			options = append(options, "quelques gros ")
			options = append(options, "des ")
			options = append(options, "des ")
		}
		if ingredient.Gender == "female" && ingredient.Multiple {
			options = append(options, "plusieurs grosses ")
			options = append(options, "quelques grosses ")
			options = append(options, "des ")
		}

		for _, beginning := range []string{"une quantité suffisante", "pas mal", "quelques morceaux", "un bon paquet", "beaucoup", "un peu", "un tout petit peu", "une pincée"} {
			if beginsWithVoyel(ingredient.name) {
				options = append(options, fmt.Sprintf("%s d'", beginning))
			} else {
				options = append(options, fmt.Sprintf("%s de ", beginning))
			}
		}

		if len(options) > 0 {
			ingredient.quantity = options[rand.Intn(len(options))]
		}

		break
	}

	return ingredient
}

func (i MainIngredient) Kind() string { return "main-ingredient" }
func (i MainIngredient) Name() string { return i.name }
func (i MainIngredient) NameAndQuantity() string {
	return fmt.Sprintf("%s%s", i.quantity, i.name)
}

type SecondaryIngredient struct {
	name          string
	isMale        bool
	isMultiple    bool
	isUncountable bool
	isPowder      bool
	isCitrus      bool
	isSpice       bool
	isByPiece     bool
	isSpreadable  bool
}

func NewSecondaryIngredient(name string, male, multiple, uncountable, powder, citrus, spice, byPiece, spreadable bool) SecondaryIngredient {
	return SecondaryIngredient{
		name:          name,
		isMale:        male,
		isMultiple:    multiple,
		isUncountable: uncountable,
		isPowder:      powder,
		isCitrus:      citrus,
		isSpice:       spice,
		isByPiece:     byPiece,
		isSpreadable:  spreadable,
	}
}

func (i SecondaryIngredient) Kind() string { return "secondary-ingredient" }
func (i SecondaryIngredient) Name() string { return i.name }
func (i SecondaryIngredient) NameAndQuantity() string {
	return fmt.Sprintf("du %s", i.name)
}

func NewPool(rnd *rand.Rand) *IngredientsPool {
	var pool IngredientsPool
	pool.rand = rnd

	pool.MainIngredients.rand = rnd
	pool.MainIngredients.append(NewMainIngredient("agneau", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("autruche", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("calamar", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("canard", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("carpe", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("cheval", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("chips", "female", true, rnd))
	pool.MainIngredients.append(NewMainIngredient("dinde", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("foie d'oie", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("foie gras", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("gambas", "female", true, rnd))
	pool.MainIngredients.append(NewMainIngredient("jambon", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("langouste", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("langoustine", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("lapin", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("lardons", "male", true, rnd))
	pool.MainIngredients.append(NewMainIngredient("lièvre", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("lotte", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("mouette", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("nems", "male", true, rnd))
	pool.MainIngredients.append(NewMainIngredient("oie", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("pieuvre", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("poney", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("poulet", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("requin", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("salamandre", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("sanglier", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("saucisse", "female", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("saucisses Knacki®", "female", true, rnd))
	pool.MainIngredients.append(NewMainIngredient("soja", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("surimi", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("veau", "male", false, rnd))
	pool.MainIngredients.append(NewMainIngredient("âne", "male", false, rnd))
	// pool.MainIngredients.append(NewMainIngredient("", "", false, rnd))

	pool.SecondaryIngredients.rand = rnd
	return &pool
}
