# -*- coding: utf-8 -*-


from .item import __all_items__ as items

# imports files so items can register by themselves in __all_items__
from .main_ingredients import all_items as main_ingredients_items


class ItemGroup(object):
    def __init__(self, items):
        self.items = items


def main_ingredients():
    items = main_ingredients_items()
    group = ItemGroup(items)

    for item in group.items:
        instance = item()
        print('{} - {}'.format(instance.ingredient_list_str(), instance))
    import sys
    sys.exit(0)

    return group
