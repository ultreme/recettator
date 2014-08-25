# -*- coding: utf-8 -*-

import random

from .item import Item


class Seasoning(Item):
    kind = 'seasoning'
    gender = 'any'
    quantity = 'any'

    def str_in_ingredients_list(self):
        parts = []
        if self._picked['value']:
            parts.append(self._picked['value'])
        if self._picked['unite']:
            parts.append(self._picked['unite'])
        parts.append(self.name)
        return ' '.join([str(part) for part in parts]).replace("' ", "'")

    def pick_some(self):
        value = None
        unite = None

        value = random.randrange(1, 31) / 10.0
        unite_base = 'litre'
        if value < 1:
            unite_base = 'centilitre'
            value *= 100
        unite = self._genderize(
            {'{} de'.format(unite_base): {'value': 1, '1st_voyel': False}},
            {'{} d\''.format(unite_base): {'value': 1, '1st_voyel': True}},
            {'{}s de'.format(unite_base): {'1st_voyel': False}},
            {'{}s d\''.format(unite_base): {'1st_voyel': True}},
            value=value,
        )
        self._picked['value'] = value
        self._picked['unite'] = unite

    @property
    def attrs(self):
        attrs = super(Seasoning, self).attrs
        attrs['gender'] = self.gender
        attrs['quantity'] = self.quantity
        return attrs


class Tisane(Seasoning):
    name = 'tisane'


class ExtraitDeFleurDOranger(Seasoning):
    name = 'extrait de fleur d\'oranger'


class Viandox(Seasoning):
    name = 'viandox'


class BiereDeNoel(Seasoning):
    name = 'biere de noel'


class VinRouge(Seasoning):
    name = 'vin rouge'


class VinBlanc(Seasoning):
    name = 'vin blanc'


class HuileDArachide(Seasoning):
    name = 'huile d\'arachide'


class SauceDHuitre(Seasoning):
    name = 'sauce d\'huitre'


class CremeFraiche(Seasoning):
    name = 'creme fraiche'


class Creme(Seasoning):
    name = 'creme'


class LiqueurDeRaisin(Seasoning):
    name = 'liqueur de raisin'


class GrandMarnier(Seasoning):
    name = 'grand marnier'


class Lait(Seasoning):
    name = 'lait'


class LaitFermente(Seasoning):
    name = 'lait fermente'


class HuileDOlive(Seasoning):
    name = 'huile d\'olive'


class VinaigreDeRiz(Seasoning):
    name = 'vinaigre de riz'


class VinaigreDeCidre(Seasoning):
    name = 'vinaigre de cidre'


class VinaigreDeVin(Seasoning):
    name = 'vinaigre de vin'


class JusDeCitron(Seasoning):
    name = 'jus de citron'


def all_items():
    return (
        Tisane, ExtraitDeFleurDOranger, Viandox, BiereDeNoel, VinRouge,
        VinBlanc, HuileDArachide, SauceDHuitre, CremeFraiche, Creme,
        LiqueurDeRaisin, GrandMarnier, Lait, LaitFermente, HuileDOlive,
        VinaigreDeRiz, VinaigreDeCidre, VinaigreDeVin, JusDeCitron,
    )
