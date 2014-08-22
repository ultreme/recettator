#!/usr/bin/env python

from random import randrange, seed
import pkg_resources
import os

from .utils import pick_random, pick_random_dict
from .custom_csv import CustomCSV
from .items import all_items


class Recettator:
    """ Recettator class. """

    def __init__(self, seed=None):
        self._data = None
        if not seed:
            seed = randrange(10000)

        items = all_items()
        for item in items.availables:
            instance = item()
            print('{} - {}'.format(instance.ingredient_list_str(), instance))
        import sys
        sys.exit(0)


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
        seed(self.seed)

        self._data = {
            'amount': {
                'main_ingredients': randrange(4) - 1,
                'secondary_ingredients': randrange(6) - 1,
                'seasonings': randrange(7) - 1,
                'methods': randrange(5) - 1,
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
        if randrange(100) < 20:
            parts.append('environ')
        parts.append(people)
        if randrange(100) < 20:
            parts.append('a')
            parts.append(people + randrange(1, 4))
        parts.append('personne(s)')
        return ' '.join([str(part) for part in parts])

    @property
    def infos(self):
        return {
            'people': self.people,
        }
