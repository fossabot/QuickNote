from config.yml_loader import get_config


def write(tag: str):
    with open(get_config().path.tagPath, "w") as f:
        f.write(tag)
