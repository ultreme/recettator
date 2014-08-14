#!/usr/bin/env python

import random

from .utils import pick_random, pick_random_dict
from .custom_csv import CustomCSV


class Recettator:
    """ Recettator class. """

    def __init__(self, seed=None):
        self._data = None
        self.seed = seed
        self.dbs = {}

    def db_pick(self, kind, **kwargs):
        if not kind in self.dbs:
            self.dbs[kind] = CustomCSV('db/{}.csv'.format(kind),
                                       shuffle=True)

        db = self.dbs[kind]
        return db.pick(**kwargs)

    def create(self, seed=None):
        if not seed:
            seed = self.seed
        if seed:
            random.seed(self.seed)

        self._data = {
            'ingredients': {},
            'howto': [],
            'recette': self.db_pick('recettes'),
            'method': self.db_pick('methods'),
            'people': random.randrange(1, 10)
        }

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
    def infos(self):
        return {
            'people': self.people,
        }
