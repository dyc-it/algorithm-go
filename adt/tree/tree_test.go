package tree

import "testing"


// The structure of example tree
//               A
//          B         C
//       D     E          F
//            G H        I

var nodeG Tree = Tree{nil, nil, "G"}
var nodeH Tree = Tree{nil, nil, "H"}
var nodeI Tree = Tree{nil, nil, "I"}

var nodeF Tree = Tree{&nodeI, nil, "F"}

var nodeD Tree =  Tree{nil, nil,"D"}
var nodeE Tree =  Tree{&nodeG, &nodeH,"E"}

var nodeB Tree = Tree{&nodeD, &nodeE, "B"}

var nodeC Tree = Tree{nil, &nodeF, "C"}

var nodeA Tree = Tree{&nodeB, &nodeC, "A"}

func TestPre_order(t *testing.T) {
	nodeA.Pre_order()
}

func TestIn_order(t *testing.T) {
	nodeA.In_order()
}

func TestPost_order(t *testing.T) {
	nodeA.Post_order()
}

func TestLevel_order(t *testing.T) {
	nodeA.Level_order()
}


