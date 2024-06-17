package main
import "fmt"

const NMAX int = 1000

type tanggal struct {
	tanggal, bulan, tahun int
}

type dataPasien struct{
	nama                    string
	umur                    int
	gender                  string
	TL                     tanggal
	tempatLahir				string
	beratBadan, tinggiBadan float64
	TDS, TDD                int
	kunjungan				tanggal
}

type arrayPasien {NMAX} dataPasien

func main{
	var A arrayPasien
	var pilihan,pilihan_2, pilihan_3, category, n int
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
			}else{
				fmt.Print("Kapasitas Pasien Penuh")
			}	
		}else if pilihan == 2{
			fmt.Println("Pilih data berikut:")
			fmt.Println("1. Tampilkan semua data")
			fmt.Println("2. Tampilkan berdasarkan kategori")
			fmt.Println("3. Tampilkan data dari yang terkini")
			fmt.Scan(&pilihan_2)
			if pilihan_2 == 1{
				printData(A, n)
			}else if pilihan_2 == 2{
				fmt.Println()
				fmt.Println("Pilih Kategori Data:")
				fmt.Println("1. Umur")
				fmt.Println("2. gender")
				fmt.Println("3. Tahun berkunjung")
				fmt.Print("Masukkan angka: ")
				fmt.Scan(&category)
				for category != 1 && category != 2 && category != 3 {
					fmt.Println("Masukan tidak valid")
					fmt.Scan(&category)
				}
				if category == 1{
					printDataBerdasarkanUsia(A, n)
				}else if category == 2 {
					printDataBerdasarkanGender(A,n)
				}else if category == 3 {
					printDataBerdasarkanTahun(A,n)
				}
			}else if pilihan_2 == 3{
				
			}
		}
	}	
}

func cekKesehatan(P *Pasien, n int) {
	fmt.Println("Masukkan nama lengkap:")
	fmt.Scan(%A[0].nama)
	
	fmt.Println("Masukkan gender(Male/Female):")
	fmt.Scan(&P[n].gender)

	fmt.Println("Masukkan tempat lahir:")
	fmt.Scan(&P[n].tempatLahir)

	fmt.Println("Masukkan Tanggal Lahir:")
	fmt.Println("(format: DD MM YYYY)")
	fmt.Scan(&P[n].TL.tanggal, &P[n].TL.bulan, &P[n].TL.tahun)

	if n >= 0 && !(VerifikasiTanggal(P[n].TTL)) {
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&P[n].TL.tanggal, &P[n].TL.bulan, &P[n].TL.tahun)
	}

	fmt.Println("Masukkan tanggal kedatangan:")
	fmt.Println("format: DD MM YYYY")
	fmt.Scan(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)

	if n > 0 && !(VerifikasiTanggal(P[n].kunjungan)){
		fmt.Println("Input tanggal tidak Sesuai!. Masukkan kembali")
		fmt.Scan(&P[n].kunjungan.tanggal, &P[n].kunjungan.bulan, &P[n].kunjungan.tahun)
	}
		
	fmt.Println("Masukkan berat badan, tinggi badan:")
	fmt.Print("Berat Badan (kg): ")
	fmt.Scan(&P[n].beratBadan)
	fmt.Print("Tinggi badan (m): ")
	fmt.Scan(&P[n].tinggiBadan)

	fmt.Println("Masukkan tekanan darah sistolik dan diastolik:")
	fmt.Print("Tekanan darah Sistolik (mmHg): ")
	fmt.Scan(&P[n].TDS)
	fmt.Print("Tekanan darah diastolik (mmHg): ")
	fmt.Scan(&P[n].TDD)

	P[n].umur = hitungUmur(P[n].TL.tanggal, P[n].TL.bulan, P[n].TL.tahun, P[n].kunjungan.tanggal, P[n].kunjungan.bulan, P[n].kunjungan.tahun)
}

func VerifikasiTanggal(TL tanggal)bool{
	if (TL.tanggal >= 1 && TL.tanggal <= 31) && (TL.bulan >= 1 && TL.bulan <= 12) {
		return true
	} else {
		return false
	}
}

func hitungUmur(tglLahir, blnLahir, thnLahir, tglKunjungan, blnKunjungan, thnKunjungan int)int{
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

func printData(A arrayPasien, n i int){
	if n > 0{
		fmt.Println()
		fmt.Println()
		fmt.Println()
		for i := 0; i < n; i++{
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Nama:", P[i].Name)
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Umur:", P[i].umur)
			fmt.Println("Gender:", P[i].gender)
			fmt.Println("Tanggal Lahir:", P[i].TL.tanggal, "/", P[i].TL.bulan, "/", P[i].TL.tahun)
			fmt.Println("Tanggal Kunjungan:", P[i].kunjungan.tanggal, "/", P[i].kunjungan.bulan, "/", P[i].kunjungan.tahun)
			fmt.Println("Berat Badan:", P[i].beratBadan)
			fmt.Println("Tinggi Badan:", P[i].tinggiBadan)
			fmt.Println("Tekanan Darah Sistolik: ", P[i].TDS)
			fmt.Println("Tekanan Darah Diatolik: ", P[i].TDD)
			fmt.Println("----------------------------------------------------")
		} 
	}else{
		fmt.Println("Data tidak ditemukan.")
	}
}

func printDataBerdasarkanUsia(P Pasien, n int){
	var B Pasien
	var m int
	m = 0
	var age int
	fmt.Println("Masukkan Umur:")
	fmt.Scan(&age)
	for i := 0; i < n; i++ {
		if P[i].umur == age {
			B[m] = P[i]
			m++
		}
	}
	ShowData(B, m)
}

func printDataBerdasarkanGender(P Pasien, n int) {
	var C Pasien
	var m int
	m = 0
	var gender string
	fmt.Println("Male/Female?")
	fmt.Scan(&gender)
	for i := 0; i < n; i++ {
		if P[i].gender == gender {
			C[m] = P[i]
			m++
		}
	}
	ShowData(C, m)
}

func printDataBerdasarkanTahun(P Pasien, n int) {
	var D Pasien
	var m int
	m = 0
	var year int
	fmt.Println("Masukkan Tahun:")
	fmt.Scan(&year)
	for i := 0; i < n; i++ {
		if P[i].kunjungan.tahun == year {
			D[m] = P[i]
			m++
		}
	}
	ShowData(D, m)
}

