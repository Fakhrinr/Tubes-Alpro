package main

import (
	"fmt"
)

type investasi struct {
	id    int     //Id Number investment
	nama  string  //nama investment
	tipe  string  //tipe investasi
	harga float64 //harga investasi awal
	total float64 //total investasi yang diperoleh
}

const NMAX int = 100

type TabInvesSaham [NMAX]investasi
type TabInvesObligasi [NMAX]investasi
type TabInvesReksadana [NMAX]investasi

var dataReksadana TabInvesReksadana
var dataSaham TabInvesSaham
var dataObligasi TabInvesObligasi
var iReksadana int = 0
var iSaham int = 0
var iObligasi int = 0

func main() {
	fmt.Println("Selamat datang di aplikasi investasi")
	fmt.Println("Silakan pilih menu yang tersedia")

	Menu()
}

func Menu() {
	var pilihan int
	for {
		fmt.Println("---------------------")
		fmt.Println("	MENU")
		fmt.Println("---------------------")
		fmt.Println("1. Tambah Investasi")
		fmt.Println("2. Ubah Investasi")
		fmt.Println("3. Hapus Investasi")
		fmt.Println("4. Cari Investasi")
		fmt.Println("5. Urutkan Investasi")
		fmt.Println("6. Laporan Portofolio")
		fmt.Println("7. Keluar")
		fmt.Println("---------------------")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahInvestasi()
		case 2:
			UbahInvestasi()
		case 3:
			HapusInvestasi()
		case 4:
			CariInvestasi()
		case 5:
			UrutkanInvestasi()
		case 6:
			LaporanPortofolio()
		default:
			fmt.Println("Terima kasih telah menggunakan aplikasi investasi")
		}
		if pilihan == 7 {
			break
		}
	}
}

func TambahInvestasi() {

	fmt.Println("Pilih tipe investasi yang ingin ditambahkan:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	TAMBAH INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("1. Saham")
		fmt.Println("2. Obligasi")
		fmt.Println("3. Reksadana")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------")

		var pilihan int
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahInvesSaham(&dataSaham)
		case 2:
			tambahInvesObligasi(&dataObligasi)
		case 3:
			tambahInvesReksadana(&dataReksadana, &iReksadana)
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			break
		}
		fmt.Println("Apakah anda ingin menambah investasi lagi? (y/n)")
		var lagi string
		fmt.Scan(&lagi)
		if lagi == "n" {
			break
		}
	}
}

func UbahInvestasi() {
	fmt.Println("Pilih tipe investasi yang ingin diubah:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	UBAH INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("1. Saham")
		fmt.Println("2. Obligasi")
		fmt.Println("3. Reksadana")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------")

		var pilihan int
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			ubahInvesSaham()
		case 2:
			ubahInvesObligasi()
		case 3:
			ubahInvesReksadana(&dataReksadana, iReksadana)
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			Menu()
		}
		fmt.Println("Apakah anda ingin mengubah investasi lagi? (y/n)")
		var lagi string
		fmt.Scan(&lagi)
		if lagi == "n" {
			break
		}
	}
}

func HapusInvestasi() {
	fmt.Println("Pilih tipe investasi yang ingin dihapus:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	HAPUS INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("1. Saham")
		fmt.Println("2. Obligasi")
		fmt.Println("3. Reksadana")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------")

		var pilihan int
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			hapusInvesSaham()
		case 2:
			hapusInvesObligasi()
		case 3:
			hapusInvesReksadana()
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			Menu()
		}
		fmt.Println("Apakah anda ingin menghapus investasi lagi? (y/n)")
		var lagi string
		fmt.Scan(&lagi)
		if lagi == "n" {
			break
		}
	}
}

func CariInvestasi() {
	fmt.Println("Pilih tipe investasi yang ingin dicari:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	CARI INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("1. Saham")
		fmt.Println("2. Obligasi")
		fmt.Println("3. Reksadana")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------")

		var pilihan int
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			cariInvesSaham()
		case 2:
			cariInvesObligasi()
		case 3:
			cariInvesReksadana()
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			Menu()
		}
		fmt.Println("Apakah anda ingin mencari investasi lagi? (y/n)")
		var lagi string
		fmt.Scan(&lagi)
		if lagi == "n" {
			break
		}
	}
}

