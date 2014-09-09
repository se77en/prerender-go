package prerender

import (
	"net/http"
)

type EventHandler interface {
	BeforeRender(req *http.Request) string
	AfterRender(req *http.Request, res http.ResponseWriter, prerenderRes http.ResponseWriter, resHtml string) string
	Destroy()
}
