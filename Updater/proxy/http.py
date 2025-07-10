from functools import lru_cache

from config.yml_loader import get_config


@lru_cache()
def get_proxies():
    proxy_url = get_config().proxy.url
    if proxy_url == "env":
        return None
    return {"http": proxy_url, "https": proxy_url} if proxy_url else {"http": None, "https": None}
