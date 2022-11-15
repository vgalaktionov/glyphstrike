from typing import Type
from .component import Component
from collections import defaultdict

AnyComponent = Type[Component] | Component | str


class World:
    entities: set[int] = set()
    last_entity = 0
    component_indices: dict[str, int] = defaultdict(lambda: 0)
    components: dict[str, dict[int, Component]] = defaultdict(dict)

    def __init__(self) -> None:
        pass

    def add_entity(self, *components: Component) -> int:
        self.last_entity += 1
        self.entities.add(self.last_entity)
        for component in components:
            self.components[self.component_name(component)][self.last_entity] = component
            self.component_indices[self.component_name(component)] |= 1 << self.last_entity
        return self.last_entity

    def has_entity(self, entity: int) -> bool:
        return entity in self.entities

    def remove_entity_component(self, entity: int, component_class: AnyComponent) -> None:
        self.component_indices[self.component_name(component_class)] &= ~(1 << entity)
        del self.components[self.component_name(component_class)][entity]

    def remove_entity(self, entity: int) -> None:
        self.entities.remove(entity)
        for component_name in self.components.keys():
            if self.entity_has_component(entity, component_name):
                self.remove_entity_component(entity, component_name)

    def entity_has_component(self, entity: int, component_class: AnyComponent) -> bool:
        return 1 == (self.component_indices[self.component_name(component_class)] >> entity) & 1

    def component_name(self, component_class: AnyComponent) -> str:
        match component_class:
            case str():
                return component_class
            case Component():
                return component_class.__class__.__name__
            case _:
                return component_class.__name__
