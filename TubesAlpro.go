package main

import "fmt"

type saham struct {
	id         string  //Id investment
	nama       string  //nama investment
	sektor     string  //tipe investasi
	perusahaan string  //nama perusahaan
	danaAwal   float64 //harga investasi awal
	danaAkhir  float64 //total investasi yang diperoleh
	keuntungan float64 //keuntungan investasi
}

const maxSaham int = 50

type tabSaham [maxSaham]saham

func main() {
	var nData, pilihan int
	var inves tabSaham
	fmt.Println("Selamat datang di aplikasi investasi Saham")
	fmt.Println("Silakan pilih menu yang tersedia")
	for {
		Menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			MulaiInvestasi(&inves, &nData)
		case 2:
			LihatInvestasi(&inves, &nData)
		case 3:
			editInvestasi(&inves, &nData)
		case 4:
			LaporanInvestasi(&inves, &nData)
		default:
			fmt.Println("Terima kasih telah menggunakan aplikasi investasi")
		}
		if pilihan == 5 {
			break
		}
	}
}

func Menu() {
	fmt.Println("---------------------")
	fmt.Println("	MENU UTAMA")
	fmt.Println("---------------------")
	fmt.Println("1. Mulai Investasi")
	fmt.Println("2. Lihat Investasi")
	fmt.Println("3. Edit Investasi")
	fmt.Println("4. Laporan Portofolio")
	fmt.Println("5. Keluar")
	fmt.Println("---------------------")
	fmt.Println("silah pilih menu yang tersedia (1/2/3/4/5):")
}

