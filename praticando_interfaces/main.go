package main

import "fmt"

type Veiculo interface {
	VelocidadeMaxima() int
	MeioDeLocomocao() string
	MostraNome() string
}

type Trem struct {
	Nome    string
	MaxVelo int
}

type Carro struct {
	Nome    string
	MaxVelo int
}

type Moto struct {
	Nome    string
	MaxVelo int
}

func (c Carro) VelocidadeMaxima() int {
	return c.MaxVelo
}

func (c Carro) MeioDeLocomocao() string {
	return "Rodas"
}

func (c Carro) MostraNome() string {
	return c.Nome
}

func (t Trem) VelocidadeMaxima() int {
	return t.MaxVelo
}

func (t Trem) MeioDeLocomocao() string {
	return "Trilhos"
}

func (t Trem) MostraNome() string {
	return t.Nome
}

func (m Moto) VelocidadeMaxima() int {
	return m.MaxVelo
}

func (m Moto) MeioDeLocomocao() string {
	return "Rodas"
}

func (m Moto) MostraNome() string {
	return m.Nome
}

func DescreverVeiculo(v Veiculo) {
	fmt.Println("Veiculo atual:", v.MostraNome())
	fmt.Printf("Meio de locomoção: %s, Velocidade Máxima: %d km/h \n", v.MeioDeLocomocao(), v.VelocidadeMaxima())
}

func main() {

	meuCarro := Carro{MaxVelo: 130, Nome: "Gol 1.0"}
	meuTrem := Trem{MaxVelo: 90, Nome: "Trem GO"}
	minhaMoto := Moto{MaxVelo: 80, Nome: "MT-03"}

	DescreverVeiculo(meuCarro)
	DescreverVeiculo(meuTrem)
	DescreverVeiculo(minhaMoto)
}
