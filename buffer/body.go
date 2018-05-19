package buffer

import (
	"errors"
)

//Buffer storage some information
type Buffer struct {
	Result             []int64
	stateResult        []bool //maybe, need to use as array enum
	GetNextIndex       func() int32
	ConfirmIndexRutine func(index int32) error
	CompactResultQuery func()
	// IsReadyResult      func(bufferId int32)
	Clear func(bufferId int32)
}

func GetDefaultBeffer() *Buffer {
	var linkObject Buffer
	linkObject.GetNextIndex = func() int32 { //this method can made some thread/routine race
		for index := 0; index < len(linkObject.stateResult); index++ {
			if !linkObject.stateResult[index] {
				return int32(index)
			}
		}
		linkObject.stateResult = append(linkObject.stateResult, false)
		return int32(len(linkObject.stateResult) - 1)
	}

	linkObject.ConfirmIndexRutine = func(index int32) error {
		if linkObject.stateResult[index] {
			return errors.New("somebody get this index first")
		}
		linkObject.stateResult[index] = true
		if index > int32(len(linkObject.Result)) {
			linkObject.Result = append(linkObject.Result, 0)
			return nil
		}
		linkObject.Result[index] = 0
		return nil
	}

	linkObject.CompactResultQuery = func() {
		//how can i reduce size of arrays. need to think
	}

	// linkObject.IsReadyResult = func(bufferId int32)
	// 	if decide use stateResult as array enum, this method can come in handy
	// }

	linkObject.Clear = func(bufferId int32) {
		linkObject.stateResult[bufferId] = false
		linkObject.Result[bufferId] = 0
	}

	return &linkObject
}
