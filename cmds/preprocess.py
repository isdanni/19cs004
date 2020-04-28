import re
import os
from pathlib import Path

PRE_REGEX = r"test:(.+?[/\.]\d+)"

def get_net_loc(topo, source_path):
    Path(source_path).write_text(
        FILE(Path(topo).read_text())
    )

def check_logs(topo, source_path):
    Path(source_path).write_text(
	if not topo:
	    print("Jump...")
	    continue
        FILE(Path(topo).read_text())
    )

def main():
    topo = Path(os.environ['test_tk.json'])
    sr = PRE_REGEX + '*.json'
    get_net_loc(topo, sr)

if __name__ == '__main__':
    main()
