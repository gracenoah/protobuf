package typedeclimport

import subpkg "github.com/gracenoah/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
