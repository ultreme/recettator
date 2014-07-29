#!/usr/bin/env python


from .utils import pick_random, pick_random_dict


class Recettator:
    base = {
        'female:multiple': [
            'tranches', 'galettes', 'lasagnes', 'chips', 'cereales',
            'escalopes',
        ],
        'female:single': [
            'truffe', 'mousse', 'buche', 'puree', 'ratatouille', 'soupe',
        ],
        'male:multiple': [
            'petits gateaux', 'rochers', 'endives', 'pates', 'patates',
            'champignons',
        ],
        'male:single': [
            'parfait', 'civet', 'gateau', 'gratin', 'kebab', 'rouleau',
        ],
    }

    ingredients = {
        'female:multiple': ['noisettes',],
        'female:single': [],
        'male:multiple': [],
        'male:single': ['gui', 'houx', 'ble', 'lierre',],
    }

    method = [
        '', 'a la juive' 'a la mexicaine', 'methode traditionnelle',
        'a l\'ancienne', 'comme a la maison', 'recette originale', 'perso',
        'special grandes occasions', 'du chef', 'a la provencale',
        'recette de ma grand-mere', 'special pizzaiolo', 'premium\'s',
        'version XXL',
    ]


    def __init__(self):
        self._data = None

    def create(self, seed=None):
        if seed:
            # FIXME: initialize random with seed
            raise NotImplementedError()
        self._data = {
            'title': '',
            'people': 0,
            'ingredients': {},
            'howto': [],
        }
        self._data['people'] = 42
        title_base = pick_random_dict(self.base)
        title_method = pick_random(self.method)
        title = title_base['value'] + ' '
        title += title_method
        self._data['title'] = title
        print(title_base)

    def _create_if_not_exists(self):
        if not self._data:
            self.create()

    @property
    def title(self):
        self._create_if_not_exists()
        return self._data['title']

    @property
    def infos(self):
        self._create_if_not_exists()
        return {
            'people': self._data['people'],
        }

    @property
    def ingredients(self):
        self._create_if_not_exists()
        return self._data['ingredients']

    @property
    def howto(self):
        self._create_if_not_exists()
        return self._data['howto']
