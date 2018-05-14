package worker

import (
	"strconv"
	"strings"

	"../buffer"
)

type state int32
type calString string

//Worker - worker without functions realization
type Worker struct {
	resultID     int32
	currentState state
	// isReady            bool
	// panicEnd           bool
	processID          int32
	Do                 func(camputedString string)
	corector           func(camputedString string) bool
	retuningResultByID func() int32
	calculateResult    func(camputedString string)
	IsFinished         func() bool
	bufferRef          *buffer.Buffer
}

//GetDefaultWorker fulling instance Worker and return a link on it
func GetDefaultWorker(resultID int32) *Worker {
	var linkObject Worker

	linkObject.resultID = resultID
	linkObject.Do = func(camputedString string) {
		linkObject.calculateResult(camputedString)
	}

	linkObject.corector = func(camputedString string) bool {
		//need to add others checking
		if camputedString == "" {
			linkObject.currentState = panicEnd
			return false
		}
		return true
	}
	linkObject.retuningResultByID = func() int32 {
		if linkObject.currentState == runing {
			//return result which storage into a buffer
			return 0
		} else {
			return 0
		}
	}

	linkObject.calculateResult = func(camputedString string) {
		if linkObject.corector(camputedString) {
			linkObject.currentState = crashed
			splitStrings := calString(camputedString).split()
			if len(splitStrings) != 3 {
				linkObject.currentState = panicEnd
				return
			}
			first, errFirst := strconv.ParseInt(splitStrings[0], 0, 32)
			second, errSecond := strconv.ParseInt(splitStrings[2], 0, 32)
			if errFirst != nil || errSecond != nil {
				linkObject.currentState = panicEnd
				return
			}

			switch splitStrings[1] {
			case "+":
				{
					linkObject.bufferRef.Result[linkObject.resultID] = first + second
				}

			case "-":
				{
					linkObject.bufferRef.Result[linkObject.resultID] = first - second
				}

			case "*":
				{
					linkObject.bufferRef.Result[linkObject.resultID] = first * second
				}

			case "/":
				{
					linkObject.bufferRef.Result[linkObject.resultID] = first / second
				}
			}

		}
		linkObject.currentState = finished
	}

	linkObject.IsFinished = func() bool {
		return linkObject.currentState == finished
	}

	linkObject.currentState = sleep

	return &linkObject
}

const (
	sleep    state = 0
	runing   state = 1
	finished state = 2
	panicEnd state = 3
	crashed  state = 4
)

func (splitString calString) split() []string { //[]calString
	//need to try Bytecode pattern
	//var arrayString []calString
	// for index := 0; index < len(splitString); index++ {
	//must to write own split or use strings.FieldsFunc
	// }
	return []string(strings.Fields(string(splitString)))
	// return arrayString
}

/*
func (worker Worker) do(camputedString string) {
	worker.do(camputedString)
}

func (worker Worker) corector(camputedString string) bool {
	if camputedString == "" {
		return false
	}
	return false
}

func (worker Worker) retuningResultByID() int32 {
	if worker.isReady {
		//return result wich storage into buffer
		return 0
	} else {
		return 0
	}
}

func (worker Worker) calculateResult(camputedString string) {
	if worker.corector(camputedString) {

	}
}

func (worker Worker) isFinished() bool {
	return worker.isReady
}
*/
