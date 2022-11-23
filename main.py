import pyxel
import esper

ASSETS = 'assets.pyxres'


def init():
    pyxel.init(160, 120)
    pyxel.load(ASSETS)


def update():
    if pyxel.btnp(pyxel.KEY_Q):
        pyxel.quit()


def draw():
    pyxel.cls(0)
    pyxel.rect(10, 10, 20, 20, 11)
    pyxel.text(40, 40, "Hello World", 1)


if __name__ == '__main__':
    init()
    pyxel.run(update, draw)
