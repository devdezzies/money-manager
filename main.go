package main 
import (
  "fmt"
  "encoding/gob"
  //"io/ioutil"
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
  var found int 
  found = -1
  for i := 0; i < *n; i++ {
    if T[i].Name == x {
      found = i
    }
  }
  return found
} 

func login_page(T *tabUser, n *int) {
  var choice int
  read_data(T, n)
  fmt.Printf(`
    1 (login into an existing account); 2 (create a new account); 
    3 (show all valid account)
    `)
  fmt.Print("Choose an option: ")
  fmt.Scan(&choice)
  for (choice > 3 || choice < 1) {
    fmt.Print("Please choose a valid number: ")
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
  for i := 0; i < *n; i++ {
    fmt.Println(tab[i].Name, tab[i].Balance, tab[i].TotalTransaction)
  }
  fmt.Printf(`Tekan tombol apapun untuk kembali: `)
  fmt.Scan(&choice)
}

func add_to_history(tab *tabUser, n *int, loc int, transaction_type string) {
  var n_history, choice int
  var amount float64
  n_history = tab[loc].TotalTransaction
  tab[loc].TransactionHistory[n_history].Status = transaction_type == "income"
  fmt.Printf(`
    Masukkan jumlah %s: 
    `, transaction_type)

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
  } else {
    tab[loc].Balance -= tab[loc].TransactionHistory[n_history].Amount
  }
  tab[loc].TotalTransaction++
  fmt.Printf(`
    Transaksi anda berhasil diproses!

    TEKAN 0 UNTUK KEMBALI
    `)
  write_data(tab, n)
  fmt.Scan(&choice)
  if choice == 0 {
    user_homepage(tab, n, loc)
  }
}

func show_history(tab *tabUser, n *int, loc int) {
  for i := 0; i < tab[loc].TotalTransaction; i++ {
    fmt.Println(tab[loc].TransactionHistory[i].Amount, tab[loc].TransactionHistory[i].Category, tab[loc].TransactionHistory[i].Status)
  }
}

func add_transaction(tab *tabUser, n *int, loc int) {
  var choice int 
  fmt.Printf(`
    *************************************
    Halo, %s! 
    Apa yang Anda ingin lakukan?

    1) Tambah Pemasukan 
    2) Tambah Pengeluaran 
    3) Kembali
    `, tab[loc].Name)
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
  fmt.Println("Good Morning", tab[loc].Name)
  fmt.Println("You have", tab[loc].Balance, "in your balance")
  fmt.Println(`
    Type 1) to delete this account; 2) to add transaction; 3) show history
    `) 

  fmt.Scan(&choice)
  if (choice == 1) {
    delete_data(tab, n, loc)
  } else if (choice == 2) {
    add_transaction(tab, n, loc)
  } else if (choice == 3) {
    show_history(tab, n, loc)
  } else {
    return
  }

}

func validation(tab *tabUser, n *int) {
  var username, password string
  var location int
  fmt.Print("Enter the username: ")
  fmt.Scan(&username)
  location = sequential_search(tab, n, username)
  if (location != -1) {
    fmt.Print("Please enter your password: ")
    fmt.Scan(&password)
    for (password != tab[location].Password) {
      fmt.Print("Your password is incorrect, please enter the correct pass: ")
      fmt.Scan(&password)
    }
    user_homepage(tab, n, location)
  } else {
    fmt.Println("the user doesn't exist!")
  }
}

func add_new_profile(T *tabUser, n *int) {
  var location int 
  location = *n
  fmt.Print("Enter the username: ")
  fmt.Scan(&T[*n].Name)
  for (sequential_search(T, n, T[*n].Name) != -1) {
    fmt.Print(T[*n].Name, " already exists! Please find another name: ")
    fmt.Scan(&T[*n].Name)
  }
  fmt.Print("Enter a password: ")
  fmt.Scan(&T[*n].Password)
  fmt.Print("Enter the first balance: ")
  fmt.Scan(&T[*n].Balance)
  for (T[*n].Balance < 0) {
    fmt.Print("The balance cannot be negative, please enter a valid balance: ")
    fmt.Scan(&T[*n].Balance)
  }
  *n++
  write_data(T, n)
  fmt.Println("A new user has been created!")
  fmt.Println("Tekan enter untuk melanjutkan")
  fmt.Scan()
  user_homepage(T, n, location)
}

func main() {
  var t_user tabUser 
  var n int 
  login_page(&t_user, &n)
}
