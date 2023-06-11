package main

import (
	"fmt"
)

const NMAX int = 100

type SparePart struct {
	name, idSpare        string
	stok, harga, counter int
}

type Transaksi struct {
	qty, total       int
	idTrans, tanggal string
	pelanggan        [NMAX]Customer
	sparepart        [NMAX]SparePart
}
type Customer struct {
	name, idCust string
}

type arrSparePart [NMAX]SparePart
type arrTransaksi [NMAX]Transaksi
type arrCustomer [NMAX]Customer

func main() {
	menu()
}
func menu() {
	var sparepart arrSparePart
	var cust arrCustomer
	var transaction arrTransaksi
	var n, j, x, pilihan, pilihan2, pilihan3, pilihan4, pilihcari int
	var ulang bool
	ulang = true

	for ulang && pilihan == 0 {
		fmt.Println(" ")
		fmt.Println("========================")
		fmt.Println(" Aplikasi Service Motor ")
		fmt.Println("========================")
		fmt.Println("1.Data Sparepart")
		fmt.Println("2.Data Customer")
		fmt.Println("3.Data Transaksi")
		fmt.Println("4.Exit")
		fmt.Print("Pilihan Anda : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Println(" ")
			fmt.Println("====Data Sparepart====")
			fmt.Println("1.Tambah Data")
			fmt.Println("2.Edit Data")
			fmt.Println("3.Hapus Data")
			fmt.Println("4.Lihat Data")
			fmt.Println("5.Urut Data yang paling laku")
			fmt.Print("Pilihan Anda : ")
			fmt.Scan(&pilihan2)
			fmt.Println(" ")
			if pilihan2 == 1 {
				addSparepart(&sparepart, &n)
			} else if pilihan2 == 2 {
				editSparepart(&sparepart, n)
			} else if pilihan2 == 3 {
				deleteSparepart(&sparepart, &n)
			} else if pilihan2 == 4 {
				printSparePart(sparepart, n)
			} else if pilihan2 == 5 {
				insertionSortLaku(&sparepart, n)
				printSparePart(sparepart, n)
			} else {
				fmt.Println("Cek kembali input anda")
			}

			ulang = true
			pilihan = 0
		} else if pilihan == 2 {
			fmt.Println(" ")
			fmt.Println("====Data Customer====")
			fmt.Println("1.Tambah Data")
			fmt.Println("2.Edit Data")
			fmt.Println("3.Hapus Data")
			fmt.Println("4.Lihat Data")
			fmt.Print("Pilihan Anda : ")
			fmt.Scan(&pilihan3)
			fmt.Println(" ")

			if pilihan3 == 1 {
				addCust(&cust, &j)
			} else if pilihan3 == 2 {
				editCust(&cust, j)
			} else if pilihan3 == 3 {
				deleteCust(&cust, &j)
			} else if pilihan3 == 4 {
				printCust(cust, j)
			} else {
				fmt.Println("Cek kembali input anda")
			}

			ulang = true
			pilihan = 0
		} else if pilihan == 3 {
			fmt.Println(" ")
			fmt.Println("====Data Transaksi====")
			fmt.Println("1.Tambah Data")
			fmt.Println("2.Edit Data")
			fmt.Println("3.Hapus Data")
			fmt.Println("4.Lihat Data")
			fmt.Println("5.Cari Data")
			fmt.Print("Pilihan Anda : ")
			fmt.Scan(&pilihan4)
			fmt.Println(" ")

			if pilihan4 == 1 {
				addTransaksi(&transaction, &x, &sparepart, cust)
			} else if pilihan4 == 2 {
				editTransaksi(&transaction, x, &sparepart, cust)
			} else if pilihan4 == 3 {
				deleteTransaksi(&transaction, &x, &sparepart, cust)
			} else if pilihan4 == 4 {
				printTransaksi(transaction, x)
			} else if pilihan4 == 5 {
				fmt.Println("1.Cari Berdasarkan Waktu")
				fmt.Println("2.Cari Berdasarkan Spare-part")
				fmt.Print("Pilihan Anda : ")
				fmt.Scan(&pilihcari)
				if pilihcari == 1 {
					searchtgl(transaction, x)
				} else if pilihcari == 2 {
					searchSparepart(transaction, x)
				}
			} else {
				fmt.Println("Cek kembali input anda")
			}

			ulang = true
			pilihan = 0
		} else if pilihan == 4 {
			ulang = false
		} else {
			fmt.Println("Maaf inputan salah")
			ulang = true
			pilihan = 0
		}
	}
}

