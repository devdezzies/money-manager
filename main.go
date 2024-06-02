package main 
import (
  "fmt"
  "encoding/gob"
  "os"
)

const mxN int = 1000

type Transaction struct {
  Date, Category string
  Amount float64
  Status bool // true for income, false for outcome
}

type User struct {
  Name, Password string 
  Balance float64
  Income, Expense float64
  TotalTransaction int
  TransactionHistory [mxN]Transaction
}

type tabUser [mxN]User

func write_data(arr *tabUser, n *int) {
  file, err := os.Create("user.dat")
  if err != nil {
    fmt.Println("Error creating file:", err)
    return 
  }
  defer file.Close()

  encode := gob.NewEncoder(file)
  err = encode.Encode(*arr) 
  if err != nil {
    fmt.Println("Error encoding array:", err)
    return
  }

  err = encode.Encode(*n)
  if err != nil {
    fmt.Println("Error encoding n:", err)
    return
  }
    //fmt.Println("array and n have been endocoded and written to file")
  }

func read_data(arr *tabUser, n *int) {
  file, err := os.Open("user.dat")
  if err != nil {
    fmt.Println("Error opening file:", err)
  }
  defer file.Close()

  //fileInfo, err := file.Stat()

  //if fileInfo.Size() == 0 {
  //  fmt.Println("File is empty, starting with an empty user list.")
  //  return
  //}

  decoder := gob.NewDecoder(file)
  err = decoder.Decode(arr)
  if err != nil {
    fmt.Println("Error decoding array:", err)
    return 
  }

  err = decoder.Decode(n)
  if err != nil {
    fmt.Println("Error decoding n:", err)
    return
  }

  //fmt.Println("Decoded array from file")
}

// WRITE-READ FUNCTIONALITY

func delete_data(T *tabUser, n *int, xLoc int) {
  for i := xLoc; i < *n-1; i++ {
    T[i] = T[i+1]
  }
  *n--
  write_data(T, n)
  fmt.Println("This account has been deleted!")
}

func sequential_search(T *tabUser, n *int, x string) int {
  var found, i int 
  i = 0
  found = -1
  for (i < *n && found == -1) {
    if T[i].Name == x {
      found = i
    }
    i++
  }
  return found
} 

func login_page(T *tabUser, n *int) {
  var choice int
  read_data(T, n)
  fmt.Printf(`
  ***********************************************
  *                                             *
  *           WELCOME TO MONEYGUARD             *
  *                                             *
  ***********************************************
  *                                             *
  *         Track your expenses and income      *
  *          for better financial health        *
  *                                             *
  ***********************************************
  *                 Main Menu                   *
  ***********************************************
  *                                             *
  *  1. Login                                   *
  *  2. Create a New Profile                    *
  *  3. Show All My Profiles                    *
  *                                             *
  ***********************************************
  *   Please enter your choice (1-3):           *
  ***********************************************
    `)
  fmt.Println()
  fmt.Printf("%2s%s", "", "Choose an option: ")
  fmt.Scan(&choice)
  for (choice > 3 || choice < 1) {
    fmt.Printf("%2s%s", "", "Please choose a valid number: ")
    fmt.Scan(&choice)
  }

  if choice == 1 {
    validation(T, n)
  } else if choice == 2 {
    add_new_profile(T, n)
  } else if choice == 3 {
    show_list(T, n)
  }
}

func show_list(tab *tabUser, n *int) {
  var choice int
  fmt.Println(`
  ***********************************************
  *              All My Profiles                *
  ***********************************************
  `)
  for i := 0; i < *n; i++ {
    fmt.Printf("  %v) %s %f %v\n", i+1, tab[i].Name, tab[i].Balance, tab[i].TotalTransaction)
  }
  fmt.Printf(`
  ***********************************************
  *    Type anything to return to Main Menu     *
  ***********************************************
    `)
  fmt.Printf("Tekan 0 untuk kembali: ")
  fmt.Scan(&choice)
  for (choice != 0) {
    fmt.Printf("%4s%s", "","Tekan 0 untuk kembali: ")
    fmt.Scan(&choice)
  }
  login_page(tab, n)
}

func add_to_history(tab *tabUser, n *int, loc int, transaction_type string) {
  var n_history, choice int
  var amount float64
  n_history = tab[loc].TotalTransaction
  tab[loc].TransactionHistory[n_history].Status = transaction_type == "income"
  fmt.Printf("%4s", "")
  fmt.Printf("Masukkan jumlah %s: ", transaction_type)
  fmt.Scan(&amount)
  for (amount > tab[loc].Balance && transaction_type == "outcome") {
    fmt.Println(`
      Maaf saldo Anda tidak mencukupi
      
      1) Kembali
      2) Masukkan nominal kembali
      `)
    fmt.Scan(&choice) 
    if choice == 1 {
      user_homepage(tab, n, loc)
    } else if choice == 2 {
      fmt.Printf(`
        Masukkan nominal %s
        `, transaction_type)
      fmt.Scan(&amount)
    }
  }
  tab[loc].TransactionHistory[n_history].Amount = amount
  fmt.Printf(`
    Masukkan tipe %s: 
    
    1) Pembayaran 
    2) lain-lain
    `, transaction_type)
  fmt.Scan(&tab[loc].TransactionHistory[n_history].Category) 
  if tab[loc].TransactionHistory[n_history].Status {
    tab[loc].Balance += tab[loc].TransactionHistory[n_history].Amount
    tab[loc].Income += tab[loc].TransactionHistory[n_history].Amount
  } else {
    tab[loc].Balance -= tab[loc].TransactionHistory[n_history].Amount
    tab[loc].Expense += tab[loc].TransactionHistory[n_history].Amount
  }
  tab[loc].TotalTransaction++
  fmt.Printf("%2s%s\n", "", "Transaksi anda berhasil diproses!")
  fmt.Printf("%2s%s", "", "Tekan 0 untuk kembali: ") 
  write_data(tab, n)
  fmt.Scan(&choice)
  if choice == 0 {
    user_homepage(tab, n, loc)
  }
}

