# Go-MerklePatriciaTree
This is a go implementation of Merkle Patricia Tree (MPT) similar to that used in go-Ethereum. The reason for re-implementing it is to ease the adoption in other projects, as the MPT in go-Ethereum is coupled with other Ethereum modules and its interface is not clean.

For the detailed explanation of MPT, please refer to [Ethereum wiki](https://github.com/ethereum/wiki/wiki/Patricia-Treel).


# Usage
The usage of this library is simple.
#### 1. Create a key-value store as backend
MPT uses a key-value store as backend to store the tree nodes. Therefore, a key-value store instance is required before creating an MPT instance.

This library provides two implementaion of key-value store in ./kvstore
1. MemKVStore: essentially a go map which stores key-value pairs in memory
2. LevelDB: Google leveldb that persist key-value pairs in disk

```
//Creating a MemKVStore instance
memKv := kvstore.NewMemKVStore()

//Creating a LevelDB instance
leveldb, err := kvstore.NewLevelDB("path/to/leveldb/file")
```
#### 2. Create an MPT instance
```
// Creating an empty MPT
tree := mpt.New(nil, leveldb)

// Create an MPT from existing tree node
tree := mpt.New(node, leveldb)

// Creating an MPT from a known root hash
tree := mpt.New(mpt.HashNode(root_hash), leveldb)
```
Now you are good to go!
```
// Put and get
tree.Put([]byte("A"), []byte("a"))
tree.Get([]byte("A"))
```

#### 3. Commit and abort
Any put operation won't be persisted in key-value store backend until you call:
```
tree.Commit()
```
If you want to cancel all the operations until last commit:
```
tree.Abort()
```

#### 4. Serialize and deserialize
```
// serialize the mpt
data, err := tree.Serialize()

// deserialze the mpt
tree := mpt.New(nil, leveldb)
tree.Deserialize(data)
```