func konfirmasi() bool {
	var konf string
	fmt.Print("Apakah Anda yakin ? (Y/N) ")
	fmt.Scan(&konf)
	if konf == "Y" || konf == "y" {
		return true
	} else {
		return false
	}

}
func cariDataSparepartSeq(S arrSparePart, n int, cariId string) int {
	for i := 0; i < n; i++ {
		if S[i].idSpare == cariId {
			return i
		}
	}
	return -1
}

func cariDataTransSeq(T arrTransaksi, n int, cariId string) int {
	for i := 0; i < n; i++ {
		if T[i].idTrans == cariId {
			return i
		}
	}
	return -1
}
func cariDataCustBin(C arrCustomer, n int, cariId string) int {
	var kr, kn, med int
	var found int
	kr = 0
	kn = n - 1
	found = -1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if C[med].idCust < cariId {
			kr = med + 1
		} else if C[med].idCust > cariId {
			kn = med
		} else {
			found = med
		}
	}

	return found

}


func addSparepart(S *arrSparePart, n *int) {
	var id, name string
	var qty, price int 
	fmt.Print("Masukkan ID Sparepart : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Sparepart : ")
	fmt.Scan(&name)
	fmt.Print("Jumlah Sparepart : ")
	fmt.Scan(&qty)
	fmt.Print("Harga Sparepart / Pcs : ")
	fmt.Scan(&price)          
	if konfirmasi() == true { 
		S[*n].idSpare = id
		S[*n].name = name
		S[*n].stok = qty
		S[*n].harga = price                      
		fmt.Println("Data berhasil ditambahkan") 
		*n++                                     
	} else {
		fmt.Print("Data Gagal ditambahkan")
	}

}
func editSparepart(S *arrSparePart, n int) {
	var name string
	var qty, price int
	var cariId string 
	fmt.Print("ID Sparepart yang ingin dirubah : ")
	fmt.Scan(&cariId)                             
	if cariDataSparepartSeq(*S, n, cariId) != -1 { 
		i := cariDataSparepartSeq(*S, n, cariId)
		fmt.Print("Masukkan Nama Sparepart : ")
		fmt.Scan(&name)
		fmt.Print("Jumlah Sparepart : ")
		fmt.Scan(&qty)
		fmt.Print("Harga Sparepart / Pcs : ")
		fmt.Scan(&price)          
		if konfirmasi() == true { 
			S[i].name = name
			S[i].stok = qty
			S[i].harga = price                     
			fmt.Println("Data Berhasil di update")
		} else {
			fmt.Println("Data gagal di update") 
		}
	} else {
		fmt.Println("Data yang dicari tidak ada")
		fmt.Println(" ")
	}
}
func deleteSparepart(S *arrSparePart, n *int) {
	var cariId string 
	fmt.Print("ID Sparepart yang ingin dihapus : ")
	fmt.Scan(&cariId)
	if cariDataSparepartSeq(*S, *n, cariId) != -1 {
		idx := cariDataSparepartSeq(*S, *n, cariId)
		if konfirmasi() {
			for j := idx; j < *n-1; j++ { 
				S[j] = S[j+1] 

			}
			fmt.Println("Data Berhasil Di hapus") 
			*n--
		} else {
			fmt.Println("Data gagal di hapus") 
		}
	} else {
		fmt.Println("Data yang dicari tidak ada") 
		fmt.Println(" ")
	}
}


func addCust(C *arrCustomer, n *int) {
	var id, name string
	fmt.Print("Masukkan ID Customer : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Customer : ")
	fmt.Scan(&name)
	if konfirmasi() == true {
		C[*n].idCust = id
		C[*n].name = name
		fmt.Println("Data berhasil ditambahkan")
		fmt.Println("")
		*n++
	} else {
		fmt.Print("Data Gagal ditambahkan")
		fmt.Println("")
	}
}

func editCust(C *arrCustomer, n int) {
	var name string
	var cariId string
	fmt.Print("ID Customer yang ingin dirubah : ")
	fmt.Scan(&cariId)

	if cariDataCustBin(*C, n, cariId) != -1 {
		fmt.Print("Masukkan Nama Customer : ")
		fmt.Scan(&name)
		i := cariDataCustBin(*C, n, cariId)
		if konfirmasi() == true {
			C[i].name = name
			fmt.Println("Data Berhasil di update")
		} else {
			fmt.Println("Data gagal di update")
		}
	} else {
		fmt.Println("Data yang dicari tidak ada")
		fmt.Println(" ")
	}

}
func deleteCust(C *arrCustomer, n *int) {
	var cariId string
	fmt.Print("ID Customer yang ingin dihapus : ")
	fmt.Scan(&cariId)

	if cariDataCustBin(*C, *n, cariId) != -1 {
		idx := cariDataCustBin(*C, *n, cariId)
		if konfirmasi() {
			for j := idx; j < *n-1; j++ {
				C[j] = C[j+1]

			}
			fmt.Println("Data Berhasil di hapus")
			*n--
		} else {
			fmt.Println("Data gagal di hapus")
		}
	} else {
		fmt.Println("Data yang dicari tidak ada")
		fmt.Println(" ")
	}
}


