package util

type Role uint16 // would use uint8 here but the smallest serial type in pgsql is uint16 compatible

const (
	User      Role = 1 << iota // 1
	Tutor                      // 2
	Moderator                  // 4
	Admin                      // 8
)

func (mask *Role) Add(toAdd Role) {
	*mask |= toAdd
}

func (mask *Role) Remove(toRemove Role) {
	*mask &= ^toRemove
}

func (mask *Role) Has(wanted Role) bool {
	return (*mask)&(wanted) == wanted
}

func Roles() []Role {
	return []Role{
		User,
		Tutor,
		Moderator,
		Admin,
	}
}

func (mask Role) Labels() []string {
	labels := []string{}
	roleMap := map[Role]string{
		User:      "User",
		Tutor:     "Tutor",
		Moderator: "Moderator",
		Admin:     "Admin",
	}
	for role, label := range roleMap {
		if mask.Has(role) {
			labels = append(labels, label)
		}
	}
	return labels
}
