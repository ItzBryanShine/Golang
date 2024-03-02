package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tambahan struct {
	KodeBuku      string
	JudulBuku     string
	Pengarang     string
	Penerbit      string
	JumlahHalaman int
	TahunTerbit   int
}

var ListBuku []Tambahan

func InputBukuBaru() {
	userInput := bufio.NewReader(os.Stdin)

	kodeBukuTambahan := ""
	judulBukuTambahan := ""
	pengarangBukuTambahan := ""
	penerbitBukuTambahan := ""
	jumlahHalamanBukuTambahan := 0
	tahunTerbitBukuTambahan := 0

	//Kode Buku Tambahan
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=   Tambah Buku Baru  =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	fmt.Print("| Silahkan Input Kode Buku : ")
	kodeBukuTambahan, err := userInput.ReadString('\n')
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	kodeBukuTambahan = strings.Replace(kodeBukuTambahan, "\n", "", 1)

	//Judul Buku Tambahan
	fmt.Print("| Silahkan Input Judul Buku : ")
	judulBukuTambahan, err = userInput.ReadString('\n')
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	judulBukuTambahan = strings.Replace(judulBukuTambahan, "\n", "", 1)

	//Pengarang Buku Tambahan
	fmt.Print("| Silahkan Input Pengarang Buku : ")
	pengarangBukuTambahan, err = userInput.ReadString('\n')
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	pengarangBukuTambahan = strings.Replace(pengarangBukuTambahan, "\n", "", 1)

	//Penerbit Buku Tambahan
	fmt.Print("| Silahkan Input Penerbit Buku : ")
	penerbitBukuTambahan, err = userInput.ReadString('\n')
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	penerbitBukuTambahan = strings.Replace(penerbitBukuTambahan, "\n", "", 1)

	//Jumlah Halaman Buku Tambahan
	fmt.Print("| Silahkan Input Jumlah Halaman Buku : ")
	_, err = fmt.Scanln(&jumlahHalamanBukuTambahan)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	//Tahun Terbit Buku Tambahan
	fmt.Print("| Silahkan Input Tahun Terbit Buku : ")
	_, err = fmt.Scanln(&tahunTerbitBukuTambahan)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	ListBuku = append(ListBuku, Tambahan{
		KodeBuku:      kodeBukuTambahan,
		JudulBuku:     judulBukuTambahan,
		Pengarang:     pengarangBukuTambahan,
		Penerbit:      penerbitBukuTambahan,
		JumlahHalaman: jumlahHalamanBukuTambahan,
		TahunTerbit:   tahunTerbitBukuTambahan,
	})
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=    Buku Berhasil Terinput   =+=+=+= |")
	fmt.Println("+=============================================+")

}

func LiatBuku() {
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=      Lihat Buku     =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Tahun Terbit : %d, Kode Buku : %s\n",
			urutan+1,
			buku.TahunTerbit,
			buku.KodeBuku,
		)

		fmt.Printf("- Judul Buku : %s\n",
			buku.JudulBuku,
		)

		fmt.Printf("- Pengarang : %s\n",
			buku.Pengarang,
		)

		fmt.Printf("- Penerbit : %s\n",
			buku.Penerbit,
		)

		fmt.Printf("- Jumlah Halaman : %d\n",
			buku.JumlahHalaman,
		)

	}
}

func BukuList() {
	for urutan, buku := range ListBuku {
		fmt.Printf("%d. Tahun Terbit : %d, Kode Buku : %s\n",
			urutan+1,
			buku.TahunTerbit,
			buku.KodeBuku,
		)

		fmt.Printf("- Judul Buku : %s\n",
			buku.JudulBuku,
		)

		fmt.Printf("- Pengarang : %s\n",
			buku.Pengarang,
		)

		fmt.Printf("- Penerbit : %s\n",
			buku.Penerbit,
		)

		fmt.Printf("- Jumlah Halaman : %d\n",
			buku.JumlahHalaman,
		)

	}
}

