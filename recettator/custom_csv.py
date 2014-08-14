# -*- coding: utf-8 -*-

import re
import random


class CustomCSV(object):
    """ CustomCSV class. """

    def __init__(self, path, sep=';', shuffle=False):
        self.entries = []
        with open(path) as fh:
            lines = fh.read().split('\n')
            header = lines[0].lstrip('#').strip().split(sep)
            for line in lines[1:]:
                if line.strip().startswith('#'):
                    continue
                if not len(line.strip()):
                    continue
                entry = line.strip().split(sep)
                self.entries.append(dict(zip(header, entry)))
        if shuffle:
            random.shuffle(self.entries)

    def pick(self, **kwargs):
        for item in self.entries:
            found = True
            for k, v in kwargs.items():
                item_value = item.get(k, None)
                if not item_value:
                    break
                if item_value != v and \
                   not re.match(item_value, v) and \
                   not re.match(v, item_value):
                    found = False
                    break
            if found:
                self.entries.remove(item)
                self.entries.append(item)
                return item
        return None
