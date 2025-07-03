package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	lulusUjian := nilaiAkhir > 80
	lulusAbsensi := absensi > 75

	lulus := lulusUjian && lulusAbsensi

	fmt.Println("Lulus Ujian:", lulusUjian)
	fmt.Println("Lulus Absensi:", lulusAbsensi)
	fmt.Println("Lulus:", lulus)
}
