#!/usr/bin/env python

import random
import pkg_resources
import os

from .utils import pick_random, pick_random_dict
from .custom_csv import CustomCSV


class Recettator:
    """ Recettator class. """

    def __init__(self, seed=None):
        self._data = None
        if not seed:
            seed = random.randrange(10000)

        self.seed = seed
        self.dbs = {}

    def db_pick(self, kind, **kwargs):
        if not kind in self.dbs:
            dirpath = pkg_resources.resource_filename('recettator', 'db')
            path = os.path.join(dirpath, '{}.csv'.format(kind))
            self.dbs[kind] = CustomCSV(path, shuffle=True)

        db = self.dbs[kind]
        return db.pick(**kwargs)

    def create(self):
        random.seed(self.seed)

        self._data = {
            'amount': {
                'main_ingredients': random.randrange(0, 4) - 1,
                'secondary_ingredients': random.randrange(0, 6) - 1,
                'seasonings': random.randrange(0, 7) - 1,
                'methods': random.randrange(0, 5) - 1,
            },
            'ingredients': {},
            'howto': [],
            'recette': self.db_pick('recettes'),
            'method': self.db_pick('methods'),
            'seed': self.seed,
        }

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
                        rand = random.randrange(3)
                        # FIXME: gender
                        if rand == 0:
                            ingredient['quantity'] = {
                                'value': unite,
                                'unite': 'grammes de',
                            }
                        elif rand == 1:
                            value = random.randrange(0, 5) + 1
                            unite = 'tranches de'
                            ingredient['quantity'] = {
                                'value': value,
                                'unite': unite,
                            }
                        elif rand == 2:
                            ingredient['quantity'] = {
                                'value': 1,
                                'unite': 'bon gros',
                                }
                    if 'quantity' in ingredient:
                        ingredient['quantity']['str'] = '{} {}'.format(
                            ingredient['quantity']['value'],
                            ingredient['quantity']['unite'],
                        )
                    self._data['ingredients'][k].append(ingredient)

    def _create_if_not_exists(self):
        if not self._data:
            self.create()

    def __getattr__(self, name):
        self._create_if_not_exists()
        if name in self._data:
            return self._data[name]
        raise KeyError('Unknown key: {}'.format(name))

    @property
    def title(self):
        title_parts = []
        title_parts.append(self.recette['name'])
        title_parts.append(self.method['name'])
        return ' '.join(title_parts)

    @property
    def _people(self):
        amounts = self.amount
        stuff_amount = sum(amounts.values())
        return max(int(stuff_amount / 2), 1)

    @property
    def people(self):
        people = self._people
        parts = ['Pour']
        if random.randrange(0, 100) < 20:
            parts.append('environ')
        parts.append(people)
        if random.randrange(0, 100) < 20:
            parts.append('a')
            parts.append(people + random.randrange(1, 4))
        parts.append('personne(s)')
        return ' '.join([str(part) for part in parts])

    @property
    def infos(self):
        return {
            'people': self.people,
        }
