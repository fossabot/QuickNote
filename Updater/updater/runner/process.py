import os
import signal

import psutil


def find_processes_by_path(target_path):
    matches = []
    for proc in psutil.process_iter(['pid', 'exe', 'cmdline']):
        try:
            exe = proc.info['exe']
            if not exe and proc.info['cmdline']:
                exe = proc.info['cmdline'][0]
            if exe and os.path.abspath(exe) == os.path.abspath(target_path):
                matches.append(proc)
        except (psutil.NoSuchProcess, psutil.AccessDenied):
            continue
    return matches

def try_terminate(proc):
    print(f"Trying to gracefully terminate PID {proc.pid} ({proc.name()})...")
    try:
        if os.name == 'nt':
            os.kill(proc.pid, signal.CTRL_BREAK_EVENT)
        else:
            proc.send_signal(signal.SIGINT)
        proc.wait(timeout=5)
        print(f"Process {proc.pid} exited gracefully.")
    except (psutil.TimeoutExpired, psutil.NoSuchProcess):
        print(f"Process {proc.pid} did not exit. Killing...")
        proc.kill()