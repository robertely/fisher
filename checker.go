package main

// Checker - checks...
type Checker interface {
	Check() (bool, error)
}
