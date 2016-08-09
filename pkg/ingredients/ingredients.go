package ingredients

var RegisteredIngredients Ingredients

func (i *Ingredients) append(ingredient Ingredient) { *i = append(*i, ingredient) }

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
	RegisteredIngredients.append(NewMainIngredient("agneau", "male", false))
	RegisteredIngredients.append(NewMainIngredient("autruche", "female", false))
	RegisteredIngredients.append(NewMainIngredient("carpe", "female", false))
	RegisteredIngredients.append(NewMainIngredient("cheval", "male", false))
	RegisteredIngredients.append(NewMainIngredient("chips", "female", true))
	RegisteredIngredients.append(NewMainIngredient("dinde", "female", false))
	RegisteredIngredients.append(NewMainIngredient("foie d'oie", "male", false))
	RegisteredIngredients.append(NewMainIngredient("foie gras", "male", false))
	RegisteredIngredients.append(NewMainIngredient("jambon", "male", false))
	RegisteredIngredients.append(NewMainIngredient("lardons", "male", true))
	RegisteredIngredients.append(NewMainIngredient("lièvre", "male", false))
	RegisteredIngredients.append(NewMainIngredient("lotte", "female", false))
	RegisteredIngredients.append(NewMainIngredient("oie", "female", false))
	RegisteredIngredients.append(NewMainIngredient("poney", "male", false))
	RegisteredIngredients.append(NewMainIngredient("poulet", "male", false))
	RegisteredIngredients.append(NewMainIngredient("requin", "male", false))
	RegisteredIngredients.append(NewMainIngredient("surimi", "male", false))
	RegisteredIngredients.append(NewMainIngredient("veau", "male", false))
	// RegisteredIngredients.append(NewMainIngredient("", "", false))
}