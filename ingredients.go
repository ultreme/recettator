package recettator

var RegisteredIngredients Ingredients

type Ingredient interface {
	Name() string
	Kind() string
	// Gender() string
	// Quantity() string
}

type Ingredients []Ingredient

type StandardIngredient struct {
	name string
	kind string

	Gender   string
	Multiple bool
}

func NewMainIngredient(name, gender string, multiple bool) StandardIngredient {
	return StandardIngredient{
		name: name,
		kind: "main",

		Gender:   gender,
		Multiple: multiple,
	}
}

func (i StandardIngredient) Name() string { return i.name }
func (i StandardIngredient) Kind() string { return i.kind }

func init() {
	RegisteredIngredients = append(RegisteredIngredients, NewMainIngredient("jambon", "male", false))
}
