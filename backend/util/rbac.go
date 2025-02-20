package util

type (
	Role uint16 // would use uint8 here but the smallest serial type in pgsql is uint16 compatible
)

const (
	User      Role = 1 << iota // 1
	Tutor                      // 2
	Moderator                  // 4
	Admin                      // 8
)

func (r Role) Add(toAdd Role) Role {
	return r | toAdd
}

func (r Role) Remove(toRemove Role) Role {
	return r & ^toRemove
}

func (r Role) Has(wanted Role) bool {
	return (r)&(wanted) == wanted
}
