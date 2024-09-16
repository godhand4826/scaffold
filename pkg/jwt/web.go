package jwt

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.New("").Parse(`
<html>
<head>
    <script type="text/javascript">
        localStorage.setItem('token', '{{.JWT}}');
        window.location.href = '{{.URL}}';
    </script>
</head>
<body>Redirecting...</body>
</html>
`))

func ReturnTokenAndRedirect(w http.ResponseWriter, jwt, url string) {
	data := struct {
		JWT string
		URL string
	}{
		JWT: jwt,
		URL: url,
	}

	w.Header().Set("Content-Type", "text/html")
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
