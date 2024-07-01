package slist

type Element struct {
	Id    int
	Name  string
	Email string
	Next  *Element
}

type List struct {
	Head *Element
	Size int
}

func (l *List) Push(id int, name string, email string) {
	n := &Element{Id: id, Name: name, Email: email}
	if l.Size == 0 {
		l.Head = n
		l.Size = 1
		return
	}
	s := l.Head

	for s != nil {
		if s.Next == nil {
			s.Next = n
			break
		}
		s = s.Next
	}
	l.Size++
}

func (l *List) Pop(id int) {
	elem := l.Head

	for elem != nil {
		if elem.Next.Id == id {
			elem.Next = elem.Next.Next
			l.Size--

		}
		elem = elem.Next
	}
}
