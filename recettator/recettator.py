#!/usr/bin/env python

from collections import OrderedDict
from math import ceil
from random import randrange, seed as set_seed
import os
import pkg_resources
import sys

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
                ('method', int(randrange(10) < 4)),
            ])

            for k, v in self._amounts.items():
                for i in xrange(max(v, 0)):
                    item = self._db.pick_random(kind=k)
                    if item:
                        self._items.append(item)

        return self._items

    @property
    def steps(self):
        steps = []
        for item in self.items:
            steps += item.steps
        steps.append('rassemblez tous les ingredients dans un grand plat et '
                     'consommez vite !')
        steps = [
            step.capitalize() for step in steps
        ]
        return steps

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
                ingredient = ingredient.capitalize()
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
