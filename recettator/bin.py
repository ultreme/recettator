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
    for category in ['main_ingredients', 'secondary_ingredients', 'seasonings']:
        ingredients = recettator.ingredients[category]
        for ingredient in ingredients:
            print('- {} {} {}'.format(
                ingredient.get('quantity', {}).get('str', ''),
                ingredient['kind']['name'],
                ingredient.get('attribute', ''),
            ))
    print('')

    print('How-to')
    print('-------')
    print(recettator.howto)

    print('Debug')
    print('-----')
    for k, v in recettator.amount.items():
        print('{} amount: {}'.format(k, v))

if __name__ == '__main__':
    recettator_cli()
