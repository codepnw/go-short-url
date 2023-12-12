package constant

const (
	// Path
	UrlShortnerPath = "/short"
	RedirectUrlPath = "/url/:code"

	// Database
	Database = "urlshortner"
	UrlCollection = "url"

	// Base URL
	BaseUrl = "localhost:8000/"

	// Error
	BindError = "there is something wrong while binding"
	ErrCodeInUse = "this code is already in use"
	ErrNoURL = "there is no url found"
)
