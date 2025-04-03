package main

func main() {}
func init() {
	setExecute(func(a, b int32) int32 {
		return a + b
	})
}

type executeFunc = func(a, b int32) int32

var execute executeFunc

//go:wasmexport execute
func _execute(a, b int32) int32 {
	if execute == nil {
		panic("execute function is not set")
	}
	res := execute(a, b)

	return res
}

func setExecute(fn executeFunc) {
	execute = fn
	if execute == nil {
		panic("execute function is nil")
	}
}

//go:wasmexport add
func add(a, b int32) int32 { return a + b }
