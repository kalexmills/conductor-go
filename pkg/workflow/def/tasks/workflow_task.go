package tasks

func SimpleTask(name string, taskRefName string) *simpleTask {
	return &simpleTask{
		task{
			Name:              name,
			TaskReferenceName: taskRefName,
			Type:              SIMPLE,
			Optional:          false,
			InputParameters:   struct{}{},
		},
	}
}
func Switch(taskRefName string, caseExpression string) *decision {
	return &decision{
		task:           task{},
		decisionCases:  map[string][]Task{},
		defaultCase:    []Task{},
		caseExpression: caseExpression,
		useJavascript:  false,
	}
}

func example() {
	decision := Switch("shipping", "${workflow.input.shipping")
	decision.SwitchCase("Ground",
		SimpleTask("ship", "ship"),
		SimpleTask("wait_shipping", "wait_shipping"))
	decision.SwitchCase("Air", SimpleTask("air", "air"))
}
