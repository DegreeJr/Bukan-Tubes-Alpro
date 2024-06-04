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
	var pilihan, n int
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
			if pilihan == 1{
				printData(A, n)
			}	
		}
	}	
}

func cekKesehatan(P *Pasien, n int) {
	fmt.Println("Masukkan nama:")
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
		} 
	}
}
