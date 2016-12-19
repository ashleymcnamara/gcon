package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/markbates/buffalo/render"
	"github.com/markbates/buffalo/render/resolvers"
)

var r *render.Engine
var resolver = &resolvers.RiceBox{
	Box: rice.MustFindBox("../templates"),
}

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "public/main.html",
		CacheTemplates: ENV == "production",
		FileResolver:   resolver,
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../assets")
	return box.HTTPBox()
}