func addTransaksi(T *arrTransaksi, n *int, spareParts *arrSparePart, costumers arrCustomer) {
	var idCust, idSpare, id, tgl string
	var qty int
	var idxCust, idxSpare int

	fmt.Print("Masukkan ID Transaksi : ")
	fmt.Scan(&id)
	fmt.Print("Masukkan ID Customer : ")
	fmt.Scan(&idCust)
	fmt.Print("Masukkan ID Sparepart : ")
	fmt.Scan(&idSpare)
	fmt.Print("Jumlah Sparepart : ")
	fmt.Scan(&qty)
	fmt.Print("Tanggal Transaksi : ")
	fmt.Scan(&tgl)

	for i := 0; i < len(costumers); i++ {
		if idCust == costumers[i].idCust {
			idxCust = i

		}
	}
	for i := 0; i < len(spareParts); i++ {
		if idSpare == spareParts[i].idSpare {
			idxSpare = i

		}
	}

	if konfirmasi() == true {
		T[*n].idTrans = id
		T[*n].qty = qty
		T[*n].pelanggan[*n].name = costumers[idxCust].name
		T[*n].sparepart[*n].name = spareParts[idxSpare].name
		T[*n].sparepart[*n].idSpare = spareParts[idxSpare].idSpare
		T[*n].total = qty * spareParts[idxSpare].harga
		T[*n].tanggal = tgl
		spareParts[idxSpare].stok -= qty
		spareParts[idxSpare].counter += qty
		fmt.Println("Data berhasil ditambahkan")
		*n++
	} else {
		fmt.Print("Data Gagal ditambahkan")
	}

}
func editTransaksi(T *arrTransaksi, n int, spareParts *arrSparePart, costumers arrCustomer) {
	var idCust, idSpare, tgl string
	var qty int
	var idxCust, idxSpare, found int

	var cariId string
	fmt.Print("ID Transaksi yang ingin dirubah : ")
	fmt.Scan(&cariId)

	if cariDataTransSeq(*T, n, cariId) != -1 {
		j := cariDataTransSeq(*T, n, cariId)
		fmt.Print("Masukkan ID Customer : ")
		fmt.Scan(&idCust)
		fmt.Print("Masukkan ID Sparepart : ")
		fmt.Scan(&idSpare)
		fmt.Print("Jumlah Sparepart : ")
		fmt.Scan(&qty)
		fmt.Print("Tanggal Transaksi : ")
		fmt.Scan(&tgl)

		for i := 0; i < len(costumers); i++ {
			if idCust == costumers[i].idCust {
				idxCust = i

			}
		}
		for i := 0; i < len(spareParts); i++ {
			if idSpare == spareParts[i].idSpare {
				idxSpare = i

			}
		}

		if konfirmasi() == true {
			for x := 0; x < len(spareParts); x++ {
				if spareParts[x].idSpare == T[j].sparepart[j].idSpare {
					found = x
				}
			}
			if spareParts[j].idSpare != idSpare {
				spareParts[found].stok += T[j].qty
				spareParts[found].counter -= T[j].qty
				T[j].qty = qty
				T[j].pelanggan[j].name = costumers[idxCust].name
				T[j].sparepart[j].name = spareParts[idxSpare].name
				T[j].sparepart[j].idSpare = spareParts[idxSpare].idSpare
				T[j].total = qty * spareParts[idxSpare].harga
				T[j].tanggal = tgl
				spareParts[idxSpare].stok -= qty
				spareParts[idxSpare].counter += qty

			} else {

				spareParts[found].stok += T[j].qty
				spareParts[found].counter -= T[j].qty
				T[j].qty = qty
				T[j].pelanggan[j].name = costumers[idxCust].name
				T[j].sparepart[j].name = spareParts[idxSpare].name
				T[j].total = qty * spareParts[idxSpare].harga
				T[j].tanggal = tgl
				spareParts[idxSpare].stok -= qty
				spareParts[idxSpare].counter += qty
			}
			fmt.Println("Data berhasil diupdate")
		} else {
			fmt.Print("Data Gagal diupdate")
		}
	} else {
		fmt.Println("Data yang dicari tidak ada")
		fmt.Println(" ")
	}
}
func deleteTransaksi(T *arrTransaksi, n *int, spareParts *arrSparePart, costumers arrCustomer) {

	var cariId string
	var found int
	fmt.Print("ID Transaksi yang ingin dihapus : ")
	fmt.Scan(&cariId)
	if cariDataTransSeq(*T, *n, cariId) != -1 {
		idx := cariDataTransSeq(*T, *n, cariId)
		if konfirmasi() == true {
			for x := 0; x < len(spareParts); x++ {
				if spareParts[x].idSpare == T[idx].sparepart[idx].idSpare {
					found = x
				}
			}
			spareParts[found].stok += T[idx].qty
			spareParts[found].counter -= T[idx].qty
			for j := idx; j < *n-1; j++ {

				T[j] = T[j+1]

			}
			fmt.Println("Data berhasil dihapus")
			*n--
		} else {
			fmt.Print("Data Gagal diupdate")
		}
	} else {
		fmt.Println("Data yang dicari tidak ada")
		fmt.Println(" ")
	}
}


