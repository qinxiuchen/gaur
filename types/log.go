// Copyright 2018 The Fractal Team Authors
// This file is part of the fractal project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fractalplatform/fractal/common"
)

// Log represents a contract log event. These events are generated by the LOG opcode and
// stored/indexed by the node.
type Log struct {
	Name        common.Name   // name of the contract that generated the event
	Topics      []common.Hash // list of topics provided by the contract.
	Data        []byte        // supplied by the contract, usually ABI-encoded
	BlockNumber uint64        // block in which the transaction was included
	BlockHash   common.Hash   // hash of the block in which the transaction was included
	TxHash      common.Hash   // hash of the transaction
	Index       uint          // index of the log in the receipt
	ActionIndex uint          // index of the input and output in the transaction
	TxIndex     uint          // index of the transaction in the block

}

// RPCLog that will serialize to the RPC representation of a log.
type RPCLog struct {
	Name        common.Name   `json:"name"`
	Topics      []common.Hash `json:"topics"`
	Data        hexutil.Bytes `json:"data" validate:"nonzero"`
	BlockNumber uint64        `json:"blockNumber"`
	BlockHash   common.Hash   `json:"blockHash"`
	TxHash      common.Hash   `json:"transactionHash"`
	Index       uint          `json:"logIndex"`
	ActionIndex uint          `json:"actionIndex"`
	TxIndex     uint          `json:"transactionIndex"`
}

// NewRPCLog returns a log that will serialize to the RPC.
func (l *Log) NewRPCLog() *RPCLog {
	return &RPCLog{
		Name:        l.Name,
		Topics:      l.Topics,
		Data:        hexutil.Bytes(l.Data),
		BlockNumber: l.BlockNumber,
		TxHash:      l.TxHash,
		TxIndex:     l.TxIndex,
		ActionIndex: l.ActionIndex,
		BlockHash:   l.BlockHash,
		Index:       l.Index,
	}
}
