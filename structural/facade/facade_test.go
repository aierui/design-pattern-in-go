package facade

import (
	"fmt"
	"log"
	"testing"
)

func TestFacade(t *testing.T) {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("ab", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	/*
		Starting create account
		Account created

		Starting add money to wallet
		Account Verified
		SecurityCode Verified
		Wallet balance added successfully
		Sending wallet credit notification
		Make ledger entry for accountId abc with txnType credit for amount 10
		Starting debit money from wallet
		2020/06/26 00:59:27 Error: Account Name is incorrect
	*/
}
