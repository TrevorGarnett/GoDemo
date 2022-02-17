package gofiles

type ByName []*Student

func (a ByName) Len() int { return len(a) }
func (a ByName) Less(i, j int) bool {
	if a[i].LastName == a[j].LastName {
		return a[i].FirstName < a[j].FirstName
	}
	return a[i].LastName < a[j].LastName
}
func (a ByName) Swap(i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}
