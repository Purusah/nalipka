Draft for task manager application.

Main entities:
* Board
* List
* Ticket
* Comment

API:

* /api/v1/board
    * GET - return list of boards sorted by order (sorted by name)
    * POST - create board
* /api/v1/board/:id
    * GET - get board by id
    * DELETE - delete board by id
* /api/v1/board/:id/list
    * GET - get all lists (sorted)
* /api/v1/list/:id
    * GET - get list by id
    * DELETE - delete list by id
* /api/v1/list/:id/ticket
    * POST - create ticket
* /api/v1/ticket/:id
    * GET - get ticket by id
    * PUT - update ticket (update order, update parent list)
    * DELETE - delete ticket by id (along with comments)
* /api/v1/ticket/:id/comment
    * GET - list comments sorted by their creation date
    * POST - create comment
* /api/v1/comment/:id
    * PUT - update the comment
    * DELETE - delete the comment

NOTES:
* the first column created by default when a Project created;
* the last column cannot be deleted.
* column name must be unique;
* if list is deleted its tasks are moved to the left list of the current
* move a list left or right

Requirements:
* docker
* docker-compose (>= 1.27)
* go (1.15.*)
