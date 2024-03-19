package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/jung-kurt/gofpdf"
)

type Tambahan struct {
	KodeBuku      string `json:"kode_buku"`
	JudulBuku     string `json:"judul_buku"`
	Pengarang     string `json:"pengarang"`
	Penerbit      string `json:"penerbit"`
	JumlahHalaman string `json:"jumlah_halaman"`
	TahunTerbit   string `json:"tahun_terbit"`
}

var ListBuku []Tambahan
var bukuKode = make(map[string]bool)

func InputBukuBaru() {
	userInput := bufio.NewReader(os.Stdin)

	kodeBukuTambahan := ""
	judulBukuTambahan := ""
	pengarangBukuTambahan := ""
	penerbitBukuTambahan := ""
	jumlahHalamanBukuTambahan := ""
	tahunTerbitBukuTambahan := ""

	//Kode Buku Tambahan
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=   Tambah Buku Baru  =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	fmt.Print("| Silahkan Input Kode Buku : ")
	kodeBukuTambahan, _ = userInput.ReadString('\n')
	kodeBukuTambahan = strings.TrimSpace(kodeBukuTambahan)

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodeBukuTambahan {
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=+ Maaf Kode Buku Sudah Ada! +=+=+=+= |")
			fmt.Println("+=============================================+")
			return
		}
	}

	//Judul Buku Tambahan
	fmt.Print("| Silahkan Input Judul Buku : ")
	judulBukuTambahan, _ = userInput.ReadString('\n')
	judulBukuTambahan = strings.TrimSpace(judulBukuTambahan)

	//Pengarang Buku Tambahan
	fmt.Print("| Silahkan Input Pengarang Buku : ")
	pengarangBukuTambahan, _ = userInput.ReadString('\n')
	pengarangBukuTambahan = strings.TrimSpace(pengarangBukuTambahan)

	//Penerbit Buku Tambahan
	fmt.Print("| Silahkan Input Penerbit Buku : ")
	penerbitBukuTambahan, _ = userInput.ReadString('\n')
	penerbitBukuTambahan = strings.TrimSpace(penerbitBukuTambahan)

	//Jumlah Halaman Buku Tambahan
	fmt.Print("| Silahkan Input Jumlah Halaman Buku : ")
	jumlahHalamanBukuTambahan, _ = userInput.ReadString('\n')
	jumlahHalamanBukuTambahan = strings.TrimSpace(jumlahHalamanBukuTambahan)

	//Tahun Terbit Buku Tambahan
	fmt.Print("| Silahkan Input Tahun Terbit Buku : ")
	tahunTerbitBukuTambahan, _ = userInput.ReadString('\n')
	tahunTerbitBukuTambahan = strings.TrimSpace(tahunTerbitBukuTambahan)

	ListBuku = append(ListBuku, Tambahan{
		KodeBuku:      kodeBukuTambahan,
		JudulBuku:     judulBukuTambahan,
		Pengarang:     pengarangBukuTambahan,
		Penerbit:      penerbitBukuTambahan,
		JumlahHalaman: jumlahHalamanBukuTambahan,
		TahunTerbit:   tahunTerbitBukuTambahan,
	})

	simpanBukuKeJson(Tambahan{
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

	var inputBaru string
	fmt.Print("| Apakah Anda ingin menambah buku lagi? (y/n)?: ")
	_, err := fmt.Scanln(&inputBaru)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	if strings.ToLower(inputBaru) == "y" {
		InputBukuBaru()
	} else {
		main()
	}

}

func buku(ch <-chan string, bukuCh chan Tambahan, wg *sync.WaitGroup) {
	var buku Tambahan
	for idBuku := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("Buku/%s", idBuku))
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(dataJSON, &buku)
		if err != nil {
			fmt.Println(err)
		}
		bukuCh <- buku
	}
	wg.Done()
}

