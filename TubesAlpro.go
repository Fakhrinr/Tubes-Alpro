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

const maxSaham int = 50

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
			LaporanInvestasi(&inves, &nData)
		case 5:
			MenuCariDataStatis()
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
	fmt.Println("5. Data Saham per Sektor")
	fmt.Println("---------------------")
	fmt.Println("6. Keluar")
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
		fmt.Println("4. Kembali ke menu utama")
		fmt.Println("Pilih sektor investasi (1/2/3/4):")
		fmt.Println("---------------------")
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
		}

		fmt.Println("Silahkan masukkan Nama investor:")
		fmt.Scan(&A[*nData].nama)
		fmt.Println("Silahkan masukkan id perusahaan investasi:")
		fmt.Scan(&A[*nData].id)
		fmt.Println("Silahkan masukkan nama perusahaan:")
		fmt.Scan(&A[*nData].perusahaan)
		fmt.Println("Silahkan masukkan dana awal investasi:")
		fmt.Scan(&A[*nData].danaAwal)
		fmt.Println("Silahkan masukkan dana akhir investasi:")
		fmt.Scan(&A[*nData].danaAkhir)
		A[*nData].IDMember = BuatID()
		*nData++

		fmt.Println("Apakah anda ingin menambah data investasi lagi? (yes/no):")
		fmt.Scan(&y)
	}
	fmt.Println("-----------------")
	fmt.Println("Data investasi berhasil disimpan")
	fmt.Println("-----------------")
}

