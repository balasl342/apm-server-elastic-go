package middleware

import (
	"net/http"
	"net/url"

	"github.com/balasl342/apm-server-elastic-go/config"

	"go.elastic.co/apm/module/apmhttp"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/transport"
)

var Tracer *apm.Tracer

func Newelasticapp() *apm.Tracer {
	transport, err := transport.NewHTTPTransport(transport.HTTPTransportOptions{})
	if err != nil {
		panic(err)
	}
	serverURL, err := url.Parse(config.AppConfig.ElasticAPM.APMServerURLs)
	if err != nil {
		panic(err)
	}
	transport.SetServerURL(serverURL)
	transport.SetSecretToken(config.AppConfig.ElasticAPM.APMSecretToken)

	// Create a new tracer with the specified transport
	Tracer, err = apm.NewTracerOptions(apm.TracerOptions{
		ServiceName:    config.AppConfig.ElasticAPM.APMServiceName,
		ServiceVersion: "1.0",
		Transport:      transport,
	})
	if err != nil {
		panic(err)
	}
	return Tracer
}

func WrapHandleFunc(next http.Handler) http.Handler {
	return apmhttp.Wrap(next)
}
