package gotask

import (
	"errors"
)

var (
	ErrTaskRunning error = errors.New("task already running")
)

// Task Struct for one task to handle inside the worker
type Task struct {
	name     string
	state    State
	progress Progress
	weight   Weight
	target   func()
	desc     string
}

// NewTask Factory method for creating a new task for proper initialition
func NewTask(name string, weight Weight, desc string, target func()) *Task {
	task := Task{
		name:     name,
		state:    Waiting,
		progress: MinProgress,
		weight:   weight,
		desc:     desc,
		target:   target,
	}
	return &task
}

// Run Runs task target function
func (t *Task) Run() {
	t.progress = MinProgress
	t.state = Running
	t.target()
	t.state = Finished
	t.progress = MaxProgress
}

// GetName Returns Task name
func (t *Task) GetName() string {
	return t.name
}

// GetState Returns Task state
func (t *Task) GetState() State {
	return t.state
}

// GetProgress Returns Task progress
func (t *Task) GetProgress() Progress {
	return t.progress
}

// GetWeight Returns current Task weight
func (t *Task) GetWeight() Weight {
	return t.weight
}

// GetDesc Returns task description
func (t *Task) GetDesc() string {
	return t.desc
}

// GetWorkLoad Returns task workload (progress times weight)
func (t *Task) GetWorkLoad() int {
	return int(t.progress) * int(t.weight) / int(MaxProgress)
}

// AddProgress Adds value to current Task Progress until ProgressMaxVal is reached
func (t *Task) Reset() error {
	if t.state == Running {
		return ErrTaskRunning
	}
	t.state = Waiting
	t.progress = MinProgress
	return nil
}
