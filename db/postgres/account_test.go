package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/muathendirangu/bank-api-service/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	return account
}

func TestAccountCreate(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	firstAccount := createRandomAccount(t)
	secondAccount, err := testQueries.GetAccount(context.Background(), firstAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, secondAccount)
}

func TestDeleteAccount(t *testing.T) {
	firstAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), firstAccount.ID)
	require.NoError(t, err)

	secondAccount, err := testQueries.GetAccount(context.Background(), firstAccount.ID)
	require.Error(t, err)
	require.Empty(t, secondAccount)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	args := ListAccountsParams{
		Offset: 5,
		Limit:  3,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, int(args.Limit))
}

func TestAccountUpdate(t *testing.T) {
	firstAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      firstAccount.ID,
		Balance: util.RandomMoney(),
	}
	secondAccount, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, secondAccount.ID, firstAccount.ID)
	require.Equal(t, secondAccount.Owner, firstAccount.Owner)
	require.Equal(t, args.Balance, secondAccount.Balance)
	require.Equal(t, firstAccount.Currency, secondAccount.Currency)
}
