package main

import (
	"fmt"
	"math"
)

func main() {
	// Data Uji
	// Data 1
	beratMarkData1 := 78.0  // kg
	tinggiMarkData1 := 1.69 // m
	beratJohnData1 := 92.0  // kg
	tinggiJohnData1 := 1.95 // m

	// Data 2
	beratMarkData2 := 95.0  // kg
	tinggiMarkData2 := 1.88 // m
	beratJohnData2 := 85.0  // kg
	tinggiJohnData2 := 1.76 // m

	// Hitung kedua BMI mereka menggunakan rumus
	// Rumus: BMI = massa / (tinggi * tinggi)
	bmiMarkData1 := beratMarkData1 / math.Pow(tinggiMarkData1, 2)
	bmiJohnData1 := beratJohnData1 / math.Pow(tinggiJohnData1, 2)

	bmiMarkData2 := beratMarkData2 / math.Pow(tinggiMarkData2, 2)
	bmiJohnData2 := beratJohnData2 / math.Pow(tinggiJohnData2, 2)

	// Cetak hasil dengan membatasi jumlah angka dibelakang koma
	fmt.Println("Data 1:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMarkData1)
	fmt.Printf("BMI John: %.2f\n", bmiJohnData1)
	fmt.Println("Mark memiliki BMI lebih tinggi dari John:", bmiMarkData1 > bmiJohnData1)

	fmt.Println("\nData 2:")
	fmt.Printf("BMI Mark: %.2f\n", bmiMarkData2)
	fmt.Printf("BMI John: %.2f\n", bmiJohnData2)
	fmt.Println("Mark memiliki BMI lebih tinggi dari John:", bmiMarkData2 > bmiJohnData2)
}
