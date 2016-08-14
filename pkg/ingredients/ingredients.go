package ingredients

import (
	"fmt"
	"math/rand"
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

func NewMainIngredient(name, gender string, multiple bool) MainIngredient {
	return MainIngredient{
		name:     name,
		quantity: "42 grammes d'",

		Gender:   gender,
		Multiple: multiple,
	}
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
	pool.MainIngredients.append(NewMainIngredient("agneau", "male", false))
	pool.MainIngredients.append(NewMainIngredient("autruche", "female", false))
	pool.MainIngredients.append(NewMainIngredient("canard", "male", false))
	pool.MainIngredients.append(NewMainIngredient("carpe", "female", false))
	pool.MainIngredients.append(NewMainIngredient("cheval", "male", false))
	pool.MainIngredients.append(NewMainIngredient("chips", "female", true))
	pool.MainIngredients.append(NewMainIngredient("dinde", "female", false))
	pool.MainIngredients.append(NewMainIngredient("foie d'oie", "male", false))
	pool.MainIngredients.append(NewMainIngredient("foie gras", "male", false))
	pool.MainIngredients.append(NewMainIngredient("jambon", "male", false))
	pool.MainIngredients.append(NewMainIngredient("lardons", "male", true))
	pool.MainIngredients.append(NewMainIngredient("lièvre", "male", false))
	pool.MainIngredients.append(NewMainIngredient("lotte", "female", false))
	pool.MainIngredients.append(NewMainIngredient("nems", "male", true))
	pool.MainIngredients.append(NewMainIngredient("oie", "female", false))
	pool.MainIngredients.append(NewMainIngredient("poney", "male", false))
	pool.MainIngredients.append(NewMainIngredient("poulet", "male", false))
	pool.MainIngredients.append(NewMainIngredient("requin", "male", false))
	pool.MainIngredients.append(NewMainIngredient("saucisse", "female", false))
	pool.MainIngredients.append(NewMainIngredient("saucisses Knacki®", "female", true))
	pool.MainIngredients.append(NewMainIngredient("surimi", "male", false))
	pool.MainIngredients.append(NewMainIngredient("veau", "male", false))
	// pool.MainIngredients.append(NewMainIngredient("", "", false))

	pool.SecondaryIngredients.rand = rnd
	return &pool
}
