// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pg "github.com/GoPlugin/pluginV2/core/services/pg"

	time "time"

	txmgr "github.com/GoPlugin/pluginV2/core/chains/evm/txmgr"

	types "github.com/GoPlugin/pluginV2/core/chains/evm/types"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() {
	_m.Called()
}

// DeleteInProgressAttempt provides a mock function with given fields: ctx, attempt
func (_m *ORM) DeleteInProgressAttempt(ctx context.Context, attempt txmgr.EthTxAttempt) error {
	ret := _m.Called(ctx, attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, txmgr.EthTxAttempt) error); ok {
		r0 = rf(ctx, attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EthTransactions provides a mock function with given fields: offset, limit
func (_m *ORM) EthTransactions(offset int, limit int) ([]txmgr.EthTx, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTx
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]txmgr.EthTx, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTx); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTransactionsWithAttempts provides a mock function with given fields: offset, limit
func (_m *ORM) EthTransactionsWithAttempts(offset int, limit int) ([]txmgr.EthTx, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTx
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]txmgr.EthTx, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTx); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EthTxAttempts provides a mock function with given fields: offset, limit
func (_m *ORM) EthTxAttempts(offset int, limit int) ([]txmgr.EthTxAttempt, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []txmgr.EthTxAttempt
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, int) ([]txmgr.EthTxAttempt, int, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []txmgr.EthTxAttempt); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindEthReceiptsPendingConfirmation provides a mock function with given fields: ctx, blockNum, chainID
func (_m *ORM) FindEthReceiptsPendingConfirmation(ctx context.Context, blockNum int64, chainID big.Int) ([]txmgr.EthReceiptsPlus, error) {
	ret := _m.Called(ctx, blockNum, chainID)

	var r0 []txmgr.EthReceiptsPlus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, big.Int) ([]txmgr.EthReceiptsPlus, error)); ok {
		return rf(ctx, blockNum, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, big.Int) []txmgr.EthReceiptsPlus); ok {
		r0 = rf(ctx, blockNum, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthReceiptsPlus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, big.Int) error); ok {
		r1 = rf(ctx, blockNum, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttempt provides a mock function with given fields: hash
func (_m *ORM) FindEthTxAttempt(hash common.Hash) (*txmgr.EthTxAttempt, error) {
	ret := _m.Called(hash)

	var r0 *txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*txmgr.EthTxAttempt, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *txmgr.EthTxAttempt); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptConfirmedByEthTxIDs provides a mock function with given fields: ids
func (_m *ORM) FindEthTxAttemptConfirmedByEthTxIDs(ids []int64) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(ids)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func([]int64) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int64) []txmgr.EthTxAttempt); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsByEthTxIDs provides a mock function with given fields: ids
func (_m *ORM) FindEthTxAttemptsByEthTxIDs(ids []int64) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(ids)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func([]int64) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int64) []txmgr.EthTxAttempt); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsRequiringReceiptFetch provides a mock function with given fields: chainID
func (_m *ORM) FindEthTxAttemptsRequiringReceiptFetch(chainID big.Int) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(chainID)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func(big.Int) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(big.Int) []txmgr.EthTxAttempt); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(big.Int) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxAttemptsRequiringResend provides a mock function with given fields: olderThan, maxInFlightTransactions, chainID, address
func (_m *ORM) FindEthTxAttemptsRequiringResend(olderThan time.Time, maxInFlightTransactions uint32, chainID big.Int, address common.Address) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(olderThan, maxInFlightTransactions, chainID, address)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func(time.Time, uint32, big.Int, common.Address) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(olderThan, maxInFlightTransactions, chainID, address)
	}
	if rf, ok := ret.Get(0).(func(time.Time, uint32, big.Int, common.Address) []txmgr.EthTxAttempt); ok {
		r0 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, uint32, big.Int, common.Address) error); ok {
		r1 = rf(olderThan, maxInFlightTransactions, chainID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxByHash provides a mock function with given fields: hash
func (_m *ORM) FindEthTxByHash(hash common.Hash) (*txmgr.EthTx, error) {
	ret := _m.Called(hash)

	var r0 *txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Hash) (*txmgr.EthTx, error)); ok {
		return rf(hash)
	}
	if rf, ok := ret.Get(0).(func(common.Hash) *txmgr.EthTx); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxWithAttempts provides a mock function with given fields: etxID
func (_m *ORM) FindEthTxWithAttempts(etxID int64) (txmgr.EthTx, error) {
	ret := _m.Called(etxID)

	var r0 txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (txmgr.EthTx, error)); ok {
		return rf(etxID)
	}
	if rf, ok := ret.Get(0).(func(int64) txmgr.EthTx); ok {
		r0 = rf(etxID)
	} else {
		r0 = ret.Get(0).(txmgr.EthTx)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(etxID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxWithNonce provides a mock function with given fields: fromAddress, nonce
func (_m *ORM) FindEthTxWithNonce(fromAddress common.Address, nonce uint) (*txmgr.EthTx, error) {
	ret := _m.Called(fromAddress, nonce)

	var r0 *txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, uint) (*txmgr.EthTx, error)); ok {
		return rf(fromAddress, nonce)
	}
	if rf, ok := ret.Get(0).(func(common.Address, uint) *txmgr.EthTx); ok {
		r0 = rf(fromAddress, nonce)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, uint) error); ok {
		r1 = rf(fromAddress, nonce)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxsRequiringGasBump provides a mock function with given fields: ctx, address, blockNum, gasBumpThreshold, depth, chainID
func (_m *ORM) FindEthTxsRequiringGasBump(ctx context.Context, address common.Address, blockNum int64, gasBumpThreshold int64, depth int64, chainID big.Int) ([]*txmgr.EthTx, error) {
	ret := _m.Called(ctx, address, blockNum, gasBumpThreshold, depth, chainID)

	var r0 []*txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int64, int64, int64, big.Int) ([]*txmgr.EthTx, error)); ok {
		return rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int64, int64, int64, big.Int) []*txmgr.EthTx); ok {
		r0 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, int64, int64, int64, big.Int) error); ok {
		r1 = rf(ctx, address, blockNum, gasBumpThreshold, depth, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEthTxsRequiringResubmissionDueToInsufficientEth provides a mock function with given fields: address, chainID, qopts
func (_m *ORM) FindEthTxsRequiringResubmissionDueToInsufficientEth(address common.Address, chainID big.Int, qopts ...pg.QOpt) ([]*txmgr.EthTx, error) {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, address, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, big.Int, ...pg.QOpt) ([]*txmgr.EthTx, error)); ok {
		return rf(address, chainID, qopts...)
	}
	if rf, ok := ret.Get(0).(func(common.Address, big.Int, ...pg.QOpt) []*txmgr.EthTx); ok {
		r0 = rf(address, chainID, qopts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, big.Int, ...pg.QOpt) error); ok {
		r1 = rf(address, chainID, qopts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEtxAttemptsConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *ORM) FindEtxAttemptsConfirmedMissingReceipt(chainID big.Int) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(chainID)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func(big.Int) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(chainID)
	}
	if rf, ok := ret.Get(0).(func(big.Int) []txmgr.EthTxAttempt); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(big.Int) error); ok {
		r1 = rf(chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransactionsConfirmedInBlockRange provides a mock function with given fields: highBlockNumber, lowBlockNumber, chainID
func (_m *ORM) FindTransactionsConfirmedInBlockRange(highBlockNumber int64, lowBlockNumber int64, chainID big.Int) ([]*txmgr.EthTx, error) {
	ret := _m.Called(highBlockNumber, lowBlockNumber, chainID)

	var r0 []*txmgr.EthTx
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, int64, big.Int) ([]*txmgr.EthTx, error)); ok {
		return rf(highBlockNumber, lowBlockNumber, chainID)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, big.Int) []*txmgr.EthTx); ok {
		r0 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*txmgr.EthTx)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, int64, big.Int) error); ok {
		r1 = rf(highBlockNumber, lowBlockNumber, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInProgressEthTxAttempts provides a mock function with given fields: ctx, address, chainID
func (_m *ORM) GetInProgressEthTxAttempts(ctx context.Context, address common.Address, chainID big.Int) ([]txmgr.EthTxAttempt, error) {
	ret := _m.Called(ctx, address, chainID)

	var r0 []txmgr.EthTxAttempt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, big.Int) ([]txmgr.EthTxAttempt, error)); ok {
		return rf(ctx, address, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, big.Int) []txmgr.EthTxAttempt); ok {
		r0 = rf(ctx, address, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]txmgr.EthTxAttempt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, big.Int) error); ok {
		r1 = rf(ctx, address, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertEthReceipt provides a mock function with given fields: receipt
func (_m *ORM) InsertEthReceipt(receipt *txmgr.EthReceipt) error {
	ret := _m.Called(receipt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthReceipt) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTx provides a mock function with given fields: etx
func (_m *ORM) InsertEthTx(etx *txmgr.EthTx) error {
	ret := _m.Called(etx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTx) error); ok {
		r0 = rf(etx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertEthTxAttempt provides a mock function with given fields: attempt
func (_m *ORM) InsertEthTxAttempt(attempt *txmgr.EthTxAttempt) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTxAttempt) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadEthTxAttempts provides a mock function with given fields: etx, qopts
func (_m *ORM) LoadEthTxAttempts(etx *txmgr.EthTx, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTx, ...pg.QOpt) error); ok {
		r0 = rf(etx, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoadEthTxesAttempts provides a mock function with given fields: etxs, qopts
func (_m *ORM) LoadEthTxesAttempts(etxs []*txmgr.EthTx, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, etxs)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*txmgr.EthTx, ...pg.QOpt) error); ok {
		r0 = rf(etxs, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllConfirmedMissingReceipt provides a mock function with given fields: chainID
func (_m *ORM) MarkAllConfirmedMissingReceipt(chainID big.Int) error {
	ret := _m.Called(chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(big.Int) error); ok {
		r0 = rf(chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkOldTxesMissingReceiptAsErrored provides a mock function with given fields: blockNum, finalityDepth, chainID, qopts
func (_m *ORM) MarkOldTxesMissingReceiptAsErrored(blockNum int64, finalityDepth uint32, chainID big.Int, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, blockNum, finalityDepth, chainID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, uint32, big.Int, ...pg.QOpt) error); ok {
		r0 = rf(blockNum, finalityDepth, chainID, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PreloadEthTxes provides a mock function with given fields: attempts
func (_m *ORM) PreloadEthTxes(attempts []txmgr.EthTxAttempt) error {
	ret := _m.Called(attempts)

	var r0 error
	if rf, ok := ret.Get(0).(func([]txmgr.EthTxAttempt) error); ok {
		r0 = rf(attempts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveConfirmedMissingReceiptAttempt provides a mock function with given fields: ctx, timeout, attempt, broadcastAt
func (_m *ORM) SaveConfirmedMissingReceiptAttempt(ctx context.Context, timeout time.Duration, attempt *txmgr.EthTxAttempt, broadcastAt time.Time) error {
	ret := _m.Called(ctx, timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, *txmgr.EthTxAttempt, time.Time) error); ok {
		r0 = rf(ctx, timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveFetchedReceipts provides a mock function with given fields: receipts, chainID
func (_m *ORM) SaveFetchedReceipts(receipts []types.Receipt, chainID big.Int) error {
	ret := _m.Called(receipts, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func([]types.Receipt, big.Int) error); ok {
		r0 = rf(receipts, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInProgressAttempt provides a mock function with given fields: attempt
func (_m *ORM) SaveInProgressAttempt(attempt *txmgr.EthTxAttempt) error {
	ret := _m.Called(attempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*txmgr.EthTxAttempt) error); ok {
		r0 = rf(attempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveInsufficientEthAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *ORM) SaveInsufficientEthAttempt(timeout time.Duration, attempt *txmgr.EthTxAttempt, broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *txmgr.EthTxAttempt, time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveReplacementInProgressAttempt provides a mock function with given fields: oldAttempt, replacementAttempt, qopts
func (_m *ORM) SaveReplacementInProgressAttempt(oldAttempt txmgr.EthTxAttempt, replacementAttempt *txmgr.EthTxAttempt, qopts ...pg.QOpt) error {
	_va := make([]interface{}, len(qopts))
	for _i := range qopts {
		_va[_i] = qopts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, oldAttempt, replacementAttempt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(txmgr.EthTxAttempt, *txmgr.EthTxAttempt, ...pg.QOpt) error); ok {
		r0 = rf(oldAttempt, replacementAttempt, qopts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveSentAttempt provides a mock function with given fields: timeout, attempt, broadcastAt
func (_m *ORM) SaveSentAttempt(timeout time.Duration, attempt *txmgr.EthTxAttempt, broadcastAt time.Time) error {
	ret := _m.Called(timeout, attempt, broadcastAt)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, *txmgr.EthTxAttempt, time.Time) error); ok {
		r0 = rf(timeout, attempt, broadcastAt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetBroadcastBeforeBlockNum provides a mock function with given fields: blockNum, chainID
func (_m *ORM) SetBroadcastBeforeBlockNum(blockNum int64, chainID big.Int) error {
	ret := _m.Called(blockNum, chainID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, big.Int) error); ok {
		r0 = rf(blockNum, chainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBroadcastAts provides a mock function with given fields: now, etxIDs
func (_m *ORM) UpdateBroadcastAts(now time.Time, etxIDs []int64) error {
	ret := _m.Called(now, etxIDs)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time, []int64) error); ok {
		r0 = rf(now, etxIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxForRebroadcast provides a mock function with given fields: etx, etxAttempt
func (_m *ORM) UpdateEthTxForRebroadcast(etx txmgr.EthTx, etxAttempt txmgr.EthTxAttempt) error {
	ret := _m.Called(etx, etxAttempt)

	var r0 error
	if rf, ok := ret.Get(0).(func(txmgr.EthTx, txmgr.EthTxAttempt) error); ok {
		r0 = rf(etx, etxAttempt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEthTxsUnconfirmed provides a mock function with given fields: ids
func (_m *ORM) UpdateEthTxsUnconfirmed(ids []int64) error {
	ret := _m.Called(ids)

	var r0 error
	if rf, ok := ret.Get(0).(func([]int64) error); ok {
		r0 = rf(ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewORM interface {
	mock.TestingT
	Cleanup(func())
}

// NewORM creates a new instance of ORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewORM(t mockConstructorTestingTNewORM) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
