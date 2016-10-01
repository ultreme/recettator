package ingredients

import "math/rand"

type MainIngredientMethod struct {
	singleMale     string
	multipleMale   string
	singleFemale   string
	multipleFemale string
	steps          Steps
	rand           *rand.Rand
	left           Ingredient
}

func NewMainIngredientMethod(singleMale, multipleMale, singleFemale, multipleFemale string, steps Steps, rnd *rand.Rand) *MainIngredientMethod {
	return &MainIngredientMethod{
		singleMale:     singleMale,
		multipleMale:   multipleMale,
		singleFemale:   singleFemale,
		multipleFemale: multipleFemale,
		steps:          steps,
		rand:           rnd,
	}
}

func (i MainIngredientMethod) SetLeft(left Ingredient) { i.left = left }
func (i MainIngredientMethod) GetSteps() Steps         { return i.steps }
func (i MainIngredientMethod) IsMultiple() bool {
	if i.left != nil {
		return i.left.IsMultiple()
	}
	panic("should not happen")
}
func (i MainIngredientMethod) GetGender() string {
	if i.left != nil {
		return i.left.GetGender()
	}
	panic("should not happen")
}

func (i MainIngredientMethod) TitlePart(left Ingredient) string {
	if left == nil {
		return i.singleMale
	}
	gender := left.GetGender()
	isMultiple := left.IsMultiple()
	switch {
	case gender == "male" && !isMultiple:
		return i.singleMale
	case gender == "male" && isMultiple:
		return i.multipleMale
	case gender == "female" && !isMultiple:
		return i.singleFemale
	case gender == "female" && isMultiple:
		return i.multipleFemale
	}
	panic("should not happen")
}

func (i MainIngredientMethod) Kind() string            { return "main-ingredient-method" }
func (i MainIngredientMethod) Name() string            { return i.TitlePart(i.left) }
func (i MainIngredientMethod) NameAndQuantity() string { return i.Name() }

func (i MainIngredientMethod) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.Name()
	ret["kind"] = i.Kind()
	return ret
}
