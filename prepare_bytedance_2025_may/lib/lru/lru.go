package lru

type LruCacheValue interface {
	ToInt() (int, bool)
	ToString() (string, bool)
}

type lruCacheValueImpl struct {
	intV    int
	stringV string
	isInt   bool
	isNil   bool
}

var nilLruCacheValue = lruCacheValueImpl{isNil: true}

func (r lruCacheValueImpl) ToInt() (int, bool) {
	if r.isNil {
		return 0, false
	}
	if r.isInt {
		return r.intV, true
	}
	return 0, false
}

func (r lruCacheValueImpl) ToString() (string, bool) {
	if r.isNil {
		return "", false
	}
	if !r.isInt {
		return r.stringV, true
	}
	return "", false
}

func NewStringValue(value string) LruCacheValue {
	return &lruCacheValueImpl{
		stringV: value,
		isInt:   false,
	}
}

func NewIntValue(value int) LruCacheValue {
	return &lruCacheValueImpl{
		intV:  value,
		isInt: true,
	}
}

type linkNode struct {
	Prev  *linkNode
	Next  *linkNode
	Key   string
	Value LruCacheValue
}

type LruCache interface {
	Get(key string) LruCacheValue
	Put(key string, value LruCacheValue)
	Monitor() (int, int)
}

type lruCacheImpl struct {
	store    map[string]*linkNode
	listHead *linkNode
	listTail *linkNode
	cap      int
	length   int
}

func NewLruCache(cap int) LruCache {
	return &lruCacheImpl{
		store:    make(map[string]*linkNode),
		listHead: nil,
		listTail: nil,
		cap:      cap,
		length:   0,
	}
}

func (l *lruCacheImpl) putNodeToHead(node *linkNode) {
	if l.listHead == nil {
		l.listHead = node
		l.listTail = node
		node.Prev = nil
		node.Next = nil
		return
	}

	if node != l.listHead {
		if node.Prev != nil {
			node.Prev.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Prev = node.Prev
		}

		if node == l.listTail {
			l.listTail = node.Prev
		}

		node.Next = l.listHead
		node.Prev = nil
		l.listHead.Prev = node
		l.listHead = node
	}
}

func (l *lruCacheImpl) Get(key string) LruCacheValue {
	node, ok := l.store[key]
	if !ok {
		return nilLruCacheValue
	}

	l.putNodeToHead(node)

	return node.Value
}

func (l *lruCacheImpl) Put(key string, value LruCacheValue) {
	node, ok := l.store[key]
	if ok {
		node.Value = value
		l.putNodeToHead(node)
		return
	}

	if l.length >= l.cap && l.listTail != nil {
		delKey := l.listTail.Key
		delete(l.store, delKey)

		if l.listTail.Prev != nil {
			l.listTail.Prev.Next = nil
		} else {
			// 被淘汰的是唯一节点
			l.listHead = nil
		}
		l.listTail = l.listTail.Prev
		l.length--
	}

	node = &linkNode{
		Key:   key,
		Value: value,
	}
	l.store[key] = node
	l.putNodeToHead(node)
	l.length++
}

func (l *lruCacheImpl) Monitor() (int, int) {
	return l.cap, l.length
}

var defaultLruCache = NewLruCache(100)

func Get(key string) LruCacheValue {
	return defaultLruCache.Get(key)
}

func Put(key string, value LruCacheValue) {
	defaultLruCache.Put(key, value)
}

func Monitor() (int, int) {
	return defaultLruCache.Monitor()
}
