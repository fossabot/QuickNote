import shutil
import zipfile
from pathlib import Path


def _extract_zip(zip_path: Path, extract_to: Path):
    with zipfile.ZipFile(zip_path, 'r') as zip_ref:
        for member in zip_ref.namelist():
            target_path = extract_to / member

            if member.endswith('/'):
                if target_path.exists():
                    shutil.rmtree(target_path)
                target_path.mkdir(parents=True, exist_ok=True)
            else:
                if target_path.exists():
                    target_path.unlink()
                target_path.parent.mkdir(parents=True, exist_ok=True)
                with zip_ref.open(member) as source, target_path.open('wb') as target:
                    shutil.copyfileobj(source, target)
