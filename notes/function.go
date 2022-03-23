package notes

type A struct {
	Name string
}

func DoSomething1(a *A) {
	b := a
	b.Name = "huhaolong"
}

func DoSomething2(a A) {
	b := a
	b.Name = "huhaolong"
}
func FunctionNotes() {

}
