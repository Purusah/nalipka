package core

import (
	"context"

	"github.com/purusah/nalipka/pkg/repository"
)

// Ticket ...
type Ticket struct {
	ID       int
	Name     string
	Position int
}

// CreateTicket ...
func CreateTicket(ctx context.Context, q repository.QueryableRow, t Ticket, listID int) (ticketID int, err error) {
	r := q.QueryRow(ctx, createTicket, t.Name, t.Position, listID)
	err = r.Scan(&ticketID)
	if err != nil {
		return ticketID, err
	}
	return ticketID, nil
}

// GetTicketsByListID ...
func GetTicketsByListID(ctx context.Context, q repository.Queryable, listID int) (tickets []Ticket, err error) {
	var t Ticket
	rows, err := q.Query(ctx, getTicketsByListID, listID)
	if err != nil {
		return tickets, err
	}
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.Name, &t.Position)
		if err != nil {
			return tickets, err
		}
		tickets = append(tickets, t)
	}
	if len(tickets) == 0 {
		return tickets, repository.ErrNoRowsFound
	}
	return tickets, nil
}
