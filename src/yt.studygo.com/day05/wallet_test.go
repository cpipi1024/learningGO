package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	/* 	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		//t.Helper()
		got := wallet.Balance()
		if want != got {
			t.Errorf("got %s want %s", got, want)
		}
	} */

	/* 	assertError := func(t *testing.T, err error, want string) {
		if err == nil {
			t.Fatal("want a error but got none")
		}
		if err.Error() != want {
			t.Errorf("got %s want %s", err, want)
		}
	} */
	t.Run("存比特币: ", func(t *testing.T) {
		wallet := Wallet{}
		btc := Bitcoin(10)
		wallet.Deposit(btc)
		//got := wallet.Balance()
		want := Bitcoin(10)
		//fmt.Printf("测试用例中wallet %x \n", &wallet.balance)
		// wallet.Rstring()
		/* 		if got != want {
			t.Errorf("got %s want %s ", got, want)
		} */
		assertBalance(t, wallet, want)
	})

	//测试取比特币
	t.Run("取比特币: ", func(t *testing.T) {
		wallet := Wallet{10} // 新的钱包？
		btc := Bitcoin(5)
		wallet.WithDraw(btc)
		//got := wallet.Balance()
		want := Bitcoin(5)
		/* 		if want != got {
			t.Errorf("got %s want %s", got, want)
		} */
		assertBalance(t, wallet, want)
	})

	// 取超标测试
	t.Run("取比特币超标: ", func(t *testing.T) {
		wallet := Wallet{10}
		btc := Bitcoin(20)
		err := wallet.WithDraw(btc)
		// 断言错误
		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want string) {
	if err == nil {
		t.Fatal("want a error but got none")
	}
	if err.Error() != want {
		t.Errorf("got %s want %s", err, want)
	}
}
