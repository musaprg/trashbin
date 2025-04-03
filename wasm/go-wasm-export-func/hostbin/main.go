package main

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

//go:embed reactor/reactor.wasm
var reactorWasm []byte

//go:wasmimport

func main() {
	ctx := context.Background()
	{

		// Create a Wasm runtime, set up WASI.
		r := wazero.NewRuntime(ctx)
		defer r.Close(ctx)
		wasi_snapshot_preview1.MustInstantiate(ctx, r)

		// Configure the module to initialize the reactor.
		config := wazero.NewModuleConfig().WithStartFunctions("_initialize")

		// Instantiate the module.
		wasmModule, _ := r.InstantiateWithConfig(ctx, reactorWasm, config)

		// Call the exported function.
		fn := wasmModule.ExportedFunction("add")
		var a, b int32 = 1, 2
		res, _ := fn.Call(ctx, api.EncodeI32(a), api.EncodeI32(b))
		c := api.DecodeI32(res[0])
		fmt.Printf("add(%d, %d) = %d\n", a, b, c)

		// The instance is still alive. We can call the function again.
		res, _ = fn.Call(ctx, api.EncodeI32(b), api.EncodeI32(c))
		fmt.Printf("add(%d, %d) = %d\n", b, c, api.DecodeI32(res[0]))
	}

	{
		// Create a Wasm runtime, set up WASI.
		r := wazero.NewRuntime(ctx)
		defer r.Close(ctx)
		wasi_snapshot_preview1.MustInstantiate(ctx, r)

		// Configure the module to initialize the command module.
		// Default to "_start" function, so we don't need to set WithStartFunctions.
		config := wazero.NewModuleConfig().WithStartFunctions("_initialize")
		// Instantiate the module.
		compiled, err := r.CompileModule(ctx, reactorWasm)
		if err != nil {
			panic(err)
		}
		wasmModule, _ := r.InstantiateModule(ctx, compiled, config)

		// Call the exported function.
		fn := wasmModule.ExportedFunction("execute")
		var a, b int32 = 1, 2
		res, err := fn.Call(ctx, api.EncodeI32(a), api.EncodeI32(b))
		if err != nil {
			fmt.Printf("Error calling execute: %v\n", err)
			return
		}
		c := api.DecodeI32(res[0])
		fmt.Printf("add(%d, %d) = %d\n", a, b, c)
	}
}
