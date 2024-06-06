package appform

import (
	"fmt"
	"go-quickstart/internal/httpcontext"
	"go-quickstart/internal/web"
	"net/http"
	"os"
	"time"
)

func Login(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	path := func(email string, password string, loginErr string) string {
		return fmt.Sprintf(`/?loginErr=%s&email=%s&password=%s`, loginErr, email, password)
	}
	email := web.Scrub(r.FormValue("email"))
	password := web.Scrub(r.FormValue("password"))
	if email != os.Getenv("ADMIN_EMAIL") || password != os.Getenv("ADMIN_PASSWORD") {
		http.Redirect(w, r, path(email, password, "invalid credentials"), http.StatusSeeOther)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    os.Getenv("ADMIN_SESSION_TOKEN"),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
