package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kavikkannan/go-task/pkg/models"
	"github.com/kavikkannan/go-task/pkg/utils"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUser()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBus(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllBus()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetEmploye(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllEmploye()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]

    userDetails, _ := models.GetUserByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetEmployeByEmail(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userEmail := vars["useremail"]

    userDetails, _ := models.GetEmployeByEmail(userEmail)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func GetTicketById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
    userDetails, _ := models.GetTicketById(ID)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func GetBusById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
    userDetails, _ := models.GetBusByID(ID)
    res, _ := json.Marshal(userDetails)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func UpdateSeats(w http.ResponseWriter, r *http.Request) {
    var updateUser = &models.Busseat{}
    utils.ParseBody(r, updateUser)
    vars := mux.Vars(r)
    userId := vars["id"] 
    ID, err := strconv.ParseInt(userId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
    }
    userDetails, db := models.GetSeatByID(ID)

    if userDetails != nil {
        
        if updateUser.Bookedstatus != userDetails.Bookedstatus {

            userDetails.Bookedstatus = updateUser.Bookedstatus
        }
        db.Save(&userDetails)
        res, _ := json.Marshal(userDetails)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
    }
}

func GetallTicketById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
    // Assume userDetails is a slice of structs or a list of data associated with the given email
    // Fetch all rows associated with the email
    allDetails, err := models.GetAllTicketById(ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(allDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetBookedUserById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
    // Assume userDetails is a slice of structs or a list of data associated with the given email
    // Fetch all rows associated with the email
    allDetails, err := models.GetBookedUserById(ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(allDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}

func GetallBookedTicketById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId := vars["id"]
    ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
    // Assume userDetails is a slice of structs or a list of data associated with the given email
    // Fetch all rows associated with the email
    allDetails, err := models.GetAllBookedTicketById(ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(allDetails)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(res)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateBus(w http.ResponseWriter, r *http.Request) {
	CreateBus := &models.Bus{}
	utils.ParseBody(r, CreateBus)
	c := CreateBus.CreateBus()
	res, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func Createbookdetails(w http.ResponseWriter, r *http.Request) {
	CreateBookdetails := &models.Bookdetails{}
	utils.ParseBody(r, CreateBookdetails)
	c := CreateBookdetails.Createbookdetails()
	res, _ := json.Marshal(c)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userid"]
	fmt.Println(userId)
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsinf")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/* func UpdateUser(w http.ResponseWriter, r *http.Request) {
    var updateUser = &models.User{}
    utils.ParseBody(r, updateUser)
    vars := mux.Vars(r)
    userEmail := vars["useremail"] // Update the route variable name
    userDetails, db := models.GetUserByEmail(userEmail) // Update the function name

    if userDetails != nil {
        if updateUser.Userid != "" {
            userDetails.Userid = updateUser.Userid
        }

        if updateUser.Busid != "" {
            userDetails.Busid = updateUser.Busid
        }

        if updateUser.Email != "" {
            userDetails.Email = updateUser.Email
        }

        if updateUser.Deadline != "" {
            userDetails.Deadline = updateUser.Deadline
        }

        db.Save(&userDetails)
        res, _ := json.Marshal(userDetails)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(res)
    } else {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte("User not found"))
    }
}

 */