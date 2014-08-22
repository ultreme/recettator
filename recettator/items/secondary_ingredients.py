# -*- coding: utf-8 -*-

from .item import Item


class SecondaryIngredient(Item):
    kind = 'secondary_ingredient'

    @property
    def attrs(self):
        attrs = super(SecondaryIngredient, self).attrs
        attrs['name'] = self.name
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs

    def ingredient_list_str(self):
        return '42'


class Noisettes(SecondaryIngredient):
    name = 'noisettes'
    gender = 'female'
    quantity = 'multiple'


class Amandes(SecondaryIngredient):
    name = 'amandes'
    gender = 'female'
    quantity = 'multiple'


class Noix(SecondaryIngredient):
    name = 'noix'
    gender = 'female'
    quantity = 'multiple'


class PommesDeTerre(SecondaryIngredient):
    name = 'pommes de terre'
    gender = 'female'
    quantity = 'multiple'


class Dattes(SecondaryIngredient):
    name = 'dattes'
    gender = 'female'
    quantity = 'multiple'


class GraintesDePavot(SecondaryIngredient):
    name = 'graintes de pavot'
    gender = 'female'
    quantity = 'multiple'


class Epices(SecondaryIngredient):
    name = 'epices'
    gender = 'female'
    quantity = 'multiple'


class Tomates(SecondaryIngredient):
    name = 'tomates'
    gender = 'female'
    quantity = 'multiple'


class GoussesDeVanille(SecondaryIngredient):
    name = 'gousses de vanille'
    gender = 'female'
    quantity = 'multiple'


class Canelle(SecondaryIngredient):
    name = 'canelle'
    gender = 'female'
    quantity = 'single'


class NoixDeCoco(SecondaryIngredient):
    name = 'noix de coco'
    gender = 'female'
    quantity = 'single'


class Mascarpone(SecondaryIngredient):
    name = 'mascarpone'
    gender = 'female'
    quantity = 'single'


class ConfigureDOrangeAmeres(SecondaryIngredient):
    name = 'confiture d\'orange ameres'
    gender = 'female'
    quantity = 'single'


class Orange(SecondaryIngredient):
    name = 'orange'
    gender = 'female'
    quantity = 'single'


class Pamplemousse(SecondaryIngredient):
    name = 'pamplemousse'
    gender = 'female'
    quantity = 'single'


class Farine(SecondaryIngredient):
    name = 'farine'
    gender = 'female'
    quantity = 'single'


class Moutarde(SecondaryIngredient):
    name = 'moutarde'
    gender = 'female'
    quantity = 'single'


class Gui(SecondaryIngredient):
    name = 'gui'
    gender = 'male'
    quantity = 'single'


class Houx(SecondaryIngredient):
    name = 'houx'
    gender = 'male'
    quantity = 'single'


class Ble(SecondaryIngredient):
    name = 'ble'
    gender = 'male'
    quantity = 'single'


class Lierre(SecondaryIngredient):
    name = 'lierre'
    gender = 'male'
    quantity = 'single'


class Anis(SecondaryIngredient):
    name = 'anis'
    gender = 'male'
    quantity = 'single'


class Citron(SecondaryIngredient):
    name = 'citron'
    gender = 'male'
    quantity = 'single'


class Mais(SecondaryIngredient):
    name = 'mais'
    gender = 'male'
    quantity = 'single'


class Oeuf(SecondaryIngredient):
    name = 'oeuf'
    gender = 'male'
    quantity = 'single'


class Beurre(SecondaryIngredient):
    name = 'beurre'
    gender = 'male'
    quantity = 'single'


class Sel(SecondaryIngredient):
    name = 'sel'
    gender = 'male'
    quantity = 'single'


class Riz(SecondaryIngredient):
    name = 'riz'
    gender = 'male'
    quantity = 'single'


class Cacao(SecondaryIngredient):
    name = 'cacao'
    gender = 'male'
    quantity = 'single'


class FromageRape(SecondaryIngredient):
    name = 'fromage rape'
    gender = 'male'
    quantity = 'single'


class CubeDeKubor(SecondaryIngredient):
    name = 'cube de Kubor'
    gender = 'male'
    quantity = 'single'


class Reblochon(SecondaryIngredient):
    name = 'reblochon'
    gender = 'male'
    quantity = 'single'


class FloconsDAvoine(SecondaryIngredient):
    name = 'flocons d\'avoine'
    gender = 'male'
    quantity = 'multiple'


class Fruits(SecondaryIngredient):
    name = 'fruits'
    gender = 'male'
    quantity = 'multiple'


class FruitsSecs(SecondaryIngredient):
    name = 'fruits secs'
    gender = 'male'
    quantity = 'multiple'


class ClousDeGirofle(SecondaryIngredient):
    name = 'clous de girofle'
    gender = 'male'
    quantity = 'multiple'


class PetitsPois(SecondaryIngredient):
    name = 'petits pois'
    gender = 'male'
    quantity = 'multiple'


class Oeufs(SecondaryIngredient):
    name = 'oeufs'
    gender = 'male'
    quantity = 'multiple'


class BlancsDOeuf(SecondaryIngredient):
    name = 'blancs d\'oeuf'
    gender = 'male'
    quantity = 'multiple'


class JaunesDOeuf(SecondaryIngredient):
    name = 'jaunes d\'oeuf'
    gender = 'male'
    quantity = 'multiple'


class MorceauxDeSucre(SecondaryIngredient):
    name = 'morceaux de sucre'
    gender = 'male'
    quantity = 'multiple'


class ChampignonsDeParis(SecondaryIngredient):
    name = 'champignons de Paris'
    gender = 'male'
    quantity = 'multiple'


def all_items():
    return (
        Noisettes, Amandes, Noix, PommesDeTerre, Dattes, GraintesDePavot,
        Epices, Tomates, GoussesDeVanille,
        Canelle, NoixDeCoco, Mascarpone, ConfigureDOrangeAmeres, Orange,
        Pamplemousse, Farine, Moutarde,
        Gui, Houx, Ble, Lierre, Anis, Citron, Mais, Oeuf, Beurre, Sel, Riz,
        Cacao, FromageRape, CubeDeKubor, Reblochon,
        FloconsDAvoine, Fruits, FruitsSecs, ClousDeGirofle, PetitsPois, Oeufs,
        BlancsDOeuf, JaunesDOeuf, MorceauxDeSucre, ChampignonsDeParis,
    )
