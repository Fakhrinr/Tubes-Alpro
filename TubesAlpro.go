package main

import "fmt"

type saham struct {
	IDMember   int     //id member
	id         string  //Id investment
	nama       string  //nama investor
	sektor     string  //tipe investasi
	perusahaan string  //nama perusahaan
	danaAwal   float64 //harga investasi awal
	danaAkhir  float64 //total investasi yang diperoleh
	keuntungan float64 //keuntungan investasi
}

const maxSaham int = 3

type tabSaham [maxSaham]saham

var ID int = 100

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
			laporanPortofolio(&inves, &nData)
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
	fmt.Print("silahkan pilih menu yang tersedia (1/2/3/4/5):")
}

func MulaiInvestasi(A *tabSaham, nData *int) {
	var x int
	var y string
	for y != "no" {
		if *nData >= maxSaham {
			fmt.Println("Data investasi sudah penuh, tidak bisa menambah data lagi.")
			y = "no"
		} else {
			fmt.Println("---------------------")
			fmt.Println("MULAI INVESTASI")
			fmt.Println("---------------------")
			fmt.Println("Silahkan pilih sektor investasi saham")
			fmt.Println("1. Teknologi")
			fmt.Println("2. Keuangan")
			fmt.Println("3. Konsumsi")
			fmt.Println("4. Kembali ke menu utama")
			fmt.Println("---------------------")
			fmt.Print("Pilih sektor investasi (1/2/3/4):")
			fmt.Scan(&x)

			switch x {
			case 1:
				A[*nData].sektor = "Teknologi"
			case 2:
				A[*nData].sektor = "Keuangan"
			case 3:
				A[*nData].sektor = "Konsumsi"
			case 4:
				fmt.Println("Kembali ke menu utama")
				return
			default:
				fmt.Println("Input Tidak Valid")
				return
			}

			fmt.Print("Silahkan masukkan Nama investor : ")
			fmt.Scan(&A[*nData].nama)
			fmt.Print("Silahkan masukkan ID perusahaan investasi : ")
			fmt.Scan(&A[*nData].id)
			fmt.Print("Silahkan masukkan nama perusahaan : ")
			fmt.Scan(&A[*nData].perusahaan)
			fmt.Print("Silahkan masukkan dana awal investasi : ")
			fmt.Scan(&A[*nData].danaAwal)
			fmt.Print("Silahkan masukkan dana akhir investasi : ")
			fmt.Scan(&A[*nData].danaAkhir)
			A[*nData].IDMember = BuatID()
			*nData++

			fmt.Print("Apakah anda ingin menambah data investasi lagi? (yes/no) : ")
			fmt.Scan(&y)
		}
		fmt.Println("-----------------")
		fmt.Println("Data investasi berhasil disimpan")
		fmt.Println("-----------------")
	}
}

