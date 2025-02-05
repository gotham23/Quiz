package main

import (
	"fmt"
)

type Ban struct {
	Jenis string
}

type Mobil struct {
	Roda [4]Ban
}

type Pintu struct {
	Posisi string
}

func (m *Mobil) GantiBan(jenis string) {
	for i := range m.Roda {
		m.Roda[i] = Ban{Jenis: jenis}
	}
	fmt.Println("Semua roda diganti ban", jenis)
}

func (p Pintu) Ketuk() {
	if p.Posisi == "kanan" {
		fmt.Println("Knock")
	} else {
		fmt.Println("beep")
	}
}

func (p Pintu) Buka() {
	if p.Posisi == "kanan" {
		fmt.Println("beep")
	} else {
		fmt.Println("Knock")
	}
}

func main() {
	mobil := Mobil{}
	mobil.GantiBan("ban karet")

	pintuKanan := Pintu{Posisi: "kanan"}
	pintuKiri := Pintu{Posisi: "kiri"}

	fmt.Println("Mengetuk kanan:")
	pintuKanan.Ketuk()

	fmt.Println("Membuka kanan:")
	pintuKanan.Buka()

	fmt.Println("Mengetuk kiri:")
	pintuKiri.Ketuk()

	fmt.Println("Membuka kiri:")
	pintuKiri.Buka()
}
