package api

import (
    "net/http"
    "encoding/json"
    "github.com/AbdullahOztoprak/go-backend-project/internal/service"
)

type Router struct {
    UserService service.UserService
}

func NewRouter(userService service.UserService) http.Handler {
    r := &Router{UserService: userService}
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/users", r.handleUsers)
    return mux
}

func (r *Router) handleUsers(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        users, err := r.UserService.List()
        if err != nil {
            http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
    case http.MethodPost:
        var user service.UserCreateRequest // veya models.User
        if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        err := r.UserService.Register(&user)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message":"User created"}`))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}