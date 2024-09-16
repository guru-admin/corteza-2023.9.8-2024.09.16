package store

import "github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/dal"

// DAL uses given store as DAL connection
//
// This is mainly used to wrap primary store connection with DAL connection wrap
// and use it to interact with records in a primary DB
func DAL(s Storer) dal.Connection {
	return s.ToDalConn()
}
