// Code generated by "stringer -type=nodeKind -linecomment"; DO NOT EDIT.

package art

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[nodeKind4-0]
	_ = x[nodeKind16-1]
	_ = x[nodeKind48-2]
	_ = x[nodeKind256-3]
	_ = x[nodeKindLeaf-4]
}

const _nodeKind_name = "NODE_4NODE_16NODE_48NODE_256NODE_LEAF"

var _nodeKind_index = [...]uint8{0, 6, 13, 20, 28, 37}

func (i nodeKind) String() string {
	if i >= nodeKind(len(_nodeKind_index)-1) {
		return "nodeKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _nodeKind_name[_nodeKind_index[i]:_nodeKind_index[i+1]]
}
