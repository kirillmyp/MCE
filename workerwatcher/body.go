package workerwatcher

import (
	"errors"

	"../worker"
)

type WorkerWatcher struct {
	workers       []*worker.Worker
	PutToSleep    func(worker int32) (bool, error)
	PutToContinue func(worker int32) (bool, error)
	CheckResult   func(worker int32) (bool, error)
	GetResult     func(worker int32) (int64, error)
}

//GetDefaultRoutineManager return fulling instance WorkerWatcher and return a link on it
func GetDefaultRoutineManager() *WorkerWatcher {
	var linkObject WorkerWatcher
	linkObject.PutToSleep = func(worker int32) (bool, error) {
		if linkObject.workers[worker].Wait() {
			return true, nil
		}
		return false, errors.New("worker doesn't stoped")
	}

	linkObject.PutToContinue = func(worker int32) (bool, error) {
		if linkObject.workers[worker].Continue() {
			return true, nil
		}
		return false, errors.New("worker doesn't continue")
	}

	linkObject.CheckResult = func(worker int32) (bool, error) {
		if linkObject.workers[worker].IsFinished() {
			return true, nil
		}
		return false, errors.New("worker doesn't finished")
	}

	linkObject.GetResult = func(worker int32) (int64, error) {
		result, err := linkObject.workers[worker].GetResult()

		return result, err
		//return 0, errors.New("worker doesn't finished")
	}

	return &linkObject
}

func (workerWatcher *WorkerWatcher) Copy() *WorkerWatcher {
	var linkObject WorkerWatcher
	linkObject.CheckResult = workerWatcher.CheckResult
	linkObject.GetResult = workerWatcher.GetResult
	linkObject.PutToContinue = workerWatcher.PutToContinue
	linkObject.PutToSleep = workerWatcher.PutToSleep
	linkObject.workers = workerWatcher.workers
	return &linkObject
}
