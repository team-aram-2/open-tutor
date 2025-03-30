package util

type RoleMask uint16 // would use uint8 here but the smallest serial type in pgsql is uint16 compatible

const (
	User      RoleMask = 1 << iota // 1
	Tutor                          // 2
	Moderator                      // 4
	Admin                          // 8
)

func (mask *RoleMask) Add(toAdd RoleMask) {
	*mask |= toAdd
}

func (mask *RoleMask) Remove(toRemove RoleMask) {
	*mask &= ^toRemove
}

func (mask *RoleMask) Has(wanted RoleMask) bool {
	return (*mask)&(wanted) == wanted
}

func Roles() []RoleMask {
	return []RoleMask{
		User,
		Tutor,
		Moderator,
		Admin,
	}
}

func (mask RoleMask) Labels() []string {
	labels := []string{}
	roleMap := map[RoleMask]string{
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
