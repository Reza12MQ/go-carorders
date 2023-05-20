package model

type Cars struct {
	CarId     string `form:"id" json:"id"`
	Name      string `form:"name" json:"name"`
	CarType   string `form:"type" json:"type"`
	Rating    string `form:"rating" json:"rating"`
	Fuel      string `form:"fuel" json:"fuel"`
	Image     string `form:"image" json:"image"`
	HourRate  string `form:"hourRate" json:"hourRate"`
	DayRate   string `form:"dayRate" json:"dayRate"`
	MonthRate string `form:"monthRate" json:"monthRate"`
}

type ResponseCars struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Cars
}

type Users struct {
	UserId      string `form:"id" json:"id"`
	Email       string `form:"email" json:"email"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber"`
	City        string `form:"city" json:"city"`
	Zip         string `form:"zip" json:"zip"`
	Message     string `form:"message" json:"message"`
	Password    string `form:"password" json:"password"`
	Username    string `form:"username" json:"username"`
	Address     string `form:"address" json:"address"`
}

type ResponseUsers struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type Admin struct {
	Id       string `form:"id" json:"id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type ResponseAdmin struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Admin
}

type Orders struct {
	OrderId     string `form:"id" json:"id"`
	PickUpLoc   string `form:"pickUpLoc" json:"pickUpLoc"`
	DropOffLoc  string `form:"dropOffLoc" json:"dropOffLoc"`
	PickUpDate  string `form:"pickUpDate" json:"pickUpDate"`
	DropOffDate string `form:"dropOffDate" json:"dropOffDate"`
	PickUpTime  string `form:"pickUpTime" json:"pickUpTime"`
	CarId       string `form:"carId" json:"carId"`
	UserId      string `form:"userId" json:"userId"`
	AdminId     string `form:"adminId" json:"adminId"`
}

type ResponseOrders struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Orders
}
