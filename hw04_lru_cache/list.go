package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	list  *list
}

type list struct {
	item ListItem
	len  int
}

func (item *ListItem) NextItem() *ListItem {
	if p := item.Next; item.list != nil && p != &item.list.item {
		return p
	}
	return nil
}

func (item *ListItem) PrevItem() *ListItem {
	if p := item.Prev; item.list != nil && p != &item.list.item {
		return p
	}
	return nil
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.item.Next
}

func (l list) Back() *ListItem {
	if l.Len() == 0 {
		return nil
	}
	return l.item.Prev
}

func (l *list) PushFront(v interface{}) *ListItem {

	var el = &ListItem{Value: v}
	el.list = l
	el.Next = l.item.Next
	el.Prev = &l.item
	el.Next.Prev = el
	el.Prev.Next = el
	l.len++
	return el
}

func (l *list) PushBack(v interface{}) *ListItem {
	var el = &ListItem{Value: v}
	el.list = l
	el.Next = l.item.Prev.Next
	el.Prev = l.item.Prev
	el.Prev.Next = el
	el.Next.Prev = el
	l.len++
	return el
}

func NewList() *list {
	var l = new(list)
	l.item.Next = &l.item
	l.item.Prev = &l.item
	l.len = 0
	return l
}

func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Next = nil
	i.Prev = nil
	i.list = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.item.Next == i {
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next

	i.Next = l.item.Next
	i.Prev = &l.item

	i.Next.Prev = i
	i.Prev.Next = i
}
