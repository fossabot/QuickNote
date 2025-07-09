import platform


def get_goos_goarch():
    machine = platform.machine().lower()
    if machine in ("x86_64", "amd64"):
        goarch = "amd64"
    elif machine in ("aarch64", "arm64"):
        goarch = "arm64"
    elif machine in ("i386", "i686", "x86"):
        goarch = "386"
    elif machine.startswith("arm"):
        goarch = "arm"
    else:
        goarch = machine

    return platform.system().lower(), goarch


def get_package_name_from_current_machine():
    goos, goarch = get_goos_goarch()

    ext = ".zip" if goos == "windows" else ".tar.gz"
    return f"QuickNote_{goos}_{goarch}{ext}"
