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

package rpcapi

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/fractalplatform/fractal/accountmanager"
	"github.com/fractalplatform/fractal/asset"
	"github.com/fractalplatform/fractal/common"
)

type AccountAPI struct {
	b Backend
}

func NewAccountAPI(b Backend) *AccountAPI {
	return &AccountAPI{b}
}

var (
	ErrGetAccounManagerErr = errors.New("get account manager failure")
)

//AccountIsExist
func (aapi *AccountAPI) AccountIsExist(ctx context.Context, acctName common.Name) (bool, error) {
	acct, err := aapi.b.GetAccountManager()
	if err != nil {
		return false, err
	}
	return acct.AccountIsExist(acctName)
}

//GetAccountByID
func (aapi *AccountAPI) GetAccountByID(ctx context.Context, accountID uint64) (*accountmanager.Account, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return am.GetAccountById(accountID)
}

//GetAccountByName
func (aapi *AccountAPI) GetAccountByName(ctx context.Context, accountName common.Name) (*accountmanager.Account, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return am.GetAccountByName(accountName)
}

// GetAccountBalanceByAssetID get account asset balance by assetID,
// if typeID is 0 not contain subaccount balance; if typeID is 1 contain subaccount balance; if typeID is other return error
func (aapi *AccountAPI) GetAccountBalanceByAssetID(ctx context.Context, accountName common.Name, assetID uint64, typeID uint64) (*big.Int, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return am.GetAccountBalanceByAssetID(accountName, assetID, typeID)
}

//GetCode
func (aapi *AccountAPI) GetCode(ctx context.Context, accountName common.Name) (hexutil.Bytes, error) {
	acct, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	result, err := acct.GetCode(accountName)
	if err != nil {
		return nil, err
	}
	return (hexutil.Bytes)(result), nil

}

//GetNonce
func (aapi *AccountAPI) GetNonce(ctx context.Context, accountName common.Name) (uint64, error) {
	acct, err := aapi.b.GetAccountManager()
	if err != nil {
		return 0, err
	}
	return acct.GetNonce(accountName)

}

//GetAssetInfoByName
func (aapi *AccountAPI) GetAssetInfoByName(ctx context.Context, assetName string) (*asset.AssetObject, error) {
	acct, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return acct.GetAssetInfoByName(assetName)
}

//GetAssetInfoByID
func (aapi *AccountAPI) GetAssetInfoByID(ctx context.Context, assetID uint64) (*asset.AssetObject, error) {
	acct, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return acct.GetAssetInfoByID(assetID)
}

//GetAssetAmountByTime
func (aapi *AccountAPI) GetAssetAmountByTime(ctx context.Context, assetID uint64, time uint64) (*big.Int, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return am.GetAssetAmountByTime(assetID, time)
}

//GetAccountBalanceByTime
func (aapi *AccountAPI) GetAccountBalanceByTime(ctx context.Context, accountName common.Name, assetID uint64, typeID uint64, time uint64) (*big.Int, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return nil, err
	}
	return am.GetBalanceByTime(accountName, assetID, typeID, time)
}

// GetSnapshotTime get snapshot time
// if TypeID = 0, time must be 0, return last snapshot time;
// if TypeID = 1, return previous snapshot time;
// if TypeID = 2, return next snapshot time.
func (aapi *AccountAPI) GetSnapshotTime(typeID uint64, time uint64) (uint64, error) {
	am, err := aapi.b.GetAccountManager()
	if err != nil {
		return 0, err
	}
	return am.GetSnapshotTime(typeID, time)
}
