package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "restaurant-order-management/db"
    "restaurant-order-management/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)

    result, err := db.DB.Exec("INSERT INTO orders (customerName, item, quantity, status) VALUES (?, ?, ?, ?)",
        order.CustomerName, order.Item, order.Quantity, order.Status)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, err := result.LastInsertId()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    order.ID = int(id)
    json.NewEncoder(w).Encode(order)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    var order models.Order
    err = db.DB.QueryRow("SELECT id, customerName, item, quantity, status FROM orders WHERE id = ?", id).Scan(
        &order.ID, &order.CustomerName, &order.Item, &order.Quantity, &order.Status)
    if err == sql.ErrNoRows {
        http.Error(w, "Order not found", http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    var order models.Order
    json.NewDecoder(r.Body).Decode(&order)
    order.ID = id

    _, err = db.DB.Exec("UPDATE orders SET customerName = ?, item = ?, quantity = ?, status = ? WHERE id = ?",
        order.CustomerName, order.Item, order.Quantity, order.Status, order.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid order ID", http.StatusBadRequest)
        return
    }

    _, err = db.DB.Exec("DELETE FROM orders WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
