//go:generate go run assets_generate.go

package runtime

import "github.com/CarlosMontilla/femto"

var Files = femto.NewRuntimeFiles(files)
