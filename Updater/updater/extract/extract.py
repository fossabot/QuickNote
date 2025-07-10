from pathlib import Path

from .tar_gz import _extract_tar
from .zip import _extract_zip


def extract_and_replace(archive_path: Path, extract_to: Path):
    if not archive_path.exists():
        raise FileNotFoundError(f"Archive not found: {archive_path}")

    if not extract_to.exists():
        extract_to.mkdir(parents=True)

    if archive_path.suffix == '.zip':
        _extract_zip(archive_path, extract_to)
    elif archive_path.suffixes[-2:] == ['.tar', '.gz'] or archive_path.suffix == '.tgz':
        _extract_tar(archive_path, extract_to)
    else:
        raise ValueError(f"Unsupported archive format: {archive_path.suffix}")