func show_history(tab *tabUser, n *int, loc int) {
  var status string
  var input int
  for i := 0; i < tab[loc].TotalTransaction; i++ {
    if (tab[loc].TransactionHistory[i].Status) {
      status = "Pemasukan"
    } else {
      status = "Pengeluaran"
    }
    fmt.Printf("%2s", "")
    fmt.Printf("%v). %s sebesar Rp%.f kategori-%s\n", i+1, status, tab[loc].TransactionHistory[i].Amount, tab[loc].TransactionHistory[i].Category)
  }
  fmt.Printf("%2s%s", "", "Tekan 0 untuk kembali: ")
  fmt.Scan(&input)
}

func add_transaction(tab *tabUser, n *int, loc int) {
  var choice int 
  fmt.Printf(`
  ***********************************************
                                               
                    Halo, %s!                  
           Apa yang Anda ingin lakukan?        
                                               
  ***********************************************
  *                                             *
  *         1) Tambah Pemasukan                 *
  *         2) Tambah Pengeluaran               *
  *         3) Kembali                          *
  *                                             *
  ***********************************************
    `, tab[loc].Name)
  fmt.Printf("Masukkan pilihan: ")
  fmt.Scan(&choice) 
  if choice == 1 {
    add_to_history(tab, n, loc, "income")
  } else if choice == 2 {
    add_to_history(tab, n, loc, "outcome")
  } else if choice == 3 {
    user_homepage(tab, n, loc)
  }
}

func user_homepage(tab *tabUser, n *int, loc int) {
  var choice int
  fmt.Printf(`
  ***********************************************
                                               
              WELCOME BACK, %s!                 
                                               
  ***********************************************
                                               
         Your Financial Overview at a Glance   
                                               
  ***********************************************
                                               
     Total Income:        Rp%.2f                
     Total Expenses:      Rp%.2f              
     Current Balance:     Rp%.2f               
                                               
  ***********************************************
  *                 Actions                     *
  ***********************************************
  *                                             *
  *  1. Add Transaction                         *
  *  2. View Detailed Summary                   *
  *  3. Edit Profile                            *
  *  4. Log Out                                 *
  *  5. Delete Account                          *
  *                                             *
  ***********************************************           
    `, tab[loc].Name, tab[loc].Income, tab[loc].Expense, tab[loc].Balance)
  fmt.Printf("Please enter your choice (1-5): ")
  fmt.Scan(&choice)
  if (choice == 1) {
    add_transaction(tab, n, loc)
  } else if (choice == 2) {
    show_history(tab, n, loc)
  } else if (choice == 3) {
    return
  } else if (choice == 4) {
    login_page(tab, n)
  } else if (choice == 5) { 
    delete_data(tab, n, loc)
  }
}

func validation(tab *tabUser, n *int) {
  var username, password string
  var location int
  fmt.Printf("%2s%s", "", "Enter the username: ")
  fmt.Scan(&username)
  location = sequential_search(tab, n, username)
  if (location != -1) {
    fmt.Printf("%2s%s", "", "Please enter your password: ")
    fmt.Scan(&password)
    for (password != tab[location].Password) {
      fmt.Printf("%2s%s", "", "Your password is incorrect, please enter the correct pass:")
      fmt.Scan(&password)
    }
    user_homepage(tab, n, location)
  } else {
    fmt.Printf("%2s%s", "", "the user doesn't exist!")
  }
}

func add_new_profile(T *tabUser, n *int) {
  var location int 
  location = *n
  fmt.Printf("%2s%s", "", "Enter the username: ")
  fmt.Scan(&T[*n].Name)
  for (sequential_search(T, n, T[*n].Name) != -1) {
    fmt.Printf("%2s%s%s", "", T[*n].Name, " already exists! Please find another name: ")
    fmt.Scan(&T[*n].Name)
  }
  fmt.Printf("%2s%s", "", "Enter a password: ")
  fmt.Scan(&T[*n].Password)
  fmt.Printf("%2s%s", "", "Enter the first balance: ")
  fmt.Scan(&T[*n].Balance)
  for (T[*n].Balance < 0) {
    fmt.Print("The balance cannot be negative, please enter a valid balance: ")
    fmt.Scan(&T[*n].Balance)
  }
  *n++
  write_data(T, n)
  fmt.Println()
  fmt.Printf("%2s%s\n", "", "A new user has been created!")
  fmt.Printf("%2s%s", "", "Tekan enter untuk melanjutkan")
  fmt.Scan()
  user_homepage(T, n, location)
}

func main() {
  var t_user tabUser 
  var n int 
  login_page(&t_user, &n)
}
