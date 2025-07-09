from functools import lru_cache
from pathlib import Path

import yaml

from Updater.config.model import Config


@lru_cache()
def get_config() -> Config:
    with open(Path(__file__).parent.parent / "./data/config.yml", "r") as f:
        return Config(**yaml.safe_load(f) or {})
