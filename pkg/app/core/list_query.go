package core

const getListByID = `
	SELECT list.id, list.name, list.position FROM list WHERE list.id = $1;`

// getListForDelete select 2 rows and total amount of lists refferenced to
// same board
const getListForDelete = `
	WITH
		board_ids AS (SELECT list.board_id FROM list WHERE list.id = $1)
	SELECT list.id, list.position, clist.count
	FROM list
	JOIN (
		SELECT COUNT(list.id) count
		FROM list
		WHERE list.board_id IN (SELECT board_id FROM board_ids)
	) clist
		ON list.id = list.id
	WHERE
		list.id <= $1 AND list.board_id IN (SELECT board_id FROM board_ids)
	ORDER BY
		list.id DESC, position DESC
	LIMIT 2;`

const updateTicketsList = `
	UPDATE ticket
	SET list_id = $1
	WHERE list_id = $2;
`

const deleteList = `
	DELETE FROM list WHERE id = $1;`

const createList = `
	INSERT INTO list (name, position, board_id) VALUES ($1, $2, $3) RETURNING id;`
