import esper
from enum import Enum
import pyxel

from .components import Position, Renderable


class GamePhase(Enum):
    INIT = 0
    UPDATE = 1
    DRAW = 2


class RenderProcessor(esper.Processor):

    def process(self, phase: GamePhase):
        if phase != GamePhase.DRAW:
            return

        for ent, (pos, render) in self.world.get_components(Position, Renderable):
            pyxel.image(0).blt(pos.x, pos.y, render.image_bank,
                               render.u, render.v, render.w, render.h)
