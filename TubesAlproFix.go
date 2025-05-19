package main

import (
	"fmt"
)

type investasi struct {
	id         int     //Id Number investment
	nama       string  //nama investment
	tipe       string  //tipe investasi
	harga      float64 //harga investasi awal
	total      float64 //total investasi yang diperoleh
	keuntungan float64 //keuntungan investasi
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
			ubahInvesSaham(&dataSaham, iSaham)
		case 2:
			ubahInvesObligasi(&dataObligasi, iObligasi)
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
			hapusInvesSaham(&dataSaham, &iSaham)
		case 2:
			hapusInvesObligasi(&dataObligasi, &iObligasi)
		case 3:
			hapusInvesReksadana(&dataReksadana, &iReksadana)
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
			cariInvesSaham(&dataSaham, iSaham)
		case 2:
			cariInvesObligasi(&dataObligasi, iObligasi)
		case 3:
			cariInvesReksadana(&dataReksadana, iReksadana)
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
			urutInvesSaham(&dataSaham, iSaham)
		case 2:
			urutInvesObligasi(&dataObligasi, iObligasi)
		case 3:
			urutInvesReksadana(&dataReksadana, iReksadana)
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
		laporanSaham(&dataSaham, iSaham)
	case 2:
		laporanObligasi(&dataObligasi, iObligasi)
	case 3:
		laporanReksadana(&dataReksadana, iReksadana)
	case 4:
		fmt.Println("Kembali ke menu utama")
	}
	if pilihan == 4 {
		Menu()
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
	if !found {
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
		if !found {
			fmt.Println("Data investasi tidak ditemukan")
		}
	}
}
func cariInvesSaham(A *TabInvesSaham, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau dicari:")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan ID")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		seqInvesSaham(A, x)
	case 2:
		binaryInvesSaham(A, x)
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pencarian yang tesedia")
	}
}
func seqInvesSaham(A *TabInvesSaham, x int) {
	var i int
	var ketemu bool
	var nama string

	ketemu = false
	fmt.Println("Masukkan nama investasi yang ingin dicari: ")
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
func binaryInvesSaham(A *TabInvesSaham, x int) {
	var i, mid, left, right int
	var ketemu bool

	left = 0
	right = x - 1
	ketemu = false

	for left <= right && !ketemu {
		mid = (left + right) / 2
		ketemu = (A[mid].id == A[i].id)
		if A[i].id < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !ketemu {
		fmt.Println("Data investasi reksadana berhasil ditemukan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[mid].id, A[mid].nama, A[mid].tipe, A[mid].harga, A[mid].total)
	} else {
		fmt.Println("Data investasi tidak ditemukan")
	}
}
func urutInvesSaham(A *TabInvesSaham, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau diurutkan:")
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Persentase Keuntungan")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		var pass, j, idx int
		var temp investasi
		var i int
		for pass = 0; pass < x-1; pass++ {
			idx = pass
			for j = pass + 1; j < x; j++ {
				if A[j].id < A[idx].id {
					idx = j
				}
			}
			temp = A[pass]
			A[pass] = A[idx]
			A[idx] = temp
		}

		fmt.Println("Data investasi reksadana berhasil diurutkan berdasarkan ID")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].total)
		}
		fmt.Printf("--------------------------------------------------------------------------\n")
	case 2:
		var pass, i int
		var temp investasi
		var j int
		pass = 1
		for pass <= x-1 {
			i = pass
			for i > 0 && A[i].keuntungan < A[i-1].keuntungan {
				temp = A[i]
				A[i] = A[i-1]
				A[i-1] = temp
				i--
			}
			pass++
		}
		fmt.Println("Data investasi reksadana berhasil diurutkan berdasarkan keuntungan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for j = 0; j < x; j++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[j].id, A[j].nama, A[j].tipe, A[j].harga, A[j].total)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pengurutan yang tesedia")
	}
}
func laporanSaham(A *TabInvesSaham, x int) {
	var pilihan int
	fmt.Scan(&pilihan)

	fmt.Println("Laporan Saham")
	fmt.Println("Silahkan pilih laporan yang ingin ditampilkan:")
	fmt.Println("1. Laporan Keuntungan Pribandi")
	fmt.Println("2. Laporan Keuntungan Keseluruhan")
	fmt.Println("3. Kembali ke menu utama")

	switch pilihan {
	case 1:
		var i, id int
		var total float64
		total = 0
		fmt.Println("Laporan Keuntungan Pribadi")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Println("Masukkan ID investasi yang ingin dihitung: ")
		fmt.Scan(&id)
		for i = 0; i < x; i++ {
			if A[i].id == id {
				total = (A[i].total / A[i].harga) * 100
				A[i].keuntungan = total
				fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
				fmt.Printf("--------------------------------------------------------------------------\n")
			}
		}
	case 2:
		var i int
		var total float64
		fmt.Println("Laporan Keuntungan Keseluruhan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			total = (A[i].total / A[i].harga) * 100
			A[i].keuntungan = total
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}
	case 3:
		fmt.Println("Kembali ke menu utama")
	default:
		fmt.Println("Pilihan tidak valid")
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
func tambahInvesObligasi(A *TabInvesObligasi, i *int) {
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
func ubahInvesObligasi(A *TabInvesObligasi, x int) {
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
func hapusInvesObligasi(A *TabInvesObligasi, x *int) {
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
func cariInvesObligasi(A *TabInvesObligasi, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau dicari:")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan ID")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		seqInvesObligasi(A, x)
	case 2:
		binaryInvesObligasi(A, x)
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pencarian yang tesedia")
	}
}
func seqInvesObligasi(A *TabInvesObligasi, x int) {
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
func binaryInvesObligasi(A *TabInvesObligasi, x int) {
	var i, mid, left, right int
	var ketemu bool

	left = 0
	right = x - 1
	ketemu = false

	for left <= right && !ketemu {
		mid = (left + right) / 2
		ketemu = (A[mid].id == A[i].id)
		if A[i].id < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !ketemu {
		fmt.Println("Data investasi reksadana berhasil ditemukan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[mid].id, A[mid].nama, A[mid].tipe, A[mid].harga, A[mid].total)
	} else {
		fmt.Println("Data investasi tidak ditemukan")
	}
}
func urutInvesObligasi(A *TabInvesObligasi, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau diurutkan:")
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Persentase Keuntungan")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		var pass, j, idx int
		var temp investasi
		var i int
		for pass = 0; pass < x-1; pass++ {
			idx = pass
			for j = pass + 1; j < x; j++ {
				if A[j].id < A[idx].id {
					idx = j
				}
			}
			temp = A[pass]
			A[pass] = A[idx]
			A[idx] = temp
		}

		fmt.Println("Data investasi obligasi berhasil diurutkan berdasarkan ID")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].total)
		}
		fmt.Printf("--------------------------------------------------------------------------\n")
	case 2:
		var pass, i int
		var temp investasi
		var j int
		pass = 1
		for pass <= x-1 {
			i = pass
			for i > 0 && A[i].keuntungan < A[i-1].keuntungan {
				temp = A[i]
				A[i] = A[i-1]
				A[i-1] = temp
				i--
			}
			pass++
		}
		fmt.Println("Data investasi obligasi berhasil diurutkan berdasarkan keuntungan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for j = 0; j < x; j++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[j].id, A[j].nama, A[j].tipe, A[j].harga, A[j].total)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pengurutan yang tesedia")
	}
}
func laporanObligasi(A *TabInvesObligasi, x int) {
	var pilihan int
	fmt.Scan(&pilihan)

	fmt.Println("Laporan Obligasi")
	fmt.Println("Silahkan pilih laporan yang ingin ditampilkan:")
	fmt.Println("1. Laporan Keuntungan Pribandi")
	fmt.Println("2. Laporan Keuntungan Keseluruhan")
	fmt.Println("3. Kembali ke menu utama")

	switch pilihan {
	case 1:
		var i, id int
		var total float64
		total = 0
		fmt.Println("Laporan Keuntungan Pribadi")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Println("Masukkan ID investasi yang ingin dihitung: ")
		fmt.Scan(&id)
		for i = 0; i < x; i++ {
			if A[i].id == id {
				total = (A[i].total / A[i].harga) * 100
				A[i].keuntungan = total
				fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
				fmt.Printf("--------------------------------------------------------------------------\n")
			}
		}
	case 2:
		var i int
		var total float64
		fmt.Println("Laporan Keuntungan Keseluruhan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			total = (A[i].total / A[i].harga) * 100
			A[i].keuntungan = total
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}
	case 3:
		fmt.Println("Kembali ke menu utama")
	default:
		fmt.Println("Pilihan tidak valid")
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
	var i, mid, left, right int
	var ketemu bool

	left = 0
	right = x - 1
	ketemu = false

	for left <= right && !ketemu {
		mid = (left + right) / 2
		ketemu = (A[mid].id == A[i].id)
		if A[i].id < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !ketemu {
		fmt.Println("Data investasi reksadana berhasil ditemukan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[mid].id, A[mid].nama, A[mid].tipe, A[mid].harga, A[mid].total)
	} else {
		fmt.Println("Data investasi tidak ditemukan")
	}
}

func urutInvesReksadana(A *TabInvesReksadana, x int) {
	var pilihan int

	fmt.Println("Pilih mana yang mau diurutkan:")
	fmt.Println("1. Berdasarkan ID")
	fmt.Println("2. Persentase Keuntungan")

	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		var pass, j, idx int
		var temp investasi
		var i int
		for pass = 0; pass < x-1; pass++ {
			idx = pass
			for j = pass + 1; j < x; j++ {
				if A[j].id < A[idx].id {
					idx = j
				}
			}
			temp = A[pass]
			A[pass] = A[idx]
			A[idx] = temp
		}

		fmt.Println("Data investasi reksadana berhasil diurutkan berdasarkan ID")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].total)
		}
		fmt.Printf("--------------------------------------------------------------------------\n")

	case 2:
		var pass, i int
		var temp investasi
		var j int
		pass = 1
		for pass <= x-1 {
			i = pass
			for i > 0 && A[i].keuntungan < A[i-1].keuntungan {
				temp = A[i]
				A[i] = A[i-1]
				A[i-1] = temp
				i--
			}
			pass++
		}
		fmt.Println("Data investasi reksadana berhasil diurutkan berdasarkan keuntungan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Total         |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for j = 0; j < x; j++ {
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[j].id, A[j].nama, A[j].tipe, A[j].harga, A[j].total)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}
	default:
		fmt.Println("Pilihan tidak valid")
		fmt.Println("Silahkan pilih metode pengurutan yang tesedia")
	}
}
func laporanReksadana(A *TabInvesReksadana, x int) {
	var pilihan int
	fmt.Scan(&pilihan)

	fmt.Println("Laporan Reksadana")
	fmt.Println("Silahkan pilih laporan yang ingin ditampilkan:")
	fmt.Println("1. Laporan Keuntungan Pribadi")
	fmt.Println("2. Laporan Keuntungan Keseluruhan")
	fmt.Println("3. Kembali ke menu utama")

	switch pilihan {
	case 1:
		var i, id int
		var total float64
		total = 0
		fmt.Println("Laporan Keuntungan Pribadi")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Println("Masukkan ID investasi yang ingin dihitung: ")
		fmt.Scan(&id)
		for i = 0; i < x; i++ {
			if A[i].id == id {
				total = ((A[i].total - A[i].harga) / A[i].harga) * 100
				A[i].keuntungan = total
				fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
				fmt.Printf("--------------------------------------------------------------------------\n")
			}
		}
	case 2:
		var i int
		var total float64
		fmt.Println("Laporan Keuntungan Keseluruhan")
		fmt.Printf("--------------------------------------------------------------------------\n")
		fmt.Printf("| ID | Nama                 | Tipe       | Harga         | Keuntungan    |\n")
		fmt.Printf("--------------------------------------------------------------------------\n")
		for i = 0; i < x; i++ {
			total = ((A[i].total - A[i].harga) / A[i].harga) * 100
			A[i].keuntungan = total
			fmt.Printf("| %2d | %-20s | %-10s | %-12.2f | %-12.2f |\n", A[i].id, A[i].nama, A[i].tipe, A[i].harga, A[i].keuntungan)
			fmt.Printf("--------------------------------------------------------------------------\n")
		}

	case 3:
		fmt.Println("Kembali ke menu utama")
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func hitungObligasi() {
	var id int
	fmt.Println("Masukkan ID investasi obligasi yang ingin dihitung: ")
	fmt.Scan(&id)

	found := false
	for i := 0; i < iObligasi; i++ {
		if dataObligasi[i].id == id {
			fmt.Println("Data ditemukan. Menghitung total obligasi...")

			//dataObligasi[i].total = dataObligasi[i].harga + (dataObligasi[i].harga * dataObligasi[i].bunga * float64(dataObligasi[i].jangkawaktu))

			fmt.Println("Perhitungan berhasil:")
			fmt.Printf("-------------------------------------------------------------------\n")
			fmt.Printf("| ID | Nama                 | Tipe     | Harga       | Total       |\n")
			fmt.Printf("-------------------------------------------------------------------\n")
			fmt.Printf("| %2d | %-20s | %-8s | %-10.2f | %-10.2f |\n", dataObligasi[i].id, dataObligasi[i].nama, dataObligasi[i].tipe, dataObligasi[i].harga, dataObligasi[i].total)
			fmt.Printf("-------------------------------------------------------------------\n")
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Data investasi tidak ditemukan.")
	}
}
