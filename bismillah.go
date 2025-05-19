package main

import "fmt"

const NMAX int = 5

type arrKomentar [NMAX]comment
type comment struct {
	isi, isiLower    string
	positif, negatif string
	sentimen         string
	nilaiSentimen    int
	pjgteks          int
	hurufPertama     rune
}

func main() {
	var komentar arrKomentar
	var i, jumlahData int
	i = 0
	komentar[0].positif = ("baik")
	komentar[1].positif = ("hebat")
	komentar[2].positif = ("keren")
	komentar[3].positif = ("mantap")
	komentar[4].positif = ("cerdas")
	komentar[1].negatif = ("jelek")

	//Tampilan utama untuk menentukan pilihan menu lainnya
	var pilihanMenuUtama int
	for pilihanMenuUtama != 7 {
		tampilanUtama()
		fmt.Println("|                  MENU UTAMA                |")
		fmt.Println("==============================================")
		fmt.Println("|  1. Tambah Komentar                        |")
		fmt.Println("|  2. Ubah Komentar                          |")
		fmt.Println("|  3. Hapus Komentar                         |")
		fmt.Println("|  4. Cari Komentar                          |")
		fmt.Println("|  5. Statistik Komentar                     |")
		fmt.Println("|  6. Daftar Komentar                        |")
		fmt.Println("|  7. Keluar                                 |")
		fmt.Println("==============================================")
		fmt.Print("   Pilih menu: ")
		fmt.Scan(&pilihanMenuUtama)
		switch pilihanMenuUtama {
		case 1:
			//Menu menambah komentar
			menuTambahKomentar(&komentar, &i, &jumlahData)
		case 2:
			//Menu mengubah komentar
			menuUbahKomentar(&komentar, &i, &jumlahData)
		case 3:
			//Menu menghapus komentar
			menuHapusKomentar(&komentar, &i, &jumlahData)
		case 4:
			//Menu mencari komentar
			menuCariKomentar(&komentar, &i, &jumlahData)
		case 5:
			//Menu statistik komentar
			//statistik(&komentar, &i, &jumlahData)
			urutAlfabetAsc(&komentar, &i, &jumlahData)
		case 6:
			//Menu cetak komentar
			daftarKomentar(&komentar, &i, &jumlahData)
		case 7:
			//Menu Keluar
			menuKeluar()
		default:
			fmt.Println("Mohon masukkan angka yang benar")
		}
	}
}

func menuTambahKomentar(komentar *arrKomentar, i, jumlahData *int) {
	//Menambahkan komentar baru ke dalam array komentar
	var j, jumlahKomentar int
	tampilanUtama()
	fmt.Println("|             MENU TAMBAH KOMENTAR           |")
	fmt.Println("==============================================")
	fmt.Print("   Jumlah komentar yang ingin ditambah: ")
	fmt.Scan(&jumlahKomentar)
	if jumlahKomentar > NMAX {
		jumlahKomentar = NMAX
	}
	//Proses penginputan komentar
	for j < jumlahKomentar && *jumlahData < NMAX {
		fmt.Print("   Tambahkan komentar: ")
		fmt.Scan(&komentar[*i].isi)
		*i++
		j++
		*jumlahData++

	}
	if j == jumlahKomentar-1 {
		fmt.Println("==============================================")
		fmt.Println("|        Komentar berhasil ditambahkan       |")
	}
	if *jumlahData == NMAX || jumlahKomentar > NMAX {
		fmt.Println("|           Daftar komentar penuh!           |")
	}
	return
}

func tampilanUtama() {
	fmt.Println("==============================================")
	fmt.Printf("|                    ADMIN                   |\n")
	fmt.Printf("|          ANALISIS SENTIMEN KOMENTAR        |\n")
	fmt.Println("==============================================")
}

