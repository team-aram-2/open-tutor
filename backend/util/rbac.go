package util

type (
	RoleMask uint16
	Roles    interface {
		AddRole()
		RemoveRole()
		HasRole()
	}
	UserRole uint16 // would use uint8 here but the smallest serial type in pgsql is uint16 compatible
)

const (
	user      UserRole = 0         // iota = 0
	tutor     UserRole = 1 << iota // 1
	moderator UserRole = 1 << iota // 2
	admin     UserRole = 1 << iota // 4
)

func AddRole(toAdd UserRole, roleMask RoleMask) (RoleMask, error) {
	// TODO ret toAdd OR roleMask
}

func RemoveRole(toRemove UserRole, roleMask RoleMask) (RoleMask, error) {
	// TODO ret !toRemove AND roleMask
}

func HasRole(wants RoleMask, has RoleMask) (bool, error) {
	// TODO ret (wants AND has)
}
