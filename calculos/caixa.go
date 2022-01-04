package calculos

import (
	"fmt"

	"github.com/samuelves/correiosgo/consts"
)

type Pacote struct {
	Weight   float64
	Height   float64
	Width    float64
	Length   float64
	Sku      string
	Quantity int
}

type Caixa struct {
	Width  float64
	Height float64
	Length float64
	Weight float64
	Pacote []Pacote
}

var (
	PesoMaximoTotal float64
)

func Pacotes(pacotes []Pacote) []Caixa {
	caixas := make([]Caixa, 1)
	var numero_pacote int = 0
	for key, pacote := range pacotes {
		fmt.Println("numero_pacote: ", numero_pacote)
		fmt.Println("Key: ", key)
		copia := pacote
		copia.Quantity = 1
		copia.Weight = DimensaoMetroParaCm(pacote.Weight)
		copia.Height = DimensaoMetroParaCm(pacote.Height)
		copia.Length = DimensaoMetroParaCm(pacote.Length)
		copia.Weight = pacote.Weight
		var numero_caixa int = 0

		for i := 0; i < copia.Quantity; i++ {
			fmt.Println("Numero Caixa: ", numero_caixa)
			if !validarDimensoesDaCaixa(caixas, numero_caixa, "width") {
				copia.Width = 0
			}
			if !validarDimensoesDaCaixa(caixas, numero_caixa, "height") {
				copia.Height = 0
			}
			if !validarDimensoesDaCaixa(caixas, numero_caixa, "length") {
				copia.Length = 0
			}
			if !validarDimensoesDaCaixa(caixas, numero_caixa, "weight") {
				copia.Weight = 0
			}
			newWidth := caixas[numero_caixa].Width + copia.Width
			newHeight := caixas[numero_caixa].Height + copia.Height
			newLength := caixas[numero_caixa].Length + copia.Length
			newWeight := caixas[numero_caixa].Weight + copia.Weight

			cabeDoLado := (newWidth <= consts.MAXIMA_WIDTH) && cabeNoLimiteDoPacote(caixas, copia, numero_caixa, "width")
			cabeNaAltura := (newHeight <= consts.MAXIMA_HEIGHT) && cabeNoLimiteDoPacote(caixas, copia, numero_caixa, "height")
			cabeNaProfundidade := (newLength <= consts.MAXIMA_LENGTH) && cabeNoLimiteDoPacote(caixas, copia, numero_caixa, "length")
			cabeNoPeso := newWeight <= consts.PESO_MAXIMO_PAC_SEDEX

			if (cabeDoLado || cabeNaAltura || cabeNaProfundidade) && cabeNoPeso {
				// tenho 10 pacotes
				// na caixa cabe 3

				// for pacotes
				// caixa 3

				// criar + 4 caixas

				if len(caixas[numero_caixa].Pacote) == key {
					caixas[numero_caixa].Pacote = append(caixas[numero_caixa].Pacote, copia)
				} else {
					if caixas[numero_caixa].Pacote[numero_pacote].Sku == copia.Sku {
						caixas[numero_caixa].Pacote[numero_pacote].Quantity++
					}

				}

				caixas[numero_caixa].Weight += copia.Weight
				if cabeDoLado {
					caixas[numero_caixa].Width += copia.Width
					caixas[numero_caixa].Height = retornaValorMaior(caixas[numero_caixa].Height, copia.Height)
					caixas[numero_caixa].Length = retornaValorMaior(caixas[numero_caixa].Length, copia.Length)
				}
				if cabeNaAltura {
					caixas[numero_caixa].Height += copia.Height
					caixas[numero_caixa].Width = retornaValorMaior(caixas[numero_caixa].Width, copia.Width)
					caixas[numero_caixa].Length = retornaValorMaior(caixas[numero_caixa].Length, copia.Length)
				}
				if cabeNaProfundidade {
					caixas[numero_caixa].Length += copia.Length
					caixas[numero_caixa].Width = retornaValorMaior(caixas[numero_caixa].Width, copia.Width)
					caixas[numero_caixa].Height = retornaValorMaior(caixas[numero_caixa].Height, copia.Height)
				}
			} else {
				numero_caixa++
				i--
			}
		}
		numero_pacote++
	}
	return caixas
}

func cabeNoLimiteDoPacote(caixas []Caixa, copia Pacote, numero_caixa int, orientacao string) bool {
	var width, height, length float64
	if orientacao == "width" {
		width = caixas[numero_caixa].Width + copia.Width
		length = retornaValorMaior(caixas[numero_caixa].Length, copia.Length)
		height = retornaValorMaior(caixas[numero_caixa].Height, copia.Height)
	}
	if orientacao == "length" {
		length = caixas[numero_caixa].Length + copia.Length
		width = retornaValorMaior(caixas[numero_caixa].Width, copia.Width)
		height = retornaValorMaior(caixas[numero_caixa].Height, copia.Height)
	}
	if orientacao == "height" {
		height = caixas[numero_caixa].Height + copia.Height
		length = retornaValorMaior(caixas[numero_caixa].Length, copia.Length)
		width = retornaValorMaior(caixas[numero_caixa].Width, copia.Width)
	}
	return (width + height + length) <= consts.DIMENSAO_MAXIMA_SOMA
}

func retornaValorMaior(valor1 float64, valor2 float64) float64 {
	if valor1 > valor2 {
		return valor1
	}
	return valor2
}

func validarDimensoesDaCaixa(caixa []Caixa, index int, valor string) bool {
	if valor == "width" {
		if caixa[index].Width > 0 {
			return true
		}
	}
	if valor == "height" {
		if caixa[index].Height > 0 {
			return true
		}
	}
	if valor == "length" {
		if caixa[index].Length > 0 {
			return true
		}
	}
	return false
}
