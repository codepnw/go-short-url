package router

import (
	"net/http"

	"github.com/codepnw/go-short-url/internal/constant"
	"github.com/codepnw/go-short-url/internal/controllers"
)

var urlShortner = Routes{
	Route{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controllers.ShortTheUrl},
	Route{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controllers.RedirectURL},
}
