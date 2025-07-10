from config.yml_loader import get_config


def read():
    try:
        with open(get_config().path.tagPath, "r") as f:
            return f.read().strip()
    except FileNotFoundError:
        return "v0.0.0"
