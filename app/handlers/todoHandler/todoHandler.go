package todoHandler

import (
	"log"
	"net/http"
	"text/template"
	"time"
	"todo-list/app/repositories/todoRepository"

	"github.com/gorilla/mux"
)

// DB名の定数
const (
	DBName = "root:pass@(127.0.0.1:3306)/playground?parseTime=true"
)

type todo struct {
	Id          int
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	// クエリ実行
	rows, err := db.Query(`SELECT id, title, description, done, created_at FROM todos`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 結果を取得
	var todos []todo
	for rows.Next() {
		var t todo

		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Done, &t.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, t)
	}

	View(w, todos, "resources/view/layout.html", "resources/view/todos/index.html")
}

func TodoNew(w http.ResponseWriter, r *http.Request) {
	View(w, nil, "resources/view/layout.html", "resources/view/todos/new.html")
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	title := r.FormValue("title")
	description := r.FormValue("description")
	createdAt := time.Now()

	_, err := db.Exec(`INSERT INTO todos (title, description, created_at) VALUES (?, ?, ?)`, title, description, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	// indexにリダイレクトする
	http.Redirect(w, r, "/", http.StatusFound)
}

func TodoEdit(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	vars := mux.Vars(r)
	id := vars["todoId"]

	// クエリ実行
	row := db.QueryRow(`SELECT id, title, description, done, created_at FROM todos WHERE id = ?`, id)

	// 結果を取得
	var t todo
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Done, &t.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	View(w, t, "resources/view/layout.html", "resources/view/todos/edit.html")
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	vars := mux.Vars(r)
	id := vars["todoId"]
	title := r.FormValue("title")
	description := r.FormValue("description")

	_, err := db.Exec(`UPDATE todos SET title = ?, description = ? WHERE id = ?`, title, description, id)
	if err != nil {
		log.Fatal(err)
	}

	// indexにリダイレクトする
	http.Redirect(w, r, "/", http.StatusFound)
}

func TodoDone(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	vars := mux.Vars(r)
	id := vars["todoId"]

	_, err := db.Exec(`UPDATE todos SET done = true WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}

	// indexにリダイレクトする
	http.Redirect(w, r, "/", http.StatusFound)
}

func TodoUndone(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	vars := mux.Vars(r)
	id := vars["todoId"]

	_, err := db.Exec(`UPDATE todos SET done = false WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}

	// indexにリダイレクトする
	http.Redirect(w, r, "/", http.StatusFound)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	db := todoRepository.Connect(DBName)

	vars := mux.Vars(r)
	id := vars["todoId"]

	_, err := db.Exec(`DELETE FROM todos WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}

	// indexにリダイレクトする
	http.Redirect(w, r, "/", http.StatusFound)
}

func View(w http.ResponseWriter, data any, tmpl ...string) {
	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
