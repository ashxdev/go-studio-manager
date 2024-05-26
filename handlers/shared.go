package handlers

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

func CacheControlFileServer(h http.Handler, cacheControlHeader string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", cacheControlHeader)
		h.ServeHTTP(w, r)
	})
}

func render(r *http.Request, w http.ResponseWriter, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func hxRedirect(w http.ResponseWriter, r *http.Request, to string) error {
	if len(r.Header.Get("HX-Request")) > 0 {
		w.Header().Set("HX-Redirect", to)
		w.WriteHeader(http.StatusSeeOther)
		return nil
	}
	http.Redirect(w, r, to, http.StatusSeeOther)
	return nil
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("internal server error", "err", err, "path", r.URL.Path)
		}
	}
}