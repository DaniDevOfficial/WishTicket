package ticket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Tasks

func GetAllOwnedTickets(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// TODO: Get id from jwt
	userId := 1
	tickets, err := GetAllOwnedTicketsFromDB(userId, db)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
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

	// TODO: Get user id not from header, but from jwt token

	ticketDataInsert := TicketForInsert{
		title:       ticketData.Title,
		description: ticketData.Description,
		creator_id:  1,
	}
	err = CreateNewTicketInDB(ticketDataInsert, db)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
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
	creatorId := 1

	ticket, err := GetTicketById(newStatus.TicketId, db)
	// TODO: Either has to be creator or Assignee
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	if ticket.CreatorId != creatorId {
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
	var addAssignee AddAssigneeRequest
	err := json.NewDecoder(r.Body).Decode(&addAssignee)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	// later get creator id from getTicket, but for now this is fine
	// and check it if the jwt creator id is the same as the creator id from the ticket
	creatorId := 1

	// TODO: Either has to be creator or Assignee
	ticket, err := GetTicketById(addAssignee.TicketId, db)

	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	if ticket.CreatorId != creatorId {
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