func BuatID() int {

	ID++
	return ID
}
func LihatInvestasi(A *tabSaham, nData *int) {
	var x int
	fmt.Println("---------------------")
	fmt.Println("	DATA INVESTASI")
	fmt.Println("---------------------")
	cetakData(A, nData)
	fmt.Println("---------------------")
	fmt.Println("	 INVESTASI")
	fmt.Println("---------------------")
	fmt.Println(" Silahkan pilih sektor investasi saham")
	fmt.Println("1. Cari data berdasarkan Id/ Perusahaan")
	fmt.Println("2. Mengurutkan data berdasarkan Keutungan")
	fmt.Println("4. Kembali ke menu utama")
	fmt.Println("Silahkan pilih menu (1/2/3/4):")
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
		fmt.Println("(1/2/3/4/5)")
		fmt.Scan(&y)
		if y == 1 {
			sektor = "Teknologi"
		} else if y == 2 {
			sektor = "Keuangan"
		} else if y == 3 {
			sektor = "Konsumsi"
		} else if y == 4 {
			CariDataMember(A, nData)
		} else if y == 5 {
			fmt.Println("Kembali ke menu utama")
			return
		} else {
			fmt.Println("Input Tidak Valid")
			return
		}
		CariDataSektor(A, nData, sektor)
	case 2:
		var y int
		fmt.Println("Silahkan pilih")
		fmt.Println("1. Mengurutkan berdasarkan nominal / keuntungan dari terkecil")
		fmt.Println("2. Mengurutkan berdasarkan nominal / keuntungan dari terbesar")
		fmt.Println("3. Kembali ke menu utama")
		fmt.Println("Silahkan pilih:")
		fmt.Println("(1/2/3)")
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
		var y int
		fmt.Println("Silahkan masukkan ID Member yang ingin di edit:")
		fmt.Scan(&y)
		ubahInvestasi(A, nData, y)
	case 2:
		var y int
		fmt.Println("Silahkan masukkan ID Member yang ingin di hapus:")
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
	fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "ID", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	for i = 0; i < *nData; i++ {
		if A[i].IDMember == y {
			found = true
			fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n",
				A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
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
func ubahInvestasi(A *tabSaham, nData *int, y int) {
	var i int
	var found bool
	fmt.Println("---------------------")
	fmt.Println("	EDIT DATA INVESTASI")
	fmt.Println("---------------------")
	fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "ID", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	found = false
	for i = 0; i < *nData; i++ {
		if A[i].IDMember == y {
			found = true
			fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n",
				A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
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
	fmt.Println(" DATA SEKTOR")
	fmt.Println("---------------------")
	fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "ID", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	found = false
	for i = 0; i < *nData; i++ {
		if A[i].sektor == y || A[i].perusahaan == y || A[i].id == y || A[i].nama == y {
			found = true
			fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n",
				A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
			fmt.Println("-----------------------------------------------------------------------------------------------")
			fmt.Println("Data investasi berhasil ditemukan")
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}
func CariDataMember(A *tabSaham, nData *int) {
	var i int
	var found bool
	var y int
	fmt.Println("---------------------")
	fmt.Println("	CARI DATA MEMBER")
	fmt.Println("---------------------")
	fmt.Printf("| ID | Nama                 | Sektor       | Perusahaan         | danaAwal         | danaAkhir         |\n")
	found = false
	fmt.Println("Silahkan masukkan ID Member yang ingin dicari:")
	fmt.Scan(&y)
	for i = 0; i < *nData; i++ {
		if A[i].IDMember == y {
			found = true
			fmt.Printf("%s %s %s %s %.2f %.2f\n", A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan")
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
	fmt.Println("	 DATA INVESTASI")
	fmt.Println("---------------------")
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %-5s | %-15s | %-12s | %-18s | %12.2f | %12.2f | %14.2f |\n",
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
		fmt.Printf("| %-5s | %-15s | %-12s | %-18s | %12.2f | %12.2f | %14.2f |\n",
			A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, A[i].keuntungan)
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
	fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %-12s | %-12s |\n", "ID", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir")
	fmt.Println("-----------------------------------------------------------------------------------------------")
	for i = 0; i < *nData; i++ {
		fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f |\n",
			A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir)
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
	fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %-12s | %-12s | %-17s |\n",
		"ID", "Nama", "Sektor", "Perusahaan", "danaAwal", "danaAkhir", "Keuntungan Kotor")
	fmt.Println("-----------------------------------------------------------------------------------------------")

	hitungUntung(A, nData)

	for i = 0; i < *nData; i++ {
		keuntunganKotor = A[i].danaAkhir - A[i].danaAwal
		keuntunganBersih = A[i].keuntungan
		totalKeuntunganKotor += keuntunganKotor
		totalKeuntunganBersih += keuntunganBersih
		fmt.Printf("| %-5s | %-15s | %-10s | %-15s | %12.2f | %12.2f | %17.2f |\n",
			A[i].id, A[i].nama, A[i].sektor, A[i].perusahaan, A[i].danaAwal, A[i].danaAkhir, keuntunganKotor)
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("Total Keuntungan Kotor : %.2f\n", totalKeuntunganKotor)
	fmt.Printf("Total Keuntungan Bersih : %.2f\n", totalKeuntunganBersih)
	fmt.Println("--------------------------------------------------------------------------")
}

func MenuCariDataStatis() {
	var sektor string
	fmt.Println("---------------------")
	fmt.Println(" CONTOH DATA STATIS PER SEKTOR")
	fmt.Println("---------------------")
	fmt.Println("Silakan ketik nama sektor (Teknologi/Keuangan/Konsumsi):")
	fmt.Scan(&sektor)

	if sektor == "Keuangan" {
		fmt.Println("ID       Nama      Perusahaan")
		fmt.Println("--------------------------------------")
		fmt.Println("SahamB   Keuangan  PT_Beta")
		fmt.Println("SahamE   Keuangan  PT_Epsilon")
		fmt.Println("SahamH   Keuangan  PT_Theta")
		fmt.Println("SahamL   Keuangan  PT_Omega")
		fmt.Println("SahamO   Keuangan  PT_Gamma2")
		fmt.Println("SahamR   Keuangan  PT_Zeta2")
		fmt.Println("SahamU   Keuangan  PT_Iota2")
		fmt.Println("SahamX   Keuangan  PT_Mu2")
		fmt.Println("SahamAA  Keuangan  PT_Omicron2")
		fmt.Println("SahamAD  Keuangan  PT_Sigma2")
		fmt.Println("SahamAG  Keuangan  PT_Phi2")
		fmt.Println("SahamAJ  Keuangan  PT_Omega2")
		fmt.Println("SahamAM  Keuangan  PT_Gamma3")
		fmt.Println("SahamAP  Keuangan  PT_Zeta3")
		fmt.Println("SahamAS  Keuangan  PT_Iota3")
		fmt.Println("SahamAV  Keuangan  PT_Mu3")
	} else if sektor == "Teknologi" {
		fmt.Println("ID       Nama       Perusahaan")
		fmt.Println("--------------------------------------")
		fmt.Println("SahamA   Teknologi  PT_Alpha")
		fmt.Println("SahamD   Teknologi  PT_Delta")
		fmt.Println("SahamG   Teknologi  PT_Eta")
		fmt.Println("SahamJ   Teknologi  PT_Kappa")
		fmt.Println("SahamK   Teknologi  PT_Sigma")
		fmt.Println("SahamN   Teknologi  PT_Beta2")
		fmt.Println("SahamQ   Teknologi  PT_Epsilon2")
		fmt.Println("SahamT   Teknologi  PT_Theta2")
		fmt.Println("SahamW   Teknologi  PT_Lambda2")
		fmt.Println("SahamZ   Teknologi  PT_Xi2")
		fmt.Println("SahamAC  Teknologi  PT_Rho2")
		fmt.Println("SahamAF  Teknologi  PT_Upsilon2")
		fmt.Println("SahamAI  Teknologi  PT_Psi2")
		fmt.Println("SahamAL  Teknologi  PT_Beta3")
		fmt.Println("SahamAO  Teknologi  PT_Epsilon3")
		fmt.Println("SahamAR  Teknologi  PT_Theta3")
		fmt.Println("SahamAU  Teknologi  PT_Lambda3")
		fmt.Println("SahamAX  Teknologi  PT_Xi3")
	} else if sektor == "Konsumsi" {
		fmt.Println("ID       Nama      Perusahaan")
		fmt.Println("--------------------------------------")
		fmt.Println("SahamC   Konsumsi  PT_Gamma")
		fmt.Println("SahamF   Konsumsi  PT_Zeta")
		fmt.Println("SahamI   Konsumsi  PT_Iota")
		fmt.Println("SahamM   Konsumsi  PT_Alpha2")
		fmt.Println("SahamP   Konsumsi  PT_Delta2")
		fmt.Println("SahamS   Konsumsi  PT_Eta2")
		fmt.Println("SahamV   Konsumsi  PT_Kappa2")
		fmt.Println("SahamY   Konsumsi  PT_Nu2")
		fmt.Println("SahamAB  Konsumsi  PT_Pi2")
		fmt.Println("SahamAE  Konsumsi  PT_Tau2")
		fmt.Println("SahamAH  Konsumsi  PT_Chi2")
		fmt.Println("SahamAK  Konsumsi  PT_Alpha3")
		fmt.Println("SahamAN  Konsumsi  PT_Delta3")
		fmt.Println("SahamAQ  Konsumsi  PT_Eta3")
		fmt.Println("SahamAT  Konsumsi  PT_Kappa3")
		fmt.Println("SahamAW  Konsumsi  PT_Nu3")
	} else {
		fmt.Println("Sektor tidak ditemukan. Masukkan sesuai format.")
	}
}

