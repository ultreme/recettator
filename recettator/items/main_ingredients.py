# -*- coding: utf-8 -*-

from .item import Item


class MainIngredient(Item):
    kind = 'main_ingredient'

    @property
    def attrs(self):
        attrs = super(MainIngredient, self).attrs
        attrs['name'] = self.name
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs

    def ingredient_list_str(self):
        return '42'


class FoieGras(MainIngredient):
    name = 'foie gras'
    gender = 'male'
    quantity = 'single'


class FoieDOie(MainIngredient):
    name = 'foie d\'oie'
    gender = 'male'
    quantity = 'single'


class Lievre(MainIngredient):
    name = 'lievre'
    gender = 'male'
    quantity = 'single'


class Jambon(MainIngredient):
    name = 'jambon'
    gender = 'male'
    quantity = 'single'


class Poulet(MainIngredient):
    name = 'poulet'
    gender = 'male'
    quantity = 'single'


class Surimi(MainIngredient):
    name = 'surimi'
    gender = 'male'
    quantity = 'single'


class Requin(MainIngredient):
    name = 'requin'
    gender = 'male'
    quantity = 'single'


class Cheval(MainIngredient):
    name = 'cheval'
    gender = 'male'
    quantity = 'single'


class Veau(MainIngredient):
    name = 'veau'
    gender = 'male'
    quantity = 'single'


class Lotte(MainIngredient):
    name = 'lotte'
    gender = 'female'
    quantity = 'single'


class Oie(MainIngredient):
    name = 'oie'
    gender = 'female'
    quantity = 'single'


class Carpe(MainIngredient):
    name = 'carpe'
    gender = 'female'
    quantity = 'single'


class Dinde(MainIngredient):
    name = 'dinde'
    gender = 'female'
    quantity = 'single'


class Autruche(MainIngredient):
    name = 'autruche'
    gender = 'female'
    quantity = 'single'


class Lardons(MainIngredient):
    name = 'lardons'
    gender = 'male'
    quantity = 'multiple'


def all_items():
    return (
        FoieGras, FoieDOie, Lievre, Jambon, Poulet, Surimi, Requin, Cheval,
        Veau, Lotte, Oie, Carpe, Dinde, Autruche, Lardons,
    )