func liatBuku() {

	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=      Lihat Buku     =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	ListBuku = []Tambahan{}

	ListBukuJson, err := os.ReadDir("Buku")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	bukuCh := make(chan Tambahan, len(ListBukuJson))

	jumlahBuku := 5

	for i := 0; i < jumlahBuku; i++ {
		wg.Add(1)
		go buku(ch, bukuCh, &wg)
	}

	for _, fileBuku := range ListBukuJson {
		fmt.Println(fileBuku.Name())
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(bukuCh)

	for dataBuku := range bukuCh {
		ListBuku = append(ListBuku, dataBuku)
	}

	for _, buku := range ListBuku {

		fmt.Printf("| Tahun Terbit : %s, Kode Buku : %s\n",
			buku.TahunTerbit,
			buku.KodeBuku,
		)

		fmt.Printf("| - Judul Buku : %s\n",
			buku.JudulBuku,
		)

		fmt.Printf("| - Pengarang : %s\n",
			buku.Pengarang,
		)

		fmt.Printf("| - Penerbit : %s\n",
			buku.Penerbit,
		)

		fmt.Printf("| - Jumlah Halaman : %s\n",
			buku.JumlahHalaman,
		)
		fmt.Println("-----------------------------------------------")
	}
}

// Untuk Daftar Buku di Menu Hapus Buku
func bukuList() {
	ListBuku = []Tambahan{}

	ListBukuJson, err := os.ReadDir("Buku")
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	bukuCh := make(chan Tambahan, len(ListBukuJson))

	jumlahBuku := 5

	for i := 0; i < jumlahBuku; i++ {
		wg.Add(1)
		go buku(ch, bukuCh, &wg)
	}

	for _, fileBuku := range ListBukuJson {
		fmt.Println(fileBuku.Name())
		ch <- fileBuku.Name()
	}

	close(ch)
	wg.Wait()
	close(bukuCh)

	for dataBuku := range bukuCh {
		ListBuku = append(ListBuku, dataBuku)
	}

	for _, buku := range ListBuku {

		fmt.Printf("| Tahun Terbit : %s, Kode Buku : %s\n",
			buku.TahunTerbit,
			buku.KodeBuku,
		)

		fmt.Printf("| - Judul Buku : %s\n",
			buku.JudulBuku,
		)

		fmt.Printf("| - Pengarang : %s\n",
			buku.Pengarang,
		)

		fmt.Printf("| - Penerbit : %s\n",
			buku.Penerbit,
		)

		fmt.Printf("| - Jumlah Halaman : %s\n",
			buku.JumlahHalaman,
		)
		fmt.Println("-----------------------------------------------")
	}
}

func main() {
	os.Mkdir("Buku", 0777)
	os.Mkdir("pdf", 0777)

	loadBuku()

	var sistemMenu int
	fmt.Println("+=============================================+")
	fmt.Println("| Aplikasi Manajemen Daftar Buku Perpustakaan |")
	fmt.Println("+=============================================+")
	fmt.Println("| Silahkan Pilih Menu : ")
	fmt.Println("| 1. Tambah Buku Baru")
	fmt.Println("| 2. Lihat Buku")
	fmt.Println("| 3. Hapus Buku")
	fmt.Println("| 4. Edit Buku")
	fmt.Println("| 5. Print Buku")
	fmt.Println("| 6. Keluar")
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
		liatBuku()
	case 3:
		HapusBuku()
	case 4:
		EditBuku()
	case 5:
		PrintBuku()
	case 6:
		os.Exit(0)

	}

	main()

}

