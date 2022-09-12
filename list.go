package toolkit

type ListNode[T any] struct {
	Data *T
	Next *ListNode[T]
	Pre  *ListNode[T]
}

func (list *ListNode[T]) PushFront(data *ListNode[T]) {
	if data == nil {
		return
	}
	data.Next = list
	data.Pre = list.Pre
	if list.Pre != nil {
		list.Pre.Next = data
	}
	list.Pre = data
}

func (list *ListNode[T]) PushBack(data *ListNode[T]) {
	if data == nil {
		return
	}
	data.Next = list.Next
	data.Pre = list
	if list.Next != nil {
		list.Next.Pre = data
	}
	list.Next = data
}

func (list *ListNode[T]) Pop() *ListNode[T] {
	if list.Pre != nil {
		list.Pre.Next = list.Next
	}

	if list.Next != nil {
		list.Next.Pre = list.Pre
	}
	return list
}
