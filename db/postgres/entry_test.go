package db

import (
	"context"
	"testing"

	"github.com/muathendirangu/bank-api-service/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	firstEntry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, firstEntry)
	require.Equal(t, args.AccountID, firstEntry.AccountID)
	require.Equal(t, args.Amount, firstEntry.Amount)
	require.NotZero(t, firstEntry.Amount)
	require.NotZero(t, firstEntry.CreatedAt)
	return firstEntry
}

func TestEntryCreate(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	firstEntry := createRandomEntry(t, account)
	secondEntry, err := testQueries.GetEntry(context.Background(), firstEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, secondEntry)
	require.Equal(t, firstEntry.AccountID, secondEntry.AccountID)
}
