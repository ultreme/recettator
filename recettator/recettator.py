#!/usr/bin/env python

from collections import OrderedDict
from math import ceil
from random import randrange, seed as set_seed
import os
import pkg_resources
import sys

from .custom_csv import CustomCSV
from .items import all_items
from .utils import parts_to_string


class Recettator:
    """ Recettator class. """

    def __init__(self, seed=None):
        self._data = None
        self._items = []
        self._db = None
        self._amounts = {}

        # seed
        if not seed:
            seed = randrange(10000)
        self.seed = seed
        set_seed(seed)

    @property
    def items(self):
        if not len(self._items):
            self._db = all_items()
            self._amounts = OrderedDict([
                ('recette', 1),
                ('main_ingredient', randrange(5) - 1),
                ('secondary_ingredient', randrange(5) - 1),
                ('seasoning', randrange(6) - 1),
                ('method', randrange(5) - 1),
            ])

            for k, v in self._amounts.items():
                for i in xrange(max(v, 0)):
                    item = self._db.pick_random(kind=k)
                    if item:
                        self._items.append(item)

        return self._items

    def create(self):
        # Picking ingredients
        for k, v in self._data['amount'].items():
            # v = max(0, v)
            self._data['amount'][k] = v = max(0, v)
            if not k in self._data['ingredients']:
                self._data['ingredients'][k] = []

                for i in xrange(v):
                    ingredient = {
                        'kind': self.db_pick(k)
                    }
                    if k == 'main_ingredients':
                        rand = randrange(3)
                        # FIXME: gender
                        if rand == 0:
                            ingredient['quantity'] = {
                                'value': randrange(1, 51) * 10,
                                'unite': 'grammes de',
                            }
                        elif rand == 1:
                            value = randrange(5) + 1
                            unite = 'tranches de'
                            ingredient['quantity'] = {
                                'value': value,
                                'unite': unite,
                            }
                        elif rand == 2:
                            ingredient['quantity'] = {
                                'value': None,
                                'unite': 'un bon gros',
                                }
                    elif k == 'secondary_ingredients':
                        rand = randrange(5)
                        # FIXME: check kind instead of random + randomize
                        if rand == 0:
                            ingredient['quantity'] = {
                                'value': None,
                                'unite': 'des',
                            }
                        elif rand == 1:
                            ingredient['quantity'] = {
                                'value': randrange(1, 51) * 10,
                                'unite': 'grammes de',
                            }
                        elif rand == 2:
                            value = randrange(1, 21)
                            # if single -> value = 1
                            ingredient['quantity'] = {
                                'value': value,
                                'unite': None,
                            }
                        elif rand == 3:
                            ingredient['quantity'] = {
                                'value': None,
                                'unite': 'un zeste de',
                            }
                        elif rand == 4:
                            ingredient['quantity'] = {
                                'value': None,
                                'unite': 'une poignee de',
                            }
                        if randrange(100) < 10:
                            ingredient['attribute'] = 'frais'

                    elif k == 'seasonings':
                        quantity = (float(randrange(31)) + 1) / 10
                        if quantity > 1:
                            unite = 'litres de'
                        else:
                            quantity *= 100
                            unite = 'centilitres de'
                        ingredient['quantity'] = {
                            'value': quantity,
                            'unite': unite,
                        }

                    if 'quantity' in ingredient:
                        quantity = ingredient['quantity']
                        if quantity['value'] and \
                           quantity['unite']:
                            string =  '{} {}'.format(
                                quantity['value'],
                                quantity['unite'],
                            )
                        elif quantity['value']:
                            string = str(quantity['value'])
                        elif quantity['unite']:
                            string = str(quantity['unite'])
                        quantity['str'] = string
                    self._data['ingredients'][k].append(ingredient)

    """
    def get_title(self):
        for i in xrange(42):
            recette = self.pick_random(kind='recette', recycle_item=True)
            main_ingredient = self.pick_random(kind='main_ingredient',
                                                recycle_item=True)
            secondary_ingredient = self.pick_random(
                kind='secondary_ingredient',
                recycle_item=True
            )
            seasoning = self.pick_random(kind='seasoning',
                                          recycle_item=True)
            method = self.pick_random(kind='method',
                                       recycle_item=True)

            parts = []
            parts += recette.str_in_title()
            parts += main_ingredient.str_in_title(recette)
            parts += secondary_ingredient.str_in_title(main_ingredient)

            print(parts_to_string(parts))
    """

    @property
    def title(self):
        title = []
        left = None
        for item in self.items:
            title += item.str_in_title(left)
            left = item
        title = parts_to_string(title)
        title = title.capitalize()
        return title

    @property
    def ingredients(self):
        ingredients = []
        for item in self.items:
            ingredient = item.str_in_ingredients_list()
            if ingredient and len(ingredient):
                ingredient = parts_to_string(ingredient)
                ingredients.append(ingredient)
        return ingredients

    @property
    def _people(self):
        people = 0
        for item in self.items:
            people += item.people
        return int(ceil(max(people, 1)))

    @property
    def people(self):
        people = self._people
        parts = ['Pour']
        if randrange(100) < 20:
            parts.append('environ')
        parts.append(people)
        if randrange(100) < 20:
            parts.append('a')
            parts.append(people + randrange(1, 4))
        parts.append('personne(s)')
        return ' '.join([str(part) for part in parts])

    """
    @property
    def infos(self):
        return {
            'people': self.people,
        }
    """
