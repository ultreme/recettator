# -*- coding: utf-8 -*-


from .item import __all_items__ as items

# imports files so items can register by themselves in __all_items__
from .main_ingredients import all_items as main_ingredients_items
from .secondary_ingredients import all_items as secondary_ingredients_items
from .seasonings import all_items as seasonings_items


class ItemGroup(object):
    def __init__(self, items):
        self.availables = items


def all_items():
    items = []
    items += main_ingredients_items()
    items += secondary_ingredients_items()
    items += seasonings_items()
    return ItemGroup(items)
