package group

type Group struct {
	ID          string
	Name        string
	GroupOwner  GroupMember
	GroupMember []GroupMember
}

func (g Group) addMember(userId string) {
}

func (g Group) changeName(name string) {
	g.Name = name
}
