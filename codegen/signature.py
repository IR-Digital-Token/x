import json
import os

from web3 import Web3

current_dir = os.getcwd()
ABIs = [
    current_dir + '/bindings/clipper/abi.json',
    current_dir + '/bindings/vat/abi.json',
]

for f in ABIs:
    res = []
    with open(f) as file:
        abi = json.load(file)
        for item in abi:
            if item['type'] == 'event':
                s = item['name'] + '(' + ','.join(map(lambda x: x['type'], item['inputs'])) + ')'
                print(s)
                print(Web3.keccak(text=s).hex())
                print("------")
                res.append({
                    'name': item['name'],
                    'signature': s,
                    'hash': Web3.keccak(text=s).hex(),
                })
    with open(f.replace('abi', 'events'), 'w') as file:
        file.write(json.dumps(res, indent=4))
