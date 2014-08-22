# -*- coding: utf-8 -*-

from .item import Item


class Seasoning(Item):
    kind = 'seasoning'


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
