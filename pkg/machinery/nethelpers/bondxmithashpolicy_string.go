// Code generated by "stringer -type=BondXmitHashPolicy -linecomment"; DO NOT EDIT.

package nethelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BondXmitPolicyLayer2-0]
	_ = x[BondXmitPolicyLayer34-1]
	_ = x[BondXmitPolicyLayer23-2]
	_ = x[BondXmitPolicyEncap23-3]
	_ = x[BondXmitPolicyEncap34-4]
}

const _BondXmitHashPolicy_name = "layer2layer3+4layer2+3encap2+3encap3+4"

var _BondXmitHashPolicy_index = [...]uint8{0, 6, 14, 22, 30, 38}

func (i BondXmitHashPolicy) String() string {
	if i >= BondXmitHashPolicy(len(_BondXmitHashPolicy_index)-1) {
		return "BondXmitHashPolicy(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BondXmitHashPolicy_name[_BondXmitHashPolicy_index[i]:_BondXmitHashPolicy_index[i+1]]
}