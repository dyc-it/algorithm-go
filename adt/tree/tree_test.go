package tree

import "testing"


// The structure of example tree
//               A
//          B         C
//       D     E          F
//      G H

var nodeG Tree = Tree{nil, nil, "G"}
var nodeH Tree = Tree{nil, nil, "H"}
var nodeF Tree = Tree{nil, nil, "F"}

var nodeD Tree = Tree{&nodeG, &nodeH, "D"}
var nodeE Tree = Tree{nil, nil, "E"}

var nodeB Tree = Tree{&nodeD, &nodeE, "B"}

var nodeC Tree = Tree{nil, &nodeF, "C"}

var nodeA Tree = Tree{&nodeB, &nodeC, "A"}

func TestPre_order(t *testing.T) {
	nodeA.pre_order()
}

func TestIn_order(t *testing.T) {
	nodeA.in_order()
}

func TestPost_order(t *testing.T) {
	nodeA.post_order()
}

func TestLevel_order(t *testing.T) {
	nodeA.level_order()
}


