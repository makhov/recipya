package server

import (
	"github.com/reaper47/recipya/internal/templates"
	"net/http"
)

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	page := templates.LandingPage
	isAuth := isAuthenticated(r, s.Repository.GetAuthToken)
	if isAuth {
		page = templates.HomePage
	}

	templates.Render(w, page, templates.Data{
		IsAuthenticated: isAuth,
		Title:           page.Title(),
	})
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templates.Render(w, templates.Simple, templates.PageNotFound)
}

func (s *Server) settingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Hx-Request") == "true" {
		templates.RenderComponent(w, "core", "settings", nil)
	} else {
		page := templates.SettingsPage
		templates.Render(w, page, templates.Data{
			IsAuthenticated: true,
			Title:           page.Title(),
		})
	}
}

func (s *Server) userInitialsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID")
	if userID == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	w.Write([]byte(s.Repository.UserInitials(userID.(int64))))
}
