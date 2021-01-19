package core

const createTicket = `
	INSERT INTO ticket (name, position, list_id) VALUES ($1, $2, $3) RETURNING id;`

const getTicketsByListID = `
	SELECT ticket.id, ticket.name, ticket.position
	FROM ticket
	JOIN list ON ticket.list_id = list.id
	WHERE list.id = $1;`
