package main

import "fmt"

const NMAX int = 1000

type tanggal struct {
	tanggal, bulan, tahun int
}

type dataPasien struct {
	hemoglobinA1c           float64
	paket                   string
	namaDepan, namaAkhir    string
	umur                    int
	gender                  string
	TL                      tanggal
	tempatLahir             string
	beratBadan, tinggiBadan float64
	TDS, TDD                int
	kunjungan               tanggal
}

type arrayPasien [NMAX]dataPasien

func main() {
	var A arrayPasien
	var pilihan, pilihan_2, pilihan_3, category, n int
	for pilihan == 1 || pilihan == 2 || pilihan == 0 || pilihan == 3 {
		fmt.Println("================================================")
		fmt.Println()
		fmt.Println("	  SELAMAT DATANG DI DATABASE TELKOMEDIKA	 ")
		fmt.Println()
		fmt.Println("================================================")
		fmt.Println()
		fmt.Println("Silahkan Pilih menu berikut: ")
		fmt.Println("1. Cek Kesehatan")
		fmt.Println("2. Tampilkan Data Pasien")
		fmt.Println("3. Edit Data Pasien")
		fmt.Println()
		fmt.Println("Program akan berhenti jika memilih angka selain di menu")
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			if n < NMAX {
				cekKesehatan(&A, n)
				n++
			} else {
				fmt.Print("Kapasitas Pasien Penuh")
			}
		} else if pilihan == 2 {
			fmt.Println("Pilih data berikut:")
			fmt.Println("1. Tampilkan semua data")
			fmt.Println("2. Tampilkan berdasarkan kategori yang sama")
			fmt.Println("3. Tampilkan data dari yang terkini")
			fmt.Scan(&pilihan_2)
			if pilihan_2 == 1 {
				printSemuaTerurut(A, n)
			} else if pilihan_2 == 2 {
				fmt.Println()
				fmt.Println("Pilih Kategori Data:")
				fmt.Println("1. Umur")
				fmt.Println("2. gender")
				fmt.Println("3. Tahun berkunjung")
				fmt.Println("4. Nama")
				fmt.Print("Masukkan angka: ")
				fmt.Scan(&category)
				for category != 1 && category != 2 && category != 3 && category != 4 {
					fmt.Println("Masukan tidak valid")
					fmt.Scan(&category)
				}
				if category == 1 {
					printDataBerdasarkanUsia(A, n)
				} else if category == 2 {
					printDataBerdasarkanGender(A, n)
				} else if category == 3 {
					printDataBerdasarkanTahun(A, n)
				} else if category == 4 {
					printDataBerdasarkanNama(A, n)
				}
			} else if pilihan_2 == 3 {
				sortingTerbaru(A, n)
			}
		} else if pilihan == 3 {
			fmt.Println("Pilih jenis rekayasa data: ")
			fmt.Println("1. Edit Data Pasien")
			fmt.Println("2. Hapus data Pasien")
			fmt.Print("Masukkan angka: ")
			fmt.Scan(&pilihan_3)
			for pilihan_3 != 1 && pilihan != 2 && pilihan != 3 {
				fmt.Println("Masukan tidak valid")
				fmt.Scan(&pilihan_3)
			}
			if pilihan_3 == 1 {
				EditData(&A, n)
			} else if pilihan_3 == 2 {
				DeleteData(&A, &n)
			}
		}
	}
	fmt.Println("Program sudah selesai")
}

