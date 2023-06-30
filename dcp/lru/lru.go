package lur

// least recently used cache problem
// daily coding problem 697

type Pair struct {
	value int
	node  *Node
}

type LRUcache struct {
	Max   int
	Count int
	Cache map[int]Pair
	Dll   *DoublyLinkedList
}

// initializer a new LRU cache
func (l *LRUcache) New(max int) {
	l.Max = max
	l.Cache = make(map[int]Pair)
	l.Dll = l.Dll.New()
}

// get the vlues if exists and return null if does not
func (l *LRUcache) Get(key int) int {
	if pair, ok := l.Cache[key]; ok {
		l.Dll.Update(pair.node)
		return pair.value
	} else {
		return 0
	}

}

// sets the key to a value
func (l *LRUcache) Set(key int, value int) {
	newNode := new(Node)
	l.Dll.Add(newNode)

}
