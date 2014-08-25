# -*- coding: utf-8 -*-

import random

from .item import GenderizedItem


class MainIngredient(GenderizedItem):
    kind = 'main_ingredient'

    def str_in_ingredients_list(self):
        parts = []
        if self._picked['value']:
            parts.append(self._picked['value'])
        if self._picked['unite']:
            parts.append(self._picked['unite'])
        parts.append(self.name)
        return ' '.join([str(part) for part in parts]).replace("' ", "'")

    def str_in_title(self):
        # FIXME: bad genderization
        left = self._genderize(
            {'a l\'': {'1st_voyel': True}},
            {'au': {'gender': 'male', '1st_voyel': False}},
            {'a la': {'gender': 'female'}},
            {'aux': {'quantity': 'multiple'}},
        )
        # FIXME: too much spaces
        return '{} {}'.format(left, self.name)

    def pick_some(self):
        rand = random.randrange(3)

        value = None
        unite = None

        if rand == 0:
            value = random.randrange(1, 51) * 10
            unite = self._genderize(
                {'gramme de': {'value': 1, '1st_voyel': False}},
                {'gramme d\'': {'value': 1, '1st_voyel': True}},
                {'grammes de': {'1st_voyel': False}},
                {'grammes d\'': {'1st_voyel': True}},
                value=value,
            )

        elif rand == 1:
            value = random.randrange(1, 6) + 1
            unite = self._genderize(
                {'tranche de': {'value': 1, '1st_voyel': False}},
                {'tranche d\'': {'value': 1, '1st_voyel': True}},
                {'tranches de': {'1st_voyel': False}},
                {'tranches d\'': {'1st_voyel': True}},
                value=value,
            )

        elif rand == 2:
            value = None
            options = [
                {'un bon gros': {'gender': 'male', 'quantity': 'single'}},
                {'une assez grosse': {'gender': 'female',
                                      'quantity': 'single'}},
                {'plusieurs gros': {'gender': 'male', 'quantity': 'multiple'}},
            ]

            for beginning in ('une quantite suffisante', 'pas mal',
                              'quelques morceaux', 'un bon paquet', 'beaucoup'):
                for ending, constraints in {
                        'de': {'1st_voyel': False},
                        'd\'': {'1st_voyel': True},
                }.items():
                    key = '{} {}'.format(beginning, ending)
                    option = {}
                    option[key] = constraints
                    options.append(option)

            unite = self._genderize(*options, shuffle=True)

        self._picked['value'] = value
        self._picked['unite'] = unite


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
