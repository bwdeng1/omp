// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"omp/internal/dao/internal"
)

// userDao is the data access object for the table user.
// You can define custom methods on it to extend its functionality as needed.
type userDao struct {
	*internal.UserDao
}

var (
	// User is a globally accessible object for table user operations.
	User = userDao{internal.NewUserDao()}
)

// Add your custom methods and functionality below.
