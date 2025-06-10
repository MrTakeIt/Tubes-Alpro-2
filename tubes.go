package main

import "fmt"

const NMAX int = 5

type arrKomentar [NMAX]comment
type comment struct {
	isi string
}

var komentar arrKomentar

func main() {
	var i, jumlahData int
	i = 0
	var pilihanMenuUtama int
	tampilanUtama()

	for pilihanMenuUtama != 7 {
		fmt.Println("1. Tambah Komentar")
		fmt.Println("2. Ubah Komentar")
		fmt.Println("3. Hapus Komentar")
		fmt.Println("4. Cari Komentar")
		fmt.Println("5. Statistik Komentar")
		fmt.Println("6. Daftar Komentar")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu: ")
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

		case 5:
			//Menu statistik komentar

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
	var pilihanMenuTambah int
	var j, jumlahKomentar int
	tampilanUtama()
	fmt.Print("Jumlah komentar yang ingin ditambah: ")
	fmt.Scan(&jumlahKomentar)
	//Proses penginputan komentar

	for j < jumlahKomentar {
		fmt.Print("Tambahkan komentar: ")
		fmt.Scan(&komentar[*i].isi)
		*i++
		j++
		*jumlahData++
	}
	fmt.Println("1. Menu")
	for {
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihanMenuTambah)
		switch pilihanMenuTambah {
		case 1:
			return
		default:
			fmt.Println("Masukkan angka yang benar!")
		}
	}
}

func tampilanUtama() {
	fmt.Printf("\nADMIN\n")
	fmt.Println("ANALISIS SENTIMEN KOMENTAR")
}

func daftarKomentar(komentar *arrKomentar, i, jumlahData *int) {
	var pilihanMenuDaftar int
	cetakKomentar(komentar, i, jumlahData)
	fmt.Println("1. Menu")
	fmt.Println("2. Keluar")
	for pilihanMenuDaftar != 2 {
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihanMenuDaftar)
		switch pilihanMenuDaftar {
		case 1:
			return
		case 2:
			menuKeluar()
		default:
			fmt.Println("Masukkan angka yang benar!")
		}
	}
}

func cariKomentarBinary(komentar *arrKomentar, i *int, kataKunci *string) {

}

func menuUbahKomentar(komentar *arrKomentar, i, jumlahData *int) {
	var ubahIdx, pilihanMenuUbah int
	var tempKomentar string
	tampilanUtama()
	cetakKomentar(komentar, i, jumlahData)
	fmt.Print("Ubah komentar nomor: ")
	fmt.Scan(&ubahIdx)
	fmt.Print("Ubah komentar menjadi: ")
	fmt.Scan(&tempKomentar)
	komentar[ubahIdx-1].isi = tempKomentar
	cetakKomentar(komentar, i, jumlahData)
	fmt.Println("1. Menu")
	for {
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihanMenuUbah)
		switch pilihanMenuUbah {
		case 1:
			return
		default:
			fmt.Println("Masukkan angka yang benar!")
		}
	}
}

func cetakKomentar(komentar *arrKomentar, i, jumlahData *int) {
	if komentar[0].isi == "" {
		fmt.Println("Komentar tidak tersedia!")
	} else {
		for *i = 0; *i < *jumlahData; *i++ {
			fmt.Printf("%d. %s\n", *i+1, komentar[*i].isi)
		}
	}
}
func menuKeluar() {
	fmt.Println("Terima Kasih!")
}

func menuHapusKomentar(komentar *arrKomentar, i, jumlahData *int) {
	var idxHapus, pilihanMenuHapus int
	tampilanUtama()
	cetakKomentar(komentar, i, jumlahData)
	fmt.Print("Hapus komen no: ")
	fmt.Scan(&idxHapus)
	for *i = idxHapus; *i < *jumlahData; *i++ {
		komentar[*i-1].isi = komentar[*i].isi
	}
	*jumlahData = *jumlahData - 1
	cetakKomentar(komentar, i, jumlahData)
	fmt.Println("1. ")
	for {
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihanMenuHapus)
		switch pilihanMenuHapus {
		case 1:
			return
		default:
			fmt.Println("Masukkan angka yang benar!")
		}
	}
}
