package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/nghia220800/golang-blockchain/wallet"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	ID     []byte
	Out    int
	Sig    []byte
	PubKey []byte
}

type TXO_set struct {
	Output_set []TxOutput
}

func NewTXO(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.LockOut([]byte(address))
	return txo
}

func (in *TxInput) PubHashCmp(pubKeyHash []byte) bool {
	lockHash := wallet.PublicKeyHash(in.PubKey)
	return bytes.Compare(lockHash, pubKeyHash) == 0
}

func (out *TxOutput) LockOut(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TxOutput) isLockedWKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func (txos TXO_set) Serialize() []byte {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(txos)
	Handle(err)
	return buffer.Bytes()
}

func DeserializeOut(data []byte) TXO_set {
	var out TXO_set
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&out)
	Handle(err)
	return out
}
