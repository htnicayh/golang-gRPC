package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/htnicayh/golang-gRPC/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomBalance(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account_one := createRandomAccount(t)
	account_two, err := testQueries.GetAccount(context.Background(), account_one.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account_two)

	require.Equal(t, account_one.ID, account_two.ID)
	require.Equal(t, account_one.Owner, account_two.Owner)
	require.Equal(t, account_one.Balance, account_two.Balance)
	require.Equal(t, account_one.Currency, account_two.Currency)

	require.WithinDuration(t, account_one.CreatedAt, account_two.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account_one := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account_one.ID,
		Balance: utils.RandomBalance(),
	}

	err := testQueries.UpdateAccount(context.Background(), arg)
	account_two, err := testQueries.GetAccount(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account_two)

	require.Equal(t, account_one.ID, account_two.ID)
	require.Equal(t, account_one.Owner, account_two.Owner)
	require.Equal(t, arg.Balance, account_two.Balance)
	require.Equal(t, account_one.Currency, account_two.Currency)
	require.WithinDuration(t, account_one.CreatedAt, account_two.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account_one := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account_one.ID)
	require.NoError(t, err)

	account_two, err := testQueries.GetAccount(context.Background(), account_one.ID)
	require.Empty(t, account_two)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