func HapusBuku() {
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=      Hapus Buku     =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	bukuList()
	fmt.Println("+=============================================+")
	var kodeHapusBuku string
	fmt.Print("| Masukkan Kode Buku yang akan dihapus: ")
	_, err := fmt.Scanln(&kodeHapusBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	for _, buku := range ListBuku {
		if kodeHapusBuku == buku.KodeBuku {
			err := os.Remove(fmt.Sprintf("Buku/book-%s.json", buku.KodeBuku))
			if err != nil {
				fmt.Println("Terjadi Error : ", err)
				return
			}

			//ListBuku = append(ListBuku[:i], ListBuku[i+1:]...)
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=+=  Buku Berhasil Dihapus  =+=+=+=+= |")
			fmt.Println("+=============================================+")

		} else {
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=+=  Buku Tidak Ditemukan!  =+=+=+=+= |")
			fmt.Println("+=============================================+")
		}

		return
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
	bukuList()
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

			simpanHasilEditBuku(ListBuku[i], fmt.Sprintf("buku/book-%s.json", buku.KodeBuku))

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

func simpanBukuKeJson(book Tambahan) {
	encoded, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println("Terjadi Error :", err)
		return
	}

	fileName := fmt.Sprintf("buku/book-%s.json", book.KodeBuku)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Terjadi Error :", err)
		return
	}
	defer file.Close()

	_, err = file.Write(encoded)
	if err != nil {
		fmt.Println("Terjadi Error:", err)
		return
	}

}

func loadBuku() {
	files, err := os.ReadDir("buku")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := fmt.Sprintf("buku/%s", file.Name())
			jsonFile, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}
			defer jsonFile.Close()

			var buku Tambahan
			err = json.NewDecoder(jsonFile).Decode(&buku)
			if err != nil {
				fmt.Println("Telah Terjadi Error : ", err)
				return
			}

			ListBuku = append(ListBuku, buku)

		}

	}

}

func simpanHasilEditBuku(bukuTeredit Tambahan, fileName string) {
	encoded, err := json.MarshalIndent(bukuTeredit, "", "    ")
	if err != nil {
		fmt.Println("Terjadi Error : ", err)

		return
	}

	err = ioutil.WriteFile(fileName, encoded, 0644)
	if err != nil {
		fmt.Println("Terjadi Error : ", err)

		return
	}

}

func PrintBuku() {
	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=+=     Print Buku     =+=+=+=+=+= |")
	fmt.Println("+=============================================+")
	bukuList()
	fmt.Println("+=============================================+")
	var pilihan string
	fmt.Print("| Apakah Anda ingin print semua buku (y/n)?: ")
	_, err := fmt.Scanln(&pilihan)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	if strings.ToLower(pilihan) == "y" {
		printSemuaBuku()
	} else {
		printBukuTunggal()
	}
}

func printBukuTunggal() {
	var kodePrintBuku string
	fmt.Print("| Masukkan Kode Buku yang ingin di-print: ")
	_, err := fmt.Scanln(&kodePrintBuku)
	if err != nil {
		fmt.Println("Telah Terjadi Error : ", err)
		return
	}

	for _, buku := range ListBuku {
		if buku.KodeBuku == kodePrintBuku {
			buatPdf(buku)
			fmt.Println("+=============================================+")
			fmt.Println("+=============================================+")
			fmt.Println("| =+=+=+=   Buku Berhasil Di-Print!   =+=+=+= |")
			fmt.Println("+=============================================+")
			return
		}
	}

	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=+=+=  Buku Tidak Ditemukan!  =+=+=+=+= |")
	fmt.Println("+=============================================+")
}

func printSemuaBuku() {
	for _, buku := range ListBuku {
		buatPdf(buku)
	}

	fmt.Println("+=============================================+")
	fmt.Println("+=============================================+")
	fmt.Println("| =+=+=  Semua Buku Berhasil Di-Print!  =+=+= |")
	fmt.Println("+=============================================+")
}

func buatPdf(buku Tambahan) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Informasi Buku")

	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Kode Buku: %s", buku.KodeBuku))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Judul Buku: %s", buku.JudulBuku))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Pengarang: %s", buku.Pengarang))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Penerbit: %s", buku.Penerbit))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Jumlah Halaman: %s", buku.JumlahHalaman))
	pdf.Ln(5)
	pdf.Cell(40, 10, fmt.Sprintf("Tahun Terbit: %s", buku.TahunTerbit))

	err := pdf.OutputFileAndClose(fmt.Sprintf("pdf/book-%s.pdf", buku.KodeBuku))
	if err != nil {
		fmt.Println("Terjadi Error :", err)
		return
	}
}