func daftarKomentar(komentar *arrKomentar, i, jumlahData *int) {
	fmt.Println("==============================================")
	fmt.Println("|               DAFTAR KOMENTAR              |")
	fmt.Println("==============================================")
	if *jumlahData == 0 {
		fmt.Println("|          Komentar tidak tersedia!          |")
	} else {
		for *i = 0; *i < *jumlahData; *i++ {
			fmt.Printf("   %d. %s %s\n", *i+1, komentar[*i].isi, komentar[*i].sentimen)
		}
	}
	fmt.Println("==============================================")
}
func menuCariKomentar(komentar *arrKomentar, i, jumlahData *int) {
	tampilanUtama()
	fmt.Println("|             MENU CARI KOMENTAR             |")
	fmt.Println("==============================================")
	var kataKunci string
	var pilihanMenuCari, ketemu int
	fmt.Println("|  1. Sequential Search                      |")
	fmt.Println("|  2. Binary Search                          |")
	fmt.Println("==============================================")
	fmt.Print("   Masukkan kata kunci: ")
	fmt.Scan(&kataKunci)
	fmt.Print("   Pilih menu (1/2): ")
	fmt.Scan(&pilihanMenuCari)
	switch pilihanMenuCari {
	case 1:
		ketemu = cariKomentarSequential(*komentar, *i, *jumlahData, kataKunci)
		if ketemu != -1 {
			fmt.Println("   Komentar ditemukan pada nomor: ", ketemu+1)
			fmt.Printf("   %d. %s\n", ketemu+1, komentar[ketemu].isi)
			fmt.Println("==============================================")
		} else {
			fmt.Println("   Komentar tidak ditemukan          ")
			fmt.Println("==============================================")
		}
	case 2:
		urutAlfabetAsc(komentar, i, jumlahData)
		ketemu = cariKomentarBinary(*komentar, *i, *jumlahData, kataKunci)
		if ketemu != -1 {
			fmt.Println("   Komentar ditemukan pada nomor: ", ketemu+1)
			fmt.Printf("   %d. %s\n", ketemu+1, komentar[ketemu].isi)
			fmt.Println("==============================================")
		} else {
			fmt.Println("   Komentar tidak ditemukan          ")
			fmt.Println("==============================================")
		}
	}
}

func cariKomentarSequential(komentar arrKomentar, i, jumlahData int, kataKunci string) int {
	var found int
	found = -1
	i = 0
	for i < jumlahData && found == -1 {
		if komentar[i].isi == kataKunci {
			found = i
			return found
		}
		i++
	}
	return found
}
func cariKomentarBinary(komentar arrKomentar, i, jumlahData int, kataKunci string) int {
	var left, mid, right int
	var found int
	found = -1
	left = 0
	right = jumlahData - 1
	for left <= right && found == -1 {
		mid = (left + right) / 2
		if kataKunci > komentar[mid].isi {
			left = mid + 1
		} else if kataKunci < komentar[mid].isi {
			right = mid - 1
		} else {
			found = mid
		}
	}
	return found
}

