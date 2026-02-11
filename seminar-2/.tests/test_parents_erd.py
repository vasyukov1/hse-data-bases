import pytest
import re
import os

def read_puml_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        return file.read()

@pytest.fixture
def puml_content():
    return read_puml_file("./src/relationships_parents.puml")

def test_parents_entity_count(puml_content):
    """Проверка количества сущностей в диаграмме родственных связей"""
    entities = re.findall(r'entity\s+[""]?\w+[""]?', puml_content, re.IGNORECASE)
    # Допустимо либо 1 сущность (например, Person), либо 2 сущности (например, Man и Woman)
    entity_count = len(entities)
    assert entity_count in [1, 2], f"Ожидается либо 1 сущность, либо 2 сущности, найдено {entity_count}"

def test_parents_relationship_types(puml_content):
    """Проверка связей в зависимости от количества сущностей в диаграмме родственных связей"""
    # Универсальный паттерн для one-to-many связей (без разницы, обязательность или опциональность)
    one_to_many_pattern = r'(?:\|\||o\|)--(?:o\{|\|\{)'
    entities = re.findall(r'entity\s+[""]?\w+[""]?', puml_content, re.IGNORECASE)
    entity_count = len(entities)

    if entity_count == 2:
        connections = len(re.findall(one_to_many_pattern, puml_content))
        assert connections == 4, (
            f"Для варианта с двумя сущностями ожидается 4 связи one-to-many, найдено {connections}"
        )
    elif entity_count == 1:
        connections = len(re.findall(one_to_many_pattern, puml_content))
        assert connections == 2, (
            f"Для варианта с одной сущностью ожидается 2 связи one-to-many, найдено {connections}"
        )
    else:
        pytest.skip("Неверное количество сущностей для проверки родственных связей")