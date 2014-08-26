# -*- coding: utf-8 -*-

import random
from ..utils import parts_to_string


def genderization(options, constraints):
    for option in options:
        checks = option.values()[0]
        matches = True
        for k, v in checks.items():
            if constraints[k] != v:
                matches = False
                break
        if matches:
            return option.keys()[0]


class Item(object):
    name = None
    kind = None

    def __init__(self):
        self._picked = {}
        self.pick_some()

    def pick_some(self):
        pass

    def str_in_ingredients_list(self):
        return []

    def str_in_title(self, left):
        return []

    @property
    def attrs(self):
        attrs = {
            'kind': self.kind,
            'name': self.kind,
            '1st_voyel': self._first_voyel,
        }
        return attrs

    def __repr__(self):
        return "<{}:{}>".format(self.kind, type(self).__name__)
        """
        More detail:

        return "<{} {}>".format(
            type(self).__name__,
            ', '.join(['{}={}'.format(k, v) for k, v in self.attrs.items()])
        )
        """

    @property
    def _first_voyel(self):
        return self.name[0] in ('a', 'e', 'i', 'o', 'u', 'y')

    @property
    def people(self):
        return 0


class GenderizedItem(Item):
    gender = 'any'
    quantity = 'any'

    @property
    def attrs(self):
        attrs = super(GenderizedItem, self).attrs
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs

    def _genderize(self, *args, **kwargs):
        shuffle = 'shuffle' in kwargs and kwargs['shuffle']

        options = list(args)
        if shuffle:
            random.shuffle(options)

        constraints = self.attrs
        for k, v in kwargs.items():

            if k in ('shuffle',):
                continue

            constraints[k] = v

        return genderization(options, constraints)
