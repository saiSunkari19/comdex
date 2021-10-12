package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (m *Deposit) Validate() error {

	if m.VaultID == 0 {
		return fmt.Errorf("VaultID cannot be empty")
	}
	if m.Depositor == "" {
		return fmt.Errorf("depositor's address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.Depositor); err != nil {
		return errors.Wrapf(err, "invalid depositor %s", m.Depositor)
	}
	if m.AmountDeposited.IsNil() {
		return fmt.Errorf("amount_deposited cannot be nil")
	}
	if m.AmountDeposited.IsNegative() {
		return fmt.Errorf("amount_deposited cannot be negative")
	}

	return nil
}

func (m *Deposits) Validate() error {

	return nil
}

func (ds Deposits) SumCollateral() (sum sdk.Int) {
	sum = sdk.ZeroInt()
	for _, d := range ds.Deposits {
		if !d.AmountDeposited.IsZero() {
			sum = sum.Add(d.AmountDeposited)
		}
	}
	return
}
