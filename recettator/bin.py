# -*- coding: utf-8 -*-

import sys

from recettator import Recettator


def recettator_cli():
    seed = None
    if len(sys.argv) > 1:
        seed = sys.argv[1]

    recettator = Recettator(seed=seed)

    print(recettator.title)
    print(len(recettator.title) * '=')
    print('')

    print(recettator.infos)
    print('')

    print('Ingredients')
    print('-----------')
    print(recettator.ingredients)
    print('')

    print('How-to')
    print('-------')
    print(recettator.howto)

if __name__ == '__main__':
    recettator_cli()
