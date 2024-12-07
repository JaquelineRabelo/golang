package main

import "fmt"

func main() {
//somaDosValores := soma(42, 13)
	//fmt.Println(somaDosValores)

	//sub := subtracao(10, 5)
	//fmt.Println(sub)

	//nome1, nome2 := printaNome ("Jaqueline")
	//fmt.Println(nome1)
	//fmt.Println(nome2)

	//nome, sobrenome := printaNomeCompleto("Jaqueline", "Cavalcanti")
	//fmt.Println(nome)
	//fmt.Println(sobrenome)

	nome, idade := informacoesPessoais("Jaqueline", 31)
	fmt.Println(nome)
	fmt.Println(idade)

	informacoes, contatos := menuPrincipal("empresa", "contatos")
	fmt.Println(informacoes)
	fmt.Println(contatos)


}

//func printaNome(nome string) (string, string) {
	//return nome, nome}

//func printaNomeCompleto(nome, sobrenome string) (string, string) {
	//return nome,sobrenome}

func informacoesPessoais(nome string, idade int) (string, int) {
	return nome, idade
}

func menuPrincipal(informacoes, contatos string) (string, string) {
	return informacoes, contatos
}


// Função começando com letra maiuscula:
// Função é PÚBLICA
// Função ela só pode ser utilizada fora do próprio pacote
// como utilizaria ela fora do pacote: main.PrintaNome(nome)


//func soma (x int, y int) int {
//return x + y }

//func subtracao(x int, y int) int { return x - y }


