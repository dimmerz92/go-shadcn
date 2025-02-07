package lib

import (
	"net/http"

	"github.com/a-h/templ"
)

// Render writes the status code and zero to many templ components to the response writer.
func Render(w http.ResponseWriter, r *http.Request, status int, components ...templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	for _, component := range components {
		if err := component.Render(r.Context(), buf); err != nil {
			return err
		}
	}

	w.WriteHeader(status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// IsHtmx returns true if the request was made by HTMX, otherwise false.
func IsHtmx(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

// HxRender is a HTMX aware renderer that conditionally renders the component if the request was made by HTMX, otherwise
// the page component is rendered.
// Writes the status code and component to the response writer.
func HxRender(w http.ResponseWriter, r *http.Request, status int, page, component templ.Component) error {
	if IsHtmx(r) {
		return Render(w, r, status, component)
	}
	return Render(w, r, status, page)
}

// Redirect is a HTMX aware renderer that performs a HTTP redirect.
// Writes the status code and redirect headers to the response writer.
//
// If the request was made by HTMX, the "HX-Redirect" header is set with the route and is returned with a 200 OK.
// HTMX does not handle 3xx redirects, it only sees the final 2xx, 4xx, or 5xx response.
// For non-HTMX requests, a standard HTTP redirect is performed with the provided status.
// See github.com/bigskysoftware/htmx/issues/2052#issuecomment-1979805051 for more details.
func Redirect(w http.ResponseWriter, r *http.Request, status int, route string) {
	if IsHtmx(r) {
		w.Header().Add("HX-Redirect", route)
		status = http.StatusOK
	}
	http.Redirect(w, r, route, status)
}
