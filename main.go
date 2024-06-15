package main

import "fmt"

type car struct {
	nama      string
	merk      string
	penjualan int
	tahun     int
}

type factory struct {
	nama      string
	produksi  int
	kendaraan car
}

type tabMobil [5]factory

func main() {
	var mobil tabMobil
	var n, pilihan, banyakData int
	for pilihan != 6 {
		clear()
		menu(&pilihan)
		if pilihan == 1 {
			panggil(&mobil, &n, &banyakData)
		} else if pilihan == 2 {
			cetak(&mobil, &n, &banyakData)
		} else if pilihan == 3 {
			edit(&mobil, &n, &banyakData)
		} else if pilihan == 4 {
			hapus(&mobil, &n, &banyakData)
		} else if pilihan == 5 {
			carimobil(&mobil, &n, &banyakData)
		}
		tunggu()
	}
}

func panggil(mobil *tabMobil, n *int, banyakData *int) {
	var m int
	fmt.Print("Masukan banyak data mobil yang ingin di masukan: ")
	fmt.Scan(&m)
	*n += m
	for *banyakData < *n {
		fmt.Println("========================")
		fmt.Println("Data Mobil ke-", *banyakData+1)
		fmt.Print("Masukan nama mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].kendaraan.nama)
		fmt.Print("Masukan merk mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].kendaraan.merk)
		fmt.Print("Masukan jumlah penjualan mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].kendaraan.penjualan)
		fmt.Print("Masukan tahun keluaran mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].kendaraan.tahun)
		fmt.Print("Masukan nama pabrik mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].nama)
		fmt.Print("Masukan jumlah produksi mobil ", *banyakData+1, ":")
		fmt.Scan(&mobil[*banyakData].produksi)
		*banyakData++
	}
	fmt.Println("========================")
	fmt.Println("data berhasil di inputkan")
}

func cetak(m *tabMobil, n *int, banyakData *int) {
	var or string
	if *n == 0 {
		fmt.Println("Data mobil tidak tersedia.")
		fmt.Println("Apakah anda ingin menginput data mobil terlebih dahulu ? (Y/N)")
		fmt.Scan(&or)
		if or == "N" {

		} else if or == "Y" {
			panggil(m, n, banyakData)

		}
	} else {
		fmt.Println("========================")
		fmt.Printf("%-3s %-13s %-13s %-17s %-17s %-15s %-17s\n", "No", "Nama Mobil", "Merk Mobil", "Jumlah Penjualan", "Tahun Keluaran", "Nama Pabrik", "Jumlah Produksi")
		for i := 0; i < *n; i++ {
			fmt.Printf("%-3d %-13s %-13s %-17d %-17d %-15s %-17d\n", i+1, m[i].kendaraan.nama, m[i].kendaraan.merk, m[i].kendaraan.penjualan, m[i].kendaraan.tahun, m[i].nama, m[i].produksi)
		}
	}
	fmt.Println("========================")
}

func menu(pilihan *int) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Println("        M E N U         ")
	fmt.Println("      DEALER MOBIL      ")
	fmt.Println("========================")
	fmt.Println("1. Input Data Mobil")
	fmt.Println("2. Cetak Data Mobil")
	fmt.Println("3. Edit data Mobil")
	fmt.Println("4. Hapus Data Mobil")
	fmt.Println("5. Cari Data Mobil")
	fmt.Println("6. Keluar")
	fmt.Println("========================")
	fmt.Print("Pilihan: ")
	fmt.Scan(pilihan)
}

func edit(mobil *tabMobil, n *int, banyakData *int) {
	var pilihData int
	var or string
	var pilihan string
	if *n == 0 {
		fmt.Println("Data mobil tidak tersedia.")
		fmt.Println("Apakah anda ingin menginput data mobil terlebih dahulu ? (Y/N)")
		fmt.Scan(&or)
		if or == "Y" {
			panggil(mobil, n, banyakData)
		}
	} else {
		cetak(mobil, n, banyakData)
		fmt.Print("Data mobil keberapa yang ingin di edit: ")
		fmt.Scan(&pilihData)
		if pilihData > *n {
			fmt.Println("Data tidak ditemukan")
		} else {
			pilihData--
			for i := 0; i < *n; i++ {
				if i == pilihData {
					fmt.Println("========================")
					fmt.Println("Data Mobil ke-", i+1)
					fmt.Println("a. Nama Mobil:", mobil[i].kendaraan.nama)
					fmt.Println("b. Merk Mobil:", mobil[i].kendaraan.merk)
					fmt.Println("c. Jumlah Penjualan:", mobil[i].kendaraan.penjualan)
					fmt.Println("d. Tahun Keluaran:", mobil[i].kendaraan.tahun)
					fmt.Println("e. Pabrik:", mobil[i].nama)
					fmt.Println("f. Jumlah Produksi:", mobil[i].produksi)
					fmt.Println("========================")
					fmt.Print("Pilih data yang ingin di edit: ")
					fmt.Scan(&pilihan)
					if pilihan == "a" {
						fmt.Print("Masukan data pengganti untuk nama mobil:")
						fmt.Scan(&mobil[i].kendaraan.nama)
					} else if pilihan == "b" {
						fmt.Print("Masukan data pengganti untuk merk mobil:")
						fmt.Scan(&mobil[i].kendaraan.merk)
					} else if pilihan == "c" {
						fmt.Print("Masukan data pengganti untuk jumlah penjualan mobil:")
						fmt.Scan(&mobil[i].kendaraan.penjualan)
					} else if pilihan == "d" {
						fmt.Print("Masukan data pengganti untuk tahun keluaran mobil:")
						fmt.Scan(&mobil[i].kendaraan.tahun)
					} else if pilihan == "e" {
						fmt.Print("Masukan data pengganti untuk pabrik mobil:")
						fmt.Scan(&mobil[i].nama)
					} else if pilihan == "f" {
						fmt.Print("Masukan data pengganti untuk jumlah produksi mobil:")
						fmt.Scan(&mobil[i].produksi)
					} else {
						fmt.Println("Pilihan tidak tersedia")
					}
				}
			}
			fmt.Println("========================")
			fmt.Println("Data berhasil di edit")
		}
	}
}

func hapus(mobil *tabMobil, n *int, banyakData *int) {
	var or string
	var pilihData int
	if *n == 0 {
		fmt.Println("Data mobil tidak tersedia.")
		fmt.Println("Apakah anda ingin menginput data mobil terlebih dahulu ? (Y/N)")
		fmt.Scan(&or)
		if or == "Y" {
			panggil(mobil, n, banyakData)
		}
	} else {
		cetak(mobil, n, banyakData)
		fmt.Print("Data mobil keberapa yang ingin di hapus: ")
		fmt.Scan(&pilihData)
		if pilihData > *n {
			fmt.Println("Data tidak ditemukan")
		} else {
			pilihData--
			for i := pilihData; i < *n-1; i++ {
				(*mobil)[i] = (*mobil)[i+1]
			}
			*n--
			*banyakData--
			fmt.Println("Data berhasil di hapus")
		}
	}
}

func carimobil(mobil *tabMobil, n *int, banyakData *int) {
	var pilihana int
	var nama string
	var merk string
	var penjualan int
	var tahun int
	var pabrik string
	var produksi int
	var entah string
	if *n == 0 {
		fmt.Println("Data mobil tidak tersedia.")
		fmt.Println("Apakah anda ingin menginput data mobil terlebih dahulu ? (Y/N)")
		fmt.Scan(&entah)
		if entah == "Y" {
			panggil(mobil, n, banyakData)
		}
	} else {
		fmt.Println("Ingin mencari data mobil berdasarkan apa ?")
		fmt.Println("1. Nama Mobil")
		fmt.Println("2. Merk Mobil")
		fmt.Println("3. Jumlah Penjualan")
		fmt.Println("4. Tahun Keluaran")
		fmt.Println("5. Nama Pabrik")
		fmt.Println("6. Jumlah Produksi")
		fmt.Println("========================")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihana)
		if pilihana == 1 {
			fmt.Println("Masukan nama mobil yang ingin dicari: ")
			fmt.Scan(&nama)
			for i := 0; i < *n; i++ {
				if mobil[i].kendaraan.nama == nama {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}
		} else if pilihana == 2 {
			fmt.Println("Masukan merk mobil yang ingin dicari: ")
			fmt.Scan(&merk)
			for i := 0; i < *n; i++ {
				if mobil[i].kendaraan.merk == merk {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}
		} else if pilihana == 3 {
			fmt.Println("Masukan jumlah penjualan mobil yang ingin dicari: ")
			fmt.Scan(&penjualan)
			for i := 0; i < *n; i++ {
				if mobil[i].kendaraan.penjualan == penjualan {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}

		} else if pilihana == 4 {
			fmt.Println("Masukan tahun keluaran mobil yang ingin dicari: ")
			fmt.Scan(&tahun)
			for i := 0; i < *n; i++ {
				if mobil[i].kendaraan.tahun == tahun {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}

		} else if pilihana == 5 {
			fmt.Println("Masukan nama pabrik mobil yang ingin dicari: ")
			fmt.Scan(&pabrik)
			for i := 0; i < *n; i++ {
				if mobil[i].nama == pabrik {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}
		} else if pilihana == 6 {
			fmt.Println("Masukan jumlah produksi mobil yang ingin dicari: ")
			fmt.Scan(&produksi)
			for i := 0; i < *n; i++ {
				if mobil[i].produksi == produksi {
					cetak(mobil, n, banyakData)
				} else {
					fmt.Println("Data mobil tidak ditemukan")
				}
			}

		} else {
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
func tunggu() {
	fmt.Print("Tekan enter untuk melanjutkan")
	var s rune
	fmt.Scanf("\n%c", &s)
}
