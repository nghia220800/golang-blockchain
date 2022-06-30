package blockchain

import "crypto/sha256"

type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

type MerkleTree struct {
	RootNode *MerkleNode
}

func NewMerkleNode(l, r *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if l == nil && r == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		preH := append(l.Data, r.Data...)
		hash := sha256.Sum256(preH)
		node.Data = hash[:]
	}

	node.Left = l
	node.Right = r

	return &node
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var node_set []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len((data))-1])
	}

	for _, d := range data {
		node := NewMerkleNode(nil, nil, d)
		node_set = append(node_set, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var level []MerkleNode

		for j := 0; j < len(node_set); j += 2 {
			node := NewMerkleNode(&node_set[j], &node_set[j+1], nil)
			level = append(level, *node)
		}
		node_set = level
	}
	tree := MerkleTree{&node_set[0]}
	return &tree
}
