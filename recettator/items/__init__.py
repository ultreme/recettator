# -*- coding: utf-8 -*-

import random
import re

from ..utils import parts_to_string

from .main_ingredients import all_items as main_ingredients_items
from .secondary_ingredients import all_items as secondary_ingredients_items
from .seasonings import all_items as seasonings_items
from .methods import all_items as methods_items
from .recettes import all_items as recettes_items


class ItemGroup(object):
    def __init__(self, items, shuffle=True):
        self.availables = items
        if shuffle:
            random.shuffle(self.availables)

    def pick_random(self, recycle_item=False, **kwargs):
        for item in self.availables:
            found = True
            for k, v in kwargs.items():
                if not hasattr(item, k):
                    found = False
                    break
                item_value = getattr(item, k)
                if not item_value:
                    break

                # Prepare for matching
                if v == 'any':
                    v = '.*'
                if item_value == 'any':
                    item_value = '.*'
                if v[0] != '^':
                    v = '^{}$'.format(v)
                if item_value[0] != '^':
                    item_value = '^{}$'.format(item_value)

                if item_value != v and \
                   not re.match(item_value, v) and \
                   not re.match(v, item_value):
                    found = False
                    break
            if found:
                self.availables.remove(item)
                if recycle_item:
                    self.availables.append(item)
                return item()
        return None

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


def all_items():
    items = []
    items += main_ingredients_items()
    items += secondary_ingredients_items()
    items += seasonings_items()
    items += methods_items()
    items += recettes_items()
    return ItemGroup(items)
