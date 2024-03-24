package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaExemplo := contas.ContaPoupanca{}
	contaExemplo.Depositar(100)
	fmt.Println(contaExemplo.ObterSaldo())
}
