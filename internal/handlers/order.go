package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"service/internal/cache"
	"service/internal/model"
	"strconv"
)

type OrderPageData struct {
	Order    model.Order
	Delivery model.Delivery
	Payment  model.Payment
	Items    []model.Items
}

func HandlerOrder(w http.ResponseWriter, r *http.Request) {
	orderIDStr := r.URL.Query().Get("id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	orderItem, err := cache.MC.Get("order:" + strconv.Itoa(orderID))
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	order := model.Order{}
	err = json.Unmarshal(orderItem.Value, &order)
	if err != nil {
		http.Error(w, "Error unmarshalling order", http.StatusInternalServerError)
		return
	}

	delivery := order.Delivery
	payment := order.Payment
	items := order.Items

	for i, item := range items {
		item.ID = i + 1
		items[i] = item
	}

	tmpl := template.Must(template.ParseFiles(".././order.html"))
	data := OrderPageData{
		Order:    order,
		Delivery: delivery,
		Payment:  payment,
		Items:    items,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
