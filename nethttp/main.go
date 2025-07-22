package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"github.com/peteraba/go-frameworks/shared/service"
)

var (
	projectService *service.ProjectService
	listService    *service.ListService
	todoService    *service.TodoService
	userService    *service.UserService
	userRepo       *repo.InMemoryUserRepo
)

func main() {
	projectRepo := repo.NewInMemoryProjectRepo()
	projectService = service.NewProjectService(projectRepo)
	listRepo := repo.NewInMemoryListRepo()
	listService = service.NewListService(listRepo)
	todoRepo := repo.NewInMemoryTodoRepo()
	todoService = service.NewTodoService(todoRepo)
	userRepo = repo.NewInMemoryUserRepo()
	userService = service.NewUserService(userRepo)

	// Create admin user
	userCreate := model.UserCreate{
		Name:      "Peter Aba",
		Email:     "peter@example.com",
		Password:  "^uw6fJ3NY45peX",
		Password2: "^uw6fJ3NY45peX",
		Groups: []string{
			"project.read",
			"project.write",
		},
	}
	_, err := userService.Create(userCreate)
	if err != nil {
		panic(err)
	}

	// --- Project Handlers ---
	http.HandleFunc("GET /projects", handleListProjects)
	http.HandleFunc("POST /projects", handleCreateProject)
	http.HandleFunc("GET /projects/{id}", handleGetProject)
	http.HandleFunc("PUT /projects/{id}", handleUpdateProject)

	// --- List Handlers ---
	http.HandleFunc("GET /lists", handleListLists)
	http.HandleFunc("POST /lists", handleCreateList)
	http.HandleFunc("GET /lists/{id}", handleGetList)
	http.HandleFunc("PUT /lists/{id}", handleUpdateList)

	// --- Todo Handlers ---
	http.HandleFunc("GET /lists/{listId}/todos", handleListTodos)
	http.HandleFunc("POST /lists/{listId}/todos", handleCreateTodo)
	http.HandleFunc("GET /lists/{listId}/todos/{todoId}", handleGetTodo)
	http.HandleFunc("PUT /lists/{listId}/todos/{todoId}", handleUpdateTodo)

	// --- User Handlers ---
	http.HandleFunc("GET /users", handleListUsers)
	http.HandleFunc("POST /users", handleCreateUser)
	http.HandleFunc("GET /users/{userId}", handleGetUser)
	http.HandleFunc("PUT /users/{userId}", handleUpdateUser)
	http.HandleFunc("DELETE /users/{userId}", handleDeleteUser)
	http.HandleFunc("PUT /users/{userId}/passwords", handleUpdateUserPassword)
	http.HandleFunc("POST /logins", handleLoginUser)
	http.HandleFunc("GET /health", handleHealth)

	log.Println("Serving API at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// Helper to write JSON error responses
func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		log.Printf("Failed to write JSON error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// --- Project Handlers ---
func handleListProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := projectService.List()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to list projects")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(projects); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleCreateProject(w http.ResponseWriter, r *http.Request) {
	var pc model.ProjectCreate
	if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	project, err := projectService.Create(pc)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(project); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleGetProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	project, err := projectService.GetByID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Project not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleUpdateProject(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var pu model.ProjectUpdate
	if err := json.NewDecoder(r.Body).Decode(&pu); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	project, err := projectService.Update(id, pu)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(project); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

// --- List Handlers ---
func handleListLists(w http.ResponseWriter, r *http.Request) {
	lists, err := listService.List()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to list lists")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lists); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleCreateList(w http.ResponseWriter, r *http.Request) {
	var lc model.ListCreate
	if err := json.NewDecoder(r.Body).Decode(&lc); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	list, err := listService.Create(lc)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(list); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleGetList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	list, err := listService.GetByID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "List not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleUpdateList(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var lu model.ListUpdate
	if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	list, err := listService.Update(id, lu)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(list); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

// --- Todo Handlers ---
func handleListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todoService.List()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to list todos")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	listId := r.PathValue("listId")
	var tc model.TodoCreate
	if err := json.NewDecoder(r.Body).Decode(&tc); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	tc.ListID = listId
	todo, err := todoService.Create(tc)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleGetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := r.PathValue("todoId")
	todo, err := todoService.GetByID(todoId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "Todo not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := r.PathValue("todoId")
	var tu model.TodoUpdate
	if err := json.NewDecoder(r.Body).Decode(&tu); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	todo, err := todoService.Update(todoId, tu)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

// --- User Handlers ---
func handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userRepo.List()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to list users")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var uc model.UserCreate
	if err := json.NewDecoder(r.Body).Decode(&uc); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	user, err := userService.Create(uc)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	user, err := userRepo.GetByID(userId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "User not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	var uu model.UserUpdate
	if err := json.NewDecoder(r.Body).Decode(&uu); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	user, err := userRepo.Update(userId, uu)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	err := userRepo.Delete(userId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "User not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func handleUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	var up model.UserPasswordUpdate
	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	user, err := userRepo.UpdatePassword(userId, []byte(up.Password))
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	var ul model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&ul); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}
	token, err := userService.Login(ul)
	if err != nil {
		writeJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"token": token}); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "ok"}); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
