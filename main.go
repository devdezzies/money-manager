package main 
import (
  "fmt"
  "encoding/gob"
  //"io/ioutil"
  "os"
)

const mxN int = 100

type Transaction struct {
  Date, Category string
  Amount float64 
}

type User struct {
  Name, Password string 
  Balance float64 
  Transaction_history [mxN]Transaction
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
    fmt.Println("array and n have been endocoded and written to file")
  }

  func read_data(arr *tabUser, n *int) {
    file, err := os.Open("user.dat")
    if err != nil {
      fmt.Println("Error opening file:", err)
    }
  defer file.Close()

  fileInfo, err := file.Stat()

  if fileInfo.Size() == 0 {
    fmt.Println("File is empty, starting with an empty user list.")
    return
  }

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

  fmt.Println("Decoded array from file")
}

// WRITE-READ FUNCTIONALITY

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
  fmt.Println("1 (login into an existing account); 2 (create a new account) 3 (show all valid account)")
  fmt.Print("Choose an option: ")
  fmt.Scan(&choice)
  for (choice > 3 && choice < 1) {
    fmt.Print("Please choose a valid number: ")
    fmt.Scan(&choice)
  }

  switch choice {
    case 1: 
      validation(T, n)
    case 2: 
      add_new_profile(T, n)
    case 3: 
      //account_list()
      show_list(T, n)
    default: 
      fmt.Println("error")
  }
}

func show_list(tab *tabUser, n *int) {
  for i := 0; i < *n; i++ {
    fmt.Println(tab[i].Name, tab[i].Balance)
  }
}

func user_homepage(tab *tabUser, loc int) {
  fmt.Println("Good Morning", tab[loc].Name)
  fmt.Println("You have", tab[loc].Balance, "balance")
}

func validation(tab *tabUser, n *int) {
  var username, password string
  fmt.Print("Enter the username: ")
  fmt.Scan(&username)
  if (sequential_search(tab, n, username) != -1) {
    fmt.Print("Please enter your password: ")
    fmt.Scan(&password)
    for (password != tab[sequential_search(tab, n, username)].Password) {
      fmt.Print("Your password is incorrect, please enter the correct one: ")
      fmt.Scan(&password)
    }
    user_homepage(tab, sequential_search(tab, n, username))
  } else {
    fmt.Println("the user doesn't exist!")
  }
}

func add_new_profile(T *tabUser, n *int) {
  var location int 
  location = *n
  fmt.Print("Enter the username: ")
  fmt.Scan(&T[*n].Name)
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
  user_homepage(T, location)
}

func main() {
  var t_user tabUser 
  var n int 
  login_page(&t_user, &n)
}
