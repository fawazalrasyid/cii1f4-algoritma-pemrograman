package main

import "fmt"

func main() {
	// deklarasi variabel
	var hari1, hari2, bulanInput string
	var tanggal1, tanggal2, bulan1, bulan2, tahun1, tahun2 int

	// pembacaan masukan
	fmt.Scan(&hari1, &tanggal1, &bulanInput, &tahun1)

	bulan1 = angkaBulan(bulanInput)

	// panggil subprogram untuk penentuan tanggal pengambilan
	pengambilan(tanggal1, bulan1, tahun1, hari1, &tanggal2, &bulan2, &tahun2, &hari2)

	// tampilkan tanggal pengambilan visa
	fmt.Println(hari2, tanggal2, bulanAngka(bulan2), tahun2)

}

func kabisat(tahun int) bool {
	// Mengembalikan true apabila tahun adalah kabisat, false apabila sebaliknya.

	if tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0 {
		return true
	}

	return false
}

func angkaBulan(bulan string) int {
	/* Mengembalikan angka berdasarkan urutan nama bulan pada kalender masehi (1 hingga 12).
	   0 untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua. Contoh: "januari" menjadi 1 */

	var namaBulan [12]string = [12]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	for i := 0; i < 12; i++ {
		if bulan == namaBulan[i] {
			return i
		}
	}

	return 0
}

func bulanAngka(angka int) string {
	/* Mengembalikan nama bulan berdasarkan urutan angka bulan pada kalender masehi (1 hingga 12).
	   "invalid" untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua. Contoh: 1 menjadi "januari" */

	var namaBulan [12]string = [12]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

	if angka <= 12 {
		return namaBulan[angka]
	}

	return "invalid"
}
func jumlahHari(bulan, tahun int) int {
	/* Mengembalikan jumlah hari berdasarkan bulan dan tahun yang terdefinisi,
	   hati-hati pada bulan februari pada saat kabisat. -1 apabila bulan tidak valid */

	var jumlahHari [12]int = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var isKabisat bool = kabisat(tahun)

	if isKabisat && bulan == 1 {
		return jumlahHari[bulan] + 1
	}

	return jumlahHari[bulan]
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	/* I.S. terdefinisi hari1 yang menyatakan hari pengajuan string, asumsi huruf kecil semua
	   F.S. hari2 berisi hari pengambilan dan durasi berisi lama pembuatan visa, karena sabtu dan minggu tidak dihitung */

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
	/* I.S. terdefinisi waktu pengajuan visa, yaitu tanggal1, bulan1, tahun1 dan hari1
	   F.S. tanggal2, bulan2, tahun2 dan hari2 berisi waktu pengambilan visa */

	var durasi, jmlhHari int

	// tentukan durasi, hari pengambilan, dan jumlah hari pada bulan pengajuan
	mencariDurasi(hari1, hari2, &durasi)

	// dapatkan tanggal pengambilan, asumsi bulan pengambilan dan tahun pengambilan sama dengan waktu pengajuan

	// cek apabila tanggal pengambilan melebihi lama hari, update tanggal dan bulan pengambilan dengan yang seharusnya
	// cek apabila bulan pengambilan melebihi 12, update dengan bulan dan tahun pengambilan yang seharusnya
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
