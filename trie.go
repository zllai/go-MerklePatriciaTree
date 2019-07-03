package mpt

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

type Trie struct {
	oldRoot []byte
	root    Node
	kv      KVStore
	lock    *sync.RWMutex
}

func New(root Node, kv KVStore) *Trie {
	var oldRoot []byte = nil
	if root != nil {
		root.Serialize() // update cached hash
		oldRoot = root.CachedHash()
	}
	return &Trie{
		oldRoot: oldRoot,
		root:    root,
		kv:      kv,
		lock:    &sync.RWMutex{},
	}
}

func (t *Trie) Get(key []byte) ([]byte, error) {
	node, expandedNode, err := t.get(t.root, key, 0)
	if expandedNode != nil {
		t.root = expandedNode
	}
	if err != nil {
		return nil, err
	} else if v, ok := node.(*ValueNode); ok {
		return []byte(v.Value), nil
	} else {
		if _, ok := node.(*HashNode); ok {
			fmt.Println("get a hash node")
		}
		return nil, errors.New(fmt.Sprintf("[Trie] key not found 1: %s", hex.EncodeToString(key)))
	}
}

func (t *Trie) get(node Node, key []byte, prefixLen int) (Node, Node, error) {
	if node == nil {
		return nil, node, errors.New(fmt.Sprintf("[Trie] key not found 2: %s", hex.EncodeToString(key)))
	}
	switch n := node.(type) {
	case *FullNode:
		if prefixLen > len(key) {
			return nil, node, errors.New(fmt.Sprintf("[Trie] key not found 3: %s", hex.EncodeToString(key)))
		}
		if prefixLen == len(key) {
			valueNode, newNode, err := t.get(n.Children[256], key, prefixLen)
			n.Children[256] = newNode
			return valueNode, node, err
		} else {
			valueNode, newNode, err := t.get(n.Children[key[prefixLen]], key, prefixLen+1)
			n.Children[key[prefixLen]] = newNode
			return valueNode, node, err
		}
	case *ShortNode:
		if len(key)-prefixLen < len(n.Key) || !bytes.Equal(n.Key, key[prefixLen:prefixLen+len(n.Key)]) {
			return nil, node, errors.New(fmt.Sprintf("[Trie] key not found 4: %s", hex.EncodeToString(key)))
		}
		valueNode, newNode, err := t.get(n.Value, key, prefixLen+len(n.Key))
		n.Value = newNode
		return valueNode, node, err
	case *HashNode:
		data, err := t.kv.Get([]byte(*n))
		if err != nil {
			return nil, node, err
		}
		loadedNode, err := DeserializeNode(data)
		if !bytes.Equal([]byte(*n), loadedNode.Hash()) {
			return nil, node, errors.New(fmt.Sprintf("[Trie] Cannot load node 5: hash does not match"))
		}
		valueNode, loadedNode, err := t.get(loadedNode, key, prefixLen)
		return valueNode, loadedNode, err
	case *ValueNode:
		if prefixLen == len(key) {
			return node, node, nil
		} else {
			return nil, node, errors.New(fmt.Sprintf("[Trie] key not found 4: %s", hex.EncodeToString(key)))
		}
	}
	return nil, node, errors.New("[Tire] Unknown node type")
}

func (t *Trie) Put(key, value []byte) error {
	valueNode := ValueNode{value, nil, true}
	expandedNode, err := t.put(t.root, key, &valueNode, 0)
	if expandedNode != nil {
		t.root = expandedNode
	}
	return err
}

func (t *Trie) put(node Node, key []byte, value Node, prefixLen int) (Node, error) {
	if node == nil {
		if prefixLen > len(key) {
			return node, errors.New("[Trie] Cannot insert: 1")
		} else if prefixLen == len(key) {
			return value, nil
		} else {
			shortNode := ShortNode{
				Key:   key[prefixLen:],
				Value: value,
				dirty: true,
			}
			return &shortNode, nil
		}
	}
	switch n := node.(type) {
	case *FullNode:
		n.dirty = true
		if prefixLen > len(key) {
			return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 2"))
		} else if prefixLen == len(key) {
			n.Children[256] = value
			return n, nil
		}
		// prefixLen < len(key)
		newNode, err := t.put(n.Children[key[prefixLen]], key, value, prefixLen+1)
		if err != nil {
			return node, err
		}
		n.Children[key[prefixLen]] = newNode
		return n, err
	case *ShortNode:
		n.dirty = true
		if prefixLen > len(key) {
			return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 3"))
		}
		commonLen := commonPrefix(n.Key, key[prefixLen:])
		if commonLen == len(n.Key) {
			newNode, err := t.put(n.Value, key, value, prefixLen+len(n.Key))
			if err != nil {
				return node, err
			}
			n.Value = newNode
			return n, nil
		}
		prefixLen += commonLen
		fullNode := &FullNode{dirty: true}
		newNode, err := t.put(fullNode, key, value, prefixLen)
		if err != nil {
			return node, err
		}
		newNode, err = t.put(newNode, n.Key, n.Value, commonLen)
		if err != nil {
			return node, err
		}
		if commonLen > 0 {
			shortNode := ShortNode{dirty: true}
			shortNode.Key = n.Key[:commonLen]
			shortNode.Value = newNode
			return &shortNode, nil
		} else {
			return newNode, nil
		}
	case *ValueNode:
		n.dirty = true
		if prefixLen == len(key) {
			return value, nil
		} else if prefixLen < len(key) {
			fullNode := &FullNode{dirty: true}
			newNode, err := t.put(fullNode, key, value, prefixLen)
			if err != nil {
				return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 3.1"))
			}
			newNode, err = t.put(newNode, key[:prefixLen], node, prefixLen)
			if err != nil {
				return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 3.1"))
			}
			return newNode, nil
		} else {
			return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 4"))
		}
	case *HashNode:
		if prefixLen >= len(key) {
			return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 5"))
		}
		data, err := t.Get([]byte(*n))
		if err != nil {
			return node, err
		}
		newNode, err := DeserializeNode(data)
		if err != nil {
			return node, err
		}
		newNode, err = t.put(newNode, key, value, prefixLen)
		if err != nil {
			return node, err
		}
		return newNode, nil
	}
	return node, errors.New(fmt.Sprintf("[Trie] Cannot insert: 6"))
}

func commonPrefix(a, b []byte) int {
	minLen := len(a)
	if len(b) < len(a) {
		minLen = len(b)
	}
	ret := 0
	for i := 0; i < minLen; i++ {
		if a[i] == b[i] {
			ret++
		} else {
			break
		}
	}
	return ret
}

func (t *Trie) Commit() {
	t.commit(t.root)
	t.oldRoot = t.root.CachedHash()
}

func (t *Trie) commit(node Node) {
	switch n := node.(type) {
	case *FullNode:
		for i := 0; i < len(n.Children); i++ {
			t.commit(n.Children[i])
		}
		n.Save(t.kv)
	case *ShortNode:
		t.commit(n.Value)
		n.Save(t.kv)
	case *ValueNode:
		n.Save(t.kv)
	}
}

func (t *Trie) Abort() {
	if t.oldRoot == nil {
		t.root = nil
	} else {
		hashNode := HashNode(t.oldRoot)
		t.root = &hashNode
	}
}

func (t *Trie) RootHash() []byte {
	return t.root.Hash()
}
