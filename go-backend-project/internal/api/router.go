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

func handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Example: return empty user list
        users := []interface{}{}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(users)
    case http.MethodPost:
        // Example: create user (dummy response)
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message":"User created"}`))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