func BuatID() int {

	ID++
	return ID
}
func LihatInvestasi(A *tabSaham, nData *int) {
	var x int
	fmt.Println("---------------------")
	fmt.Println("DATA INVESTASI")
	fmt.Println("---------------------")
	if *nData == 0 {
		fmt.Println("Tidak ada data investasi yang tersedia.")
		return
	}
	cetakData(A, nData)
	fmt.Println("---------------------")
	fmt.Println("	 INVESTASI")
	fmt.Println("---------------------")
	fmt.Println(" Silahkan pilih sektor investasi saham")
	fmt.Println("1. Cari data berdasarkan Id/ Perusahaan")
	fmt.Println("2. Mengurutkan data berdasarkan Keutungan")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Println("---------------------")
	fmt.Print("Silahkan pilih menu (1/2/3/4): ")
	fmt.Scan(&x)
	switch x {
	case 1:
		var sektor string
		var y int
		fmt.Println("Silahkan masukkan sektor investasi:")
		fmt.Println("1. Teknologi")
		fmt.Println("2. Keuangan")
		fmt.Println("3. Konsumsi")
		fmt.Println("4. Berdasarkan ID Member")
		fmt.Println("5. Kembali ke menu utama")
		fmt.Println("Masukan pilihan sektor")
		fmt.Print("(1/2/3/4/5) : ")
		fmt.Scan(&y)
		if y == 1 {
			sektor = "Teknologi"
			CariDataSektor(A, nData, sektor)
		} else if y == 2 {
			sektor = "Keuangan"
			CariDataSektor(A, nData, sektor)
		} else if y == 3 {
			sektor = "Konsumsi"
			CariDataSektor(A, nData, sektor)
		} else if y == 4 {
			CariDataMember(A, nData)
		} else if y == 5 {
			fmt.Println("Kembali ke menu utama")
			return
		} else {
			fmt.Println("Input Tidak Valid")
			return
		}
	case 2:
		var y int
		fmt.Println("Silahkan pilih")
		fmt.Println("1. Mengurutkan berdasarkan nominal / keuntungan dari terkecil")
		fmt.Println("2. Mengurutkan berdasarkan nominal / keuntungan dari terbesar")
		fmt.Println("3. Kembali ke menu utama")
		fmt.Println("Silahkan masukan pilihan:")
		fmt.Print("(1/2/3) : ")
		fmt.Scan(&y)
		if y == 1 {
			UrutKeuntunganKecil(A, nData, y)
		} else if y == 2 {
			urutInvestasiBesar(A, nData, y)
		} else {
			fmt.Println("Input Tidak Valid")
			return
		}
	case 3:
		fmt.Println("Kembali ke menu utama")
		return
	}
}

func editInvestasi(A *tabSaham, nData *int) {
	var x int
	if *nData == 0 {
		fmt.Println("Tidak ada data investasi yang tersedia.")
		return
	}
	cetakData(A, nData)
	fmt.Println("---------------------")
	fmt.Println("	EDIT INVESTASI")
	fmt.Println("---------------------")
	fmt.Println("Silahkan pilih data yang ingin di edit")
	fmt.Println("1. Edit investasi")
	fmt.Println("2. Hapus investasi")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Silahkan pilih menu (1/2/3) : ")
	fmt.Scan(&x)
	switch x {
	case 1:
		var y int
		fmt.Print("Silahkan masukkan ID Member yang ingin di edit : ")
		fmt.Scan(&y)
		ubahInvestasi(A, nData, y)
	case 2:
		var y int
		fmt.Print("Silahkan masukkan ID Member yang ingin di hapus : ")
		fmt.Scan(&y)
		hapusInvestasi(A, nData, y)
	case 3:
		fmt.Println("Kembali ke menu utama")
		return
	default:
		fmt.Println("Menu tidak tersedia")
	}
}

