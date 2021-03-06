// Code generated by "stringer -type=ActNoiseType"; DO NOT EDIT.

package leabra

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoNoise-0]
	_ = x[VmNoise-1]
	_ = x[GeNoise-2]
	_ = x[ActNoise-3]
	_ = x[GeMultNoise-4]
	_ = x[ActNoiseTypeN-5]
}

const _ActNoiseType_name = "NoNoiseVmNoiseGeNoiseActNoiseGeMultNoiseActNoiseTypeN"

var _ActNoiseType_index = [...]uint8{0, 7, 14, 21, 29, 40, 53}

func (i ActNoiseType) String() string {
	if i < 0 || i >= ActNoiseType(len(_ActNoiseType_index)-1) {
		return "ActNoiseType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ActNoiseType_name[_ActNoiseType_index[i]:_ActNoiseType_index[i+1]]
}

func (i *ActNoiseType) FromString(s string) error {
	for j := 0; j < len(_ActNoiseType_index)-1; j++ {
		if s == _ActNoiseType_name[_ActNoiseType_index[j]:_ActNoiseType_index[j+1]] {
			*i = ActNoiseType(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ActNoiseType")
}
