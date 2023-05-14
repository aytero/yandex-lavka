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

// todo order repository
// todo deal with pointers
// one global repo for two sub repos?

const (
    queryGetOrderById         = `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
    queryGetAllOrders         = `SELECT order_id, weight, regions, delivery_hours, cost, completed_time FROM orders LIMIT $1 OFFSET $2`
    queryCreateOrder          = `INSERT INTO orders (weight, regions, delivery_hours, cost) VALUES ($1, $2, $3, $4) RETURNING order_id`
    queryGetByIdWithCourierId = `SELECT cour_id, order_id, weight, delivery_hours, cost, completed_time FROM orders WHERE order_id = $1`
    queryUpdateOrder          = `UPDATE orders SET completed_time = $1 WHERE cour_id = $2 AND order_id = $3`
)

type OrderRepository struct {
    db *database.Database
}

func NewOrderRepository(dbConn *database.Database) *OrderRepository {
    return &OrderRepository{
        db: dbConn,
    }
}

func (r *OrderRepository) GetById(ctx context.Context, id int64) (*model.OrderDto, error) {
    var order model.OrderDto

    //err := r.db.Conn.GetContext(ctx, &order, query, id)
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

func (r *OrderRepository) GetOrders(ctx context.Context, limit, offset int32) ([]model.OrderDto, error) {
    orders := make([]model.OrderDto, 0, limit)

    rows, err := r.db.Conn.QueryContext(ctx, queryGetAllOrders, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - GetOrders: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        temp := model.OrderDto{}
        err := rows.Scan(&temp.OrderId, &temp.Weight, &temp.Regions, pq.Array(&temp.DeliveryHours), &temp.Cost, &temp.CompletedTime)
        if err != nil {
            return nil, fmt.Errorf("OrderRepository - GetOrders: %w", err)
        }
        orders = append(orders, temp)
    }

    //err := rows.Scan(&temp.CourierId, &temp.CourierType, pq.Array(&temp.Regions), pq.Array(&temp.WorkingHours))
    /*
       if err := r.db.Conn.SelectContext(ctx, &orders, queryGetAllOrders, limit, offset); err != nil {
           if errors.Is(err, sql.ErrNoRows) {
               return orders, nil
           } else {
               return nil, fmt.Errorf("OrderRepository - GetOrders: %w", err)
           }
       }
    */
    return orders, nil
}

// todo create all at once or one by one; and if tx rollbacks?

func (r *OrderRepository) CreateOrders(ctx context.Context, ords *model.CreateOrderRequest) ([]model.OrderDto, error) {

    resOrd := make([]model.OrderDto, 0, len(ords.Orders))
    log.Debug(ords.Orders, "request orders")

    tx, err := r.db.Conn.BeginTxx(ctx, &sql.TxOptions{}) // nil
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - AddOrder: %w", err)
    }
    defer tx.Rollback()

    for _, o := range ords.Orders {
        var orderId int64

        err := tx.QueryRowContext(ctx, queryCreateOrder, o.Weight, o.Regions, pq.Array(&o.DeliveryHours), o.Cost).Scan(&orderId)
        if err != nil {
            return nil, fmt.Errorf("OrderRepository - CreateOrders: %w", err)
        }

        //res, err := tx.NamedExecContext(ctx, q, o)
        //if err != nil {
        //    return nil, fmt.Errorf("OrderRepository - AddOrder: %w", err)
        //}
        //id, _ := res.LastInsertId()
        temp := model.OrderDto{
            OrderId:       orderId,
            Weight:        o.Weight,
            Regions:       o.Regions,
            DeliveryHours: o.DeliveryHours,
            Cost:          o.Cost,
        }
        resOrd = append(resOrd, temp)
    }
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("OrderRepository - CreateOrders: %w", err)
    }
    log.Infof("orderRepo - Create - result %v", resOrd)
    return resOrd, nil
}

func (r *OrderRepository) UpdateOrders(ctx context.Context, ords []model.CompleteOrder) ([]model.OrderDto, error) {

    tx, err := r.db.Conn.BeginTxx(ctx, &sql.TxOptions{}) // nil
    if err != nil {
        return nil, fmt.Errorf("OrderRepository - AddOrder: %w", err)
    }
    defer tx.Rollback()

    resOrders := make([]model.OrderDto, 0, len(ords))
    var order model.OrderDto
    var id int64
    for _, o := range ords {
        err = tx.QueryRowContext(ctx, queryGetByIdWithCourierId, o.OrderId).
            Scan(&id, &order.OrderId, &order.Weight, &order.Regions, &order.DeliveryHours, &order.Cost, &order.CompletedTime)
        if err != nil {
            return nil, err
        }
        if id == 0 || id != o.CourierId {
            return nil, errors.New("invalid courier id")
        }

        if order.CompletedTime == nil {
            temp := o.CompleteTime
            order.CompletedTime = &temp
            if _, err = tx.ExecContext(ctx, queryUpdateOrder, o.CompleteTime, o.CourierId, o.OrderId); err != nil {
                return nil, fmt.Errorf("OrderRepository - UpdateOrders: %w", err)
            }
        }
        resOrders = append(resOrders, order)
    }
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("OrderRepository - Upadate: %w", err)
    }
    log.Infof("orderRepo - Update - result %v", resOrders)
    return resOrders, nil
}
