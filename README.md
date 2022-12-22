# X

x is a repository that contains common golang packages.

Generate events file from abi.
```shell
python3 codegen/signature.py chain/bindings/clipper/abi.json chain/bindings/clipper/events.json
python3 codegen/signature.py chain/bindings/vat/abi.json chain/bindings/vat/events.json
```
Generate `events.Handler` from events file.
```shell
go run codegen/events_handler.go --abi=chain/bindings/clipper/abi.json --binding-package=github.com/IR-Digital-Token/x/chain/bindings/clipper --output-dir=chain/events/clipper --contract=Clipper
go run codegen/events_handler.go --abi=chain/bindings/vat/abi.json  --binding-package=github.com/IR-Digital-Token/x/chain/bindings/vat --output-dir=chain/events/vat --contract=Vat
```