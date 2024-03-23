package ethabi

import (
	"github.com/pkg/errors"
	"github.com/wavesplatform/gowaves/pkg/ride/meta"
)

type Argument struct {
	Name string
	Type Type
}

type Arguments []Argument

func NewArgumentFromRideTypeMeta(name string, rideT meta.Type) (Argument, error) {
	t, err := AbiTypeFromRideTypeMeta(rideT)
	if err != nil {
		return Argument{}, errors.Wrapf(err,
			"failed to build ABI argument with name %q from ride type metadata", name,
		)
	}
	arg := Argument{
		Name: name,
		Type: t,
	}
	return arg, err
}

// UnpackRideValues can be used to unpack ABI-encoded hexdata according to the ABI-specification,
// without supplying a struct to unpack into. Instead, this method returns a list containing the
// values. An atomic argument will be a list with one element.
func (arguments Arguments) UnpackRideValues(output []byte) (_ []DataType, paymentsSliceIndex, slotsReadTotal int, _ error) {
	retval := make([]DataType, 0, len(arguments))
	virtualArgs := 0
	for index, arg := range arguments {
		marshalledValue, slotsRead, err := toDataType((index+virtualArgs)*abiSlotSize, arg.Type, output)
		if arg.Type.T == TupleType && !isDynamicType(arg.Type) {
			// If we have a static tuple, like (uint256, bool, uint256), these are
			// coded as uint256,bool,uint256
			tupleSize := getTypeSize(arg.Type)/abiSlotSize - 1
			virtualArgs += tupleSize
		}
		if err != nil {
			return nil, 0, 0, err
		}
		retval = append(retval, marshalledValue)
		slotsReadTotal += slotsRead
	}
	paymentsSliceIndex = (len(arguments) + virtualArgs) * abiSlotSize
	return retval, paymentsSliceIndex, slotsReadTotal, nil
}
