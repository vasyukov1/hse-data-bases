import pytest
import re
import os

def read_puml_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as file:
        return file.read()

@pytest.fixture
def puml_content():
    return read_puml_file("./src/relationships_match.puml")

def test_match_entity_count(puml_content):
    """Проверка количества сущностей в диаграмме матчей"""
    entities = re.findall(r'entity\s+[""]?\w+[""]?', puml_content, re.IGNORECASE)
    assert len(entities) == 3, \
        f"Ожидается 3 сущности, найдено {len(entities)}"

def test_match_relationship_types(puml_content):
    """Проверка количества различных типов связей"""

    one_to_many_pattern = r'(?:\|\||o\|)--(?:o\{|\|\{)'
    one_to_many = len(re.findall(one_to_many_pattern, puml_content))
    assert one_to_many == 3, f"Ожидается 5 связей один-ко-многим, найдено {one_to_many}"

    many_to_many = len(re.findall(r'\}o--o\{', puml_content))
    assert many_to_many == 0, f"Ожидается 1 связь многие-ко-многим, найдено {many_to_many}"

    one_to_one_pattern = r'(?:\|\||o\||\|o)--(?:\|\||o\||\|o)'
    one_to_one = len(re.findall(one_to_one_pattern, puml_content))
    assert one_to_one == 0, f"Ожидается 0 связей один-к-одному, найдено {one_to_one}"