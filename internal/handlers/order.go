package handlers

import (
	"html/template"
	"log"
	"net/http"
	"service/internal/model"
	"service/internal/stan"
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
	if order, ok := stan.Cache[orderID]; ok {

		delivery := order.Delivery
		payment := order.Payment
		items := order.Items

		for i, item := range items {
			item.ID = i + 1
			items[i] = item
		}

		tmpl := template.Must(template.ParseFiles("/home/user/Desktop/WBTECH/NATS-OrderStream-Service/order.html"))
		data := OrderPageData{
			Order:    order,
			Delivery: delivery,
			Payment:  payment,
			Items:    items,
		}
		// fmt.Println("-------------------------\n")
		// fmt.Println(order)
		// fmt.Println("-------------------------\n")
		// fmt.Println("-------------------------\n")
		// fmt.Println(payment)
		// fmt.Println("-------------------------\n")
		// fmt.Println("-------------------------\n")
		// fmt.Println(delivery)
		// fmt.Println("-------------------------\n")
		// fmt.Println("-------------------------\n")
		// fmt.Println(items)
		// fmt.Println("-------------------------\n")
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error executing template :", err)
			return
		}
	} else {
		http.Error(w, "Order not found", http.StatusNotFound)
	}
}
