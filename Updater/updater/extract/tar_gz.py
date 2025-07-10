import shutil
import tarfile
from pathlib import Path


def _extract_tar(tar_path: Path, extract_to: Path):
    with tarfile.open(tar_path, 'r:*') as tar:
        for member in tar.getmembers():
            target_path = extract_to / member.name

            if member.isdir():
                if target_path.exists():
                    shutil.rmtree(target_path)
                target_path.mkdir(parents=True, exist_ok=True)
            else:
                if target_path.exists():
                    target_path.unlink()
                target_path.parent.mkdir(parents=True, exist_ok=True)
                with tar.extractfile(member) as source, target_path.open('wb') as target:
                    shutil.copyfileobj(source, target)
