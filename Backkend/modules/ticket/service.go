package ticket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	if ticket.creator_id != creatorId {
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
	fmt.Fprintf(w, "Updated Status")
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

	if ticket.creator_id != creatorId {
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
