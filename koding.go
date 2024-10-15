package main

import (
	"fmt"
)

// deklarasi variabel
var (
	username = "helga"
	password = "2406499102"
	menu     int
	histori  []string
)

type pengguna struct {
	nama           string
	umur           int
	alamat         string
	jeniskelamin   string
	makananfavorit string
	minumanfavorit string
}

type buku struct {
	nama   string
	jumlah int
}

func main() {

	// Membuat data pengguna
	datapengguna := pengguna{
		nama:           "Helga",
		umur:           18,
		alamat:         "Perumahan Jatiwarna blok E",
		jeniskelamin:   "Perempuan",
		makananfavorit: "Sushi",
		minumanfavorit: "Teh Hijau",
	}

	// daftar buku
	daftarbuku := []buku{
		{nama: "Pemrograman", jumlah: 10},
		{nama: "Film", jumlah: 5},
		{nama: "Printing", jumlah: 20},
	}

	// Pesan Pembuka
	fmt.Println("=== Selamat Datang di Penyewaan Buku ===")

	// variabel input
	var inputUsername, inputPassword string

	// Input Username
	fmt.Println("Silahkan input username: ")
	fmt.Scanln(&inputUsername)

	// Input Password
	fmt.Println("Silahkan input password: ")
	fmt.Scanln(&inputPassword)

	// Validasi Username & Password
	if inputUsername != username || inputPassword != password {
		fmt.Println("Username atau Password salah.")
		return
	}

	// Loop untuk menampilkan menu terus-menerus
	for {
		// Menampilkan Menu
		fmt.Println("Menu:")
		fmt.Println("1. Lihat Informasi Pengguna Program:")
		fmt.Println("2. Lihat daftar Buku:")
		fmt.Println("3. Tambah Daftar Buku:")
		fmt.Println("4. Tambah Peminjaman Buku:")
		fmt.Println("5. Histori Peminjaman Buku:")
		fmt.Println("6. Keluar dari Program:")

		// Input Pilihan Menu
		var pilihan int
		fmt.Print("Pilih Menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			// Menampilkan informasi pengguna
			fmt.Println("Lihat Informasi Pengguna Akun")
			fmt.Printf("Nama: %s\nUmur: %d\nAlamat: %s\nJenis Kelamin: %s\nMakanan Favorit: %s\nMinuman Favorit: %s\n",
				datapengguna.nama, datapengguna.umur, datapengguna.alamat, datapengguna.jeniskelamin, datapengguna.makananfavorit, datapengguna.minumanfavorit)

		case 2:
			// Menampilkan daftar buku
			fmt.Println("Lihat Daftar Buku")
			for _, b := range daftarbuku {
				fmt.Printf("Nama: %s, Jumlah: %d\n", b.nama, b.jumlah)
			}

		case 3:
			// Tambah Daftar Buku
			var namaBuku string
			var jumlahBuku int

			fmt.Println("Masukkan nama buku yang ingin ditambah:")
			fmt.Scanln(&namaBuku)
			fmt.Println("Masukkan jumlah buku:")
			fmt.Scanln(&jumlahBuku)

			// Validasi agar jumlah buku harus bilangan bulat positif
			if jumlahBuku <= 0 {
				fmt.Println("Jumlah buku harus bilangan bulat positif.")
				break
			}

			found := false
			for i, buku := range daftarbuku {
				if buku.nama == namaBuku {
					daftarbuku[i].jumlah += jumlahBuku
					histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", buku.nama, jumlahBuku))
					fmt.Println("Buku berhasil ditambahkan.")
					found = true
					break
				}
			}
			if !found {
				daftarbuku = append(daftarbuku, buku{nama: namaBuku, jumlah: jumlahBuku})
				histori = append(histori, fmt.Sprintf("Ditambahkan: %s, Jumlah: %d", namaBuku, jumlahBuku))
				fmt.Println("Buku baru berhasil ditambahkan.")
			}

		case 4:
			// Menambahkan peminjaman buku
			fmt.Println("=== Peminjaman Buku ===")
			var namaBuku string
			var jumlahPinjam int
			fmt.Print("Masukkan nama buku yang ingin dipinjam: ")
			fmt.Scanln(&namaBuku)
			fmt.Print("Masukkan jumlah buku yang ingin dipinjam: ")
			fmt.Scanln(&jumlahPinjam)

			// Validasi agar jumlah pinjaman harus bilangan bulat positif
			if jumlahPinjam <= 0 {
				fmt.Println("Jumlah buku yang dipinjam harus bilangan bulat positif.")
				break
			}

			found := false
			for i, buku := range daftarbuku {
				if buku.nama == namaBuku {
					if buku.jumlah < jumlahPinjam {
						fmt.Println("Jumlah buku tidak mencukupi!")
						found = true
						break
					} else {
						daftarbuku[i].jumlah -= jumlahPinjam
						histori = append(histori, fmt.Sprintf("Dipinjam: %s, Jumlah: %d", buku.nama, jumlahPinjam))
						fmt.Println("Peminjaman buku berhasil.")
						found = true
						break
					}
				}
			}
			if !found {
				fmt.Println("Nama dan jumlah buku yang ingin dipinjam tidak sesuai.")
			}

		case 5:
			// Menampilkan histori peminjaman
			fmt.Println("=== Histori Peminjaman Buku ===")
			for _, transaksi := range histori {
				fmt.Println(transaksi)
			}

		case 6:
			fmt.Println("Keluar dari Program")
			fmt.Println("Terimakasih !")
			return

		default:
			fmt.Println("Menu tidak valid, silahkan pilih lagi.")
		}
	}
}
