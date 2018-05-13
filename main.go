package main

import "./worker"

//main. MCE - Managed Calculator Extensibility
func main() {
	var temp *worker.Worker = worker.GetDefaultWorker()
	temp.Do("Some string")
}
