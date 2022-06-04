package main

import "fmt"

type buku struct {
	judul, penulis string
	tahun          int
}

type tabBuku [5]buku

func tambahBuku(kardus *tabBuku, atas *int) {
	if *atas == 4 {
		fmt.Println("kardus penuh")
	} else {
		*atas++
		fmt.Scan(&kardus[*atas].judul, &kardus[*atas].penulis, &kardus[*atas].tahun)
	}
}

func ambilBuku(kardus *tabBuku, atas *int, ambil *buku) {
	if *atas == -1 {
		fmt.Println("kardus kosong")
	} else {
		*ambil = kardus[*atas]
		*atas--
	}
}

func cariBuku(kardus *tabBuku, atas *int, X string) {
	var found bool = false
	var ambil buku

	for !found && *atas != -1 {
		ambilBuku(kardus, atas, &ambil)
		fmt.Println("Judul buku yang dikeluarkan: ", ambil.judul)
		found = ambil.judul == X
	}

	if found {
		fmt.Println("Ketemu")
	} else {
		fmt.Println("Tidak Ketemu")
	}
}

func main() {
	// {Deklarasikan variabel yang dibutuhkan di sini}
	var kardus tabBuku
	var atas int

	// {1. Buatlah sebuah kardus kosong, di mana atas bernilai -1}
	atas = -1

	// {2. Tambahkan 4 buku seperti contoh gambar, data tahun dan penulis bebas}
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)

	// {3. Cari buku dengan judul “C”, harusnya yang tampil adalahK,ZZ,dan CKETEMU}
	cariBuku(&kardus, &atas, "C")

	// {4. Tambahkan buku sampai kardus penuh}
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)
	tambahBuku(&kardus, &atas)

	// {5. Cari buku dengan judul yang tidak terdapat pada kardus tersebut.}
	cariBuku(&kardus, &atas, "Sistem Digital")

}