func hapusInvestasi(A *tabSaham, nData *int, y int) {
	var i, j int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("	HAPUS DATA INVESTASI")
	fmt.Printf("| %-3s | %-10s | %-10s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "No", "IDMember", "ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	for i = 0; i < *nData; i++ {
		if A[i].IDMember == y {
			found = true
			fmt.Printf("| %-3d | %-10d | %-10s | %-15s | %-10s | %-15s | %-12.2f | %-12.2f |\n", i+1, A[i].IDMember, A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Println("Data investasi berhasil ditemukan")
			for j = i; j < *nData-1; j++ {
				A[j] = A[j+1]
			}
			*nData--
			fmt.Println("Data investasi berhasil dihapus")
		}
	}
	if !found {
		fmt.Println("Data investasi tidak ditemukan")
		fmt.Println("Mohon masukan data yang benar")
		fmt.Println("---------------------")
	}
}
func ubahInvestasi(A *tabSaham, nData *int, y int) {
	var i, left, mid, right int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("EDIT DATA INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| %-3s | %-10s | %-10s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "No", "IDMember", "ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	found = false

	left = 0
	right = *nData - 1
	for left <= right && !found {
		mid = (left + right) / 2
		if A[mid].IDMember == y {
			found = true
			i = mid
			fmt.Printf("| %-3d | %-10d | %-10s | %-15s | %-10s | %-15s | %-12.2f | %-12.2f |\n",
				i+1, A[mid].IDMember, A[mid].id, A[mid].nama, A[mid].sektor, A[mid].perusahaan, A[mid].danaAwal, A[mid].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Println("Data investasi berhasil ditemukan")
			fmt.Println("Silahkan masukkan data yang ingin diubah:")
			fmt.Print("Silahkan masukkan id investasi : ")
			fmt.Scan(&A[mid].id)
			fmt.Print("Silahkan masukkan nama investasi : ")
			fmt.Scan(&A[mid].nama)
			fmt.Print("Silahkan masukkan nama perusahaan : ")
			fmt.Scan(&A[mid].perusahaan)
			fmt.Print("Silahkan masukkan dana awal investasi : ")
			fmt.Scan(&A[mid].danaAwal)
			fmt.Print("Silahkan masukkan dana akhir investasi : ")
			fmt.Scan(&A[mid].danaAkhir)
			fmt.Println("---------------------")
			fmt.Println("Data investasi berhasil diubah")
			fmt.Println("---------------------")
		} else if A[mid].IDMember < y {
			left = mid + 1
		} else {
			right = mid - 1
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
	fmt.Println(" DATA SEKTOR")
	fmt.Println("---------------------")
	fmt.Printf("| %-10s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	found = false
	for i = 0; i < *nData; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			found = true
			fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n",
				A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
		}
	}
	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("Data investasi berhasil ditemukan")
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}
func CariDataMember(A *tabSaham, nData *int) {
	var mid, left, right int
	var found bool
	var y int

	left = 0
	right = *nData - 1

	fmt.Println("---------------------")
	fmt.Println("CARI DATA MEMBER")
	fmt.Println("---------------------")
	found = false
	fmt.Print("Silahkan masukkan ID Member yang ingin dicari : ")
	fmt.Scan(&y)
	for left <= right && !found {
		mid = (left + right) / 2
		if A[mid].IDMember == y {
			found = true
			fmt.Printf("| ID Perusahaan | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n", A[mid].id, A[mid].nama, A[mid].sektor, A[mid].perusahaan, A[mid].danaAwal, A[mid].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
		} else if A[mid].IDMember < y {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func laporanPortofolio(A *tabSaham, nData *int) {
	var x int
	fmt.Println("---------------------")
	fmt.Println(" LAPORAN PORTOFOLIO")
	fmt.Println("---------------------")
	if *nData == 0 {
		fmt.Println("Tidak ada data investasi yang tersedia.")
		return
	}
	fmt.Println("siahkan pilih menu yang ingin ditampilkan")
	fmt.Println("1. Laporan Investasi sektor")
	fmt.Println("2. Laporan Investasi keseluruhan")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Silahkan pilih menu (1/2/3/4):")
	fmt.Scan(&x)

	if x == 1 {
		var sektor string
		fmt.Println("Silahkan masukkan sektor investasi:")
		fmt.Println("1. Teknologi")
		fmt.Println("2. Keuangan")
		fmt.Println("3. Konsumsi")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Println("Masukan pilihan sektor")
		fmt.Print("(1/2/3/4) : ")
		fmt.Scan(&x)
		if x == 1 {
			sektor = "Teknologi"
		} else if x == 2 {
			sektor = "Keuangan"
		} else if x == 3 {
			sektor = "Konsumsi"
		} else if x == 4 {
			fmt.Println("Kembali ke menu utama")
			return
		}
		laporanSektor(A, nData, sektor)
	} else if x == 2 {
		LaporanInvestasi(A, nData)
	} else if x == 3 {
		fmt.Println("Kembali ke menu utama")
		return
	} else {
		fmt.Println("Input Tidak Valid")
		return
	}

}

func laporanSektor(A *tabSaham, nData *int, sektor string) {
	var i int
	var found bool
	var totalKeuntunganKotor, totalKeuntunganBersih float64
	var keuntunganKotor, keuntunganBersih float64

	fmt.Println("---------------------")
	fmt.Println(" LAPORAN INVESTASI SEKTOR")
	fmt.Println("---------------------")
	fmt.Printf("| %-10s | %-15s | %-10s | %-15s | %-12s | %-12s | %-17s |\n",
		"ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir", "Keuntungan Kotor")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	for i = 0; i < *nData; i++ {
		if A[i].sektor == sektor {
			found = true
			hitungUntung(A, nData)
			keuntunganKotor = A[i].danaAkhir - A[i].danaAwal
			keuntunganBersih = A[i].keuntungan
			totalKeuntunganKotor += keuntunganKotor
			totalKeuntunganBersih += keuntunganBersih
			fmt.Printf("| %-10s | %-15s | %-10s | %-15s | %12.2f | %12.2f | %17.2f |\n",
				A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, keuntunganKotor)
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
		fmt.Println("Mohon masukan data yang benar")
		return
	}
}
func UrutKeuntunganKecil(A *tabSaham, nData *int, y int) {
	var i, pass int
	var temp saham

	pass = 1
	hitungUntung(A, nData)
	for pass < *nData {
		i = pass
		temp = A[pass]
		for i > 0 && temp.keuntungan < A[i-1].keuntungan {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	fmt.Println("---------------------")
	fmt.Println("	DATA INVESTASI")
	fmt.Println("---------------------")
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %-10s | %-15s | %-12s | %-18s | %12.2f | %12.2f | %14.2f |\n",
			A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, A[i].keuntungan)
	}
}

func urutInvestasiBesar(A *tabSaham, nData *int, y int) {
	var i, idx, pass int
	var temp saham

	hitungUntung(A, nData)
	for pass = 1; pass < *nData-1; pass++ {
		idx = pass
		i = pass - 1
		for i = pass; i < *nData; i++ {
			if A[i].danaAwal > A[idx].danaAwal {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
	}
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %-3d | %-10d | %-10s | %-15s | %-10s | %-15s | %-12f | %-12f | %12f |\n", i+1, A[i].IDMember, A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, A[i].keuntungan)
	}
}

func hitungUntung(A *tabSaham, nData *int) {
	var i int
	var keuntunganKotor, keuntunganBersih float64
	keuntunganBersih = 0
	keuntunganKotor = 0

	for i = 0; i < *nData; i++ {
		keuntunganKotor = A[i].danaAkhir - A[i].danaAwal
		keuntunganBersih = (keuntunganKotor * 0.9) / 100
		A[i].keuntungan = keuntunganBersih
	}

}

func cetakData(A *tabSaham, nData *int) {
	var i int
	if *nData == 0 {
		fmt.Println("Tidak ada data investasi yang tersedia.")
		return
	}
	fmt.Printf("| %-3s | %-10s | %-10s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "No", "IDMember", "ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %-3d | %-10d | %-10s | %-15s | %-10s | %-15s | %-12.2f | %-12.2f |\n", i+1, A[i].IDMember, A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
	}
	fmt.Println("-----------------------------------------------------------------------------------------------")
}

func LaporanInvestasi(A *tabSaham, nData *int) {
	var i int
	var totalKeuntunganKotor, totalKeuntunganBersih float64
	var keuntunganKotor, keuntunganBersih float64

	fmt.Println("---------------------")
	fmt.Println(" LAPORAN INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| %-3s | %-10s | %-10s | %-15s | %-10s | %-15s | %-12s | %-12s | %-17s | %-17s |\n", "No", "IDMember", "ID Perusahaan", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir", "Keuntungan Kotor", "Keuntungan Bersih")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	hitungUntung(A, nData)

	for i = 0; i < *nData; i++ {
		keuntunganKotor = A[i].danaAkhir - A[i].danaAwal
		keuntunganBersih = A[i].keuntungan
		totalKeuntunganKotor += keuntunganKotor
		totalKeuntunganBersih += keuntunganBersih
		fmt.Printf("| %-3d | %-10d | %-10s | %-15s | %-10s | %-15s | %-12f | %-12f | %-17.2f | %-17.2f |\n", i+1, A[i].IDMember, A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, keuntunganKotor, keuntunganBersih)
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Total Keuntungan Kotor : %.2f\n", totalKeuntunganKotor)
	fmt.Printf("Total Keuntungan Bersih : %.2f\n", totalKeuntunganBersih)
	fmt.Println("--------------------------------------------------------------------------")
}
