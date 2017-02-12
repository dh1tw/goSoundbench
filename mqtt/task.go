package mqttserver

import (
	"log"

	test "github.com/dh1tw/goSoundbench/test"
)

// ToWireMsg contains the topic and data to be send to the MQTT Broker
type ToWireMsg struct {
	Topic string
	Data  []byte
}

// Task contains the details of the task
type Task struct {
	BaseTopic string
	Data      []byte
	ResultCh  chan<- ToWireMsg
	ErrorCh   chan<- ToWireMsg
}

// Run executes a task
func (task *Task) Run() {

	TCs, err := test.DeserializeTestCases(task.Data)
	if err != nil {
		log.Println(err)
	}

	err = task.RunTest(TCs)
	if err != nil {
		log.Println(err)
	}
}

// RunTest runs all tests in testQueue which implement the TestcaseI Interface
func (task *Task) RunTest(testQueue []test.TestcaseI) error {

	// errorTopic := task.ItemName + "/error"
	resultTopic := task.BaseTopic + "/result"

	for _, test := range testQueue {

		defer test.Cleanup()

		if err := test.Setup(); err != nil {
			return err
		}

		if err := test.Execute(); err != nil {
			return err
		}

		res, err := test.GetResult()
		if err != nil {
			return err
		}

		// only send in case the Results are not empty
		if res != nil {
			data, err := res.Serialize()
			if err != nil {
				return err
			}

			msg := ToWireMsg{
				Data:  data,
				Topic: resultTopic,
			}

			task.ResultCh <- msg
			log.Println(res.String())
		}
	}

	return nil
}
