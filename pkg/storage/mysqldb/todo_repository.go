package mysqldb

import (
	"context"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"

	v1 "github.com/cardenasrjl/emem/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//CreateTodo ...
func (db *DB) CreateTodo(ctx context.Context, td *v1.ToDo) (id int64, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	//converts the timestamp in a time.time
	reminder, err := ptypes.Timestamp(td.Reminder)
	if err != nil {
		return id, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// insert ToDo entity data
	res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
		td.Title, td.Description, reminder)
	if err != nil {
		err = status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
		return
	}

	// get ID of creates ToDo
	id, err = res.LastInsertId()
	if err != nil {
		err = status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
		return
	}

	return
}

//ReadTodo ...
func (db *DB) ReadTodo(ctx context.Context, id int64) (td *v1.ToDo, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// query ToDo by ID
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
			id))
	}

	td = &v1.ToDo{}
	// get ToDo data
	var reminder time.Time
	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
	}

	td.Reminder, err = ptypes.TimestampProto(reminder)
	if err != nil {
		return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
	}

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
			id))
	}

	return

}

//UpdateTodo ...
func (db *DB) UpdateTodo(ctx context.Context, td *v1.ToDo) (rows int64, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	reminder, err := ptypes.Timestamp(td.Reminder)
	if err != nil {
		return rows, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
	}

	// update ToDo
	res, err := c.ExecContext(ctx, "UPDATE ToDo SET `Title`=?, `Description`=?, `Reminder`=? WHERE `ID`=?",
		td.Title, td.Description, reminder, td.Id)
	if err != nil {
		return rows, status.Error(codes.Unknown, "failed to update ToDo-> "+err.Error())
	}

	rows, err = res.RowsAffected()
	if err != nil {
		return 0, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return 0, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
			td.Id))
	}
	return
}

//delete todo
func (db *DB) DeleteTodo(ctx context.Context, id int64) (rows int64, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// delete todo
	res, err := c.ExecContext(ctx, "DELETE FROM ToDo WHERE `ID`=?", id)
	if err != nil {
		return 0, status.Error(codes.Unknown, "failed to delete ToDo-> "+err.Error())
	}

	rows, err = res.RowsAffected()
	if err != nil {
		return 0, status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return 0, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
			id))
	}

	return
}

//ReadAllTodo ...`
func (db *DB) ReadAllTodo(ctx context.Context) (list []*v1.ToDo, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get ToDo list
	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()

	var reminder time.Time
	list = []*v1.ToDo{}
	for rows.Next() {
		td := new(v1.ToDo)
		if err := rows.Scan(&td.Id, &td.Title, &td.Description, &reminder); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from ToDo row-> "+err.Error())
		}
		td.Reminder, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
		}
		list = append(list, td)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
	}
	return
}
