from invoke import Collection

from . import docs
from . import image
from . import test
from .utils import set_root

ns = Collection()
ns.add_collection(Collection.from_module(image))
ns.add_collection(Collection.from_module(docs))
ns.add_collection(Collection.from_module(test))

set_root()