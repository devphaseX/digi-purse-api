// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/devphasex/cedar-bank-api/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/devphasex/cedar-bank-api/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	pgtype "github.com/jackc/pgx/v5/pgtype"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateBalanceEntry mocks base method.
func (m *MockStore) CreateBalanceEntry(arg0 context.Context, arg1 db.CreateBalanceEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBalanceEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBalanceEntry indicates an expected call of CreateBalanceEntry.
func (mr *MockStoreMockRecorder) CreateBalanceEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBalanceEntry", reflect.TypeOf((*MockStore)(nil).CreateBalanceEntry), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// GetAccountBalanceEntries mocks base method.
func (m *MockStore) GetAccountBalanceEntries(arg0 context.Context, arg1 pgtype.Int8) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountBalanceEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountBalanceEntries indicates an expected call of GetAccountBalanceEntries.
func (mr *MockStoreMockRecorder) GetAccountBalanceEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountBalanceEntries", reflect.TypeOf((*MockStore)(nil).GetAccountBalanceEntries), arg0, arg1)
}

// GetAccountByID mocks base method.
func (m *MockStore) GetAccountByID(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockStoreMockRecorder) GetAccountByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockStore)(nil).GetAccountByID), arg0, arg1)
}

// GetAccountByIDForUpdate mocks base method.
func (m *MockStore) GetAccountByIDForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByIDForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByIDForUpdate indicates an expected call of GetAccountByIDForUpdate.
func (mr *MockStoreMockRecorder) GetAccountByIDForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByIDForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountByIDForUpdate), arg0, arg1)
}

// GetAccounts mocks base method.
func (m *MockStore) GetAccounts(arg0 context.Context, arg1 db.GetAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockStoreMockRecorder) GetAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockStore)(nil).GetAccounts), arg0, arg1)
}

// GetBalanceEntry mocks base method.
func (m *MockStore) GetBalanceEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceEntry indicates an expected call of GetBalanceEntry.
func (mr *MockStoreMockRecorder) GetBalanceEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceEntry", reflect.TypeOf((*MockStore)(nil).GetBalanceEntry), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (*db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(*db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateBalance mocks base method.
func (m *MockStore) UpdateBalance(arg0 context.Context, arg1 db.UpdateBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockStoreMockRecorder) UpdateBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockStore)(nil).UpdateBalance), arg0, arg1)
}

// UpdateTransferAccountBalance mocks base method.
func (m *MockStore) UpdateTransferAccountBalance(arg0 context.Context, arg1 db.UpdateTransferAccountBalanceParams) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransferAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransferAccountBalance indicates an expected call of UpdateTransferAccountBalance.
func (mr *MockStoreMockRecorder) UpdateTransferAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransferAccountBalance", reflect.TypeOf((*MockStore)(nil).UpdateTransferAccountBalance), arg0, arg1)
}