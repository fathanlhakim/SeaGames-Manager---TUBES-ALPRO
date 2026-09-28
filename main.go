package main

import "fmt"

const NMAX int = 100

type negara struct {
	nama     string
	emas     int
	perak    int
	perunggu int
}

type tabNegara [NMAX]negara

func cariSequential(T tabNegara, n int, cari string) int {
	var idx, i int
	var ketemu bool

	idx = -1
	i = 0
	ketemu = false
	for i < n && !ketemu {
		if T[i].nama == cari {
			idx = i
			ketemu = true
		}
		i++
	}
	return idx
}

func cariBinary(T tabNegara, kiri, kanan int, cari string) int {
	var tengah int

	if kiri > kanan {
		return -1
	}
	tengah = (kiri + kanan) / 2
	if T[tengah].nama == cari {
		return tengah
	} else if T[tengah].nama > cari {
		return cariBinary(T, kiri, tengah-1, cari)
	}
	return cariBinary(T, tengah+1, kanan, cari)
}

func cariMedaliTerbanyak(T tabNegara, n int) int {
	var i, maks, total, idx int

	if n == 0 {
		return -1
	}
	idx = 0
	maks = T[0].emas + T[0].perak + T[0].perunggu
	for i = 1; i < n; i++ {
		total = T[i].emas + T[i].perak + T[i].perunggu
		if total > maks {
			maks = total
			idx = i
		}
	}
	return idx
}

func urutPeringkat(T *tabNegara, n int) {
	var i, j, maks int
	var temp negara

	for i = 0; i < n-1; i++ {
		maks = i
		for j = i + 1; j < n; j++ {
			if T[j].emas > T[maks].emas {
				maks = j
			} else if T[j].emas == T[maks].emas {
				if T[j].perak > T[maks].perak {
					maks = j
				} else if T[j].perak == T[maks].perak {
					if T[j].perunggu > T[maks].perunggu {
						maks = j
					}
				}
			}
		}
		temp = T[i]
		T[i] = T[maks]
		T[maks] = temp
	}
}

func urutNama(T *tabNegara, n int) {
	var i, j int
	var temp negara

	for i = 1; i < n; i++ {
		temp = T[i]
		j = i - 1
		for j >= 0 && T[j].nama > temp.nama {
			T[j+1] = T[j]
			j--
		}
		T[j+1] = temp
	}
}

func tambahNegara(T *tabNegara, n *int) {
	var nama string

	if *n >= NMAX {
		fmt.Println("Kapasitas peserta penuh!")
	} else {
		fmt.Print("Masukkan nama negara: ")
		fmt.Scan(&nama)

		if cariSequential(*T, *n, nama) != -1 {
			fmt.Println("Negara sudah terdaftar!")
		} else {
			T[*n].nama = nama
			T[*n].emas = 0
			T[*n].perak = 0
			T[*n].perunggu = 0
			*n++
			fmt.Println("Negara berhasil ditambahkan.")
		}
	}
}

func ubahNegara(T *tabNegara, n int) {
	var lama, baru string
	var idx int

	fmt.Print("Masukkan nama negara yang ingin diubah: ")
	fmt.Scan(&lama)

	idx = cariSequential(*T, n, lama)
	if idx == -1 {
		fmt.Println("Negara tidak ditemukan!")
	} else {
		fmt.Print("Masukkan nama negara baru: ")
		fmt.Scan(&baru)
		T[idx].nama = baru
		fmt.Println("Nama negara berhasil diubah.")
	}
}

func hapusNegara(T *tabNegara, n *int) {
	var nama string
	var idx, i int

	fmt.Print("Masukkan nama negara yang ingin dihapus: ")
	fmt.Scan(&nama)

	idx = cariSequential(*T, *n, nama)
	if idx == -1 {
		fmt.Println("Negara tidak ditemukan!")
	} else {
		for i = idx; i < *n-1; i++ {
			T[i] = T[i+1]
		}
		*n--
		fmt.Println("Negara berhasil dihapus.")
	}
}

func updateMedali(T *tabNegara, n int) {
	var nama string
	var idx int

	fmt.Print("Masukkan nama negara: ")
	fmt.Scan(&nama)

	idx = cariSequential(*T, n, nama)
	if idx == -1 {
		fmt.Println("Negara tidak ditemukan!")
	} else {
		fmt.Print("Jumlah emas: ")
		fmt.Scan(&T[idx].emas)
		fmt.Print("Jumlah perak: ")
		fmt.Scan(&T[idx].perak)
		fmt.Print("Jumlah perunggu: ")
		fmt.Scan(&T[idx].perunggu)
		fmt.Println("Data medali berhasil diperbarui.")
	}
}

func tampilPeringkat(T tabNegara, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada data negara.")
	} else {
		urutPeringkat(&T, n)
		fmt.Println("\n--- KLASEMEN MEDALI SEAGAMES ---")
		fmt.Printf("%-5s | %-15s | %-5s | %-5s | %-8s\n", "Rank", "Negara", "Emas", "Perak", "Perunggu")
		fmt.Println("------------------------------------------------------")
		for i = 0; i < n; i++ {
			fmt.Printf("%-5d | %-15s | %-5d | %-5d | %-8d\n", i+1, T[i].nama, T[i].emas, T[i].perak, T[i].perunggu)
		}
	}
}

func cariNegara(T *tabNegara, n int) {
	var nama string
	var idx int

	urutNama(T, n)
	fmt.Print("Masukkan nama negara yang dicari: ")
	fmt.Scan(&nama)

	idx = cariBinary(*T, 0, n-1, nama)
	if idx == -1 {
		fmt.Println("Negara tidak ditemukan.")
	} else {
		fmt.Printf("Negara %s ditemukan! Emas: %d, Perak: %d, Perunggu: %d\n",
			T[idx].nama, T[idx].emas, T[idx].perak, T[idx].perunggu)
	}
}

func tampilMedaliTerbanyak(T tabNegara, n int) {
	var idx, total int

	idx = cariMedaliTerbanyak(T, n)
	if idx == -1 {
		fmt.Println("Belum ada data negara.")
	} else {
		total = T[idx].emas + T[idx].perak + T[idx].perunggu
		fmt.Printf("Negara dengan medali terbanyak: %s (total %d medali)\n", T[idx].nama, total)
	}
}

func main() {
	var data tabNegara
	var n, pilihan int
	var selesai bool

	n = 0
	selesai = false
	for !selesai {
		fmt.Println("\n==================================")
		fmt.Println("     APLIKASI SEAGAMES MANAGER")
		fmt.Println("==================================")
		fmt.Println("1. Tambah negara")
		fmt.Println("2. Ubah nama negara")
		fmt.Println("3. Hapus negara")
		fmt.Println("4. Update data medali")
		fmt.Println("5. Tampilkan peringkat")
		fmt.Println("6. Cari negara")
		fmt.Println("7. Negara dengan medali terbanyak")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahNegara(&data, &n)
		} else if pilihan == 2 {
			ubahNegara(&data, n)
		} else if pilihan == 3 {
			hapusNegara(&data, &n)
		} else if pilihan == 4 {
			updateMedali(&data, n)
		} else if pilihan == 5 {
			tampilPeringkat(data, n)
		} else if pilihan == 6 {
			cariNegara(&data, n)
		} else if pilihan == 7 {
			tampilMedaliTerbanyak(data, n)
		} else if pilihan == 0 {
			selesai = true
			fmt.Println("Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
