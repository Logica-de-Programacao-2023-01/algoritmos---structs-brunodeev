package main

import "fmt"

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo  string
	Artista string
	Duracao float64
}

func main() {
	playlist := Playlist{
		Nome: "AO SOM DE LENO BREGA",
		Musicas: []Musica{
			{
				Titulo:  "Certa Noite em Campo Grande",
				Duracao: 6.50,
				Artista: "Leno Brega",
			},
			{
				Titulo:  "Something",
				Artista: "The Beatles",
				Duracao: 4.20,
			},
			{
				Titulo:  "While my Guitar Gently Weeps",
				Artista: "The Beatles",
				Duracao: 5.0,
			},
		},
	}
	fmt.Printf("Tempo da playlist %f\n", somador(playlist))
	for _, musica := range playlist.Musicas {
		fmt.Printf("Título da canção: %s Artista: %s, Duração: %f\n", musica.Titulo, musica.Artista, musica.Duracao)
	}
}

func somador(p Playlist) float64 {
	somaplay := 0.0
	for _, musica := range p.Musicas {
		somaplay += musica.Duracao
	}
	return somaplay
}