func UrutkanInvestasi() {
	fmt.Println("Pilih tipe investasi yang ingin diurutkan:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	URUTKAN INVESTASI")
		fmt.Println("---------------------")
		fmt.Println("1. Saham")
		fmt.Println("2. Obligasi")
		fmt.Println("3. Reksadana")
		fmt.Println("4. Keluar")
		fmt.Println("---------------------")

		var pilihan int
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			urutInvesSaham()
		case 2:
			urutInvesObligasi()
		case 3:
			urutInvesReksadana()
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			Menu()
		}
		fmt.Println("Apakah anda ingin mengurutkan investasi lagi? (y/n)")
		var lagi string
		fmt.Scan(&lagi)
		if lagi == "n" {
			break
		}
	}
}
func LaporanPortofolio() {
	fmt.Println("Laporan Portofolio")
	fmt.Println("---------------------")
	fmt.Println("1. Laporan Saham")
	fmt.Println("2. Laporan Obligasi")
	fmt.Println("3. Laporan Reksadana")
	fmt.Println("4. Keluar")
	fmt.Println("---------------------")

	var pilihan int
	fmt.Println("Masukkan pilihan anda: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		laporanSaham()
	case 2:
		laporanObligasi()
	case 3:
		laporanReksadana()
	case 4:
		fmt.Println("Kembali ke menu utama")
	}
	if pilihan == 4 {
		Menu()
	}
}

func tambahInvesReksadana(A *TabInvesReksadana, i *int) {
	var cek string
	cek = "yes"
	for *i < NMAX && (cek == "yes" || cek == "Yes" || cek == "YES") {
		fmt.Println("Masukkan ID investasi: ")
		fmt.Scan(&A[*i].id)
		fmt.Println("Masukkan nama investasi: ")
		fmt.Scan(&A[*i].nama)
		fmt.Println("Masukkan tipe investasi: ")
		fmt.Scan(&A[*i].tipe)
		fmt.Println("Masukkan harga investasi awal: ")
		fmt.Scan(&A[*i].harga)
		fmt.Println("Masukkan total investasi yang diperoleh: ")
		fmt.Scan(&A[*i].total)

		fmt.Println("Apakah anda ingin menambah investasi lagi? (yes/no)")
		fmt.Scan(&cek)
		*i++
	}
	fmt.Println("Data investasi reksadana berhasil ditambahkan")
}

func ubahInvesReksadana(A *TabInvesReksadana, x int) {
	fmt.Println("Masukkan ID investasi yang ingin diubah: ")
	var id, i int
	var found bool

	fmt.Scan(&id)
	found = false

	for i = 0; i < x && found == false; i++ {
		if A[i].id == id {
			fmt.Println("Data investasi reksadana berhasil ditemukan, silakan ubah data berikut:")
			fmt.Println("Masukkan nama investasi: ")
			fmt.Scan(&A[i].nama)
			fmt.Println("Masukkan tipe investasi: ")
			fmt.Scan(&A[i].tipe)
			fmt.Println("Masukkan harga investasi awal: ")
			fmt.Scan(&A[i].harga)
			fmt.Println("Masukkan total investasi yang diperoleh: ")
			fmt.Scan(&A[i].total)

			found = true
		}
	}
	if found != false {
		fmt.Println("Data investasi tidak ditemukan")
	}
}
func hapusInvesReksadana(A *TabInvesReksadana, x *int) {
	fmt.Println("Masukkan ID investasi yang ingin dihapus: ")
	var id, i, j int
	var found bool
	fmt.Scan(&id)
	found = false

	for i = 0; i < *x && found == false; i++ {
		if A[i].id == id {
			for j = i; j < *x-1; j++ {
				A[j] = A[j+1]
			}
			fmt.Println("Data investasi reksadana berhasil dihapus")
			found = true
			*x--
		}
		if found == false {
			fmt.Println("Data investasi tidak ditemukan")
		}
	}
}

