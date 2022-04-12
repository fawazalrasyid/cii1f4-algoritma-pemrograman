package main

import "fmt"

func kabisat(tahun int) bool {
	if tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0 {
		return true
	}

	return false
}

func angkaBulan(bulan string) int {
	var namaBulan [12]string = [12]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	for i := 0; i < 12; i++ {
		if bulan == namaBulan[i] {
			return i
		}
	}

	return 0
}

func bulanAngka(angka int) string {
	var namaBulan [12]string = [12]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	if angka <= 12 {
		return namaBulan[angka]
	}

	return "invalid"
}
func jumlahHari(bulan, tahun int) int {
	var jumlahHari [12]int = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var isKabisat bool = kabisat(tahun)

	if isKabisat && bulan == 1 {
		return jumlahHari[bulan] + 1
	}

	return jumlahHari[bulan]
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	var namaHari [7]string = [7]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	var hariAngka int

	for i := 0; i < 7; i++ {
		if hari1 == namaHari[i] {
			hariAngka = i
		}
	}

	if hariAngka == 4 {
		*hari2 = namaHari[hariAngka-3]
		*durasi = 4
	} else if hariAngka == 3 {
		*hari2 = namaHari[hariAngka-3]
		*durasi = 4
	} else {
		*hari2 = namaHari[hariAngka+2]
		*durasi = 2
	}

}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durasi, jmlhHari int

	mencariDurasi(hari1, hari2, &durasi)
	jmlhHari = jumlahHari(bulan1, tahun1)

	if tanggal1+durasi < jmlhHari && bulan1 != 11 {
		if tanggal1 == jmlhHari {
			*tanggal2 = durasi
		} else if tanggal1 >= jmlhHari-2 {
			*tanggal2 = (durasi - (jmlhHari - tanggal1))
		} else {
			*tanggal2 = tanggal1 + durasi
		}
		*bulan2 = bulan1
		*tahun2 = tahun1
	} else if tanggal1+durasi > jmlhHari && bulan1 != 11 {
		if tanggal1 == jmlhHari {
			*tanggal2 = durasi
		} else if tanggal1 >= jmlhHari-2 {
			*tanggal2 = (durasi - (jmlhHari - tanggal1))
		} else {
			*tanggal2 = tanggal1 + durasi
		}
		*bulan2 = bulan1 + 1
		*tahun2 = tahun1
	} else if tanggal1+durasi > jmlhHari && bulan1 == 11 {
		if tanggal1 == jmlhHari {
			*tanggal2 = durasi
		} else if tanggal1 >= jmlhHari-2 {
			*tanggal2 = (durasi - (jmlhHari - tanggal1))
		} else {
			*tanggal2 = tanggal1 + durasi
		}
		*bulan2 = 0
		*tahun2 = tahun1 + 1
	}
}

func main() {
	var hari1, hari2, bulanInput string
	var tanggal1, tanggal2, bulan1, bulan2, tahun1, tahun2 int

	fmt.Scan(&hari1, &tanggal1, &bulanInput, &tahun1)

	bulan1 = angkaBulan(bulanInput)

	pengambilan(tanggal1, bulan1, tahun1, hari1, &tanggal2, &bulan2, &tahun2, &hari2)

	fmt.Println(hari2, tanggal2, bulanAngka(bulan2), tahun2)
}
