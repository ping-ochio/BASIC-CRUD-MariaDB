//------------------------------------------------------------------------------
// USE OF MySQL/MariaDB WITH GOLANG
//------------------------------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Create the database handle, confirm driver is present
	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	// Testing connection with database
	err = db.Ping()
	if err != nil {
		fmt.Printf("Unable connect with database: %v", err)
		return
	}
	fmt.Println("Connection established....")

	//------------------------------------------------------------------------------
	// 								MENU
	//------------------------------------------------------------------------------

	title := `  ==============================================================================================
						
	                        CRUD USING MYSQL/MARIADB IN GOLANG 

                                    __  ___ ______ _   __ __  __
                                   /  |/  // ____// | / // / / /
                                  / /|_/ // __/  /  |/ // / / / 
                                 / /  / // /___ / /|  // /_/ /  
                                /_/  /_//_____//_/ |_/ \____/

  
  ==============================================================================================`

	fmt.Println(title)
	menu := ` 
   -- SELECT AN OPTION --

 [1] -  Create user
 [2] -  Update user
 [3] -  Show table
 [4] -  View single user data
 [5] -  Delete user
 [6] -  EXIT
`

	//	cmd := exec.Command("clear")
	var choice string
	us := User{}
	for choice != "6" {
		//	cmd.Stdout = os.Stdout
		//	cmd.Run()
		fmt.Println(menu)
		fmt.Scan(&choice)
		i, _ := strconv.Atoi(choice)
		if i < 1 || i > 6 {
			fmt.Printf("Checkout, %v is not a valid option\n\n", choice)
			fmt.Println("Press the ENTER key to continue")
			fmt.Scanln()
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			switch i {

			case 1:
				var next int
				fmt.Println(" [1] - Back to menu\n [2] - Continue	")
				fmt.Scan(&next)
				if next == 1 {
					break
				} else if next == 2 {
					fmt.Println("---- ENTER THE USER DATA ----")

					fmt.Printf(" ID ?: ")
					if scanner.Scan() {
						us.User_id, _ = strconv.Atoi(scanner.Text())
					}

					fmt.Printf(" Name ?: ")
					if scanner.Scan() {
						us.Name = scanner.Text()
					}

					fmt.Printf(" Surname ?: ")
					if scanner.Scan() {
						us.Surname = scanner.Text()
					}

					fmt.Printf(" Document ?: ")
					if scanner.Scan() {
						us.Id_card, _ = strconv.Atoi(scanner.Text())

					}
					insert(us)
				}

			case 2:

				fmt.Println("---- UPDATE USER DATA ----")
				qlito := User{5, "Charles", "", 5}
				update(qlito)

			case 3:

				fmt.Println("---- CONTENT OF users TABLE ----")
				allrow()

			case 4:

				fmt.Println("---- VIEW USER DATA ----")
				onerow()

			case 5:
				fmt.Println("---- DELTE USER'S RECORD ----")

				fmt.Println("---- Enter user_id to delete ----")

				qlote := User{}
				fmt.Scanln(&qlote.User_id)
				delete(qlote)

			}

		}

	}
}

// ------------------------------------------------------------------------------
//
//	FETCHING ALL DATA FROM 'users' TABLE
//
// ------------------------------------------------------------------------------
func allrow() {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	results, err := db.Query("select * from users")
	if err != nil {
		log.Fatal("error,can't read user data:", err)
	}

	defer results.Close()
	fmt.Println("-------------------------------------------------------------------------")
	for results.Next() {

		var (
			user_id int
			name    string
			surname string
			id_card int
		)

		err = results.Scan(&user_id, &name, &surname, &id_card)
		if err != nil {
			log.Fatal("no funka", err)
		}

		fmt.Printf("UserID: %d \t| Name: %s \t| Lastname: %s \t| Id: %d\n", user_id, name, surname, id_card)
		fmt.Println("-------------------------------------------------------------------------")
	}

}

// ------------------------------------------------------------------------------
// WITH THIS COMMAND WE FETCH ONLY ONE ROW, IS MORE EFFICIENT, THIS IS IMPORTANT
// MOST OF THE TIME WE NEED TO FETCH ONLY A SINGLE RECORD
// ------------------------------------------------------------------------------
func onerow() {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()
	var namee string

	fmt.Println("Enter the name of the user whose data you want to obtain: ")
	fmt.Scanf("%s\n", &namee)

	ruser := User{}
	terri := fmt.Sprintf("Select * from users where name = '%s'", namee)

	err = db.QueryRow(terri).Scan(&ruser.User_id, &ruser.Name, &ruser.Surname, &ruser.Id_card)
	if err != nil {

		log.Fatal("The required data could not be obtained: ", err)
	}
	fmt.Println("-------------------------------------------------------------------------")
	fmt.Printf("UserID: %d | Name: %s | Lastname: %s | Id: %d\n", ruser.User_id, ruser.Name, ruser.Surname, ruser.Id_card)
	fmt.Println("-------------------------------------------------------------------------")
}

//------------------------------------------------------------------------------
//                    INSERT DATA IN 'users' TABLE
//------------------------------------------------------------------------------

func insert(us User) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (user_id, name, surname, id_card) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Unable to insert data:", err)
	}
	defer stmt.Close()

	//	for _, users := range us {
	_, err = stmt.Exec(us.User_id, us.Name, us.Surname, us.Id_card)
	if err != nil {
		log.Fatal("Unable to execute statment:", err)
	}

	//	}
}

//------------------------------------------------------------------------------
//                         UPDATE DATA TO 'users' TABLE
//------------------------------------------------------------------------------

func update(u User) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//stmt_up, err := db.Prepare("UPDATE users SET user_id = ?, name = ?, surname = ?, id_card = ?, WHERE user_id = ?")
	stmt_up, err := db.Prepare("UPDATE users SET  name = ? WHERE user_id = 5")
	if err != nil {
		log.Fatal("Unable to prepare update: ", err)
	}

	defer stmt_up.Close()
	_, err = stmt_up.Exec(u.Name)
	if err != nil {
		fmt.Printf("Unable to update data: %v", err)
	} else {
		fmt.Println("Data updated succefully")
	}
}

// ------------------------------------------------------------------------------
//
//	DELETE DATA FROM 'users' TABLE
//
// ------------------------------------------------------------------------------
func delete(us User) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer db.Close()

	stmt_del, err := db.Prepare("DELETE FROM users WHERE user_id = ?")
	if err != nil {

		fmt.Printf("Unable prepa to delete register: %s ", err)
	}
	defer stmt_del.Close()

	_, err = stmt_del.Exec(us.User_id)
	if err != nil {
		fmt.Printf("Can not delete de register: %s", err)
	}

}
