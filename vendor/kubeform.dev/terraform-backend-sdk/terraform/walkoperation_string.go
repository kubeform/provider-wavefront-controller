// Code generated by "stringer -type=walkOperation graph_walk_operation.go"; DO NOT EDIT.

package terraform

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[walkInvalid-0]
	_ = x[walkApply-1]
	_ = x[walkPlan-2]
	_ = x[walkPlanDestroy-3]
	_ = x[walkValidate-4]
	_ = x[walkDestroy-5]
	_ = x[walkImport-6]
	_ = x[walkEval-7]
}

const _walkOperation_name = "walkInvalidwalkApplywalkPlanwalkPlanDestroywalkValidatewalkDestroywalkImportwalkEval"

var _walkOperation_index = [...]uint8{0, 11, 20, 28, 43, 55, 66, 76, 84}

func (i walkOperation) String() string {
	if i >= walkOperation(len(_walkOperation_index)-1) {
		return "walkOperation(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _walkOperation_name[_walkOperation_index[i]:_walkOperation_index[i+1]]
}
