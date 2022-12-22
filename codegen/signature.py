import json
import os
import sys

from web3 import Web3

current_dir = os.getcwd()
args = sys.argv
if len(args) < 3:
    print('usage: <abi_path> <output-path>')
    sys.exit(1)

path = current_dir + '/' + args[1]
out = current_dir + '/' + args[2]
res = []
with open(path) as file:
    abi = json.load(file)
    counts = {}
    for item in abi:
        if item['type'] == 'event':
            name = item['name']
            sig = name + '(' + ','.join(map(lambda x: x['type'], item['inputs'])) + ')'
            counts[name] = 0 if name not in counts else counts[name] + 1
            res.append({
                'name': name + str(counts[name] - 1) if counts[name] > 0 else name,
                'signature': sig,
                'hash': Web3.keccak(text=sig).hex(),
            })
with open(out, 'w') as f:
    f.write(json.dumps(res, indent=4))