func cariInvesReksadana(A *TabInvesReksadana, x int) {
	seqInvesReksadana(A, x)
	binaryInvesReksadana(A, x)
}

func seqInvesReksadana(A *TabInvesReksadana, x int) {
	var i int
	var ketemu bool
	var tipe string

	ketemu = false
	fmt.Println("Masukkan ID investasi yang ingin dicari: ")
	fmt.Scan(&tipe)
	for i = 0; i < x && ketemu == false; i++ {
		if A[i].tipe == tipe {
			fmt.Println("Data investasi reksadana berhasil ditemukan")
			fmt.Printf("--------------------------------------------------------------------------\n")
			fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
			fmt.Printf("--------------------------------------------------------------------------\n")

			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].total)
			fmt.Printf("--------------------------------------------------------------------------\n")
			ketemu = true
		}

	}
	if ketemu == false {
		fmt.Println("Data investasi tidak ditemukan")
	}
}
func binaryInvesReksadana(A *TabInvesReksadana, x int) {

}

func urutInvesReksadana(A *TabInvesReksadana) {
	fmt.Println("Pilih tipe pengurutan:")
	fmt.Println("1. ID")
	fmt.Println("2. Nama")
	fmt.Println("3. Tipe")
	fmt.Println("4. Harga")
	fmt.Println("5. Total")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		urutID(A)
	case 2:
		urutNama(A)
	case 3:
		urutTipe(A)
	case 4:
		urutHarga(A)
	case 5:
		urutTotal(A)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func urutID(A *TabInvesReksadana) {
}

func tambahInvesSaham(A *TabInvesSaham, i *int) {
	var cek string
	cek = "yes"
	for *i < NMAX && (cek == "yes" || cek == "Yes" || cek == "YES") {
		fmt.Println("Masukkan ID investasi: ")
		fmt.Scan(&A[*i].id)
		fmt.Println("Masukkan nama investasi: ")
		fmt.Scan(&A[*i].nama)
		fmt.Println("Masukkan tipe investasi: ")
		fmt.Scan(&A[*i].tipe)
		fmt.Println("Masukkan harga investasi awal: ")
		fmt.Scan(&A[*i].harga)
		fmt.Println("Masukkan total investasi yang diperoleh: ")
		fmt.Scan(&A[*i].total)

		fmt.Println("Apakah anda ingin menambah investasi lagi? (yes/no)")
		fmt.Scan(&cek)
		*i++
	}
	fmt.Println("Data investasi reksadana berhasil ditambahkan")
}

func ubahInvesSaham(A *TabInvesSaham, x int) {
	fmt.Println("Masukkan ID investasi yang ingin diubah: ")
	var id, i int
	var found bool

	fmt.Scan(&id)
	found = false

	for i = 0; i < x && found == false; i++ {
		if A[i].id == id {
			fmt.Println("Data investasi reksadana berhasil ditemukan, silakan ubah data berikut:")
			fmt.Println("Masukkan nama investasi: ")
			fmt.Scan(&A[i].nama)
			fmt.Println("Masukkan tipe investasi: ")
			fmt.Scan(&A[i].tipe)
			fmt.Println("Masukkan harga investasi awal: ")
			fmt.Scan(&A[i].harga)
			fmt.Println("Masukkan total investasi yang diperoleh: ")
			fmt.Scan(&A[i].total)

			found = true
		}
	}
	if found != false {
		fmt.Println("Data investasi tidak ditemukan")
	}
}

func hapusInvesSaham(A *TabInvesSaham, x *int) {
	fmt.Println("Masukkan ID investasi yang ingin dihapus: ")
	var id, i, j int
	var found bool
	fmt.Scan(&id)
	found = false

	for i = 0; i < *x && found == false; i++ {
		if A[i].id == id {
			for j = i; j < *x-1; j++ {
				A[j] = A[j+1]
			}
			fmt.Println("Data investasi reksadana berhasil dihapus")
			found = true
			*x--
		}
		if found == false {
			fmt.Println("Data investasi tidak ditemukan")
		}
	}
}