func cekKesehatan(A *arrayPasien, n int) {
	var pilihPaket int
	fmt.Println("Masukkan nama depan dan belakang:")
	fmt.Scan(&A[n].namaDepan, &A[n].namaAkhir)

	fmt.Println("Masukkan gender(Male/Female):")
	fmt.Scan(&A[n].gender)

	fmt.Println("Masukkan tempat lahir:")
	fmt.Scan(&A[n].tempatLahir)

	fmt.Println("Masukkan Tanggal Lahir:")
	fmt.Println("(format: DD MM YYYY)")
	fmt.Scan(&A[n].TL.tanggal, &A[n].TL.bulan, &A[n].TL.tahun)

	if n >= 0 && !(VerifikasiTanggal(A[n].TL)) {
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&A[n].TL.tanggal, &A[n].TL.bulan, &A[n].TL.tahun)
	}

	fmt.Println("Masukkan tanggal kedatangan:")
	fmt.Println("format: DD MM YYYY")
	fmt.Scan(&A[n].kunjungan.tanggal, &A[n].kunjungan.bulan, &A[n].kunjungan.tahun)

	if n > 0 && !(VerifikasiTanggal(A[n].kunjungan)) {
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&A[n].kunjungan.tanggal, &A[n].kunjungan.bulan, &A[n].kunjungan.tahun)
	}

	fmt.Println("Masukkan berat badan, tinggi badan:")
	fmt.Print("Berat Badan (kg): ")
	fmt.Scan(&A[n].beratBadan)
	fmt.Print("Tinggi badan (m): ")
	fmt.Scan(&A[n].tinggiBadan)

	fmt.Println("Masukkan Paket Medical Check Up:")
	fmt.Println("1. Paket Cek Tekanan Darah")
	fmt.Println("2. Paket cek Gula Darah")
	fmt.Print("Masukkan angka 1 atau 2: ")
	fmt.Scan(&pilihPaket)
	if pilihPaket == 1 {
		A[n].paket = "Cek Tekanan Darah"
		fmt.Println("Masukkan tekanan darah sistolik dan diastolik:")
		fmt.Print("Tekanan darah Sistolik (mmHg): ")
		fmt.Scan(&A[n].TDS)
		fmt.Print("Tekanan darah diastolik (mmHg): ")
		fmt.Scan(&A[n].TDD)
		A[n].hemoglobinA1c = 0.0

	} else if pilihPaket == 2 {
		A[n].paket = "Cek Gula Darah"
		fmt.Println("Masukkan angka Hemoglobin A1c: ")
		fmt.Scan(&A[n].hemoglobinA1c)
		A[n].TDS = 0
		A[n].TDD = 0
	}
	A[n].umur = hitungUmur(A[n].TL.tanggal, A[n].TL.bulan, A[n].TL.tahun, A[n].kunjungan.tanggal, A[n].kunjungan.bulan, A[n].kunjungan.tahun)
}

func VerifikasiTanggal(TL tanggal) bool {
	if (TL.tanggal >= 1 && TL.tanggal <= 31) && (TL.bulan >= 1 && TL.bulan <= 12) {
		return true
	} else {
		return false
	}
}

func hitungUmur(tglLahir, blnLahir, thnLahir, tglKunjungan, blnKunjungan, thnKunjungan int) int {
	var umur int
	umur = thnKunjungan - thnLahir
	if blnLahir > blnKunjungan {
		umur -= 1
	} else if tglLahir < tglKunjungan {
		umur -= 1
	} else {
		umur = umur
	}
	return umur
}