func printSparePart(S arrSparePart, n int) {
	fmt.Println("--------------")
	fmt.Println("Data SparePart")
	fmt.Println("--------------")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, " ")
		fmt.Print(S[i].idSpare, " ")
		fmt.Print(S[i].name, " ")
		fmt.Print(S[i].stok, " ")
		fmt.Print(S[i].counter, " ")
		fmt.Println(S[i].harga)

	}

}
func printCust(C arrCustomer, n int) {
	fmt.Println("--------------")
	fmt.Println("Data Customer")
	fmt.Println("--------------")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, " ")
		fmt.Print(C[i].idCust, " ")
		fmt.Println(C[i].name)

	}

}
func printTransaksi(T arrTransaksi, n int) {
	fmt.Println("--------------")
	fmt.Println("Data Transaksi")
	fmt.Println("--------------")
	for i := 0; i < n; i++ {
		fmt.Print(i+1, " ")
		fmt.Print(T[i].idTrans, " ")
		for j := 0; j < NMAX; j++ {
			if T[i].pelanggan[j].name != "" {
				fmt.Print(T[i].pelanggan[j].name, " ")
			}
		}

		for j := 0; j < NMAX; j++ {
			if T[i].sparepart[j].name != "" {
				fmt.Print(T[i].sparepart[j].name, " ")
			}
		}
		fmt.Print(T[i].qty, " ")
		fmt.Print(T[i].total, " ")
		fmt.Println(T[i].tanggal)
	}

}


func SelectionSortTanggal(T *arrTransaksi, n int) {

	var idx_min, i, j int
	var t Transaksi
	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = 1
		for j < n {
			if T[idx_min].tanggal < T[j].tanggal {
				idx_min = j
			}
			j += 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i += 1

	}
}
func insertionSortLaku(S *arrSparePart, n int) {
	var i, j int
	var temp SparePart
	i = 1
	for i <= n-1 {
		j = i
		temp = S[j]
		for j > 0 && temp.counter > S[j-1].counter {
			S[j] = S[j-1]
			j = j - 1
		}
		S[j] = temp
		i += 1
	}
}
func searchtgl(T arrTransaksi, n int) {
	SelectionSortTanggal(&T, n)
	var idxTanggal int
	var tgl string
	fmt.Print("Masukkan Tanggal Yang dicari : ")
	fmt.Scan(&tgl)
	for i := 0; i < n; i++ {
		if tgl == T[i].tanggal {
			idxTanggal = i
			fmt.Print(T[idxTanggal].idTrans, " ")
			for j := 0; j < NMAX; j++ {
				if T[idxTanggal].pelanggan[j].name != "" {
					fmt.Print(T[idxTanggal].pelanggan[j].name, " ")
				}
			}
			for j := 0; j < NMAX; j++ {
				if T[idxTanggal].sparepart[j].name != "" {
					fmt.Print(T[idxTanggal].sparepart[j].name, " ")
				}
			}

			fmt.Print(T[idxTanggal].qty, " ")
			fmt.Print(T[idxTanggal].total, " ")
			fmt.Println(T[idxTanggal].tanggal)
		}
	}
}

func searchSparepart(T arrTransaksi, n int) {
	var idxSpare int
	var searchSparepart string

	fmt.Print("Masukkan ID Spare-part Yang dicari : ")
	fmt.Scan(&searchSparepart)

	for i := 0; i < n; i++ {
		if T[i].sparepart[i].idSpare == searchSparepart {
			idxSpare = i
			fmt.Print(T[idxSpare].idTrans, " ")
			for j := 0; j < NMAX; j++ {
				if T[idxSpare].pelanggan[j].name != "" {
					fmt.Print(T[idxSpare].pelanggan[j].name, " ")
				}
			}
			for j := 0; j < NMAX; j++ {
				if T[idxSpare].sparepart[j].name != "" {
					fmt.Print(T[idxSpare].sparepart[j].name, " ")
				}
			}
			fmt.Print(T[idxSpare].qty, " ")
			fmt.Print(T[idxSpare].total, " ")
			fmt.Println(T[idxSpare].tanggal)

		}

	}
}
