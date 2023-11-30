package routes

import (
	"github.com/gorilla/mux"
	"github.com/kavikkannan/go-task/pkg/controllers"

)

var Registeruser = func (router *mux.Router)  {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/employe/", controllers.GetEmploye).Methods("GET")
	router.HandleFunc("/user/{useremail}", controllers.GetUserByEmail).Methods("GET")
	router.HandleFunc("/employe/{useremail}", controllers.GetEmployeByEmail).Methods("GET")
/* 	router.HandleFunc("/grievances/{useremail}", controllers.GetEmployeByGrievances).Methods("GET")
	router.HandleFunc("/user/{useremail}", controllers.UpdateUser).Methods("PUT")
 */	router.HandleFunc("/user/{useremail}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/bus/", controllers.GetBus).Methods("GET")

	router.HandleFunc("/allticket/{id}", controllers.GetallTicketById).Methods("GET")
	router.HandleFunc("/bookticket/{id}", controllers.UpdateSeats).Methods("PUT")
	
	router.HandleFunc("/bookedticket/{id}", controllers.GetallBookedTicketById).Methods("GET")
	router.HandleFunc("/bookeduser/{id}", controllers.GetBookedUserById).Methods("GET")

	router.HandleFunc("/bookdetails/", controllers.Createbookdetails).Methods("POST")

	router.HandleFunc("/bus/{id}", controllers.GetBusById).Methods("GET")
}