func printData(A arrayPasien, n int) {
	if n > 0 {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		for i := 0; i < n; i++ {
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Nama:", A[i].namaDepan, A[i].namaAkhir)
			fmt.Println("Paket:", A[i].paket)
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Umur:", A[i].umur)
			fmt.Println("Gender:", A[i].gender)
			fmt.Println("Tanggal Lahir:", A[i].TL.tanggal, "/", A[i].TL.bulan, "/", A[i].TL.tahun)
			fmt.Println("Tanggal Kunjungan:", A[i].kunjungan.tanggal, "/", A[i].kunjungan.bulan, "/", A[i].kunjungan.tahun)
			fmt.Println("Berat Badan:", A[i].beratBadan)
			fmt.Println("Tinggi Badan:", A[i].tinggiBadan)
			fmt.Println("Tekanan Darah Sistolik: ", A[i].TDS)
			fmt.Println("Tekanan Darah Diatolik: ", A[i].TDD)
			fmt.Println("----------------------------------------------------")
		}
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func printDataBerdasarkanUsia(A arrayPasien, n int) {
	var B arrayPasien
	var m int
	m = 0
	var age int
	fmt.Println("Masukkan Umur:")
	fmt.Scan(&age)
	for i := 0; i < n; i++ {
		if A[i].umur == age {
			B[m] = A[i]
			m++
		}
	}
	printData(B, m)
}

func printDataBerdasarkanGender(A arrayPasien, n int) {
	var C arrayPasien
	var m int
	m = 0
	var gender string
	fmt.Println("Male/Female?")
	fmt.Scan(&gender)
	for i := 0; i < n; i++ {
		if A[i].gender == gender {
			C[m] = A[i]
			m++
		}
	}
	printData(C, m)
}

func printDataBerdasarkanTahun(A arrayPasien, n int) {
	var D arrayPasien
	var m int
	m = 0
	var year int
	fmt.Println("Masukkan Tahun:")
	fmt.Scan(&year)
	for i := 0; i < n; i++ {
		if A[i].kunjungan.tahun == year {
			D[m] = A[i]
			m++
		}
	}
	printData(D, m)
}

func printDataBerdasarkanNama(A arrayPasien, n int) {
	var D arrayPasien
	var m int
	m = 0
	var namaDepan, namaAkhir string
	fmt.Println("Masukkan Tahun:")
	fmt.Scan(&namaDepan, &namaAkhir)
	for i := 0; i < n; i++ {
		if A[i].namaDepan == namaDepan && A[i].namaAkhir == namaAkhir {
			D[m] = A[i]
			m++
		}
	}
	printData(D, m)
}

func printSemuaTerurut(A arrayPasien, n int) {
	var pilihan, pilihanSorting int
	fmt.Println("Tampilkan terurut berdasarkan:")
	fmt.Println("1. Usia")
	fmt.Println("2. Tanggal Medical Checkup")
	fmt.Println("Masukkan angka: ")
	fmt.Scan(&pilihan)
	if pilihan == 1 {
		fmt.Println("Menaik Atau menurun?")
		fmt.Println("1. Menaik")
		fmt.Println("2. Menurun")
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&pilihanSorting)
		if pilihanSorting == 1 {
			ascendingUsiaInsertion(A, n)
		} else if pilihanSorting == 2 {
			descendingUsiaSelecton(A, n)
		}
	} else if pilihan == 2 {
		fmt.Println("Menaik Atau menurun?")
		fmt.Println("1. Menaik")
		fmt.Println("2. Menurun")
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&pilihanSorting)
		if pilihanSorting == 1 {
			ascendingTanggalInsertion(A, n)
		} else if pilihanSorting == 2 {
			descendingTanggalSelecton(A, n)
		}
	}

}

func sortingTerbaru(A arrayPasien, n int) {
	var E arrayPasien
	var m int
	m = 0
	for i := 0; i < n; i++ {
		E[m] = A[n-1-i]
		m++
	}
	printData(E, m)
}

