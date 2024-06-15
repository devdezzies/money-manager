package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"os/exec"
)

const mxN int = 1000

type Transaction struct {
	Category int
	Date     Date
	Amount   float64
	Status   bool // true for income, false for outcome
}

type User struct {
	Name, Password     string
	Balance            float64
	Income, Expense    float64
	TotalTransaction   int
	TransactionHistory [mxN]Transaction
}

type Date struct {
	Year, Month, Day int
}

type tabUser [mxN]User

func clear_screen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func is_leap_year(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

// func binary_search(arr *tabUser, n *int, x int) {
// 	var left, mid, right, found int
// 	left = 0
// 	right = *n-1
// 	mid = (left + right) / 2
// 	found = -1
// }

func insertion_sort_ts_descending(arr *tabUser, n *int) {
	/* */
	var pass, i int
	var move User
	for pass = 1; pass < *n; pass++ {
		move = arr[pass]
		i = pass
		for i > 0 && float64(move.TotalTransaction) > float64(arr[i-1].TotalTransaction) { // ordered Descending
			arr[i] = arr[i-1]
			i--
		}
		arr[i] = move
	}
}

func insertion_sort_ts_ascending(arr *tabUser, n *int) {
	/* */
	var pass, i int
	var move User
	for pass = 1; pass < *n; pass++ {
		move = arr[pass]
		i = pass
		for i > 0 && float64(move.TotalTransaction) < float64(arr[i-1].TotalTransaction) { // ordered Descending
			arr[i] = arr[i-1]
			i--
		}
		arr[i] = move
	}
}

func selection_sort_balance_descending(arr *tabUser, n *int) {
	/* */
	var pass, i, move int
	var temp User
	for pass = 1; pass < *n; pass++ {
		move = pass - 1
		for i = pass; i < *n; i++ {
			if arr[move].Balance < arr[i].Balance { // Descending order
				move = i
			}
		}
		temp = arr[pass-1]
		arr[pass-1] = arr[move]
		arr[move] = temp
	}
}

func selection_sort_balance_ascending(arr *tabUser, n *int) {
	/* */
	var pass, i, move int
	var temp User
	for pass = 1; pass < *n; pass++ {
		move = pass - 1
		for i = pass; i < *n; i++ {
			if arr[move].Balance > arr[i].Balance { // Descending order
				move = i
			}
		}
		temp = arr[pass-1]
		arr[pass-1] = arr[move]
		arr[move] = temp
	}
}

func write_data(arr *tabUser, n *int) {
	/*I.S
	  F.S */
	file, _ := os.Create("user.dat")
	defer file.Close()

	encode := gob.NewEncoder(file)
	_ = encode.Encode(*arr)
	_ = encode.Encode(*n)

}

func read_data(arr *tabUser, n *int) {
	/* */
	file, _ := os.Open("user.dat")
	defer file.Close()
	decoder := gob.NewDecoder(file)
	_ = decoder.Decode(arr)
	_ = decoder.Decode(n)
}

// WRITE-READ FUNCTIONALITY

func delete_data(T *tabUser, n *int, xLoc int) {
	/* */
	for i := xLoc; i < *n-1; i++ {
		T[i] = T[i+1]
	}
	*n--
	write_data(T, n)
	fmt.Printf("%2s%s\n", "", "Akun ini berhasil dihapus!")
	clear_screen()
	user_homepage(T, n, xLoc)
}

func sequential_search(T *tabUser, n *int, x string) int {
	/* */
	var found, i int
	i = 0
	found = -1
	for i < *n && found == -1 {
		if T[i].Name == x {
			found = i
		}
		i++
	}
	return found
}

func sort_by_date_newest(arr *tabUser, loc int) {
	var pass, i int
	var move Transaction
	for pass = 1; pass < arr[loc].TotalTransaction; pass++ {
		move = arr[loc].TransactionHistory[pass]
		i = pass
		for i > 0 && ((move.Date.Year > arr[loc].TransactionHistory[i-1].Date.Year) || (move.Date.Year == arr[loc].TransactionHistory[i-1].Date.Year && move.Date.Month > arr[loc].TransactionHistory[i-1].Date.Month) || (move.Date.Year == arr[loc].TransactionHistory[i-1].Date.Year && move.Date.Month == arr[loc].TransactionHistory[i-1].Date.Month && move.Date.Day > arr[loc].TransactionHistory[i-1].Date.Day)) {
			arr[loc].TransactionHistory[i] = arr[loc].TransactionHistory[i-1]
			i--
		}
		arr[loc].TransactionHistory[i] = move
	}
}

func sort_by_date_oldest(arr *tabUser, loc int) {
	var pass, i int
	var move Transaction
	for pass = 1; pass < arr[loc].TotalTransaction; pass++ {
		move = arr[loc].TransactionHistory[pass]
		i = pass
		for i > 0 && ((move.Date.Year < arr[loc].TransactionHistory[i-1].Date.Year) || (move.Date.Year == arr[loc].TransactionHistory[i-1].Date.Year && move.Date.Month < arr[loc].TransactionHistory[i-1].Date.Month) || (move.Date.Year == arr[loc].TransactionHistory[i-1].Date.Year && move.Date.Month == arr[loc].TransactionHistory[i-1].Date.Month && move.Date.Day < arr[loc].TransactionHistory[i-1].Date.Day)) {
			arr[loc].TransactionHistory[i] = arr[loc].TransactionHistory[i-1]
			i--
		}
		arr[loc].TransactionHistory[i] = move
	}
}

func is_valid_date(year, month, day int) bool {
	if year < 1 || month > 12 || month < 1 || day < 1 || day > 31 {
		return false
	}
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		if day < 1 || day > 31 {
			return false
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		if day < 1 || day > 30 {
			return false
		}
	} else if month == 2 {
		if is_leap_year(year) {
			if day < 1 || day > 29 {
				return false
			}
		} else {
			if day < 1 || day > 28 {
				return false
			}
		}
	}
	return true
}

func login_page(T *tabUser, n *int) {
	/* */
	var choice int
	read_data(T, n)
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "      SELAMAT DATANG DI APLIKASI KEUANGAN      ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "LOGIN <<                                       ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "Apa yang Anda ingin lakukan?                   ")
	fmt.Printf("%2s%s\n", "", "1). Login (tabungan yang sudah terdaftar)      ")
	fmt.Printf("%2s%s\n", "", "2). Buat Profil Tabungan Baru                  ")
	fmt.Printf("%2s%s\n", "", "3). Perlihatkan Semua Tabungan yang Terdaftar  ")
	fmt.Printf("%2s%s\n", "", "4). Keluar                                     ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "Made by:                                       ")
	fmt.Printf("%2s%s\n", "", "(1) Abdullah - 103012330146                    ")
	fmt.Printf("%2s%s\n", "", "(2) Muammar Irza Choirruzzat - 103012330146    ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s", "", "Masukkan pilihan: ")
	fmt.Scan(&choice)
	for choice > 4 || choice < 1 {
		fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
		fmt.Scan(&choice)
	}

	if choice == 1 {
		clear_screen()
		validation(T, n)
	} else if choice == 2 {
		clear_screen()
		add_new_profile(T, n)
	} else if choice == 3 {
		clear_screen()
		show_list(T, n)
	} else if choice == 4 {
		return
	}
}

func show_list(tab *tabUser, n *int) {
	/* */
	var choice int
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "LOGIN << Daftar Tabungan                       ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "Semua profil tabungan yang terdaftar           ")
	if *n == 0 {
		fmt.Printf("%2s%s\n", "", "Belum Ada Akun yang Terdaftar!             ")
	} else {
		for i := 0; i < *n; i++ {
			fmt.Printf("  %v) %s - Rp%.2f\n", i+1, tab[i].Name, tab[i].Balance)
		}
	}
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "1) Urutkan Berdasarkan Saldo                   ")
	fmt.Printf("%2s%s\n", "", "2) Urutkan Berdasarkan Banyak Aktivitas        ")
	fmt.Printf("%2s%s\n", "", "3) Kembali                                     ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s", "", "→ ")
	fmt.Scan(&choice)
	for choice > 3 || choice < 1 {
		fmt.Printf("%2s%s", "", "Masukkan angka yang tertera → ")
		fmt.Scan(&choice)
	}
	switch choice {
	case 1:
		fmt.Printf("%2s%s", "", "(1) Terurut menaik (2) Terurut menurun: ")
		fmt.Scan(&choice)
		for choice > 2 || choice < 1 {
			fmt.Printf("%2s%s", "", "(1) Terurut menaik (2) Terurut menurun: ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			selection_sort_balance_ascending(tab, n) // ascending
		case 2:
			selection_sort_balance_descending(tab, n) // descending
		}
		clear_screen()
		show_list(tab, n)
	case 2:
		fmt.Printf("%2s%s", "", "(1) Terurut menaik (2) Terurut menurun: ")
		fmt.Scan(&choice)
		for choice > 2 || choice < 1 {
			fmt.Printf("%2s%s", "", "(1) Terurut menaik (2) Terurut menurun: ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			insertion_sort_ts_ascending(tab, n)
		case 2:
			insertion_sort_ts_descending(tab, n)
		}
		clear_screen()
		show_list(tab, n)
	case 3:
		clear_screen()
		login_page(tab, n)
	}

}

func add_to_history(tab *tabUser, n *int, loc int, transaction_type string) {
	/* */
	var n_history, choice int
	var amount float64
	var dd Date
	n_history = tab[loc].TotalTransaction
	tab[loc].TransactionHistory[n_history].Status = transaction_type == "income"
	fmt.Printf("%2s%s%s%s", "", "Masukkan jumlah ", transaction_type, ": ")
	fmt.Scan(&amount)
	for amount > tab[loc].Balance && transaction_type == "outcome" {
		fmt.Printf("%2s%s\n", "", "Maaf Saldo Anda tidak Mencukupi!")
		fmt.Printf("%2s%s\n", "", "(1) Kembali (2) Masukkan nominal kembali")
		fmt.Printf("%2s%s", "", "→ ")
		fmt.Scan(&choice)
		if choice == 1 {
			clear_screen()
			user_homepage(tab, n, loc)
		} else if choice == 2 {
			fmt.Printf("%2s%s%s%s", "", "Masukkan nominal ", transaction_type, ":")
			fmt.Scan(&amount)
		}
	}
	tab[loc].TransactionHistory[n_history].Amount = amount
	fmt.Printf("%2s%s", "", "Masukkan Tanggal (DD MM YY): ")
	fmt.Scan(&dd.Day, &dd.Month, &dd.Year)
	for !(is_valid_date(dd.Year, dd.Month, dd.Day)) {
		fmt.Printf("%2s%s", "", "Tanggal yang dimasukkan tidak valid! Masukkan Tanggal (DD MM YY): ")
		fmt.Scan(&dd.Day, &dd.Month, &dd.Year)
	}
	tab[loc].TransactionHistory[n_history].Date.Year = dd.Year
	tab[loc].TransactionHistory[n_history].Date.Month = dd.Month
	tab[loc].TransactionHistory[n_history].Date.Day = dd.Day
	if transaction_type == "income" {
		fmt.Printf("%2s%s%s\n", "", "Masukkan tipe ", transaction_type)
		fmt.Printf("%2s%s\n", "", "(1) Gaji (2) Pendapatan Lain-Lain")
		fmt.Printf("%2s%s", "", "→ ")
		fmt.Scan(&tab[loc].TransactionHistory[n_history].Category)
		for tab[loc].TransactionHistory[n_history].Category > 2 && tab[loc].TransactionHistory[n_history].Category < 1 {
			fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
			fmt.Scan(&tab[loc].TransactionHistory[n_history].Category)
		}
	} else {
		fmt.Printf("%2s%s%s\n", "", "Masukkan tipe ", transaction_type)
		fmt.Printf("%2s%s\n", "", "(1) Belanja Harian (2) Transportasi (3) Tagihan")
		fmt.Printf("%2s%s", "", "→ ")
		fmt.Scan(&tab[loc].TransactionHistory[n_history].Category)
		for tab[loc].TransactionHistory[n_history].Category > 3 && tab[loc].TransactionHistory[n_history].Category < 1 {
			fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
			fmt.Scan(&tab[loc].TransactionHistory[n_history].Category)
		}
	}
	if tab[loc].TransactionHistory[n_history].Status {
		tab[loc].Balance += tab[loc].TransactionHistory[n_history].Amount
		tab[loc].Income += tab[loc].TransactionHistory[n_history].Amount
	} else {
		tab[loc].Balance -= tab[loc].TransactionHistory[n_history].Amount
		tab[loc].Expense += tab[loc].TransactionHistory[n_history].Amount
	}
	tab[loc].TotalTransaction++
	fmt.Printf("%2s%s\n", "", "Transaksi anda berhasil diproses!")
	fmt.Printf("%2s%s", "", "Tekan (0) untuk kembali (1) Tambah transaksi: ")
	write_data(tab, n)
	fmt.Scan(&choice)
	for choice > 1 || choice < 0 {
		fmt.Printf("%2s%s", "", "Tekan (0) untuk kembali (1) Tambah transaksi: ")
		fmt.Scan(&choice)
	}
	switch choice {
	case 0:
		clear_screen()
		user_homepage(tab, n, loc)
	case 1:
		clear_screen()
		add_to_history(tab, n, loc, transaction_type)
	}
}

func add_transaction(tab *tabUser, n *int, loc int) {
	/* */
	var choice int
	fmt.Printf("%2s%s\n", "", "LOGIN << DASHBOARD << TABUNGAN")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s%s%s\n", "", "Halo! ", "", tab[loc].Name)
	fmt.Printf("%2s%s\n", "", "Apa yang Anda ingin Lakukan?                   ")
	fmt.Printf("%2s%s\n", "", "1). Tambah Pemasukan                           ")
	fmt.Printf("%2s%s\n", "", "2). Tambah Pengeluaran                         ")
	fmt.Printf("%2s%s\n", "", "3). Kembali                                    ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s", "", "→ ")
	fmt.Scan(&choice)
	for choice > 3 || choice < 1 {
		fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
		fmt.Scan(&choice)
	}
	if choice == 1 {
		clear_screen()
		add_to_history(tab, n, loc, "income")
	} else if choice == 2 {
		clear_screen()
		add_to_history(tab, n, loc, "outcome")
	} else if choice == 3 {
		clear_screen()
		user_homepage(tab, n, loc)
	}
}

func user_homepage(tab *tabUser, n *int, loc int) {
	/* */
	var choice int
	var status string
	tipePengeluaran := [3]string{"Belanja Harian", "Tagihan", "Transportasi"}
	tipePemasukan := [2]string{"Gaji", "Pendapatan Lain-Lain"}
	fmt.Printf("%2s%s\n", "", "LOGIN << DASHBOARD                             ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s%s%s\n", "", "Selamat Datang! ", "", tab[loc].Name)
	fmt.Println()
	fmt.Printf("%2s%s\n", "", "Berikut adalah laporan keuangan Anda           ")
	fmt.Printf("%2s%s%s%.2f\n", "", "Total Pemasukan: ", "Rp", tab[loc].Income)
	fmt.Printf("%2s%s%s%.2f\n", "", "Total Pengeluaran: ", "Rp", tab[loc].Expense)
	fmt.Printf("%2s%s%s%.2f\n", "", "Total Saldo: ", "Rp", tab[loc].Balance)
	fmt.Printf("%2s%s\n", "", "===============================================")
	if tab[loc].TotalTransaction == 0 {
		fmt.Printf("%2s%s\n", "", "Belum Ada Aktivitas!                       ")
	} else {
		fmt.Printf("%2s%s\n", "", "Daftar Aktivitas:                          ")
		for i := 0; i < tab[loc].TotalTransaction; i++ {
			if tab[loc].TransactionHistory[i].Status {
				status = "Pemasukan"
				fmt.Printf("%2s", "")
				fmt.Printf("(%v/%v/%v) %s sebesar Rp%.f #%s\n", tab[loc].TransactionHistory[i].Date.Day, tab[loc].TransactionHistory[i].Date.Month, tab[loc].TransactionHistory[i].Date.Year, status, tab[loc].TransactionHistory[i].Amount, tipePemasukan[tab[loc].TransactionHistory[i].Category-1])
			} else {
				status = "Pengeluaran"
				fmt.Printf("%2s", "")
				fmt.Printf("(%v/%v/%v) %s sebesar Rp%.f #%s\n", tab[loc].TransactionHistory[i].Date.Day, tab[loc].TransactionHistory[i].Date.Month, tab[loc].TransactionHistory[i].Date.Year, status, tab[loc].TransactionHistory[i].Amount, tipePengeluaran[tab[loc].TransactionHistory[i].Category-1])
			}
		}
	}
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s\n", "", "Aksi:                                          ")
	fmt.Printf("%2s%s\n", "", "1). Tambah Catatan Keuangan                    ")
	fmt.Printf("%2s%s\n", "", "2). Edit Profil Keuangan                       ")
	fmt.Printf("%2s%s\n", "", "3). Urutkan catatan berdasarkan tanggal        ")
	fmt.Printf("%2s%s\n", "", "4). Keluar Akun                                ")
	fmt.Printf("%2s%s\n", "", "5). Hapus Akun                                 ")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s", "", "→ ")
	fmt.Scan(&choice)
	for choice > 4 || choice < 1 {
		fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
		fmt.Scan(&choice)
	}
	if choice == 1 {
		clear_screen()
		add_transaction(tab, n, loc)
	} else if choice == 2 {
		clear_screen()
		edit_user_profile(tab, n, loc)
	} else if choice == 3 {
		fmt.Printf("%2s%s", "", "→ (1) Ascending (2) Descending: ")
		fmt.Scan(&choice)
		for choice > 2 || choice < 1 {
			fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 1:
			sort_by_date_newest(tab, loc)
		case 2:
			sort_by_date_oldest(tab, loc)
		}
		clear_screen()
		user_homepage(tab, n, loc)
	} else if choice == 4 {
		clear_screen()
		login_page(tab, n)
	} else if choice == 5 {
		delete_data(tab, n, loc)
	}
}

func edit_user_profile(tab *tabUser, n *int, loc int) {
	/* */
	var choice int
	var temp string
	fmt.Printf("%2s%s\n", "", "LOGIN << DASHBOARD << EDIT PROFIL")
	fmt.Printf("%2s%s\n", "", "===============================================")
	fmt.Printf("%2s%s%s\n", "", "Nama Akun: ", tab[loc].Name)
	fmt.Printf("%2s%s\n", "", "(1) Edit Nama Akun")
	fmt.Printf("%2s%s\n", "", "(2) Ganti Password")
	fmt.Printf("%2s%s\n", "", "(3) Kembali")
	fmt.Printf("%2s%s\n", "", "===============================================")
	for choice > 3 || choice < 1 {
		fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
		fmt.Scan(&choice)
	}
	switch choice {
	case 1:
		fmt.Printf("%2s%s", "", "Masukkan nama akun yang baru: ")
		fmt.Scan(&temp)
		for sequential_search(tab, n, temp) != -1 {
			fmt.Printf("%2s%s%s", "", temp, " sudah dipakai! Mohon gunakan nama yang lain: ")
			fmt.Scan(&temp)
		}
		tab[loc].Name = temp
		fmt.Printf("%2s%s\n", "", "Nama akun berhasil diubah!")
		fmt.Printf("%2s%s", "", "Ketik (0) untuk kembali → ")
		fmt.Scan(&choice)
		for choice != 0 {
			fmt.Printf("%2s%s", "", "Ketik (0) untuk kembali → ")
			fmt.Scan(&choice)
		}
		write_data(tab, n)
		clear_screen()
		edit_user_profile(tab, n, loc)
	case 2:
		fmt.Printf("%2s%s", "", "Masukkan password yang baru: ")
		fmt.Scan(&tab[loc].Password)
		fmt.Printf("%2s%s\n", "", "Password akun berhasil diubah!")
		fmt.Printf("%2s%s", "", "Ketik (0) untuk kembali → ")
		fmt.Scan(&choice)
		for choice != 0 {
			fmt.Printf("%2s%s", "", "Ketik (0) untuk kembali → ")
			fmt.Scan(&choice)
		}
		write_data(tab, n)
		clear_screen()
		edit_user_profile(tab, n, loc)
	case 3:
		clear_screen()
		user_homepage(tab, n, loc)
	}
}

func validation(tab *tabUser, n *int) {
	/* */
	var username, password string
	var location, choice int
	fmt.Printf("%2s%s", "", "Masukkan nama profil tabungan: ")
	fmt.Scan(&username)
	location = sequential_search(tab, n, username)
	if location != -1 {
		fmt.Printf("%2s%s", "", "Masukkan password: ")
		fmt.Scan(&password)
		for password != tab[location].Password {
			fmt.Printf("%2s%s", "", "Maaf password Anda salah, mohon coba lagi:")
			fmt.Scan(&password)
		}
		clear_screen()
		user_homepage(tab, n, location)
	} else {
		fmt.Printf("%2s%s\n", "", "Profil tabungan tidak terdaftar!")
		fmt.Printf("%2s%s\n", "", "ketik (0) kembali (1) login ke akun lain")
		fmt.Printf("%2s%s", "", "→ ")
		fmt.Scan(&choice)
		for choice > 2 || choice < 0 {
			fmt.Printf("%2s%s", "", "Masukkan angka sesuai dengan yang tertera → ")
			fmt.Scan(&choice)
		}
		switch choice {
		case 0:
			clear_screen()
			login_page(tab, n)
		case 1:
			clear_screen()
			validation(tab, n)
		}
	}
}

func add_new_profile(T *tabUser, n *int) {
	/* */
	location := *n
	fmt.Printf("%2s%s", "", "Masukkan nama profil tabungan: ")
	fmt.Scan(&T[*n].Name)
	for sequential_search(T, n, T[*n].Name) != -1 {
		fmt.Printf("%2s%s%s", "", T[*n].Name, " sudah dipakai! Mohon gunakan nama yang lain: ")
		fmt.Scan(&T[*n].Name)
	}
	fmt.Printf("%2s%s", "", "Masukkan password: ")
	fmt.Scan(&T[*n].Password)
	fmt.Printf("%2s%s", "", "Masukkan jumlah saldo awal: ")
	fmt.Scan(&T[*n].Balance)
	for T[*n].Balance < 0 {
		fmt.Print("Saldo tidak bisa bernilai negatif, masukkan kembali saldo yang valid: ")
		fmt.Scan(&T[*n].Balance)
	}
	*n++
	write_data(T, n)
	fmt.Println()
	fmt.Printf("%2s%s\n", "", "Akun berhasil dibuat!")
	clear_screen()
	user_homepage(T, n, location)
}

func main() {
	var t_user tabUser
	var n int
	login_page(&t_user, &n)
}
