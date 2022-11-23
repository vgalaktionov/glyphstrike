from dataclasses import dataclass
from typing import Union


@dataclass
class Position:
    x: int
    y: int


@dataclass
class Renderable:
    image_bank: Union[1, 2, 3]
    u: int
    v: int
    w: int
    h: int