func EditData(A *arrayPasien, n int) {
	var namaDepan, namaAkhir string
	fmt.Println("Masukkan nama lengkap pasien yang ingin diedit:")
	fmt.Scan(&namaDepan, &namaAkhir)

	found := false // Variabel flag untuk menandai apakah nama ditemukan atau tidak

	for i := 0; i < n; i++ {
		if (A[i].namaDepan == namaDepan) && (A[i].namaAkhir == namaAkhir) {
			// Menampilkan dan memperbarui data pasien
			found = true
			fmt.Println("Data pasien yang akan diedit:")
			fmt.Println("Nama:", A[i].namaDepan, A[i].namaAkhir)
			fmt.Println("Umur:", A[i].umur)
			fmt.Println("Gender:", A[i].gender)
			fmt.Println("Tanggal Lahir:", A[i].TL.tanggal, "/", A[i].TL.bulan, "/", A[i].TL.tahun)
			fmt.Println("Tanggal Kunjungan:", A[i].kunjungan.tanggal, "/", A[i].kunjungan.bulan, "/", A[i].kunjungan.tahun)
			fmt.Println("Berat Badan:", A[i].beratBadan)
			fmt.Println("Tinggi Badan:", A[i].tinggiBadan)
			fmt.Println("---------------------")

			// Memperbarui data pasien
			var choice int
			fmt.Println("Apa yang ingin Anda edit?")
			fmt.Println("1. Nama")
			fmt.Println("2. Gender")
			fmt.Println("3. Tanggal Lahir")
			fmt.Println("4. Tanggal Kunjungan")
			fmt.Println("5. Berat Badan")
			fmt.Println("6. Tinggi Badan")

			fmt.Scan(&choice)

			if choice == 1 {
				fmt.Println("Masukkan nama baru (terdiri dari nama depan dan nama belakang):")
				fmt.Scan(&A[i].namaDepan, &A[i].namaAkhir)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 2 {
				fmt.Println("Masukkan gender baru:")
				fmt.Scan(&A[i].gender)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 3 {
				fmt.Println("Masukkan Tanggal Lahir baru (format: DD MM YYYY):")
				fmt.Scan(&A[i].TL.tanggal, &A[i].TL.bulan, &A[i].TL.tahun)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 4 {
				fmt.Println("Masukkan tanggal kunjungan baru (format: DD MM YYYY):")
				fmt.Scan(&A[i].kunjungan.tanggal, &A[i].kunjungan.bulan, &A[i].kunjungan.tahun)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 5 {
				fmt.Println("Masukkan berat badan baru:")
				fmt.Scan(&A[i].beratBadan)
				fmt.Println("Data berhasil diupdate!")
			} else if choice == 6 {
				fmt.Println("Masukkan tinggi badan baru:")
				fmt.Scan(&A[i].tinggiBadan)
				fmt.Println("Data berhasil diupdate!")
			} else {
				fmt.Println("Pilihan tidak valid")
			}
		}
	}

	if !found {
		fmt.Println("Pasien tidak ditemukan.")
	}
}
func DeleteData(A *arrayPasien, n *int) {
	var namaDepan, namaAkhir string
	fmt.Println("Masukkan nama lengkao pasien yang ingin dihapus:")
	fmt.Scan(&namaDepan, &namaAkhir)

	found := false // Variabel flag untuk menandai apakah ID ditemukan atau tidak

	for i := 0; i < *n; i++ {
		if A[i].namaDepan == namaDepan && A[i].namaAkhir == namaAkhir {
			found = true
			// Geser data ke kiri untuk menutup celah
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("Data berhasil dihapus.")
			break // Keluar dari loop setelah data ditemukan dan dihapus
		}
	}

	if !found {
		fmt.Println("nama pasien tidak ditemukan.")
	}
}

func ascendingUsiaInsertion(A arrayPasien, k int) {
	n := len(A)
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && A[j].umur > key.umur {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
	printData(A, k)
}

func descendingUsiaSelecton(A arrayPasien, k int) {
	n := len(A)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if A[j].umur < A[minIndex].umur {
				minIndex = j
			}
		}
		A[i], A[minIndex] = A[minIndex], A[i]
	}
	printData(A, k)
}

func ascendingTanggalInsertion(A arrayPasien, k int) {
	n := len(A)
	for i := 1; i < n; i++ {
		key := A[i]
		j := i - 1

		// Pindahkan elemen arr[0..i-1], yang lebih besar dari key,
		// ke satu posisi di depan posisi sekarang
		for j >= 0 && kurangDari(key.kunjungan, A[j].kunjungan) {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
	printData(A, k)
}

func descendingTanggalSelecton(A arrayPasien, k int) {
	n := len(A)
	for i := 0; i < n-1; i++ {
		// Temukan elemen maksimum dalam array yang belum diurutkan
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if lebihDari(A[j].kunjungan, A[maxIdx].kunjungan) {
				maxIdx = j
			}
		}
		// Tukar elemen maksimum yang ditemukan dengan elemen pertama
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}
	printData(A, k)
}

func lebihDari(a, b tanggal) bool {
	if a.tahun != b.tahun {
		return a.tahun > b.tahun
	}
	if a.bulan != b.bulan {
		return a.bulan > b.bulan
	}
	return a.tanggal > b.tanggal
}

func kurangDari(a, b tanggal) bool {
	if a.tahun != b.tahun {
		return a.tahun < b.tahun
	}
	if a.bulan != b.bulan {
		return a.bulan < b.bulan
	}
	return a.tanggal < b.tanggal
}
