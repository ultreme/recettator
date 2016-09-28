package ingredients

import (
	"fmt"
	"math/rand"
)

type SecondaryIngredient struct {
	name          string
	gender        string
	quantity      string
	isMultiple    bool
	isUncountable bool
	isPowder      bool
	isCitrus      bool
	isSpice       bool
	isByPiece     bool
	isSpreadable  bool
}

func NewSecondaryIngredient(name string, gender string, isMultiple bool, rnd *rand.Rand) *SecondaryIngredient {
	ingredient := SecondaryIngredient{
		name:       name,
		gender:     gender,
		isMultiple: isMultiple,
		/*
			isMultiple:
			isUncountable:
			isPowder:
			isCitrus:
			isSpice:
			isByPiece:
			isSpreadable:
		*/
	}
	// FIXME: compute quantity
	return &ingredient
}

func (i SecondaryIngredient) Kind() string { return "secondary-ingredient" }
func (i SecondaryIngredient) Name() string { return i.name }
func (i SecondaryIngredient) NameAndQuantity() string {
	return fmt.Sprintf("%s%s", i.quantity, i.name)
}
func (i SecondaryIngredient) GetGender() string { return i.gender }
func (i SecondaryIngredient) IsMultiple() bool  { return i.isMultiple }
func (i SecondaryIngredient) TitlePart(left Ingredient) string {
	// FIXME: implement
	return ""
}

func (i SecondaryIngredient) ToMap() map[string]interface{} {
	ret := make(map[string]interface{}, 0)
	ret["name"] = i.name
	ret["kind"] = i.Kind()
	ret["name-and-quantity"] = i.NameAndQuantity()
	ret["quantity"] = i.quantity
	ret["is-multiple"] = i.isMultiple
	ret["gender"] = i.gender
	ret["is-by-piece"] = i.isByPiece
	ret["is-uncountable"] = i.isUncountable
	ret["is-powder"] = i.isPowder
	ret["is-citrus"] = i.isCitrus
	ret["is-spice"] = i.isSpice
	ret["is-spreadable"] = i.isSpreadable
	return ret
}

func (i *SecondaryIngredient) SetIsByPiece() *SecondaryIngredient {
	i.isByPiece = true
	return i
}
func (i *SecondaryIngredient) SetIsSpreadable() *SecondaryIngredient {
	i.isSpreadable = true
	return i
}
func (i *SecondaryIngredient) SetIsPowder() *SecondaryIngredient {
	i.isPowder = true
	return i
}
func (i *SecondaryIngredient) SetIsUncountable() *SecondaryIngredient {
	i.isUncountable = true
	return i
}
func (i *SecondaryIngredient) SetIsSpice() *SecondaryIngredient {
	i.isSpice = true
	return i
}
func (i *SecondaryIngredient) SetIsCitrus() *SecondaryIngredient {
	i.isCitrus = true
	return i
}

//, uncountable, powder, citrus, spice, byPiece, spreadable bool) SecondaryIngredient {
