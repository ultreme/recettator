# -*- coding: utf-8 -*-

import sys

from recettator import Recettator


def recettator_cli():
    seed = None
    if len(sys.argv) > 1:
        seed = sys.argv[1]

    recettator = Recettator(seed=seed)

    print('#{} - {}'.format(recettator.seed, recettator.title))
    print(len(recettator.title) * '=')
    print('')

    for k, v in recettator.infos.items():
        print('{}: {}'.format(k, v))
    print('')

    print('Ingredients')
    print('-----------')
    for category, ingredients in recettator.ingredients.items():
        for ingredient in ingredients:
            print('- {} {}'.format(
                ingredient.get('unite', '...'),
                ingredient['kind']['name'],
            ))
    print('')

    print('How-to')
    print('-------')
    print(recettator.howto)

if __name__ == '__main__':
    recettator_cli()
