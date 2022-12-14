package targetstack

import (
	"bitbucket.org/exotel/ent-leadassist-test/baremetrix/internal/models"
)

type FlowStack []models.TargetParameters

// Push adds a target on top of the FlowStack.
func (s FlowStack) Push(target models.TargetParameters) FlowStack {
	return append(s, target)
}

// Pop removes a target from the top of the FlowStack.
func (s FlowStack) Pop() (FlowStack, models.TargetParameters) {
	l := len(s)
	return s[:l-1], s[l-1]
}

// IsEmpty checks whether stack is empty.
func (s FlowStack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}
