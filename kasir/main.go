package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct untuk menyimpan informasi menu
type Menu struct {
	Nama  string
	Harga float64
}

func main() {
	// Inisialisasi daftar menu
	daftarMenu := []Menu{
		{"Bakso Mercon		", 20000},
		{"Bakso Telor		", 20000},
		{"Bakso Urat		", 25000},
		{"Bakso Keju		", 25000},
		{"Es Teh Manis		", 5000},
		{"Es Alpukat		", 5000},

		// Tambahkan menu lain sesuai kebutuhan
	}

	// Membuat map untuk menyimpan pesanan
	pesanan := make(map[string]int)

	// Memulai pemesanan
	fmt.Println("------Selamat datang di restoran kami!------")
	for {
		// Menampilkan daftar menu
		fmt.Println("\n ------------Daftar Menu:------------")
		for i, menu := range daftarMenu {
			fmt.Printf("%d. %s - Harga: %.2f\n", i+1, menu.Nama, menu.Harga)
		}

		// Meminta input pengguna
		fmt.Print("\n Pilih nomor menu yang ingin dipesan (atau ketik 'selesai' untuk selesai): ")
		input := getUserInput()

		// Memeriksa apakah pengguna ingin selesai
		if strings.ToLower(input) == "selesai" {
			break
		}

		// Memeriksa nomor menu yang dipilih
		menuIndex, err := strconv.Atoi(input)
		if err != nil || menuIndex < 1 || menuIndex > len(daftarMenu) {
			fmt.Println("\n Input tidak valid. Pilih nomor menu yang sesuai.")
			continue
		}

		// Meminta jumlah menu yang ingin dipesan
		fmt.Print("\n Masukkan jumlah menu yang ingin dipesan: ")
		jumlah := getUserInput()

		// Mengubah jumlah menjadi integer
		jumlahMenu, err := strconv.Atoi(jumlah)
		if err != nil || jumlahMenu < 1 {
			fmt.Println("\n Jumlah menu yang dipesan tidak valid.")
			continue
		}

		// Menambahkan menu yang dipesan ke dalam map pesanan
		menu := daftarMenu[menuIndex-1]
		pesanan[menu.Nama] += jumlahMenu
	}

	// Menghitung total biaya pesanan
	total := 0.0
	fmt.Println("\nPesanan Anda:")
	pesananText := "	"
	for menu, jumlah := range pesanan {
		for _, m := range daftarMenu {
			if m.Nama == menu {
				harga := m.Harga * float64(jumlah)
				fmt.Printf("%s - Jumlah: %d - Harga: %.2f\n", menu, jumlah, harga)
				total += harga
				pesananText += fmt.Sprintf("%s - Jumlah: %d - Harga: %.2f\n", menu, jumlah, harga)
			}
		}
	}

	// Menampilkan total biaya pesanan
	fmt.Printf("\nTotal Biaya Pesanan : %.2f\n", total)

	// Meminta jumlah uang yang diberikan oleh pelanggan
	fmt.Print("\nMasukkan jumlah uang yang diberikan oleh pelanggan : ")
	uangDiberikan := getUserInput()

	// Mengubah jumlah uang yang diberikan menjadi float64
	uangPelanggan, err := strconv.ParseFloat(uangDiberikan, 64)
	if err != nil || uangPelanggan < total {
		fmt.Println("\nJumlah uang tidak valid atau kurang dari total biaya pesanan.")
		return
	}

	// Menghitung kembalian
	kembalian := uangPelanggan - total
	fmt.Printf("\nKembalian : %.2f\n", kembalian)

	// ...

	// Menyimpan pesanan, total biaya pesanan, dan kembalian dalam file di folder kasir/cetakstruk
	folderPath := "kasir/cetakinvoice" // Ubah sesuai dengan path folder yang Anda inginkan
	fmt.Print("\n Masukkan nama file untuk menyimpan pesanan, total, dan kembalian (misalnya, 'pesanan.txt') : ")
	namaFile := getUserInput()
	filePath := fmt.Sprintf("%s/%s", folderPath, namaFile)

	err = saveToFile(filePath, fmt.Sprintf("Pesanan Anda:\n\n%s\nTotal Biaya Pesanan: %.2f\nTotal Bayar: %.2f\nKembalian: %.2f\n", pesananText, total, uangPelanggan, kembalian))
	if err != nil {
		fmt.Println("\n Gagal menyimpan pesanan, total, dan kembalian ke dalam file.")
		return
	}

	fmt.Printf("\n Pesanan, total biaya pesanan, dan kembalian telah disimpan dalam file '%s'. Terima kasih!", filePath)

}

// Fungsi untuk mendapatkan input dari pengguna
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk menyimpan data ke dalam file
func saveToFile(namaFile, data string) error {
	file, err := os.Create(namaFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
