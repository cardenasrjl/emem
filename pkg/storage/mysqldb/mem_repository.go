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
func (db *DB) CreateMem(ctx context.Context, m *v1.NewMemRequest) (id int64, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// insert mem data
	res, err := c.ExecContext(ctx, "INSERT INTO `mems` (`title`, `description`) VALUES(?, ?)",
		m.Title, m.Description)
	if err != nil {
		err = status.Error(codes.Unknown, "failed to insert a mem -> "+err.Error())
		return
	}

	// get the ID
	id, err = res.LastInsertId()
	if err != nil {
		err = status.Error(codes.Unknown, "failed to retrieve id for created mem-> "+err.Error())
		return
	}

	return
}

//GetMem ...
func (db *DB) GetMem(ctx context.Context, id int64) (m *v1.Mem, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// query mem
	rows, err := c.QueryContext(ctx, "SELECT `id`, `title`, `description`, `created_at`,  `updated_at` FROM mems WHERE `id`=?",
		id)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from mem-> "+err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve data from mems-> "+err.Error())
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
			id))
	}

	m = &v1.Mem{}

	var createdAt time.Time
	var updatedAt time.Time

	if err := rows.Scan(&m.Id, &m.Title, &m.Description, &createdAt, &updatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from mem row-> "+err.Error())
	}

	m.CreatedAt, err = ptypes.TimestampProto(createdAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "created at field has invalid format-> "+err.Error())
	}

	m.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
	if err != nil {
		return nil, status.Error(codes.Unknown, "updated at field has invalid format-> "+err.Error())
	}

	return
}

//UpdateTodo ...
func (db *DB) UpdateMem(ctx context.Context, m *v1.UpdateMemRequest) (err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// update ToDo
	_, err = c.ExecContext(ctx, "UPDATE mems SET `title`=?, `description`=?  WHERE `id`=?",
		m.Title, m.Description, m.Id)
	if err != nil {
		return status.Error(codes.Unknown, "failed to update mem-> "+err.Error())
	}

	return
}

//GetMems ...`
func (db *DB) GetMems(ctx context.Context) (list []*v1.Mem, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get ToDo list
	rows, err := c.QueryContext(ctx, "SELECT `id`, `title`, `description`, `created_at`, `updated_at` FROM mems order by id desc")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()

	list = make([]*v1.Mem, 0)
	for rows.Next() {
		m := new(v1.Mem)
		// get data
		var createdAt time.Time
		var updatedAt time.Time
		if err := rows.Scan(&m.Id, &m.Title, &m.Description, &createdAt, &updatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from mem row-> "+err.Error())
		}

		m.CreatedAt, err = ptypes.TimestampProto(createdAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "created at field has invalid format-> "+err.Error())
		}

		m.UpdatedAt, err = ptypes.TimestampProto(updatedAt)
		if err != nil {
			return nil, status.Error(codes.Unknown, "updated at field has invalid format-> "+err.Error())
		}

		list = append(list, m)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from mem-> "+err.Error())
	}

	return
}

//delete todo
func (db *DB) DeleteMem(ctx context.Context, id int64) (err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return
	}
	defer c.Close()

	// delete todo
	res, err := c.ExecContext(ctx, "DELETE FROM `mems` WHERE `id`=?", id)
	if err != nil {
		return status.Error(codes.Unknown, "failed to delete ToDo-> "+err.Error())
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return status.Error(codes.Unknown, "failed to retrieve rows affected value-> "+err.Error())
	}

	if rows == 0 {
		return status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
			id))
	}

	return
}
