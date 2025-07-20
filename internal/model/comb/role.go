package comb

import "omp/internal/model/entity"

type Role struct {
	*entity.Role
	Permissions []*Permission
}
