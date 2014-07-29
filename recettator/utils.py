# -*- coding: utf-8 -*-


import random


def pick_random(lst):
    return lst[int(random.random() * len(lst))]


def pick_random_dict(dct):
    values = []
    for value in dct.values():
        values += value

    max_len = 0
    for k in dct.keys():
        max_len = max(max_len, len(k.split(':')))

    selected = pick_random(values)
    for k, v in dct.items():
        if selected in v:
            group = k.split(':')
            group += [None] * max_len
            group = group[:max_len]
            return {
                'group': group,
                'value': selected,
            }
