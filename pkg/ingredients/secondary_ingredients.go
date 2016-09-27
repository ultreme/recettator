package ingredients

import "fmt"

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
