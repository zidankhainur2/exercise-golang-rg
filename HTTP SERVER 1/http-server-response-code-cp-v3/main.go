package main

import (
	"net/http"
)

// Daftar nama siswa
var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

// Fungsi untuk memeriksa apakah nama siswa ada dalam daftar
func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}
	return false
}

// Handler untuk memproses permintaan pada endpoint /students
func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Hanya izinkan metode GET
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		// Ambil parameter "name" dari query
		name := r.URL.Query().Get("name")
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
			return
		}

		if IsNameExists(name) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Name is exists"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Data not found"))
		}
	}
}

// Fungsi untuk mengatur routing dan mengembalikan Mux
func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
