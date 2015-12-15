package apps

import (
	"net/http"

	"github.com/mholt/binding"
	"github.com/mix3/go-todo-webapp/db"
	"github.com/mix3/go-todo-webapp/forms"
	"github.com/mix3/go-todo-webapp/options"
	"github.com/unrolled/render"
)

type App struct {
	mux *http.ServeMux
	db  *db.DB
	r   *render.Render
}

func New(opts options.Options, db *db.DB) *App {
	a := &App{
		mux: http.NewServeMux(),
		db:  db,
		r:   render.New(render.Options{}),
	}
	a.mux.HandleFunc("/api/list", a.ListHandler)
	a.mux.HandleFunc("/api/create", a.CreateHandler)
	a.mux.HandleFunc("/api/switch", a.SwitchHandler)
	a.mux.HandleFunc("/api/delete", a.DeleteHandler)
	a.mux.Handle("/", http.FileServer(http.Dir(opts.Static)))
	return a
}

func (a *App) Close() {
	a.db.Close()
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func (a *App) ListHandler(w http.ResponseWriter, r *http.Request) {
	res, err := a.db.TodoList()
	if err != nil {
		a.renderErr(w, err)
		return
	}
	a.renderOK(w, res)
}

func (a *App) CreateHandler(w http.ResponseWriter, r *http.Request) {
	createForm := new(forms.CreateForm)
	errs := binding.Bind(r, createForm)
	if 0 < errs.Len() {
		a.renderErr(w, errs)
		return
	}

	err := a.db.TodoCreate(createForm.Text)
	if err != nil {
		a.renderErr(w, err)
		return
	}

	a.renderOK(w, nil)
}

func (a *App) SwitchHandler(w http.ResponseWriter, r *http.Request) {
	switchForm := new(forms.SwitchForm)
	errs := binding.Bind(r, switchForm)
	if 0 < errs.Len() {
		a.renderErr(w, errs)
		return
	}

	err := a.db.TodoSwitch(switchForm.Id)
	if err != nil {
		a.renderErr(w, err)
		return
	}

	a.renderOK(w, nil)
}

func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	deleteForm := new(forms.DeleteForm)
	errs := binding.Bind(r, deleteForm)
	if 0 < errs.Len() {
		a.renderErr(w, errs)
		return
	}

	err := a.db.TodoDelete(deleteForm.Id)
	if err != nil {
		a.renderErr(w, err)
		return
	}

	a.renderOK(w, nil)
}

func (a *App) renderOK(w http.ResponseWriter, res interface{}) {
	a.r.JSON(w, http.StatusOK, Response{
		Success: true,
		Result:  res,
	})
}

func (a *App) renderErr(w http.ResponseWriter, err error) {
	a.r.JSON(w, http.StatusOK, Response{
		Success: false,
		Message: err.Error(),
	})
}
