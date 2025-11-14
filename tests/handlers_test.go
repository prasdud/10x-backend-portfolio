package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prasdud/10x-backend-portfolio/internal/handlers"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/ping", handlers.Ping)
	v1.GET("/resume", handlers.GetResume)
	v1.GET("/skills", handlers.GetSkills)
	v1.GET("/projects", handlers.GetProjects)
	v1.GET("/experience", handlers.GetExperience)
	v1.GET("/contact", handlers.GetContact)
	v1.GET("/blog", handlers.GetBlog)
	v1.GET("/about", handlers.GetAbout)
	v1.GET("/easteregg1", handlers.GetEasterEgg1)
	v1.GET("/easteregg2", handlers.GetEasterEgg2)
	return r
}

func TestPing(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/ping")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetResume(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/resume")
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected status 200 or 404, got %d", w.Code)
	}
}

func TestGetSkills(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/skills")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetProjects(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/projects")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetExperience(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/experience")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetContact(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/contact")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetBlog(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/blog")
	if w.Code != http.StatusFound && w.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected status 302 or 307, got %d", w.Code)
	}
}

func TestGetAbout(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/about")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetEasterEgg1(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/easteregg1")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestGetEasterEgg2(t *testing.T) {
	r := setupRouter()
	w := performRequest(r, "GET", "/api/v1/easteregg2")
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
