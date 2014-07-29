#!/usr/bin/env python


class Recettator:
    def __init__(self):
        self._data = None

    def create(self):
        self._data = {}
        self._data['title'] = 'title'
        self._data['infos'] = 'infos'
        self._data['ingredients'] = 'ingredients'
        self._data['howto'] = 'howto'

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
        return self._data['infos']

    @property
    def ingredients(self):
        self._create_if_not_exists()
        return self._data['ingredients']

    @property
    def howto(self):
        self._create_if_not_exists()
        return self._data['howto']
