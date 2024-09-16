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
