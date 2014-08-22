# -*- coding: utf-8 -*-


class Item(object):
    name = None
    kind = None

    def __init__(self):
        pass

    @property
    def attrs(self):
        attrs = {
            'kind': self.kind,
            'name': self.kind,
        }
        return attrs

    def __repr__(self):
        return "<{} {}>".format(
            type(self).__name__,
            ', '.join(['{}={}'.format(k, v) for k, v in self.attrs.items()])
        )

    @property
    def begins_with_voyel(self):
        return self.name[0] in ('a', 'e', 'i', 'o', 'u', 'y')


    def ingredient_list_str(self):
        return self.name
