package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rysmaadit/go-template/app"
	"github.com/rysmaadit/go-template/external/gorm_client"
	"gorm.io/gorm"
)

// func CobaRun() {
// 	fmt.Println("coba jalankan fungsi dari handler")
// }

func InitMigration() (db *gorm.DB, err error) {
	db, err = gorm_client.Connect(app.Init())
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(gorm_client.Employee)
	return db, nil
}

// Create
func CreateEmployee(db *gorm.DB, name, email, adress string, age int) {
	new_emp :=
		gorm_client.Employee{Name: name, Email: email, Address: adress, Age: age}
	db.Create(&new_emp)
	fmt.Println("Data berhasil masuk !")
}

// Read
func ReadEmployee(db *gorm.DB) {
	var emp []gorm_client.Employee
	db.Find(&emp)

	for _, e := range emp {
		fmt.Println("Name \t : ", e.Name, "\n Email\t : ", e.Email, "\n Address\t : ", e.Address, "\n Age\t :", e.Age)
	}
}

// Update
func UpdateEmployee(db *gorm.DB, email string, ent string, sel int) {
	var emp gorm_client.Employee
	selection := [4]string{"name", "email", "age", "address"}
	db.Model(&emp).Where("email = ?", email).Update(selection[sel+1], ent)
	info := fmt.Sprintf("Sesuatu berhasil di update !", selection[sel-1])
	fmt.Println(info)
}

// Delete
func DeleteEmployee(db *gorm.DB, email string) {
	var emp gorm_client.Employee
	db.Model(&emp).Where("email = ?", email).Delete(&emp)
	fmt.Println("Data berhasil dihapus")
}

func getInputData() (name, email, address string, age int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama : ")
	name, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Email : ")
	email, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Alamat : ")
	address, _ = reader.ReadString('\n')
	fmt.Print("Masukkan Usia : ")
	fmt.Scanf("%d", &age)
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	address = strings.TrimSpace(address)
	return name, email, address, age
}

func MainCLI() {
	var pilihan int

	db, err := InitMigration()
	if err != nil {
		print(err.Error())
	}

	fmt.Println("Menu : ")
	fmt.Println("1. Create Data \n 2. Read Data \n 3. Update Data \n 4. Delete Data")
	fmt.Print("Masukkan pilihan : ")
	// reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d", &pilihan)
	if pilihan == 1 {
		name, email, address, age := getInputData()
		CreateEmployee(db, name, email, address, age)
	} else if pilihan == 2 {
		ReadEmployee(db)
	} else if pilihan == 3 {
		fmt.Print("Masukkan email yang ingin di update : ")
		email, _ := reader.ReadString('\n')
		email = strings.TrimSpace(email)
		fmt.Println("Data yang ingin diubah")
		fmt.Println("1. Nama \n 2. Email \n 3. Usia \n 4. Alamat")
		fmt.Scanf("%d", &pilihan)
		fmt.Print("Masukkan perubahan : ")
		input_mod, _ := reader.ReadString('\n')
		input_mod = strings.TrimSpace(input_mod)
		UpdateEmployee(db, email, input_mod, pilihan)
	} else if pilihan == 4 {
		fmt.Print("Hapus Data : ")
		delete_data, _ := reader.ReadString('\n')
		delete_data = strings.TrimSpace(delete_data)
		DeleteEmployee(db, delete_data, pilihan)
	}

}