func MulaiInvestasi(A *tabSaham, nData *int) {
	var x int
	var y string
	y = "yes"
	for y != "no" {
		fmt.Println("---------------------")
		fmt.Println("	MULAI INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("Silahkan pilih sektor investasi saham")
		fmt.Println("1. Teknologi")
		fmt.Println("2. Keuangan")
		fmt.Println("3. Konsumsi")
		fmt.Println("Pilih sektor investasi (1/2/3):")
		fmt.Println("---------------------")
		fmt.Scan(&x)
		switch x {
		case 1:
			A[*nData].sektor = "Teknologi"
		case 2:
			A[*nData].sektor = "Keuangan"
		case 3:
			A[*nData].sektor = "Konsumsi"

		}

		fmt.Println("Silahkan masukkan id investasi:")
		fmt.Scan(&A[*nData].id)
		fmt.Println("Silahkan masukkan nama investasi:")
		fmt.Scan(&A[*nData].nama)
		fmt.Println("Silahkan masukkan nama perusahaan:")
		fmt.Scan(&A[*nData].perusahaan)
		fmt.Println("Silahkan masukkan dana awal investasi:")
		fmt.Scan(&A[*nData].danaAwal)
		fmt.Println("Silahkan masukkan dana akhir investasi:")
		fmt.Scan(&A[*nData].danaAkhir)

		*nData++

		fmt.Println("Apakah anda ingin menambah data investasi lagi? (yes/no):")
		fmt.Scan(&y)
	}
	fmt.Println("-----------------")
	fmt.Println("Data investasi berhasil disimpan")
	fmt.Println("-----------------")
}
func LihatInvestasi(A *tabSaham, nData *int) {
	var x, i int
	cetakData(A, nData)
	fmt.Println("---------------------")
	fmt.Println("	DATA INVESTASI")
	fmt.Println("---------------------")
	for i = 0; i < *nData; i++ {
		fmt.Printf("%s %s %s %s %.2f %.2f %.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
	}
	fmt.Println("---------------------")
	fmt.Println("	 INVESTASI")
	fmt.Println("---------------------")
	fmt.Println(" Silahkan pilih sektor investasi saham")
	fmt.Println("1. Cari data berdasarkan Id/ Perusahaan")
	fmt.Println("2. Mengurutkan data berdasarkan Keutungan")
	fmt.Println("3. Mengurutkan data berdasarkan investasi terbesar")
	fmt.Println("Silahkan pilih menu (1/2/3):")
	fmt.Scan(&x)
	switch x {
	case 1:
		var y string
		fmt.Println("Silahkan masukkan sektor investasi:")
		fmt.Scan(&y)
		CariDataSektor(A, nData, y)
	case 2:
		var y string
		hitungUntung(A, nData)
		fmt.Println("Silahkan masukkan sektor investasi:")
		fmt.Scan(&y)
		UrutKeuntungan(A, nData, y)
	case 3:
		var y string
		fmt.Println("Silahkan masukkan sektor investasi yang ingin dicari :")
		fmt.Scan(&y)

	}
}

func editInvestasi(A *tabSaham, nData *int) {
	var x int
	cetakData(A, nData)
	fmt.Println("---------------------")
	fmt.Println("	EDIT INVESTASI")
	fmt.Println("---------------------")
	fmt.Println("Silahkan pilih data yang ingin di edit")
	fmt.Println("1. Edit investasi")
	fmt.Println("2. Hapus investasi")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Println("Silahkan pilih menu (1/2/3):")
	fmt.Scan(&x)
	switch x {
	case 1:
		var y string
		fmt.Println("Silahkan masukkan id / sektor investasi yang ingin di edit:")
		fmt.Scan(&y)
		ubahInvestasi(A, nData, y)
	case 2:
		var y string
		fmt.Println("Silahkan masukkan id / sektor investasi yang ingin di hapus:")
		fmt.Scan(&y)
		hapusInvestasi(A, nData, y)
	case 3:
		fmt.Println("Kembali ke menu utama")
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func hapusInvestasi(A *tabSaham, nData *int, y string) {
	var i, j int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("	HAPUS DATA INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	for i = 0; i < *nData; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			found = true
			fmt.Printf("%s %s %s %s %.2f %.2f %.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil ditemukan")
			for j = i; j < *nData-1; j++ {
				A[j] = A[j+1]
			}
			*nData--
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil dihapus")
		}
	}
	if !found {
		fmt.Println("Data investasi tidak ditemukan")
		fmt.Println("Mohon masukan data yang benar")
		fmt.Println("---------------------")
	}
}
func ubahInvestasi(A *tabSaham, nData *int, y string) {
	var i int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("	EDIT DATA INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	found = false
	for i = 0; i < *nData; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			found = true
			fmt.Printf("%s %s %s %s %.2f %.2f %.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil ditemukan")
			fmt.Println("Silahkan masukkan data yang ingin diubah:")
			fmt.Println("Silahkan masukkan id investasi:")
			fmt.Scan(&A[i].id)
			fmt.Println("Silahkan masukkan nama investasi:")
			fmt.Scan(&A[i].nama)
			fmt.Println("Silahkan masukkan nama perusahaan:")
			fmt.Scan(&A[i].perusahaan)
			fmt.Println("Silahkan masukkan dana awal investasi:")
			fmt.Scan(&A[i].danaAwal)
			fmt.Println("Silahkan masukkan dana akhir investasi:")
			fmt.Scan(&A[i].danaAkhir)
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil diubah")
			fmt.Println("---------------------")
		}
	}
	if !found {
		fmt.Println("Data investasi tidak ditemukan")
		fmt.Println("Mohon masukan data yang benar")
		fmt.Println("---------------------")
	}
}
func CariDataSektor(A *tabSaham, nData *int, y string) {
	var i int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("	 DATA SEKTOR")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	found = false
	for i = 0; i < *nData && found == false; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			found = true
			fmt.Printf("%s %s %s %s %.2f %.2f %.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil ditemukan")
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func UrutKeuntungan(A *tabSaham, nData *int, y string) {
	var i, pass int
	var temp saham
	pass = 1
	for pass = 0; pass <= *nData-1; pass++ {
		i = pass
		for i < *nData-1 && A[i].keuntungan < A[i+1].keuntungan {
			temp = A[i]
			A[i] = A[i+1]
			A[i+1] = temp
			i--
		}
	}
}

func hitungUntung(A *tabSaham, nData *int) {
	var i int
	for i = 0; i < *nData; i++ {
		A[i].keuntungan = ((A[i].danaAkhir) * 0.001) - A[i].danaAwal
	}

}

func cetakData(A *tabSaham, nData *int) {
	var i int
	fmt.Printf("| ID   | Nama   | Sektor    | Perusahaan   | danaAwal  | danaAkhir |\n")
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %s | %s | %s | %s       | %.2f   | %.2f |\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
	}
}

func LaporanInvestasi(A *tabSaham, nData *int) {
	var i int
	var totalKeuntungan float64 = 0
	var keuntunganKotor float64 = 0

	fmt.Println("---------------------")
	fmt.Println(" LAPORAN INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| ID   | Nama   | Sektor    | Perusahaan   | danaAwal  | danaAkhir | Keuntungan Kotor |\n")
	for i = 0; i < *nData; i++ {
		keuntunganKotor = A[i].danaAkhir - A[i].danaAwal
		totalKeuntungan += keuntunganKotor
		fmt.Printf("| %s | %s | %s | %s | %.2f | %.2f | %.2f |\n",
			A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, keuntunganKotor)
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Total Keuntungan : %.2f\n", totalKeuntungan)
	fmt.Println("--------------------------------------------------------------------------")
}
