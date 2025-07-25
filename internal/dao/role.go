// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"omp/internal/dao/internal"
)

// roleDao is the data access object for the table role.
// You can define custom methods on it to extend its functionality as needed.
type roleDao struct {
	*internal.RoleDao
}

var (
	// Role is a globally accessible object for table role operations.
	Role = roleDao{internal.NewRoleDao()}
)

// Add your custom methods and functionality below.
