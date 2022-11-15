from .world import World
from .component import Component
from dataclasses import dataclass


@dataclass
class TestComponentA(Component):
    a: int


@dataclass
class TestComponentB(Component):
    b: str


def test_add_entity(benchmark):
    world = World()

    entity = world.add_entity()

    assert world.has_entity(entity)

    benchmark(world.add_entity)


def test_add_entity_with_components(benchmark):
    world = World()

    entity = world.add_entity(TestComponentA(1), TestComponentB("lorem ipsum"))

    assert world.has_entity(entity)
    assert world.entity_has_component(entity, TestComponentA)
    assert world.entity_has_component(entity, TestComponentB)

    benchmark(world.add_entity, TestComponentA(1), TestComponentB("lorem ipsum"))
