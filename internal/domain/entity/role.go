package entity

type Role struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Permissions []string `json:"permissions"`
}

func NewRole(id int64, name string, permissions []string) *Role {
    return &Role{
        ID:          id,
        Name:        name,
        Permissions: permissions,
    }
}

func (r *Role) AddPermission(permission string) {
    r.Permissions = append(r.Permissions, permission)
}

func (r *Role) RemovePermission(permission string) {
    for i, p := range r.Permissions {
        if p == permission {
            r.Permissions = append(r.Permissions[:i], r.Permissions[i+1:]...)
            break
        }
    }
}