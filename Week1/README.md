# 本周学习小结(2020-11-16 ~ 2020-11-22)

本周是参加算法训练营的第一周，由于此前LeetCode已经刷了不少了，因此第一周感觉还是比较轻松的，参加训练营的目的是希望对算法能够形成一个系统的认识。

最近开始使用Golang，因此用Golang刷题，但是比较生疏。本周总结学习了Golang container(heap,list,ring)，记录如下：

- [1. 前言](#1-前言)
- [2. heap(堆)](#2-heap堆)
- [3. list(双向链表)](#3-list双向链表)
- [4. ring(循环链表)](#4-ring循环链表)
- [5. 总结](#5-总结)

## 1. 前言
C++的STL标准库中定义了常见的数据结构，比如stack、vector、list等等；而在Java中，也有许多集合（Collections）可供调用，比如LinkedList、ArrayList、PriorityQueue等等。而在Go中，语言本身提供的高级数据结构并不多，主要是[`container`](https://golang.org/pkg/container/)下的`heap`(堆)、`list`(双向链表)、`ring`(循环链表)。下面主要分析这3个数据结构在Go中的实现以及常用的API。

## 2. heap(堆)
"container/heap"包为任意类型的数据提供了**堆**操作的接口，但前提是该类型必须实现`heap.Interface`接口。该接口定义如下：
```golang
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
```
其中`sort.Interface`的定义如下：
```golang
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
// 为了让集合中的元素能够比较，则必须实现sort.Interface接口
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	// 如果Less()返回true，则item[i]将会排在item[j]的前面
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```
因此，为了使用`heap`，必须实现`Len()`，`Less()`，`Swap()`，`Push()`，`Pop()`这5个方法。

首先看一个例子。下面的例子中，堆中存放的元素是`int`。
```golang
package main

import (
	"container/heap"
	"fmt"
)

// 定义了一个数据类型为int的最小堆
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
// 如果是大顶堆，则return h[i] > h[j]
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小的元素排前面
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	minHeap := IntHeap{2, 1, 5}
	heap.Init(&minHeap) // 初始化最小堆
	heap.Push(&minHeap, 3) // 插入新元素，内部会动态调整，使其重新满足最小堆的特性
	fmt.Printf("min number: %d\n", minHeap[0]) // 1，堆顶元素始终为最小元素
	for minHeap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(&minHeap).(int)) // 1 2 3 5
	}
}
```

再看一个复杂点的例子，在这个例子中，堆中存放的元素是一个自定义的结构体。
```golang
package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Item struct {
	Value    string // 存放的元素
	Priority int    // 优先级
	Index    int    // 该元素在堆中的下标
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool { // 大顶堆，优先级高的排前面
	if pq[i].Priority != pq[j].Priority {
		return pq[i].Priority > pq[j].Priority
	} else {
		return strings.Compare(pq[i].Value, pq[j].Value) < 0 // 按照字典序排序
	}
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// 将新元素插入至原数组末尾
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

// 将堆顶元素弹出
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.Index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Update(old *Item, newValue string, newPriority int) {
	old.Value = newValue
	old.Priority = newPriority
	heap.Fix(pq, old.Index) // 改变数据项后并重新调整，使其满足堆结构特性
}

func main() {
	// 构造测试数据
	items := map[string]int{
		"Kobe Bryant":  99,
		"LeBron James": 95,
		"Steven Curry": 94,
	}
	pq := make(PriorityQueue, len(items))
	i := 0
	for key, value := range items {
		pq[i] = &Item{
			Value:    key,
			Priority: value,
		}
		i++
	}
	// 初始化最大堆
	heap.Init(&pq)
	// 插入一个新元素
	item := &Item{
		Value:    "Michael Jordan",
		Priority: 90,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 100)

	for pq.Len() > 0 {
		x := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d, %s\n", x.Priority, x.Value)
	}
}
```

## 3. list(双向链表)
"container/list"包实现了一个双向链表（doubly linked list）。首先，看一个简单的例子。
```golang
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New() // 创建一个新的链表
	fmt.Println(l.Len()) // 0
	e4 := l.PushBack(4) // 返回新插入元素对应的节点
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4) // 在节点e4之前插入元素3
	l.InsertAfter(2, e1)
	fmt.Println(l.Len()) // 4
	// 遍历链表
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value.(int)) // 1 2 3 4
	}
}
```
在Java的源码实现中，链表中的一个节点被称为`Node`，而在Go中，list中的节点称为`Element`，其定义如下：
```golang
// Element is an element of a linked list.
type Element struct {
	// 指向后一个节点和前一个节点的指针
	next, prev *Element

	// 该节点属于哪个链表
	list *List

	// 节点中存放的数据
	Value interface{}
}
```
而`List`的定义如下：
```golang
// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
// List的零值不是nil，而是一个空的链表
type List struct {
	root Element // 哨兵元素，作为双向链表的连接点
	len  int     // 当前链表的长度（不包括哨兵元素）
}
```

list是一个双向链表（doubly linked list），当创建一个新链表，其实现如下：
```golang
// Init initializes or clears list l.
// Init方法用于初始化链表，或者清空链表
func (l *List) Init() *List {
	l.root.next = &l.root // 向前指向自己
	l.root.prev = &l.root // 向后指向自己
	l.len = 0
	return l
}

// New returns an initialized list.
// New方法返回一个新的链表
func New() *List { 
    return new(List).Init()
}
```
`root`作为哨兵节点，**`root.next`表示链表的首节点，`root.prev`表示链表的尾结点**。比如，链表的插入操作，实现如下：
```golang
// insert inserts e after at, increments l.len, and returns e.
// 在节点 at 之后插入节点 e
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// PushFront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root) // 头插法，即在root节点之后插入新节点
}

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev) // 尾插法，即在root.prev节点之后插入新节点
}
```
`list`比较简单，不再赘述，它提供的API参考[这里](https://golang.org/pkg/container/list/)，更详细的源码实现参考[container/list](https://golang.org/src/container/list/list.go)。

## 4. ring(循环链表)
"container/ring"包实现了**循环链表**（circular lists），或者说，就是一个**环**，其定义了一个`Ring`类型，如下所示：
```golang
// A Ring is an element of a circular list, or ring.
// Rings do not have a beginning or end; a pointer to any ring element
// serves as reference to the entire ring. Empty rings are represented
// as nil Ring pointers. The zero value for a Ring is a one-element
// ring with a nil Value.
// 注意和list.Element的区别
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}
```
环的结构有点特殊，环的尾部就是头部（`Ring`与`List`不同，前者并不存在`root`这样一个哨兵元素），所以每个元素实际上就可以代表自身的这个环。 **它不需要像list一样保持`List`和`Element`两个结构，只需要保持一个`Ring`结构就行**。

一个简单的示例如下：
```golang
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	l := ring.New(5)
	fmt.Printf("len:%d\n", l.Len()) // 5
	for i := 0; i < l.Len(); i++ {
		l.Value = i
		l = l.Next()
	}
	sum := 0
	l.Do(func(v interface{}){ // Do()方法对环中的每个元素执行f方法
		fmt.Printf("%d ", v.(int)) // 0 1 2 3 4
		sum += v.(int)
	})
	fmt.Printf("\nsum:%d\n", sum) // 10
}
```
List 的零值是一个只包含了根元素（root），但不包含任何实际元素值的空链表。那么，既然 Ring 和 List 的本质上都是循环链表，它们到底有什么不同呢？

Ring 和 List 的不同有以下几种：
* Ring 类型的数据结构仅由它自身即可代表，而 List 类型则需要由它以及 Element 类型联合表示。这是表示方式上的不同，也是结构复杂度上的不同。
* 一个 Ring 类型的值严格来讲，只代表了其所属的循环链表中的一个元素，而一个 List 类型的值则代表了一个完整的链表。这是表示维度上的不同。
* 在创建并初始化一个 Ring 值的时候，我们可以指定它包含的元素数量，但是对于一个 List 值来说却不能这样做(也没必要这样做)。循环链表一旦被创建，其长度是不可变的。这是两个代码包中 New 函数在功能上的不同，也是两个类型在初始化值方面的第一个不同
* 仅通过 `var r ring.Ring` 语句声明的 r 将会是一个长度为 1 的循环链表，而 List 类型的零值则是一个长度为 0 的链表。别忘了，List 中的根元素(root)不会持有实际元素的值，因此计算长度时不会包含它。这是两个类型在初始化值方面的第二个不同。
* Ring 值的 Len 方法的算法复杂度是 O(N) 的，而 List 值的 Len 方法的算法复杂度是 O(1)的。这是两者在性能方面最显而易见的差别。

## 5. 总结
个人感觉，`list`比`ring`要更加适用；而`heap`使用起来相比于Java中的`PriorityQueue`要繁琐很多，这是因为目前Go尚不支持泛型导致的，使得我们在使用时不得不根据堆中存放的元素来实现接口。

目前也没有在实际生产代码中使用到这些，暂时留下这些笔记。有待继续深入。

---
参考：
1. https://golang.org/pkg/container/
2. https://time.geekbang.org/column/article/14117
3. list源码：https://golang.org/src/container/list/list.go
4. heap源码：https://golang.org/src/container/heap/heap.go
5. ring源码：https://golang.org/src/container/ring/ring.go
