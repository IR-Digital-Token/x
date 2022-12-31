// Code generated - DO NOT EDIT.
// This file is a generated event handler and any manual changes will be lost.

package clipper

import (
	"errors"

	"github.com/IR-Digital-Token/x/chain/events"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type File0Handler struct {
	binding  *Clipper
	callback events.CallbackFn[ClipperFile0]
}

func (h *File0Handler) ID() string {
	return "0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba"
}

func (h *File0Handler) DecodeLog(log types.Log) (interface{}, error) {
	return h.binding.ParseFile0(log)
}

func (h *File0Handler) HandleEvent(header types.Header, event interface{}) error {
	e, ok := event.(ClipperFile0)
	if !ok {
		return errors.New("event type is not ClipperFile0")
	}
	return h.callback(header, e)
}

func (h *File0Handler) DecodeAndHandle(header types.Header, log types.Log) error {
	e, err := h.binding.ParseFile0(log)
	if err != nil {
		return err
	}
	return h.callback(header, *e)
}

func NewFile0Handler(addr common.Address, eth *ethclient.Client, callback events.CallbackFn[ClipperFile0]) events.Handler {
	b, err := NewClipper(addr, eth)
	if err != nil {
		panic(err)
	}
	return &File0Handler{
		binding:  b,
		callback: callback,
	}
}
