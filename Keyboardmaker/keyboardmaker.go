package keyboardmaker

type keyboardmaker interface {
	MakeGrid(int, int)
	AddButton(string, int, int)
	ClearButton(int, int)
	GetKeyboard() interface{}
}
