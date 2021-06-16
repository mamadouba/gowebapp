package roles

import (
	"fmt"
	"strings"
)

type Role struct {
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
}

func (r *Role) Add(permission string) {
	if r.Permissions != "" {
		r.Permissions += fmt.Sprintf(",%s", permission)
	} else {
		r.Permissions = permission
	}
}

func (r *Role) ToArray() []string {
	return strings.Split(r.Permissions, ",")
}

func (r *Role) Authorized(permission string) bool {
	return strings.Contains(r.Permissions, permission)
}
