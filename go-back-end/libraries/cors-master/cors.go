package cors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"libraries/gin-1.10.1"
)

// Config represents all available options for the middleware.
type Config struct {
	AllowAllOrigins bool

	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
	AllowOrigins []string

	// AllowOriginFunc is a custom function to validate the origin. It takes the origin
	// as an argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowOrigins is ignored.
	AllowOriginFunc func(origin string) bool

	// Same as AllowOriginFunc except also receives the full request context.
	// This function should use the context as a read only source and not
	// have any side effects on the request, such as aborting or injecting
	// values on the request.
	AllowOriginWithContextFunc func(c *gin.Context, origin string) bool

	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS)
	AllowMethods []string

	// AllowPrivateNetwork indicates whether the response should include allow private network header
	AllowPrivateNetwork bool

	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// ExposeHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string

	// MaxAge indicates how long (with second-precision) the results of a preflight request
	// can be cached
	MaxAge time.Duration

	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool

	// Allows usage of popular browser extensions schemas
	AllowBrowserExtensions bool

	// Allows to add custom schema like tauri://
	CustomSchemas []string

	// Allows usage of WebSocket protocol
	AllowWebSockets bool

	// Allows usage of file:// schema (dangerous!) use it only when you 100% sure it's needed
	AllowFiles bool

	// Allows to pass custom OPTIONS response status code for old browsers / clients
	OptionsResponseStatusCode int
}

// AddAllowMethods is allowed to add custom methods
func (c *Config) AddAllowMethods(methods ...string) {
	c.AllowMethods = append(c.AllowMethods, methods...)
}

// AddAllowHeaders is allowed to add custom headers
func (c *Config) AddAllowHeaders(headers ...string) {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
}

// AddExposeHeaders is allowed to add custom expose headers
func (c *Config) AddExposeHeaders(headers ...string) {
	c.ExposeHeaders = append(c.ExposeHeaders, headers...)
}

func (c Config) getAllowedSchemas() []string {
	allowedSchemas := DefaultSchemas
	if c.AllowBrowserExtensions {
		allowedSchemas = append(allowedSchemas, ExtensionSchemas...)
	}
	if c.AllowWebSockets {
		allowedSchemas = append(allowedSchemas, WebSocketSchemas...)
	}
	if c.AllowFiles {
		allowedSchemas = append(allowedSchemas, FileSchemas...)
	}
	if c.CustomSchemas != nil {
		allowedSchemas = append(allowedSchemas, c.CustomSchemas...)
	}
	return allowedSchemas
}

var regexpBasedOrigin = regexp.MustCompile(`^\/(.+)\/[gimuy]?$`)

func (c Config) validateAllowedSchemas(origin string) bool {
	allowedSchemas := c.getAllowedSchemas()

	if regexpBasedOrigin.MatchString(origin) {
		// Normalize regexp-based origins
		origin = regexpBasedOrigin.FindStringSubmatch(origin)[1]
		origin = strings.Replace(origin, "?", "", 1)
	}

	for _, schema := range allowedSchemas {
		if strings.HasPrefix(origin, schema) {
			return true
		}
	}
	return false
}

// Validate is check configuration of user defined.
func (c Config) Validate() error {
	hasOriginFn := c.AllowOriginFunc != nil
	hasOriginFn = hasOriginFn || c.AllowOriginWithContextFunc != nil

	if c.AllowAllOrigins && (hasOriginFn || len(c.AllowOrigins) > 0) {
		originFields := strings.Join([]string{
			"AllowOriginFunc",
			"AllowOriginFuncWithContext",
			"AllowOrigins",
		}, " or ")
		return fmt.Errorf(
			"conflict settings: all origins enabled. %s is not needed",
			originFields,
		)
	}
	if !c.AllowAllOrigins && !hasOriginFn && len(c.AllowOrigins) == 0 {
		return errors.New("conflict settings: all origins disabled")
	}
	for _, origin := range c.AllowOrigins {
		if !strings.Contains(origin, "*") && !c.validateAllowedSchemas(origin) {
			return errors.New("bad origin: origins must contain '*' or include " + strings.Join(c.getAllowedSchemas(), ","))
		}
	}
	return nil
}

func (c Config) parseWildcardRules() [][]string {
	var wRules [][]string

	if !c.AllowWildcard {
		return wRules
	}

	for _, o := range c.AllowOrigins {
		if !strings.Contains(o, "*") {
			continue
		}

		if c := strings.Count(o, "*"); c > 1 {
			panic(errors.New("only one * is allowed").Error())
		}

		i := strings.Index(o, "*")
		if i == 0 {
			wRules = append(wRules, []string{"*", o[1:]})
			continue
		}
		if i == (len(o) - 1) {
			wRules = append(wRules, []string{o[:i], "*"})
			continue
		}

		wRules = append(wRules, []string{o[:i], o[i+1:]})
	}

	return wRules
}

// DefaultConfig returns a generic default configuration mapped to localhost.
func DefaultConfig() Config {
	return Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

// Default returns the location middleware with default configuration.
func Default() gin.HandlerFunc {
	config := DefaultConfig()
	config.AllowAllOrigins = true
	return New(config)
}

// New returns the location middleware with user-defined custom configuration.
func New(config Config) gin.HandlerFunc {
	cors := newCors(config)
	return func(c *gin.Context) {
		cors.applyCors(c)
	}
}
