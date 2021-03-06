package ingredients

import (
	"math/rand"
	"sort"
)

func (i *PoolCategory) append(ingredient Ingredient) {
	i.Availables = append(i.Availables, ingredient)
}

type Step struct {
	Instruction string
	Weight      int
}

type Steps []Step

func (s *Steps) Shuffle(rnd *rand.Rand) {
	for i := range *s {
		j := rnd.Intn(len(*s))
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (s *Steps) List(rnd *rand.Rand) []string {
	list := []string{}

	availableWeights := []int{}
	stepsByWeight := map[int]Steps{}
	for _, step := range *s {
		if _, found := stepsByWeight[step.Weight]; !found {
			stepsByWeight[step.Weight] = make(Steps, 0)
			availableWeights = append(availableWeights, step.Weight)
		}
		stepsByWeight[step.Weight] = append(stepsByWeight[step.Weight], step)
	}

	sort.Ints(availableWeights)

	for _, weight := range availableWeights {
		steps := stepsByWeight[weight]
		steps.Shuffle(rnd)
		for _, step := range steps {
			list = append(list, step.Instruction)
		}
	}

	return list
}

type Ingredient interface {
	Name() string
	Kind() string
	NameAndQuantity() string
	ToMap() map[string]interface{}
	TitlePart(left Ingredient) string
	IsMultiple() bool
	GetGender() string
	GetSteps() Steps
	GetMethod() Method
	SetMethod(Method)
}

type Method Ingredient

type Ingredients []Ingredient

type IngredientsPool struct {
	rand                 *rand.Rand
	MainIngredients      PoolCategory
	SecondaryIngredients PoolCategory
	IngredientMethods    PoolCategory
}

type PoolCategory struct {
	rand       *rand.Rand
	Availables Ingredients
	Picked     Ingredients
}

type IngredientMap map[string]interface{}

func (i *Ingredients) ToMap() []IngredientMap {
	ret := []IngredientMap{}
	for _, ingredient := range *i {
		ret = append(ret, ingredient.ToMap())
	}
	return ret
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
	return i.Picked[len(i.Picked)-1]
}

func (i *PoolCategory) GetSteps() Steps {
	steps := make(Steps, 0)
	for _, ingredient := range i.Picked {
		steps = append(steps, ingredient.GetSteps()...)
	}
	return steps
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
	pool.SecondaryIngredients.append(NewSecondaryIngredient("amandes", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("anis", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("betteraves", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("beurre", "male", false, rnd).SetIsSpreadable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("blancs d'oeufs", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("blé", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("cacahuètes", "female", true, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("cacao", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("camembert", "male", false, rnd).SetIsSpreadable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("canelle", "female", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("carottes", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("champignons de Paris", "male", false, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("citron", "male", false, rnd).SetIsCitrus())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("clous de girofle", "male", false, rnd).SetIsSpice())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("confifure d'oranges amères", "female", false, rnd).SetIsSpreadable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("cornichons", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("cube de Kubor®", "male", false, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("dattes", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("épices", "female", true, rnd).SetIsSpice())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("farine", "female", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("figues", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("flocons d'avoine", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("frites", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("fromage rapé", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("fruits sechés", "male", true, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("fruits", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("gousses de vanille", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("graines de pavot", "female", true, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("gui", "male", false, rnd).SetIsUncountable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("houx", "male", false, rnd).SetIsUncountable()) // can be singular or plural
	pool.SecondaryIngredients.append(NewSecondaryIngredient("haricots", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("jaunes d'oeufs", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("lierre", "male", false, rnd).SetIsUncountable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("mascarpone", "female", false, rnd).SetIsSpreadable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("maïs", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("morceaux de sucre", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("moutarde", "female", false, rnd).SetIsUncountable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("noisettes", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("noix de coco", "female", false, rnd).SetIsByPiece()) // can be singular or plural
	pool.SecondaryIngredients.append(NewSecondaryIngredient("noix", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("oeufs", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("olives", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("orange", "female", false, rnd).SetIsCitrus())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("pamplemousse", "male", false, rnd).SetIsCitrus())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("petits pois", "male", true, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("pommes de terre", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("radis", "male", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("reblochon", "male", false, rnd).SetIsSpreadable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("riz", "male", false, rnd).SetIsPowder())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("sel", "male", false, rnd).SetIsSpice())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("spaghettis", "male", true, rnd).SetIsUncountable())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("tomates", "female", true, rnd).SetIsByPiece())
	pool.SecondaryIngredients.append(NewSecondaryIngredient("vermicelles", "female", true, rnd).SetIsUncountable())

	pool.IngredientMethods.rand = rnd
	pool.IngredientMethods.append(NewIngredientMethod("farci", "farcie", "farcis", "farcies", Steps{{"remplissez %left% avec ce que vous voulez", 10}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("glacé", "glacée", "glacés", "glacées", Steps{{"mettez %left% au réfrigérateur quelques heures", 50}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("poêlé", "poêlée", "poêlés", "poêlées", Steps{{"faites revenir %left% dans une poêle", 50}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("roti", "rotie", "rotis", "roties", Steps{{"préchauffez le four pour y mettre %left% par la suite", -50}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("chaud", "chaude", "chauds", "chaudes", Steps{{"chauffez légerement %left% au four", 50}}, rnd))
	//pool.IngredientMethods.append(NewIngredientMethod("découpé", "découpée", "découpés", "découpées", Steps{{"découpez %left% en granches plutôt épaisses", -50}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("grillé", "grillée", "grillés", "grillées", Steps{{"mettez %left% sur le grill", 50}}, rnd))
	pool.IngredientMethods.append(NewIngredientMethod("battu", "battue", "battus", "battues", Steps{{"battez énergiquement %left% avec un fouet", -10}}, rnd))
	return &pool
}
