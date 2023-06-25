package main

type Viagem struct {
	origem  string
	destino string
	data    int
	preço   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	var viagemMax Viagem
	for _, viagem := range viagens {
		if viagem.preço > viagemMax.preço {
			viagemMax = viagem
		}
	}
	return viagemMax
}
