package mpt

import (
	"errors"
	fmt "fmt"

	"github.com/golang/protobuf/proto"
	"github.com/zllai/mpt/kvstore"
	"golang.org/x/crypto/sha3"
)

type (
	Node interface {
		Hash() []byte
		CachedHash() []byte
		Serialize() []byte
		Save(kvstore.KVStore)
	}
	FullNode struct {
		Children [257]Node
		cache    []byte
		dirty    bool
	}
	ShortNode struct {
		Key   []byte
		Value Node
		cache []byte
		dirty bool
	}
	HashNode  []byte
	ValueNode struct {
		Value []byte
		cache []byte
		dirty bool
	}
)

func (n *FullNode) CachedHash() []byte  { return n.cache }
func (n *ShortNode) CachedHash() []byte { return n.cache }
func (n *ValueNode) CachedHash() []byte { return n.cache }
func (n *HashNode) CachedHash() []byte  { return []byte(*n) }

func DeserializeNode(data []byte) (Node, error) {
	persistNode := &PersistNode{}
	err := proto.Unmarshal(data, persistNode)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[Node] cannot deserialize persist node: %s", err.Error()))
	}
	switch v := persistNode.Content.(type) {
	case *PersistNode_Full:
		fullNode := FullNode{}
		for i := 0; i < len(fullNode.Children); i++ {
			if len(v.Full.Children[i]) != 0 {
				child := HashNode(v.Full.Children[i])
				fullNode.Children[i] = &child
				if len([]byte(child)) == 0 {
					return nil, errors.New("[Node] nil full node child")
				}
			}
		}
		hash := sha3.Sum256(data)
		fullNode.cache = hash[:]
		return &fullNode, nil
	case *PersistNode_Short:
		shortNode := ShortNode{}
		shortNode.Key = v.Short.Key
		if len(v.Short.Value) == 0 {
			return nil, errors.New("[Node] nil short node value")
		}
		child := HashNode(v.Short.Value)
		shortNode.Value = &child
		hash := sha3.Sum256(data)
		shortNode.cache = hash[:]
		return &shortNode, nil
	case *PersistNode_Value:
		hash := sha3.Sum256(data)
		ret := ValueNode{v.Value, hash[:], false}
		return &ret, nil
	}
	return nil, errors.New("[Node] Unknown node type")
}

func (vn *ValueNode) Serialize() []byte {
	persistValueNode := PersistNode_Value{}
	persistValueNode.Value = vn.Value
	persistNode := PersistNode{
		Content: &persistValueNode,
	}
	data, _ := proto.Marshal(&persistNode)
	hash := sha3.Sum256(data)
	vn.cache = hash[:]
	vn.dirty = false
	return data
}

func (vn *ValueNode) Hash() []byte {
	if vn.dirty {
		vn.Serialize()
	}
	return vn.cache
}

func (vn *ValueNode) Save(kv kvstore.KVStore) {
	data := vn.Serialize()
	kv.Put(vn.cache, data)
}

func (fn *FullNode) Serialize() []byte {
	persistFullNode := PersistFullNode{}
	persistFullNode.Children = make([][]byte, 257)
	for i := 0; i < len(fn.Children); i++ {
		if fn.Children[i] != nil {
			persistFullNode.Children[i] = fn.Children[i].Hash()
		}
	}
	data, _ := proto.Marshal(&PersistNode{
		Content: &PersistNode_Full{Full: &persistFullNode},
	})
	hash := sha3.Sum256(data)
	fn.cache = hash[:]
	fn.dirty = false
	return data
}

func (fn *FullNode) Hash() []byte {
	if fn.dirty {
		fn.Serialize()
	}
	return fn.cache
}

func (fn *FullNode) Save(kv kvstore.KVStore) {
	data := fn.Serialize()
	kv.Put(fn.cache, data)
}

func (sn *ShortNode) Serialize() []byte {
	persistShortNode := PersistShortNode{}
	persistShortNode.Key = sn.Key
	persistShortNode.Value = sn.Value.Hash()
	data, _ := proto.Marshal(&PersistNode{
		Content: &PersistNode_Short{Short: &persistShortNode},
	})
	hash := sha3.Sum256(data)
	sn.cache = hash[:]
	sn.dirty = false
	return data
}

func (sn *ShortNode) Hash() []byte {
	if sn.dirty {
		sn.Serialize()
	}
	return sn.cache
}

func (sn *ShortNode) Save(kv kvstore.KVStore) {
	data := sn.Serialize()
	kv.Put(sn.cache, data)
}

func (hn *HashNode) Hash() []byte            { return []byte(*hn) }
func (hn *HashNode) Serialize() []byte       { return nil }
func (hn *HashNode) Save(kv kvstore.KVStore) {}
