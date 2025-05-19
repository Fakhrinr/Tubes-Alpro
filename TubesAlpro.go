package main

import "fmt"

type saham struct {
	id         string     //Id investment
	nama       string  //nama investment
	sektor       string  //tipe investasi
	perusahaan  string  //nama perusahaan
	danaAwal      float64 //harga investasi awal
	danaAkhir      float64 //total investasi yang diperoleh
	keuntungan float64 //keuntungan investasi
}

const maxSaham int = 50

type tabSaham [maxSaham]saham

func main() {
	var x,nData,pilihan int
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
		ListInvestasi()
	case 3:
		UbahInvestasi()
	case 4:
		LaporanInvestasi()
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
		fmt.Println("2. List Investasi")
		fmt.Println("3. Ubah Investasi")
		fmt.Println("4. Laporan Portofolio")
		fmt.Println("5. Keluar")
		fmt.Println("---------------------")
		fmt.Println("silah pilih menu yang tersedia (1/2/3/4/5):")
}

func MulaiInvestasi(A * tabSaham, nData *int) {
	var x int

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

	fmt.Println("-----------------")
	fmt.Println("Data investasi berhasil disimpan")
	fmt.Println("-----------------")
	*nData++
}
func ListInvestasi(A *tabSaham, nData *int) {
	var x int

	fmt.Println("---------------------")
	fmt.Println("	 INVESTASI")
	fmt.Println("---------------------")
	fmt.Println(" Silahkan pilih sektor investasi saham")
	fmt.Println("1. Cari data berdasarkan Id/ Perusahaan")
	fmt.Println("2. Mengurutkan data berdasarkan Id")
	fmt.Println("3. Mengurutkan data berdasarkan Keuntungan")
	fmt.Println("Silahkan pilih menu (1/2/3):")
	fmt.Scan(&x)
	switch x {
	case 1:
		var y string
		fmt.Println("Silahkan masukkan sektor investasi:")
		fmt.Scan(&y)
		CariDataSektor(A, nData, y)
	case 2:

	

	
}
}

func CariDataSektor(A *tabSaham, nData *int, y string) {
	var i int
	fmt.Println("---------------------")
	fmt.Println("	 DATA SEKTOR")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	for i = 0; i < *nData; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			fmt.Printf("%d\t%s\t%s\t%s\t%.2f\t%.2f\t%.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
		}
	}
	fmt.Println("---------------------")
	fmt.Println("Data investasi berhasil ditemukan")
}

Urut(A *tabSaham, nData *int, id int){
	var i,mid,left,right int
	left = 0
	right = *nData - 1
	for left <= right {
		mid = (left + right) / 2
		ketemu = (A[mid].id == A[i].id)
		if A[i].id < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("---------------------")
	fmt.Println("	 DATA ID")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	
}