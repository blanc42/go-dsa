package ll

import "errors"

type Dnode struct {
	left  *Dnode
	data  int
	right *Dnode
}

type Dll struct {
	length int
	head   *Dnode
}

func NewDll() *Dll {
	temp := new(Dll)
	temp.length = 0
	temp.head = nil
	return temp
}

// PUT FUNCTIONS : insert, prepend, append,

func (d *Dll) Insert(data int, position int) error {
	if position > d.length {
		return errors.New("out of bounds")

	}

	return nil
}

// GET FUNCTIONS : POP, PUSH, QUEUE, DEQUE, GetNode
