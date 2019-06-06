package main

// Block describing try/catch functions
type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception object
type Exception interface{}

// Throw is a set of instructions of what to do when the exception is caught
func Throw(e Exception) {
	panic(e)
}

// Do executes the function set
func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
