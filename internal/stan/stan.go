package stan

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"service/internal/config"
	"service/internal/model"
	"service/pkg"
	"service/pkg/postgres"

	"github.com/nats-io/stan.go"
)

var Cache = make(map[int]model.Order)

func Stan() {
	config, err := config.GetConfig()
	if err != nil {
		log.Println(err)
	}

	clusterID := config.Stan.ClusterID
	clientID := config.Stan.ClientID

	orders := pkg.ReadOrdersFromDirectory("/home/user/Desktop/WBTECH/NATS-OrderStream-Service/json")

	conn := postgres.ConnectToDB(&config)
	defer conn.Close(context.Background())

	sc := ConnectToStan(clusterID, clientID)
	defer sc.Close()

	for _, order := range orders {
		postgres.InsertOrderToDB(conn, &order)

		Cache[order.ID] = order
		fmt.Println(Cache)

		PublishOrder(sc, &order)

		SubscribeToOrder(sc)
	}
}

func ConnectToStan(clusterID, clientID string) stan.Conn {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatalf("Error connecting to NATS Streaming: %v", err)
	}

	return sc
}

func PublishOrder(sc stan.Conn, order *model.Order) {
	orderData, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Error marshalling Order: %v", err)
	}

	err = sc.Publish("order", orderData)
	if err != nil {
		log.Fatalf("Error publishing to channel: %v", err)
	} else {
		log.Println("Successfully published to channel")
	}
}

func SubscribeToOrder(sc stan.Conn) {
	sub, err := sc.Subscribe("order", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatalf("Error subscribing to channel: %v", err)
	} else {
		log.Println("Successfully subscribed to channel")
	}
	defer sub.Unsubscribe()
}
