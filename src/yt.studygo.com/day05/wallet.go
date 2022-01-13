package wallet

import (
	"errors"
	"fmt"
)

// 定义全局错误变量

var WithDrawError = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

// 类型别名 需要实现stringer接口

type Bitcoin int

// 创建bitcoin的返回字符串接口

type Stringer interface {
	String() string // 为什么一定要叫string?
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("Deposit方法中wallet的地址 %x \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("Balance方法中wallet的地址 %x \n", &w.balance)
	return w.balance
}

// 实现取钱方法
func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount > w.balance {
		return WithDrawError
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