func urutAlfabetAsc(komentar *arrKomentar, i, jumlahData *int) {
	var j int
	var temp comment
	for *i = 1; *i < *jumlahData; *i++ {
		j = *i
		temp = komentar[j]
		for j > 0 && temp.isi < komentar[j-1].isi {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
}

func menuUbahKomentar(komentar *arrKomentar, i, jumlahData *int) {
	var ubahIdx int
	tampilanUtama()
	fmt.Println("|             MENU UBAH KOMENTAR             |")
	daftarKomentar(komentar, i, jumlahData)
	if *jumlahData != 0 {
		fmt.Print("Ubah komentar nomor: ")
		fmt.Scan(&ubahIdx)
		for *i = 0; *i < *jumlahData; *i++ {
			if *i == ubahIdx-1 {
				fmt.Print("   Ubah komentar menjadi: ")
				fmt.Scan(&komentar[*i].isi)
				fmt.Println("   Komentar berhasil diperbarui")
				daftarKomentar(komentar, i, jumlahData)
				return
			}
		}
	}
	if *jumlahData != 0 {
		fmt.Println("   Komentar tidak ditemukan!")
		fmt.Println("==============================================")
	}
}
func menuKeluar() {
	fmt.Println("==============================================")
	fmt.Println("|                Terima Kasih!               |")
	fmt.Println("==============================================")
	return
}

func menuHapusKomentar(komentar *arrKomentar, i, jumlahData *int) {
	var idxHapus int
	tampilanUtama()
	fmt.Println("|             MENU HAPUS KOMENTAR            |")
	fmt.Println("==============================================")
	daftarKomentar(komentar, i, jumlahData)
	fmt.Print("Hapus komen no: ")
	fmt.Scan(&idxHapus)
	if idxHapus > 0 && idxHapus < *jumlahData {
		for *i = idxHapus; *i < *jumlahData; *i++ {
			komentar[*i-1].isi = komentar[*i].isi
		}
	} else {
		fmt.Println("|          Komentar tidak ditemukan          |")
		return
	}
	*jumlahData = *jumlahData - 1
	fmt.Println("   Komentar berhasil dihapus         ")
	daftarKomentar(komentar, i, jumlahData)
	return
}

func sentimen(komentar *arrKomentar, i, jumlahData *int) {
	var j int
	for *i = 0; *i < *jumlahData; *i++ {
		for j = 0; j < 5; j++ {
			if komentar[*i].isi == komentar[j].positif {
				komentar[*i].sentimen = ("positif")
			}
			if komentar[*i].isi == komentar[j].negatif {
				komentar[*i].sentimen = ("negatif")
			}
		}
		if komentar[*i].sentimen != "positif" && komentar[*i].sentimen != "negatif" {
			komentar[*i].sentimen = ("netral")
		}
		j = 0
	}
}

func panjangKomen(komentar *arrKomentar, i, jumlahData *int) {
	for *i = 0; *i < *jumlahData; *i++ {
		komentar[*i].pjgteks = len(komentar[*i].isi)
	}
}

func statistik(komentar *arrKomentar, i, jumlahData *int) {
	var pilihStatistik int
	tampilanUtama()
	fmt.Println("|               MENU STATISTIK               |")
	fmt.Println("==============================================")
	fmt.Println("   1.(Ascending) panjang teks")
	fmt.Println("   2.(Descending) panjang teks")
	fmt.Println("   3.(Ascending) sentimen")
	fmt.Println("   4.(Descending) sentimen")
	fmt.Scan(&pilihStatistik)
	switch pilihStatistik {
	case 1:
		urutPanjangAsc(komentar, i, jumlahData)
	case 2:
		urutPanjangDesc(komentar, i, jumlahData)
	case 3:
		urutSentimenAsc(komentar, i, jumlahData)
	case 4:
		urutSentimenDesc(komentar, i, jumlahData)
	}
}

func urutPanjangAsc(komentar *arrKomentar, i, jumlahData *int) {
	//Selection Sort
	panjangKomen(komentar, i, jumlahData)
	var j, idx int
	var temp comment
	for *i = 1; *i <= *jumlahData; *i++ {
		idx = *i - 1
		j = *i
		for j < *jumlahData {
			if komentar[idx].pjgteks > komentar[j].pjgteks {
				idx = j
			}
			j++
		}
		temp = komentar[idx]
		komentar[idx] = komentar[*i-1]
		komentar[*i-1] = temp
	}
	daftarKomentar(komentar, i, jumlahData)
}

func urutPanjangDesc(komentar *arrKomentar, i, jumlahData *int) {
	//Selection Sort
	panjangKomen(komentar, i, jumlahData)
	var j, idx int
	var temp comment
	for *i = 1; *i <= *jumlahData; *i++ {
		idx = *i - 1
		j = *i
		for j < *jumlahData {
			if komentar[idx].pjgteks < komentar[j].pjgteks {
				idx = j
			}
			j++
		}
		temp = komentar[idx]
		komentar[idx] = komentar[*i-1]
		komentar[*i-1] = temp
	}
	daftarKomentar(komentar, i, jumlahData)
}

func urutSentimenAsc(komentar *arrKomentar, i, jumlahData *int) {
	//Insertion Sort
	sentimen(komentar, i, jumlahData)
	nilaiSentimen(komentar, i, jumlahData)
	var j int
	var temp comment
	for *i = 0; *i < *jumlahData; *i++ {
		j = *i
		temp = komentar[j]
		for j > 0 && temp.nilaiSentimen < komentar[j-1].nilaiSentimen {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
	daftarKomentar(komentar, i, jumlahData)
}
func urutSentimenDesc(komentar *arrKomentar, i, jumlahData *int) {
	//Insertion Sort
	sentimen(komentar, i, jumlahData)
	nilaiSentimen(komentar, i, jumlahData)
	var j int
	var temp comment
	for *i = 0; *i < *jumlahData; *i++ {
		j = *i
		temp = komentar[j]
		for j > 0 && temp.nilaiSentimen > komentar[j-1].nilaiSentimen {
			komentar[j] = komentar[j-1]
			j--
		}
		komentar[j] = temp
	}
	daftarKomentar(komentar, i, jumlahData)
}

func nilaiSentimen(komentar *arrKomentar, i, jumlahData *int) {
	for *i = 0; *i < *jumlahData; *i++ {
		if komentar[*i].sentimen == "positif" {
			komentar[*i].nilaiSentimen = 3
		} else if komentar[*i].sentimen == "negatif" {
			komentar[*i].nilaiSentimen = 1
		} else {
			komentar[*i].nilaiSentimen = 2
		}
	}
}
