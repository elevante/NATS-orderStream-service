package cache

import (
	"encoding/json"
	"service/internal/model"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
)

var MC *memcache.Client

func init() {
	MC = memcache.New("127.0.0.1:11211")
}

func SaveToCache(order *model.Order) error {
	orderData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = MC.Set(&memcache.Item{Key: "order:" + strconv.Itoa(order.ID), Value: orderData})
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err := MC.Set(&memcache.Item{Key: "item:" + strconv.Itoa(item.ID), Value: []byte(strconv.Itoa(item.ChrtID))})
		if err != nil {
			return err
		}
	}

	err = MC.Set(&memcache.Item{Key: "delivery:" + strconv.Itoa(order.Delivery.ID), Value: []byte(order.Delivery.Name)})
	if err != nil {
		return err
	}

	err = MC.Set(&memcache.Item{Key: "payment:" + strconv.Itoa(order.Payment.ID), Value: []byte(order.Payment.Transaction)})
	if err != nil {
		return err
	}

	return nil
}
