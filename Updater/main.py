import json
import time

from Updater.config.yml_loader import get_config
from Updater.updater.runner.process import find_processes_by_path, try_terminate
from Updater.updater.tag.reader import read
from Updater.updater.tag.writer import write
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

        procs = find_processes_by_path(get_config().path.processPath)
        if not procs:
            print("No matching processes found.")
        for proc in procs:
            try_terminate(proc)

        write(tag)

    print("=" * 80)
    time.sleep(INTERVAL)

if __name__ == "__main__":
    main()
