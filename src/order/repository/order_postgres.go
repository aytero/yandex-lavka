package repository

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "github.com/labstack/gommon/log"
    "github.com/lib/pq"
    "yandex-team.ru/bstask/database"
    "yandex-team.ru/bstask/model"
)

const (
    queryGetOrderById         = `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
    queryGetAllOrders         = `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders LIMIT $1 OFFSET $2`
    queryCreateOrder          = `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`
    queryGetByIdWithCourierId = `SELECT cour_id, order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
    queryUpdateOrderComplete  = `UPDATE orders SET completed_time = $1 WHERE cour_id = $2 AND order_id = $3 RETURNING order_id, courier_id, region, weight, cost, completed_time`
    //queryCompletedTime = `UPDATE orders SET courier_id = $1, completed_time = $2 WHERE order_id = $3
    queryGetCourierIdFromOrder = `SELECT courier_id FROM orders WHERE order_id = $1`
)

type OrderRepository struct {
    db *database.Database
}

func NewOrderRepository(dbConn *database.Database) *OrderRepository {
    return &OrderRepository{
        db: dbConn,
    }
}

func (r *OrderRepository) GetById(ctx context.Context, id int64) (*model.Order, error) {
    var order model.Order

    err := r.db.Conn.QueryRowContext(ctx, queryGetOrderById, id).Scan(&order.OrderId, &order.Weight, &order.Regions,
        pq.Array(&order.DeliveryHours), &order.Cost, &order.CompletedTime)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        } else {
            return nil, fmt.Errorf("OrderRepository - GetByID: %w", err)
        }
    }

    return &order, nil
}

func (r *OrderRepository) GetOrders(ctx context.Context, limit, offset int32) ([]*model.Order, error) {
    orders := make([]*model.Order, 0, limit)

    rows, err := r.db.Conn.QueryContext(ctx, queryGetAllOrders, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - GetOrders: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        temp := model.Order{}
        err := rows.Scan(&temp.OrderId, &temp.Weight, &temp.Regions, pq.Array(&temp.DeliveryHours), &temp.Cost, &temp.CompletedTime)
        if err != nil {
            return nil, fmt.Errorf("OrderRepository - GetOrders: %w", err)
        }
        orders = append(orders, &temp)
    }
    return orders, nil
}

func (r *OrderRepository) CreateOrders(ctx context.Context, orders []*model.CreateOrder) ([]*model.Order, error) {

    resOrd := make([]*model.Order, 0, len(orders))

    tx, err := r.db.Conn.BeginTxx(ctx, &sql.TxOptions{}) // nil
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - AddOrder: %w", err)
    }
    defer tx.Rollback()

    for _, o := range orders {
        var orderId int64

        err := tx.QueryRowContext(ctx, queryCreateOrder, o.Weight, o.Regions, pq.Array(&o.DeliveryHours), o.Cost).Scan(&orderId)
        if err != nil {
            return nil, fmt.Errorf("OrderRepository - CreateOrders: %w", err)
        }
        temp := model.Order{
            OrderId:       orderId,
            Weight:        o.Weight,
            Regions:       o.Regions,
            DeliveryHours: o.DeliveryHours,
            Cost:          o.Cost,
        }
        resOrd = append(resOrd, &temp)
    }
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("OrderRepository - CreateOrders: %w", err)
    }
    return resOrd, nil
}

// todo UpdateOrder ___
func (r *OrderRepository) UpdateOrders(ctx context.Context, ords []*model.CompleteOrder) ([]*model.Order, error) {

    tx, err := r.db.Conn.BeginTxx(ctx, &sql.TxOptions{}) // nil
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - AddOrder: %w", err)
    }
    //defer tx.Rollback()

    resOrders := make([]*model.Order, 0, len(ords))
    var updatedOrder model.Order
    //var id int64
    for _, o := range ords {
        /*
        	// todo rename to smth like queryGetByIdAllInfo (courier_id)
        	err = tx.QueryRowContext(ctx, queryGetByIdWithCourierId, o.OrderId).
        		Scan(&id, &order.OrderId, &order.Weight, &order.Regions, pq.Array(&order.DeliveryHours), &order.Cost, &order.CompletedTime)
        	if err != nil {
        		return nil, fmt.Errorf("OrderRepository - Update: %w", err)
        		//err = tx.QueryRowContext(ctx, queryGetOrderById, o.OrderId).
        		//	Scan(&order.OrderId, &order.Weight, &order.Regions, pq.Array(&order.DeliveryHours), &order.Cost, &order.CompletedTime)
        		//id = o.CourierId
        		//if err != nil {
        		//	return nil, fmt.Errorf("OrderRepository - Update: %w", err)
        		//}

        		// todo // assign // commit new id
        		//queryUpdateCourId := `UPDATE orders SET cour_id = $1 WHERE order_id = $2`
        		//if _, err = tx.ExecContext(ctx, queryUpdateCourId, o.CourierId, o.OrderId); err != nil {
        		//	return nil, fmt.Errorf("OrderRepository - Update: %w", err)
        		//}
        	}
        */

        // todo validate delivery hours ? in assign
        var courierId int64
        err = tx.QueryRowContext(ctx, queryGetCourierIdFromOrder, o.OrderId).
            Scan(&courierId)
        if err != nil {
            return nil, fmt.Errorf("OrderRepository - Update: %w", err)
        }
        //if id == 0 {
        // complete anyway, assign in process
        //}

        // 0 or nil
        if courierId == 0 || courierId != o.CourierId {
            return nil, errors.New("invalid courier id")
        }

        //if updatedOrder.CompletedTime != nil {
        //	return nil, fmt.Errorf("already completed")
        //}
        if updatedOrder.CompletedTime == nil {
            temp := o.CompleteTime
            updatedOrder.CompletedTime = &temp
            log.Infof("update complete time")
            //-- RETURNING order_id, courier_id, region, weight, cost, completed_time`
            err := tx.QueryRowContext(ctx, queryUpdateOrderComplete, o.CompleteTime, o.CourierId, o.OrderId).
                Scan(updatedOrder.OrderId, updatedOrder.Regions, updatedOrder.Weight, updatedOrder.Cost, updatedOrder.CompletedTime)
            //_, err := tx.ExecContext(ctx, queryUpdateOrderComplete, o.CompleteTime, o.CourierId, o.OrderId).Scan()
            if err != nil {
                return nil, fmt.Errorf("OrderRepository - Update: %w", err)
            }
        }
        resOrders = append(resOrders, &updatedOrder)
    }
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("OrderRepository - Upadate: %w", err)
    }
    return resOrders, nil
}
