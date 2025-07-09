from pydantic import BaseModel


class PathConfig(BaseModel):
    workPath: str = "../"
    processName: str = "QuickNote"
    tagPath: str = "./tags"

class ProxyConfig(BaseModel):
    url: str = ""

class Config(BaseModel):
    path: PathConfig = PathConfig()
    proxy: ProxyConfig = ProxyConfig()