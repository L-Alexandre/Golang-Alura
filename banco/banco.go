package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"
)

func main() {
	clienteAlexandre := clientes.Titular{"Alexandre", "413.314.413.17", "Desenvolvedor Go"}
	contaDoAlexandre := contas.ContaCorrente{clienteAlexandre, 589, 123456, 500}

	fmt.Println(contaDoAlexandre)
}
