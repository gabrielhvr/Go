package calc

import "errors"

func Dividir2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("denominador nao pode ser zero")
	}
	return (a / b), nil // no go e possivel return mais de uma funcao. se colocar nil neste caso porque tem retorna duas variaveis
}
