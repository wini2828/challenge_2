package main

import (
	"fmt"
)

func main() {
	// Data Uji
	scores1Lumba := []int{96, 108, 89}
	scores1Koala := []int{88, 91, 110}

	scores2Lumba := []int{97, 112, 101}
	scores2Koala := []int{109, 95, 123}

	scores3Lumba := []int{97, 112, 101}
	scores3Koala := []int{109, 95, 106}

	// Fungsi untuk menghitung skor rata-rata
	average := func(scores []int) float64 {
		sum := 0
		for _, score := range scores {
			sum += score
		}
		return float64(sum) / float64(len(scores))
	}

	// Menghitung skor rata-rata untuk setiap tim
	avg1Lumba := average(scores1Lumba)
	avg1Koala := average(scores1Koala)

	avg2Lumba := average(scores2Lumba)
	avg2Koala := average(scores2Koala)

	avg3Lumba := average(scores3Lumba)
	avg3Koala := average(scores3Koala)

	// Mencetak skor rata-rata untuk setiap tim
	fmt.Printf("Data 1: Lumba-lumba: %.2f, Koala: %.2f\n", avg1Lumba, avg1Koala)
	fmt.Printf("Data Bonus 1: Lumba-lumba: %.2f, Koala: %.2f\n", avg2Lumba, avg2Koala)
	fmt.Printf("Data Bonus 2: Lumba-lumba: %.2f, Koala: %.2f\n", avg3Lumba, avg3Koala)

	// Menentukan pemenang
	winner := ""
	if avg1Lumba > avg1Koala && avg1Lumba >= 100 {
		winner = "Tim Lumba-lumba"
	} else if avg1Koala > avg1Lumba && avg1Koala >= 100 {
		winner = "Tim Koala"
	} else if avg1Lumba == avg1Koala && avg1Lumba >= 100 && avg1Koala >= 100 {
		winner = "Seri"
	} else {
		winner = "Tidak ada pemenang"
	}
	fmt.Println("Pemenang Data 1:", winner)

	if avg2Lumba > avg2Koala && avg2Lumba >= 100 {
		winner = "Tim Lumba-lumba"
	} else if avg2Koala > avg2Lumba && avg2Koala >= 100 {
		winner = "Tim Koala"
	} else if avg2Lumba == avg2Koala && avg2Lumba >= 100 && avg2Koala >= 100 {
		winner = "Seri"
	} else {
		winner = "Tidak ada pemenang"
	}
	fmt.Println("Pemenang Data Bonus 1:", winner)

	if avg3Lumba > avg3Koala && avg3Lumba >= 100 {
		winner = "Tim Lumba-lumba"
	} else if avg3Koala > avg3Lumba && avg3Koala >= 100 {
		winner = "Tim Koala"
	} else if avg3Lumba == avg3Koala && avg3Lumba >= 100 && avg3Koala >= 100 {
		winner = "Seri"
	} else {
		winner = "Tidak ada pemenang"
	}
	fmt.Println("Pemenang Data Bonus 2:", winner)
}
