package nocopy

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type UserInfo struct {
	noCopy  noCopy
	Name    string
	Address string
}

type Person struct {
	User UserInfo
}
