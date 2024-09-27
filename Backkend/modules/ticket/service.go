package ticket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wishticket/modules/user"
	"wishticket/util/auth"
)

// Tasks

func GetAllOwnedTickets(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jwtData, _ := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here

	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)
	if err != nil {
		log.Printf("Error fetching tickets for user %d: %v", userId, err)
		http.Error(w, `{"error": "Failed to retrieve tickets"}`, http.StatusInternalServerError)
		return
	}
	onlyPublic := true
	if userId == jwtData.UserId {
		onlyPublic = false
	}

	tickets, err := GetAllOwnedTicketsFromDB(userId, db, onlyPublic)
	if err != nil {
		log.Printf("Error fetching tickets for user %d: %v", userId, err)
		http.Error(w, `{"error": "Failed to retrieve tickets"}`, http.StatusInternalServerError)
		return
	}

	log.Printf("User %d has %d tickets\n", userId, len(tickets))

	// Respond with tickets in JSON format
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tickets)
	if err != nil {
		log.Println("Error encoding tickets:", err)
		http.Error(w, `{"error": "Failed to encode tickets"}`, http.StatusInternalServerError)
		return
	}

	log.Println("Successfully responded with tickets.")
}

func GetAssignedTickets(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	jwtData, _ := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here

	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)
	if err != nil {
		log.Printf("Error fetching tickets for user %d: %v", userId, err)
		http.Error(w, `{"error": "Failed to retrieve tickets"}`, http.StatusInternalServerError)
		return
	}
	onlyPublic := true
	if userId == jwtData.UserId {
		onlyPublic = false
	}

	tickets, err := GetAssignedTicketsFromDB(userId, db, onlyPublic)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func GetAllAssignedAndOwnedTicketsForUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	jwtData, _ := auth.GetJWTPayloadFromHeader(r) // TODO: do some error handeling here

	ownerUsername := r.URL.Query().Get("username")
	userId, err := user.GetUserIdByName(ownerUsername, db)

	if err != nil {
		log.Printf("Error fetching tickets for user %d: %v", userId, err)
		http.Error(w, `{"error": "Failed to retrieve tickets"}`, http.StatusInternalServerError)
		return
	}

	onlyPublic := true
	if userId == jwtData.UserId {
		onlyPublic = false
	}

	assignedTickets, err := GetAssignedTicketsFromDB(userId, db, onlyPublic)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	log.Println("assignedTickets: ")
	log.Print(assignedTickets)

	ownedTickets, err := GetAllOwnedTicketsFromDB(userId, db, onlyPublic)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	log.Println("ownedTickets: ")
	log.Print(ownedTickets)
	response := struct {
		AssignedTickets interface{} `json:"assignedTickets"`
		OwnedTickets    interface{} `json:"ownedTickets"`
	}{
		AssignedTickets: assignedTickets,
		OwnedTickets:    ownedTickets,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		log.Println("Error marshaling JSON:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func CreateNewTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: some auth before letting Creation happen but idk how to do this currenty

	var ticketData TicketRequest
	err := json.NewDecoder(r.Body).Decode(&ticketData)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	userData, err := auth.GetJWTPayloadFromHeader(r)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	_, err = user.GetUserById(userData.UserId, db)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	ticketDataInsert := TicketForInsert{
		title:       ticketData.Title,
		description: ticketData.Description,
		visibility:  ticketData.Visibility,
		creator_id:  userData.UserId,
	}
	err = CreateNewTicketInDB(ticketDataInsert, db)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	fmt.Fprintf(w, "yayyyy")
}

func ChangeTicketStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	var newStatus StatusRequest

	err := json.NewDecoder(r.Body).Decode(&newStatus)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	//TODO: get userId form jwt
	userData, err := auth.GetJWTPayloadFromHeader(r)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	userId := userData.UserId
	ticket, err := GetTicketById(newStatus.TicketId, db)
	// TODO: Either has to be creator or Assignee
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	if ticket.CreatorId != userId {
		fmt.Fprintf(w, "Authorization Error")
		log.Println("Wrong Creator id")
		return
	}

	_, err = UpdateTicketStatus(newStatus, db)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	log.Println("Updated Status")
	fmt.Fprintf(w, "Updated Status")
}

func CommentOnTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: Crud on comments
}

func AddAssigneeToTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Get all tickets youre assigned to")
	var addAssignee AddAssigneeRequest
	err := json.NewDecoder(r.Body).Decode(&addAssignee)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	userData, err := auth.GetJWTPayloadFromHeader(r)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	userId := userData.UserId

	// TODO: Either has to be creator or Assignee
	ticket, err := GetTicketById(addAssignee.TicketId, db)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	if ticket.CreatorId != userId {
		fmt.Fprintf(w, "Authorization Error")
		log.Println("Wrong Creator id")
		return
	}

	//TODO: check if someone blocked the other user
	err = CreateNewAssignment(addAssignee, db)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	//TODO: Success handler
}

func RemoveAssigneeFromTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//TODO: implement
}
