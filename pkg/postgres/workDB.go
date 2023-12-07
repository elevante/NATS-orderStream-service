package postgres

import (
	"fmt"
	"log"
	"os"
	"service/internal/config"
	"service/internal/db"
	"service/internal/model"

	"github.com/jackc/pgx/v4"
)

func ConnectToDB(config *config.AppConfig) *pgx.Conn {
	conn, err := db.ConnectDB(*config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func InsertOrderToDB(conn *pgx.Conn, order *model.Order) {
	err := db.InsertDelivery(conn, &order.Delivery)
	if err != nil {
		log.Fatalf("Error inserting delivery: %v", err)
	}

	err = db.InsertPayment(conn, &order.Payment)
	if err != nil {
		log.Fatalf("Error inserting payment: %v", err)
	}

	for _, item := range order.Items {
		err = db.InsertItem(conn, &item)
		if err != nil {
			log.Fatalf("Error inserting item: %v", err)
		}
	}

	err = db.InsertOrder(conn, order)
	if err != nil {
		log.Fatalf("Error inserting order: %v", err)
	}
}
