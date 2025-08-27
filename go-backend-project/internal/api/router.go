package api

import (
    "net/http"
    "encoding/json"
    "fmt"
    "github.com/AbdullahOztoprak/go-backend-project/internal/service"
    "github.com/AbdullahOztoprak/go-backend-project/internal/models"
)

type Router struct {
    UserService        service.UserService
    TransactionService service.TransactionService
    BalanceService     service.BalanceService
}

func NewRouter(userService service.UserService, transactionService service.TransactionService, balanceService service.BalanceService) http.Handler {
    r := &Router{
        UserService:        userService,
        TransactionService: transactionService,
        BalanceService:     balanceService,
    }
    mux := http.NewServeMux()
    mux.HandleFunc("/api/v1/users", r.handleUsers)
    mux.HandleFunc("/api/v1/transactions", r.handleTransactions)
    mux.HandleFunc("/api/v1/balances", r.handleBalances)
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
        return
    case http.MethodPost:
        var user models.User
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
        return
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func (r *Router) handleBalances(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        userIDParam := req.URL.Query().Get("user_id")
        if userIDParam == "" {
            http.Error(w, "Missing user_id", http.StatusBadRequest)
            return
        }
        var userID int64
        if _, err := fmt.Sscan(userIDParam, &userID); err != nil {
            http.Error(w, "Invalid user_id", http.StatusBadRequest)
            return
        }
        balance, err := r.BalanceService.GetByUserID(userID)
        if err != nil {
            http.Error(w, "Failed to fetch balance", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(balance)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func (r *Router) handleTransactions(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
    case http.MethodGet:
        // Example: /api/v1/transactions?user_id=123
        userIDParam := req.URL.Query().Get("user_id")
        if userIDParam == "" {
            http.Error(w, "Missing user_id", http.StatusBadRequest)
            return
        }
        // Parse user_id
        var userID int64
        if _, err := fmt.Sscan(userIDParam, &userID); err != nil {
            http.Error(w, "Invalid user_id", http.StatusBadRequest)
            return
        }
        txs, err := r.TransactionService.ListByUser(userID)
        if err != nil {
            http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(txs)
    case http.MethodPost:
        var tx models.Transaction
        if err := json.NewDecoder(req.Body).Decode(&tx); err != nil {
            http.Error(w, "Invalid request", http.StatusBadRequest)
            return
        }
        err := r.TransactionService.Create(&tx)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`{"message":"Transaction created"}`))
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}