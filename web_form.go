package main

import "net/http"

func serveForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<form method="POST">
				<p>First name: <input type="text" name="first_name"></p>
				<p>Last name: <input type="text" name="last_name"></p>
				<p>Phone: <input type="text" name="phone"></p>
				<input type="submit">
			</form>
		`))
		return
	}

	r.ParseForm()
	w.Write([]byte(r.FormValue("first_name")))
	w.Write([]byte(r.FormValue("last_name")))
	w.Write([]byte(r.FormValue("phone")))
}

func processForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte(r.FormValue("first_name")))
	w.Write([]byte(r.FormValue("last_name")))
	w.Write([]byte(r.FormValue("phone")))
}

func main() {
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/process", processForm)
	http.ListenAndServe(":8080", nil)
}
