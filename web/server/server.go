package server

import (
	"errors"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/im-jinsu/yepanmap/web/route"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/go-playground/validator"
)

// CustomValidator validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validate
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Run run api server
func Run() {
	e := route.Route()

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		return name
	})
	e.Validator = &CustomValidator{validator: validate}
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// }))

	go func() {
		e := route.Route()
		e.Pre(middleware.HTTPSRedirect())
		e.Logger.Fatal(e.Start(":80"))
	}()

	// HTTPS
	// var err error
	// var cert []byte
	// if cert, err = filepathOrContent(fmt.Sprintf("%s/web/server/cert.crt",
	// 	loadconf.WASROOTDIR)); err != nil {
	// 	return
	// }
	// var key []byte
	// if key, err = filepathOrContent(fmt.Sprintf("%s/web/server/cert.key",
	// 	loadconf.WASROOTDIR)); err != nil {
	// 	return
	// }
	// e.TLSServer.TLSConfig = new(tls.Config)
	// e.TLSServer.TLSConfig.Certificates = make([]tls.Certificate, 1)

	// if e.TLSServer.TLSConfig.Certificates[0], err = tls.X509KeyPair(cert, key); err != nil {
	// 	return
	// }
	// e.TLSServer.Addr = ":3443"
	// e.TLSServer.ReadTimeout = 5 * time.Second
	// e.TLSServer.WriteTimeout = 10 * time.Second

	// e.Logger.Fatal(grace.Serve(e.TLSServer))

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Logger.Fatal(e.StartAutoTLS(":443"))
}

func filepathOrContent(fileOrContent interface{}) (content []byte, err error) {
	switch v := fileOrContent.(type) {
	case string:
		return ioutil.ReadFile(v)
	case []byte:
		return v, nil
	default:
		return nil, errors.New("invalid cert or key type, must be string or []byte")
	}
}
