package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"test/carorders/config"
	"test/carorders/model"

	"golang.org/x/crypto/bcrypt"
)

func AllCar(w http.ResponseWriter, r *http.Request) {
	var cars model.Cars
	var response model.ResponseCars
	var arrCars []model.Cars

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM cars")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&cars.CarId, &cars.Name, &cars.CarType, &cars.Rating, &cars.Fuel, &cars.Image, &cars.HourRate, &cars.DayRate, &cars.MonthRate)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrCars = append(arrCars, cars)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrCars

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func InsertCar(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseCars

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	name := r.FormValue("name")
	carType := r.FormValue("type")
	rating := r.FormValue("rating")
	fuel := r.FormValue("fuel")
	image := r.FormValue("image")
	hourRate := r.FormValue("hourRate")
	dayRate := r.FormValue("dayRate")
	monthRate := r.FormValue("monthRate")

	_, err = db.Exec("INSERT INTO cars(name, carType, rating, fuel, image, hourRate, dayRate, monthRate) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", name, carType, rating, fuel, image, hourRate, dayRate, monthRate)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data success"
	fmt.Print("Insert data success")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseCars

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	carType := r.FormValue("type")
	rating := r.FormValue("rating")
	fuel := r.FormValue("fuel")
	image := r.FormValue("image")
	hourRate := r.FormValue("hourRate")
	dayRate := r.FormValue("dayRate")
	monthRate := r.FormValue("monthRate")

	if name != "" {
		_, err = db.Exec("UPDATE cars SET name=? WHERE carId=?", name, id)
	}
	if carType != "" {
		_, err = db.Exec("UPDATE cars SET carType=? WHERE carId=?", carType, id)
	}
	if rating != "" {
		_, err = db.Exec("UPDATE cars SET rating=? WHERE carId=?", rating, id)
	}
	if fuel != "" {
		_, err = db.Exec("UPDATE cars SET fuel=? WHERE carId=?", fuel, id)
	}
	if image != "" {
		_, err = db.Exec("UPDATE cars SET image=? WHERE carId=?", image, id)
	}
	if hourRate != "" {
		_, err = db.Exec("UPDATE cars SET hourRate=? WHERE carId=?", hourRate, id)
	}
	if dayRate != "" {
		_, err = db.Exec("UPDATE cars SET dayRate=? WHERE carId=?", dayRate, id)
	}
	if monthRate != "" {
		_, err = db.Exec("UPDATE cars SET monthRate=? WHERE carId=?", monthRate, id)
	}

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Update data success"
	fmt.Print("Update data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseCars

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM cars WHERE carId=?", id)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Delete data success"
	fmt.Print("Delete data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AllUser(w http.ResponseWriter, r *http.Request) {
	var users model.Users
	var response model.ResponseUsers
	var arrUsers []model.Users

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&users.UserId, &users.Email, &users.PhoneNumber, &users.City, &users.Zip, &users.Message, &users.Password, &users.Username, &users.Address)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrUsers = append(arrUsers, users)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrUsers

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseUsers

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	email := r.FormValue("email")
	phoneNumber := r.FormValue("phoneNumber")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	message := r.FormValue("message")

	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
	// password := r.FormValue("password")

	username := r.FormValue("username")
	address := r.FormValue("address")

	_, err = db.Exec("INSERT INTO users(email, phoneNumber, city, zip, message, password, username, address) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", email, phoneNumber, city, zip, message, password, username, address)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data success"
	fmt.Print("Insert data success")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseUsers

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	email := r.FormValue("email")
	phoneNumber := r.FormValue("phoneNumber")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	message := r.FormValue("message")
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
	username := r.FormValue("username")
	address := r.FormValue("address")

	if email != "" {
		_, err = db.Exec("UPDATE users SET email=? WHERE userId=?", email, id)
	}
	if phoneNumber != "" {
		_, err = db.Exec("UPDATE users SET phoneNumber=? WHERE userId=?", phoneNumber, id)
	}
	if city != "" {
		_, err = db.Exec("UPDATE users SET city=? WHERE userId=?", city, id)
	}
	if zip != "" {
		_, err = db.Exec("UPDATE users SET zip=? WHERE userId=?", zip, id)
	}
	if message != "" {
		_, err = db.Exec("UPDATE users SET message=? WHERE userId=?", message, id)
	}
	if password != nil {
		_, err = db.Exec("UPDATE users SET password=? WHERE userId=?", password, id)
	}
	if username != "" {
		_, err = db.Exec("UPDATE users SET username=? WHERE userId=?", username, id)
	}
	if address != "" {
		_, err = db.Exec("UPDATE users SET address=? WHERE userId=?", address, id)
	}

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Update data success"
	fmt.Print("Update data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseUsers

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM users WHERE userId=?", id)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Delete data success"
	fmt.Print("Delete data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AllAdmin(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	var response model.ResponseAdmin
	var arrAdmin []model.Admin

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM admin")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&admin.Id, &admin.Email, &admin.Password)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrAdmin = append(arrAdmin, admin)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrAdmin

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func InsertAdmin(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseAdmin

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	email := r.FormValue("email")
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)
	// password := r.FormValue("password")

	_, err = db.Exec("INSERT INTO admin(email, password) VALUES(?, ?)", email, password)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data success"
	fmt.Print("Insert data success")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseUsers

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	email := r.FormValue("email")
	password, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 12)

	if email != "" {
		_, err = db.Exec("UPDATE admin SET email=? WHERE id=?", email, id)
	}
	if password != nil {
		_, err = db.Exec("UPDATE admin SET password=? WHERE id=?", password, id)
	}

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Update data success"
	fmt.Print("Update data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseAdmin

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM admin WHERE id=?", id)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Delete data success"
	fmt.Print("Delete data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AllOrder(w http.ResponseWriter, r *http.Request) {
	var order model.Orders
	var response model.ResponseOrders
	var arrOrder []model.Orders

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM orders")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&order.OrderId, &order.PickUpLoc, &order.DropOffLoc, &order.PickUpDate, &order.DropOffDate, &order.PickUpTime, &order.CarId, &order.UserId, &order.AdminId)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrOrder = append(arrOrder, order)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrOrder

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func InsertOrder(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseOrders

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	pickUpLoc := r.FormValue("pickUpLoc")
	dropOffLoc := r.FormValue("dropOffLoc")
	pickUpDate := r.FormValue("pickUpDate")
	dropOffDate := r.FormValue("dropOffDate")
	pickUpTime := r.FormValue("pickUpTime")
	carId := r.FormValue("carId")
	userId := r.FormValue("userId")
	adminId := r.FormValue("adminId")

	_, err = db.Exec("INSERT INTO orders(pickUpLoc, dropOffLoc, pickUpDate, dropOffDate, pickUpTime, carId, userId, adminId) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", pickUpLoc, dropOffLoc, pickUpDate, dropOffDate, pickUpTime, carId, userId, adminId)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data success"
	fmt.Print("Insert data success")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseOrders

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	id := r.FormValue("id")
	pickUpLoc := r.FormValue("pickUpLoc")
	dropOffLoc := r.FormValue("dropOffLoc")
	pickUpDate := r.FormValue("pickUpDate")
	dropOffDate := r.FormValue("dropOffDate")
	pickUpTime := r.FormValue("pickUpTime")
	carId := r.FormValue("carId")
	userId := r.FormValue("userId")
	adminId := r.FormValue("adminId")

	if pickUpLoc != "" {
		_, err = db.Exec("UPDATE orders SET pickUpLoc=? WHERE orderId=?", pickUpLoc, id)
	}
	if dropOffLoc != "" {
		_, err = db.Exec("UPDATE orders SET dropOffLoc=? WHERE orderId=?", dropOffLoc, id)
	}
	if pickUpDate != "" {
		_, err = db.Exec("UPDATE orders SET pickUpDate=? WHERE orderId=?", pickUpDate, id)
	}
	if dropOffDate != "" {
		_, err = db.Exec("UPDATE orders SET dropOffDate=? WHERE orderId=?", dropOffDate, id)
	}
	if pickUpTime != "" {
		_, err = db.Exec("UPDATE orders SET pickUpTime=? WHERE orderId=?", pickUpTime, id)
	}
	if carId != "" {
		_, err = db.Exec("UPDATE orders SET carId=? WHERE orderId=?", carId, id)
	}
	if userId != "" {
		_, err = db.Exec("UPDATE orders SET userId=? WHERE orderId=?", userId, id)
	}
	if adminId != "" {
		_, err = db.Exec("UPDATE orders SET adminId=? WHERE orderId=?", adminId, id)
	}

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Update data success"
	fmt.Print("Update data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var response model.ResponseOrders

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("id")
	_, err = db.Exec("DELETE FROM orders WHERE orderId=?", id)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Delete data success"
	fmt.Print("Delete data success")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
