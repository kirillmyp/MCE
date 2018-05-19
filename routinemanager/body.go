package routinemanager

import (
	"errors"

	"../buffer"
	"../workerwatcher"
)

//RoutineManager have to controll all processes. it's most critical node
type RoutineManager struct {
	buffer              *buffer.Buffer
	workerWatchers      []*workerwatcher.WorkerWatcher
	AddNewWorkerWatcher func() bool
	//need to think about timer for worker
	CheckResult func(worker int32) bool
	GetResult   func(worker int32) (int64, error)
	KillWorker  func(work int32) (bool, error)
}

//GetDefaultRoutineManager return a filled link on the instance Worker
func GetDefaultRoutineManager() *RoutineManager {
	var linkObject RoutineManager
	linkObject.AddNewWorkerWatcher = func() bool {
		if len(linkObject.workerWatchers) > 0 {
			linkObject.workerWatchers = append(linkObject.workerWatchers, linkObject.workerWatchers[0].Copy())
			return true
		} else {
			linkObject.workerWatchers = []*workerwatcher.WorkerWatcher{workerwatcher.GetDefaultRoutineManager()}
			return true
		}
		//return false
	}

	linkObject.CheckResult = func(worker int32) bool {
		result, _ := linkObject.workerWatchers[worker].CheckResult(worker)
		if result {
			result, err := linkObject.workerWatchers[worker].GetResult(worker)
			if err != nil {
				linkObject.buffer.Result[worker] = result
				return true
			}
		} else {

		}
		return false
	}

	linkObject.GetResult = func(worker int32) (int64, error) {
		result, err := linkObject.workerWatchers[worker].CheckResult(worker)
		if result {
			result, err := linkObject.workerWatchers[worker].GetResult(worker)
			if err != nil {
				linkObject.buffer.Result[worker] = result
				return linkObject.buffer.Result[worker], nil
			}
		} else {
			if err == nil {
				return 0, err
			}
		}
		return 0, errors.New("something wrong")
	}

	linkObject.KillWorker = func(int32) (bool, error) {
		return false, nil
	}

	return &linkObject
}
