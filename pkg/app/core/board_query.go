package core

const getBoard = `
	SELECT id, name FROM board WHERE id = $1;`

// TODO Should we limit amount of boards?
const getBoardLists = `
	SELECT list.id, list.name, list.position
	FROM list
	JOIN board ON list.board_id = board.id
	WHERE board.id = $1;`

const createBoard = `
	WITH inserted_data AS (
		INSERT INTO board (name) VALUES
			($1)
		RETURNING id
	)
	INSERT INTO list (name, board_id) VALUES
		($2, (SELECT * FROM inserted_data))
	RETURNING board_id;`

const deleteBoard = `
	DELETE FROM board WHERE id = $1;`
