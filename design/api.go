package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("studiously", func() {
	Title("Studiously API")
	Scheme("http")
	Host("localhost:8080")
})
