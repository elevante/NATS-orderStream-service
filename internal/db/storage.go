package db

import (
	"context"
	"fmt"
	"service/internal/config"
	"service/internal/model"

	"github.com/jackc/pgx/v4"
)

func ConnectDB(config config.AppConfig) (*pgx.Conn, error) {
	return pgx.Connect(
		context.Background(),
		fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name))
}

func InsertItem(conn *pgx.Conn, item *model.Items) error {
	err := conn.QueryRow(context.Background(),
		`SELECT id FROM items WHERE chrt_id = $1 AND track_number = $2 AND price = $3`,
		item.ChrtID,
		item.TrackNumber,
		item.Price).Scan(&item.ID)

	if err == pgx.ErrNoRows {
		err = conn.QueryRow(context.Background(),
			`INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status).Scan(&item.ID)
	}
	return err
}

func InsertDelivery(conn *pgx.Conn, delivery *model.Delivery) error {
	err := conn.QueryRow(context.Background(),
		`INSERT INTO delivery (name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email).Scan(&delivery.ID)
	return err
}

func InsertPayment(conn *pgx.Conn, payment *model.Payment) error {
	err := conn.QueryRow(context.Background(),
		`INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		payment.Transaction,
		payment.ReqestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDT,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee).Scan(&payment.ID)
	return err
}

func InsertOrder(conn *pgx.Conn, order *model.Order) error {
	err := conn.QueryRow(context.Background(),
		`INSERT INTO orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery.ID,
		order.Payment.ID,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard).Scan(&order.ID)

	if err != nil {
		return err
	}
	return nil
}