func main() {
	var sistemMenu int
	fmt.Println("+=============================================+")
	fmt.Println("| Aplikasi Manajemen Daftar Buku Perpustakaan |")
	fmt.Println("+=============================================+")
	fmt.Println("| Silahkan Pilih Menu : ")
	fmt.Println("| 1. Tambah Buku Baru")
	fmt.Println("| 2. Lihat Buku")
	fmt.Println("| 3. Hapus Buku")
	fmt.Println("| 4. Edit Buku")
	fmt.Println("| 5. Keluar")
	fmt.Println("+=============================================+")
	fmt.Print("| Masukkan Pilihanmu : ")
	_, err := fmt.Scanln(&sistemMenu)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
	}

	switch sistemMenu {
	case 1:
		InputBukuBaru()
	case 2:
		LiatBuku()
	case 3:
		HapusBuku()
	case 4:
		EditBuku()
	case 5:
		os.Exit(0)

	}
	main()

}

func HapusBuku() {
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=      Hapus Buku     =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	BukuList()
	fmt.Println("+=============================================+")
	var kodeHapusBuku string
	fmt.Print("| Masukkan Kode Buku yang akan dihapus: ")
	_, err := fmt.Scanln(&kodeHapusBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	for i, buku := range ListBuku {
		if strings.TrimSpace(buku.KodeBuku) == strings.TrimSpace(kodeHapusBuku) {
			ListBuku = append(ListBuku[:i], ListBuku[i+1:]...)
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=+=  Buku Berhasil Dihapus  =+=+=+=+= |")
			fmt.Println("+=============================================+")
			return
		}
	}

	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=  Buku Tidak Ditemukan!  =+=+=+=+= |")
	fmt.Println("+=============================================+")
}

func EditBuku() {
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=+=    Edit Buku    =+=+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	BukuList()
	fmt.Println("+=============================================+")
	var kodeEditBuku string
	fmt.Print("| Masukkan Kode Buku yang akan diedit: ")
	_, err := fmt.Scanln(&kodeEditBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	for i, buku := range ListBuku {
		if strings.TrimSpace(buku.KodeBuku) == strings.TrimSpace(kodeEditBuku) {
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Printf("| Masukkan Informasi Baru untuk Buku dengan Kode %s\n", buku.KodeBuku)
			fmt.Println("+=============================================+")

			userInput := bufio.NewReader(os.Stdin)

			//Judul Buku Tambahan
			fmt.Print("| Silahkan Input Judul Buku : ")
			judulBukuTambahan, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			judulBukuTambahan = strings.Replace(judulBukuTambahan, "\n", "", 1)

			//Pengarang Buku Tambahan
			fmt.Print("| Silahkan Input Pengarang Buku : ")
			pengarangBukuTambahan, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			pengarangBukuTambahan = strings.Replace(pengarangBukuTambahan, "\n", "", 1)

			//Penerbit Buku Tambahan
			fmt.Print("| Silahkan Input Penerbit Buku : ")
			penerbitBukuTambahan, err := userInput.ReadString('\n')
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			penerbitBukuTambahan = strings.Replace(penerbitBukuTambahan, "\n", "", 1)

			//Jumlah Halaman Buku Tambahan
			fmt.Print("| Silahkan Input Jumlah Halaman Buku : ")
			_, err = fmt.Scanln(&buku.JumlahHalaman)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			//Tahun Terbit Buku Tambahan
			fmt.Print("| Silahkan Input Tahun Terbit Buku : ")
			_, err = fmt.Scanln(&buku.TahunTerbit)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			ListBuku[i] = Tambahan{
				KodeBuku:      buku.KodeBuku,
				JudulBuku:     judulBukuTambahan,
				Pengarang:     pengarangBukuTambahan,
				Penerbit:      penerbitBukuTambahan,
				JumlahHalaman: buku.JumlahHalaman,
				TahunTerbit:   buku.TahunTerbit,
			}

			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=+=  Buku Berhasil Diubah!  =+=+=+=+= |")
			fmt.Println("+=============================================+")
			return

		}
	}

	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=  Buku Tidak Ditemukan!  =+=+=+=+= |")
	fmt.Println("+=============================================+")

}
