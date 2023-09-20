# Program-Kasir-Cetak-Invoice

Ini adalah program sederhana untuk mengelola pesanan dan pembayaran di sebuah restoran. Program ini dibuat dalam bahasa pemrograman Go (Golang) dan memiliki fitur berikut:

## Deskripsi Program

Program ini adalah aplikasi kasir restoran yang memungkinkan pengguna untuk:
- Melihat daftar menu restoran.
- Memilih menu yang ingin dipesan dan menentukan jumlahnya.
- Menghitung total biaya pesanan.
- Memasukkan jumlah uang yang diberikan oleh pelanggan.
- Menghitung kembalian yang harus dikembalikan kepada pelanggan.
- Menyimpan pesanan, total biaya pesanan, dan kembalian dalam sebuah file.

## Fitur Utama

1. **Daftar Menu**: Program menyediakan daftar menu restoran beserta harga masing-masing menu.

2. **Pemesanan**: Pengguna dapat memilih menu dengan memasukkan nomor menu yang diinginkan dan jumlah yang ingin dipesan.

3. **Perhitungan Biaya**: Program akan menghitung total biaya pesanan berdasarkan menu yang dipilih dan jumlahnya.

4. **Pembayaran**: Pengguna diminta untuk memasukkan jumlah uang yang diberikan oleh pelanggan. Program akan menghitung kembalian yang harus dikembalikan.

5. **Simpan ke File**: Program akan menyimpan pesanan, total biaya pesanan, dan kembalian dalam sebuah file teks yang dapat ditentukan oleh pengguna.

## Cara Menggunakan Program

1. Jalankan program dengan menjalankan perintah `go run nama_program.go` di terminal.

2. Ikuti instruksi yang muncul di layar untuk memilih menu, menentukan jumlah pesanan, dan memasukkan jumlah uang yang diberikan oleh pelanggan.

3. Program akan menghitung total biaya pesanan dan kembalian.

4. Program akan meminta nama file untuk menyimpan rincian pesanan, total biaya, dan kembalian.

5. Hasil pesanan akan disimpan dalam file tersebut.

## Struktur Program

- Program terdiri dari satu file sumber utama (`main.go`).
- Ada beberapa fungsi yang digunakan untuk membantu proses, seperti `getUserInput` untuk mendapatkan input dari pengguna dan `saveToFile` untuk menyimpan data ke dalam file.

## Pengembangan Selanjutnya

Program ini adalah versi dasar dan dapat ditingkatkan dengan fitur-fitur tambahan seperti:
- Menggunakan basis data untuk menyimpan riwayat pesanan.
- Menambahkan opsi diskon atau promosi.
- Membuat antarmuka pengguna grafis (GUI).
- Menambahkan manajemen stok menu.
- Mengoptimalkan tampilan daftar menu.

## Catatan

- Pastikan Anda memiliki Go (Golang) terinstal di komputer Anda sebelum menjalankan program ini.
- Pastikan untuk mengganti folder tempat file disimpan sesuai dengan kebutuhan Anda.
- Program ini hanya merupakan contoh sederhana dan dapat disesuaikan dengan kebutuhan bisnis restoran Anda.

