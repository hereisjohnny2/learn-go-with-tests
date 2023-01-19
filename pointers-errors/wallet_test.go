package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("should deposit bitcoin amount in a wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertEquals(t, Bitcoin(10), wallet)
	})

	t.Run("should with draw a valid amount of bitcoin from a wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertEquals(t, Bitcoin(10), wallet)
		assetErrorNil(t, err)
	})

	t.Run("should error if the are no funds available to withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertEquals(t, Bitcoin(20), wallet)
		assetError(t, ErrInsufficientFunds, err)
	})
}

func assetErrorNil(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("expected error to nil but it is not")
	}
}

func assertEquals(t testing.TB, want Bitcoin, wallet Wallet) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
}

func assetError(t testing.TB, want error, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error but got none")
	}
	if got != want {
		t.Errorf("got %s, expected %s", got, want)
	}
}
