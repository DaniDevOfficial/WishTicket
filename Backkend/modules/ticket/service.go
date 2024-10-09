package ticket

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"wishticket/modules/user"
	"wishticket/util/auth"
	"wishticket/util/error"
	"wishticket/util/responses"
)

// Tickets

func GetAllOwnedTickets(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jwtData, err := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here
	requesterId := -1

	if err == nil {
		requesterId = jwtData.UserId
	}
	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)
	if err != nil {
		error.HttpResponse(w, "Error fetching userData", http.StatusBadRequest)
		return
	}

	tickets, err := GetAllOwnedTicketsFromDB(userId, requesterId, db)
	if err != nil {
		error.HttpResponse(w, "Error while getting All owned Tickets From DB", http.StatusBadRequest)
		return
	}

	log.Printf("User %d has %d tickets\n", userId, len(tickets))

	responses.ResponseWithJSON(w, tickets, http.StatusOK)
}

func GetAssignedTickets(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	jwtData, err := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here
	requesterId := -1
	if err == nil {
		requesterId = jwtData.UserId
	}
	log.Println(requesterId)
	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)
	if err != nil {
		error.HttpResponse(w, "Error fetching userData", http.StatusInternalServerError)
		return
	}

	tickets, err := GetAssignedTicketsFromDB(userId, requesterId, db)

	if err != nil {
		error.HttpResponse(w, "Error while getting All assigned Tickets From DB", http.StatusInternalServerError)
		return
	}

	responses.ResponseWithJSON(w, tickets, http.StatusOK)

}

func GetAllAssignedAndOwnedTicketsForUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jwtData, err := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here
	requesterId := -1
	if err == nil {
		requesterId = jwtData.UserId
	}

	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)
	if err != nil {
		error.HttpResponse(w, "Error fetching userData", http.StatusInternalServerError)
		return
	}

	assignedTickets, err := GetAssignedTicketsFromDB(userId, requesterId, db)
	if err != nil {
		log.Println(err)
		error.HttpResponse(w, "Error while getting All assigned Tickets From DB", http.StatusInternalServerError)
		return
	}

	ownedTickets, err := GetAllOwnedTicketsFromDB(userId, requesterId, db)
	if err != nil {
		error.HttpResponse(w, "Error while getting All owned Tickets From DB", http.StatusInternalServerError)
		return
	}

	response := struct {
		Assigned interface{} `json:"assigned"`
		Owned    interface{} `json:"owned"`
	}{
		Assigned: assignedTickets,
		Owned:    ownedTickets,
	}

	responses.ResponseWithJSON(w, response, http.StatusOK)
}

func GetTicketById(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jwtData, err := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here
	requesterId := -1
	if err == nil {
		requesterId = jwtData.UserId
	}

	ticketIdStr := r.URL.Query().Get("ticketId")

	if ticketIdStr == "" {
		error.HttpResponse(w, "TicketId is missing from query params", http.StatusBadRequest)
		return
	}

	ticketId, err := strconv.Atoi(ticketIdStr)
	if err != nil {
		error.HttpResponse(w, "TicketId must be a valid integer", http.StatusBadRequest)
		return
	}

	ticketData, err := GetTicketFromDB(ticketId, requesterId, db)

	if err != nil {
		log.Println(err)
		error.HttpResponse(w, "Ticket Does not exist", http.StatusBadRequest)
		return
	}

	response := ticketData

	responses.ResponseWithJSON(w, response, http.StatusOK)
}

func CreateNewTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: some auth before letting Creation happen but idk how to do this currenty

	userData, err := auth.GetJWTPayloadFromHeader(r)
	if err != nil {
		error.HttpResponse(w, "Error while Authenticating", http.StatusInternalServerError)
		return
	}

	var ticketData TicketRequest
	err = json.NewDecoder(r.Body).Decode(&ticketData)
	if err != nil {
		error.HttpResponse(w, "Error decoding RequestBody", http.StatusBadRequest)
		return
	}

	_, err = user.GetUserById(userData.UserId, db)
	if err != nil {
		error.HttpResponse(w, "No user with this Id Found", http.StatusBadRequest)
		return
	}

	ticketDataInsert := TicketForInsert{
		title:       ticketData.Title,
		description: ticketData.Description,
		visibility:  ticketData.Visibility,
		creator_id:  userData.UserId,
		dueDate:     ticketData.DueDate,
	}
	lastId, err := CreateNewTicketInDB(ticketDataInsert, db)

	if err != nil {
		error.HttpResponse(w, "Error while creating ticket", http.StatusBadRequest)
		return
	}
	response := struct {
		Message  string `json:"message"`
		TicketId int    `json:"ticketId"`
	}{
		Message:  "Successfully Created ticket",
		TicketId: lastId,
	}
	responses.ResponseWithJSON(w, response, http.StatusCreated)
}

func ChangeTicketStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	userData, err := auth.GetJWTPayloadFromHeader(r)

	if err != nil {
		error.HttpResponse(w, "Error Authenticating User", http.StatusUnauthorized)
		return
	}

	requesterId := userData.UserId

	var newStatus StatusRequest

	err = json.NewDecoder(r.Body).Decode(&newStatus)

	if err != nil {
		error.HttpResponse(w, "Error Decoding Body", http.StatusInternalServerError)
		return
	}
	userId := userData.UserId
	ticket, err := GetTicketFromDB(newStatus.TicketId, requesterId, db)
	// TODO: Either has to be creator or Assignee
	if err != nil {
		error.HttpResponse(w, "Error Fetching Ticket", http.StatusBadRequest)
		return
	}

	if ticket.CreatorId != userId {
		error.HttpResponse(w, "Not Allowed User", http.StatusUnauthorized)
		return
	}

	_, err = UpdateTicketStatus(newStatus, db)
	if err != nil {
		error.HttpResponse(w, "Error Updating Status", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message  string `json:"message"`
		TicketId int    `json:"ticketId"`
	}{
		Message:  "Successfully Updated Status",
		TicketId: newStatus.TicketId,
	}
	responses.ResponseWithJSON(w, response, http.StatusOK)
}

func CommentOnTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: Crud on comments
}

func AddAssigneeToTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Add assignee to ticket")

	userData, err := auth.GetJWTPayloadFromHeader(r)
	if err != nil {
		error.HttpResponse(w, "Error Authenticating User", http.StatusUnauthorized)
		return
	}
	requesterId := userData.UserId
	var addAssignee AddAssigneeRequest
	err = json.NewDecoder(r.Body).Decode(&addAssignee)
	if err != nil {
		error.HttpResponse(w, "Error Decoding Body", http.StatusInternalServerError)
		return
	}

	userId := userData.UserId

	// TODO: Either has to be creator or Assignee
	ticket, err := GetTicketFromDB(addAssignee.TicketId, requesterId, db)

	if err != nil {
		error.HttpResponse(w, "Error Fetching Ticket", http.StatusBadRequest)
		return
	}

	if ticket.CreatorId != userId {
		error.HttpResponse(w, "Not Allowed User", http.StatusUnauthorized)
		return
	}

	//TODO: check if someone blocked the other user
	err = CreateNewAssignment(addAssignee, db)
	if err != nil {
		error.HttpResponse(w, "Error Creating assignment", http.StatusInternalServerError)
		return
	}
	response := struct {
		Message  string `json:"message"`
		TicketId int    `json:"ticketId"`
	}{
		Message:  "Successfully Added Assignee to Ticket",
		TicketId: addAssignee.TicketId,
	}
	responses.ResponseWithJSON(w, response, http.StatusOK)
}

func RemoveAssigneeFromTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//TODO: implement
}
