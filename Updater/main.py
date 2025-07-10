import json
import time
from pathlib import Path

from Updater.config.yml_loader import get_config
from Updater.updater.assets.name import get_package_name_from_current_machine
from Updater.updater.downloader.temp import download_to_temp
from Updater.updater.extract.extract import extract_and_replace
from Updater.updater.runner.process import find_processes_by_path, try_terminate
from Updater.updater.tag.reader import read
from fetch.github_release_api import fetch_latest_release
from proxy.http import get_proxies

OWNER = "Sn0wo2"
REPO = "QuickNote"
INTERVAL = 5.0


def main():
    result = fetch_latest_release(OWNER, REPO, get_proxies())
    print(json.dumps(
        result,
        indent=2,
        ensure_ascii=False
    ))
    tag = result.get("tag_name")
    if read() != tag:
        print("=" * 80)
        print(f"New tag: {tag}")

        procs = find_processes_by_path(Path(get_config().path.workPath, get_config().path.processName))
        if not procs:
            print("No matching processes found.")
        for proc in procs:
            try_terminate(proc)

        download_url = None

        for asset in result.get("assets"):
            if asset.get("name") == get_package_name_from_current_machine():
                download_url = asset.get("browser_download_url")
                break

        if download_url is None:
            print(
                f"Could not find asset for {get_package_name_from_current_machine()}\nPlease update manually or send an issue.")
            return

        downloaded_path = download_to_temp(download_url)

        extract_and_replace(downloaded_path, Path(get_config().path.workPath))

        print(Path(get_config().path.workPath))

        print(downloaded_path)

    print("=" * 80)
    time.sleep(INTERVAL)


if __name__ == "__main__":
    main()
