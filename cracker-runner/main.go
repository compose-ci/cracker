package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

// store all your plugins in a normal Go hash map, protected by a Mutex
// (reproduce something like the node.js event loop)
// to avoid "memory collision 💥"
var m sync.Mutex
var plugins = make(map[string]*extism.Plugin)

func StorePlugin(plugin *extism.Plugin) {
	plugins["code"] = plugin
}

func GetPlugin() (extism.Plugin, error) {
	if plugin, ok := plugins["code"]; ok {
		return *plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}

func GetBytesBody(request *http.Request) []byte {
	body := make([]byte, request.ContentLength)
	request.Body.Read(body)
	return body
}

func main() {

	// test the number of arguments
	if len(os.Args) < 3 {
		log.Println("👋 Cracker Runner Demo 🚀")
		os.Exit(0)
	}

	wasmFilePath := os.Args[1:][0]
	wasmFunctionName := os.Args[1:][1]

	//httpPort := os.Args[1:][2]
	httpPort := "8080" // Default value
	if len(os.Args) > 3 {
		httpPort = os.Args[3]
	}

	ctx := context.Background()

	config := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
		AllowedHosts: []string{"*"},
		Config:       map[string]string{},
	}

	pluginInst, err := extism.NewPlugin(ctx, manifest, config, nil) // new
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	StorePlugin(pluginInst)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /", func(response http.ResponseWriter, request *http.Request) {

		params := GetBytesBody(request)
		// unmarshal the json data
		//var data map[string]string

		//err := json.Unmarshal(body, &data)
		//if err != nil {
		//	response.Write([]byte("😡 Error: " + err.Error()))
		//}

		//model := data["model"]
		//systemContent := data["system"]
		//userContent := data["user"]
		m.Lock()
		// don't forget to release the lock on the Mutex
		defer m.Unlock()

		pluginInst, err := GetPlugin()

		if err != nil {
			log.Println("🔴 !!! Error when getting the plugin", err)
			response.Write([]byte("😡 Error: " + err.Error()))

		}

		_, out, err := pluginInst.Call(wasmFunctionName, params)

		if err != nil {
			fmt.Println(err)
			response.Write([]byte("😡 Error: " + err.Error()))

		} else {
			//c.Status(http.StatusOK)
			response.Write(out)

			//return c.SendString(string(out))
		}

	})

	var errListening error
	log.Println("🌍 http server is listening on: " + httpPort)
	errListening = http.ListenAndServe(":"+httpPort, mux)

	log.Fatal(errListening)
}
