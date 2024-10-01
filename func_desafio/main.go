package main

import "fmt"

// FORMA COMO O PROFESSOR FEZ:

type Aluno struct {
	Nome  string
	Nota1 float64
	Nota2 float64
}

func calculaMedia(aluno Aluno) (string, float64) {
	media := (aluno.Nota1 + aluno.Nota2) / 2

	return aluno.Nome, media
}

func verifiaAprovacao(media float64) string {
	if media >= 7 {
		return "Parabéns aluno aprovado"
	}

	return "Ops, você está de exame"
}

func main() {

	aluno1 := Aluno{Nome: "Mateus", Nota1: 6.0, Nota2: 9.5}

	fmt.Printf("%s:\n Nota1: %.2f;\n Nota2: %.2f;\n", aluno1.Nome, aluno1.Nota1, aluno1.Nota2)

	nome, media := calculaMedia(aluno1)

	fmt.Printf("Média do %s: %.2f \n", nome, media)

	statusAluno1 := verifiaAprovacao(media)
	fmt.Println(statusAluno1)

}

// FORMA COMO EU TENTEI FAZER LOGO ABAIXO:

/* func media(nomAluno string, nota1 float64, N2 float64) (string, float64) {
	medAluno := (nota1 + N2) / 2

	return nomAluno, medAluno
}

func aprovado(nomaluno string, media float64) (float64, string) {
	if media >= 7 {
		fmt.Println("ALUNO APROVADO")
	}
	return media, "ALUNO ESTÁ DE EXAME"
}

func main() {
	aluno1 := "Walacy"
	notaAluno1 := 6.0
	notaAluno2 := 9.5

	resultado1, mediaAluno1 := media(aluno1, notaAluno1, notaAluno2)
	aprovacao1 := aprovado(resultado 1)
	fmt.Printf("A média do aluno: %s, \né: %f", resultado1, mediaAluno1)
	fmt.Print("O aluno: ", aprovacao1)
} */
