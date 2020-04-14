package mysqldb

import (
	"context"
	"fmt"

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
func (db *DB) GetMem(ctx context.Context, id int64) (m *v1.MemInfo, err error) {
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

	m = &v1.MemInfo{}
	if err := rows.Scan(&m.Mem.Id, &m.Mem.Title, &m.Mem.Description, &m.CreatedAt, &m.UpdatedAt); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve field values from mem row-> "+err.Error())
	}

	//m.Reminder, err = ptypes.TimestampProto(reminder)
	//if err != nil {
		//return nil, status.Error(codes.Unknown, "reminder field has invalid format-> "+err.Error())
	//}

	//if rows.Next() {
	//	return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple ToDo rows with ID='%d'",
	//		id))
	//}

	return

}

//UpdateTodo ...
func (db *DB) UpdateMem(ctx context.Context, m *v1.Mem) ( err error) {
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
		return  status.Error(codes.Unknown, "failed to update mem-> "+err.Error())
	}

	return
}

//GetMems ...`
func (db *DB) GetMems(ctx context.Context) (list []*v1.MemInfo, err error) {
	// get SQL connection from pool
	c, err := db.Connect(ctx)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// get ToDo list
	rows, err := c.QueryContext(ctx, "SELECT `id`, `title`, `description`, `created_at`, `updated_at` FROM mems")
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
	}
	defer rows.Close()
	
	list = []*v1.MemInfo{}
	for rows.Next() {
		m := new(v1.MemInfo)
		m.Mem = new(v1.Mem)
		if err := rows.Scan(&m.Mem.Id, &m.Mem.Title, &m.Mem.Description, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, status.Error(codes.Unknown, "failed to retrieve field values from mem row-> "+err.Error())
		}
		
		list = append(list, m)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed to retrieve data from mem-> "+err.Error())
	}
	return
}