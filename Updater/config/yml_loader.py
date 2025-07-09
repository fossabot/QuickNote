from functools import lru_cache
from pathlib import Path

import yaml
from pydantic import BaseModel


class PathConfig(BaseModel):
    processPath: str = "../QuickNote"
    tagPath: str = "./tags"

class ProxyConfig(BaseModel):
    url: str = ""

class Config(BaseModel):
    path: PathConfig = PathConfig()
    proxy: ProxyConfig = ProxyConfig()

@lru_cache()
def get_config() -> Config:
    with open(Path(__file__).parent.parent / "./data/config.yml", "r") as f:
        return Config(**yaml.safe_load(f) or {})