package main

type Trie struct {
	//用于承接子节点
	children [26]*Trie
	//用于判断是否结束
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		node = node.children[index]
	}
	node.isEnd = true

}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	if node.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true

}
