package main

import "fmt"

func tebakan(player rune, nilai int) int {
	/* mengembalikan angka tebakan dari player yang diperoleh dari masukan, masukan
	akan dilakukan berulang HINGGA tebakan player benar (sama dengan nilai) atau
	telah  menebak  sebanyak  3x  kesempatan.  Catatan:  perhatikan  string  yang
	ditampilkan seperti pada contoh*/

	var jawaban, i int
	i = 1
	fmt.Printf("%c - masukkan angaka tebakan ke-%d: ", player, i)
	fmt.Scan(&jawaban)
	for i < 3 && jawaban != nilai {
		i += 1
		fmt.Printf("%c - masukkan angka tebakan ke-%d: ", player, i)
		fmt.Scan(&jawaban)
	}

	return jawaban
}

func tukerPemenang(benar bool, winner, player *rune) {
	/* 	I.S. terdefinisi benar yang menyatakan tebakan benar atau salah
	F.S. nilai winner dan player saling bertukar apabila benar adalah true */

	if benar {
		var temp rune = *winner
		*winner = *player
		*player = temp
	}

}

func mulaiRonde(ronde int, winner rune, nilai *int) {
	/*  I.S. terdefinisi nomor ronde dan pemenang (winner) dari permainan
		F.S. menampilkan ronde, sedangkan nilai berisi angka rahasia yang diperoleh
	dari masukan, perhatikan string yang ditampilkan seperti pada contoh*/

	fmt.Println("Ronde", ronde, ":")
	fmt.Printf("%c - masukkan angka rahasia: ", winner)
	fmt.Scan(nilai)

}

func main() {
	// deklarasi variabel
	var winner, player rune
	var ronde, nilai, answer int

	// inisialisasi
	winner = 'A'
	player = 'B'
	ronde = 1

	// memulai ronde pertama
	mulaiRonde(ronde, winner, &nilai)
	for nilai != -101 {
		// pemain menebak
		answer = tebakan(player, nilai)
		// tuker pemenang apabila syarat memenuhi
		tukerPemenang(answer == nilai, &winner, &player)
		// tampilakan pemenangnya dan increment ronde
		fmt.Printf("%c adalah pemenangnya\n\n", winner)
		ronde += 1
		// mulai ronde selanjutnya
		mulaiRonde(ronde, winner, &nilai)
	}

	fmt.Println("Permainan Selesai")
}
