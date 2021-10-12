package types

import (
	"encoding/binary"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName   = "vault"
	QuerierRoute = ModuleName
	RouterKey    = ModuleName
	StoreKey     = ModuleName
)

var (
	TypeMsgCreateRequest    = ModuleName + ":create"
	TypeMsgDepositRequest   = ModuleName + ":deposit"
	TypeMsgWithdrawRequest  = ModuleName + ":withdraw"
	TypeMsgDrawRequest      = ModuleName + ":draw"
	TypeMsgRepayRequest     = ModuleName + ":repay"
	TypeMsgLiquidateRequest = ModuleName + ":liquidate"
)

var (
	IDKey                          = []byte{0x00}
	VaultKeyPrefix                 = []byte{0x01}
	VaultForAddressByPairKeyPrefix = []byte{0x02}
	DepositKeyPrefix               = []byte{0x03}
)

var splitter = []byte(":")

func VaultKey(id uint64) []byte {
	return append(VaultKeyPrefix, sdk.Uint64ToBigEndian(id)...)
}

func VaultForAddressByPair(address sdk.AccAddress, pairID uint64) []byte {
	v := append(VaultForAddressByPairKeyPrefix, address.Bytes()...)
	if len(v) != 1+20 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+20))
	}

	return append(v, sdk.Uint64ToBigEndian(pairID)...)
}

func DepositKey(vaultID uint64, depositor sdk.AccAddress) []byte {
	return generateKey(DepositKeyPrefix, VaultIDBytes(vaultID), splitter, depositor)
}

func VaultIDBytes(vaultID uint64) (vaultIDBz []byte) {
	vaultIDBz = make([]byte, 16)
	binary.BigEndian.PutUint64(vaultIDBz, vaultID)
	return
}

func generateKey(bytes ...[]byte) (r []byte) {
	for _, b := range bytes {
		r = append(r, b...)
	}
	return
}
