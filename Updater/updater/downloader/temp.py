import atexit
import tempfile
from datetime import datetime
from pathlib import Path
from urllib.parse import urlparse

import requests


def download_to_temp(url: str) -> Path:
    original_name = Path(urlparse(url).path).name

    filename = f"{Path(original_name).stem}_{datetime.now().strftime("%Y%m%d_%H%M%S")}{Path(original_name).suffix}"

    temp_path = Path(tempfile.gettempdir()) / filename

    with requests.get(url, stream=True) as r:
        r.raise_for_status()
        with open(temp_path, 'wb') as f:
            for chunk in r.iter_content(chunk_size=1024):
                if chunk:
                    f.write(chunk)

    atexit.register(lambda: temp_path.exists() and temp_path.unlink())

    return temp_path
