package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	controller "test/carorders/Controller"
)

func main() {
	router := mux.NewRouter()

	// router for Cars
	router.HandleFunc("/cars", controller.AllCar).Methods("GET")
	router.HandleFunc("/insert-car", controller.InsertCar).Methods("POST")
	router.HandleFunc("/update-car", controller.UpdateCar).Methods("PUT")
	router.HandleFunc("/delete-car", controller.DeleteCar).Methods("DELETE")

	// router for Users
	router.HandleFunc("/users", controller.AllUser).Methods("GET")
	router.HandleFunc("/insert-user", controller.InsertUser).Methods("POST")
	router.HandleFunc("/update-user", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/delete-user", controller.DeleteUser).Methods("DELETE")

	// router for Admin
	router.HandleFunc("/admin", controller.AllAdmin).Methods("GET")
	router.HandleFunc("/insert-admin", controller.InsertAdmin).Methods("POST")
	router.HandleFunc("/update-admin", controller.UpdateAdmin).Methods("PUT")
	router.HandleFunc("/delete-admin", controller.DeleteAdmin).Methods("DELETE")

	// router for Orders
	router.HandleFunc("/orders", controller.AllOrder).Methods("GET")
	router.HandleFunc("/insert-order", controller.InsertOrder).Methods("POST")
	router.HandleFunc("/update-order", controller.UpdateOrder).Methods("PUT")
	router.HandleFunc("/delete-order", controller.DeleteOrder).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
