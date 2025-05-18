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
		fmt.Println("1. Mulai Investasi")
		fmt.Println("2. Tambah Investasi")
		fmt.Println("3. Ubah Investasi")
		fmt.Println("4. Hapus Investasi")
		fmt.Println("5. Cari Investasi")
		fmt.Println("6. Urutkan Investasi")
		fmt.Println("7. Laporan Portofolio")
		fmt.Println("8. Keluar")
		fmt.Println("---------------------")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			MulaiInvestasi()
		case 2:
			TambahInvestasi()
		case 3:
			UbahInvestasi()
		case 4:
			HapusInvestasi()
		case 5:
			CariInvestasi()
		case 6:
			UrutkanInvestasi()
		case 7:
			LaporanPortofolio()
		default:
			fmt.Println("Terima kasih telah menggunakan aplikasi investasi")
		}
		if pilihan == 8 {
			break
		}
	}
}
func MulaiInvestasi() {
	fmt.Println("Pilih tipe investasi yang ingin dimulai:")
	for {
		fmt.Println("---------------------")
		fmt.Println("	MULAI INVESTASI")
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
			BacaInvesSaham(&dataSaham, &iSaham)
		case 2:
			BacaInvesObligasi(&dataObligasi, &iObligasi)
		case 3:
			BacaInvesReksadana(&dataReksadana, &iReksadana)
		case 4:
			fmt.Println("Kembali ke menu utama")
		}
		if pilihan == 4 {
			break
		}
	}
}
func BacaInvesSaham(A *TabInvesSaham, x *int) {
	var cek bool
	var pilihan string
	cek = true
	for cek != false {
		fmt.Println("Silahkan masukkan data investasi saham")
		fmt.Println("Masukkan ID investasi: ")
		fmt.Scan(&A[*x].id)
		fmt.Println("Masukkan nama investasi: ")
		fmt.Scan(&A[*x].nama)
		fmt.Println("Masukkan tipe investasi: ")
		fmt.Scan(&A[*x].tipe)
		fmt.Println("Masukkan harga investasi awal: ")
		fmt.Scan(&A[*x].harga)
		fmt.Println("Masukkan total investasi yang diperoleh: ")
		fmt.Scan(&A[*x].total)
		*x++
		fmt.Println("Apakah anda ingin menambah investasi lagi? (yes/no)")
		fmt.Scan(&pilihan)
		if pilihan == "no" {
			cek = false
		}
	}
}
func BacaInvesObligasi(A *TabInvesObligasi, x *int) {
	var cek bool
	var pilihan string
	cek = true
	for cek != false {
		fmt.Println("Silahkan masukkan data investasi obligasi")
		fmt.Println("Masukkan ID investasi: ")
		fmt.Scan(&A[*x].id)
		fmt.Println("Masukkan nama investasi: ")
		fmt.Scan(&A[*x].nama)
		fmt.Println("Masukkan tipe investasi: ")
		fmt.Scan(&A[*x].tipe)
		fmt.Println("Masukkan harga investasi awal: ")
		fmt.Scan(&A[*x].harga)
		fmt.Println("Masukkan total investasi yang diperoleh: ")
		fmt.Scan(&A[*x].total)
		*x++
		fmt.Println("Apakah anda ingin menambah investasi lagi? (yes/no)")
		fmt.Scan(&pilihan)
		if pilihan == "no" {
			cek = false
		}
	}
}
func BacaInvesReksadana(A *TabInvesReksadana, x *int) {
	var cek bool
	var pilihan string
	cek = true
	for cek != false {
		fmt.Println("Silahkan masukkan data investasi reksadana")
		fmt.Println("Masukkan ID investasi: ")
		fmt.Scan(&A[*x].id)
		fmt.Println("Masukkan nama investasi: ")
		fmt.Scan(&A[*x].nama)
		fmt.Println("Masukkan tipe investasi: ")
		fmt.Scan(&A[*x].tipe)
		fmt.Println("Masukkan harga investasi awal: ")
		fmt.Scan(&A[*x].harga)
		fmt.Println("Masukkan total investasi yang diperoleh: ")
		fmt.Scan(&A[*x].total)
		*x++
		fmt.Println("Apakah anda ingin menambah investasi lagi? (yes/no)")
		fmt.Scan(&pilihan)
		if pilihan == "no" {
			cek = false
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
			tambahInvesSaham(&dataSaham, &iSaham)
		case 2:
			tambahInvesObligasi(&dataObligasi, &iObligasi)
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
	var pilihan int

	fmt.Println("Pilih mana yang mau dicari:")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan ID")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		seqInvesReksadana(A, x)
	case 2:
		binaryInvesReksadana(A, x)
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pencarian yang tesedia")
	}
}

func seqInvesReksadana(A *TabInvesReksadana, x int) {
	var i int
	var ketemu bool
	var nama string

	ketemu = false
	fmt.Println("Masukkan ID investasi yang ingin dicari: ")
	fmt.Scan(&nama)
	for i = 0; i < x && ketemu == false; i++ {
		if A[i].nama == nama {
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

func urutInvesReksadana(A *TabInvesReksadana, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau diurutkan:")
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Persentase Keuntungan")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		urutID(A, x)
	case 2:
		urutPersen(A, x)
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pengurutan yang tesedia")
	}
}

func urutID(A *TabInvesReksadana, x int) {
	var i, j, idx, temp int
	for i = 0; i < x-1; i++ {
		idx = i
		for j = i; j < x; j++ {
			if A[j].id < A[idx].id {
				idx = j
			}
		}
		temp = A[i].id
		A[i].id = A[idx].id
		A[idx].id = temp
	}
}
func urutPersen(A *TabInvesReksadana, x int) {

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
