package repository

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "github.com/labstack/gommon/log"
    "github.com/lib/pq"
    "time"
    "yandex-team.ru/bstask/database"
    "yandex-team.ru/bstask/model"
)

const (
    queryGetCourierById = `SELECT courier_id, courier_type, regions, working_hours FROM couriers WHERE courier_id = $1`
    queryGetAllCouriers = `SELECT courier_id, courier_type, regions, working_hours FROM couriers LIMIT $1 OFFSET $2`
    queryCreateCourier  = `INSERT INTO couriers (courier_type, regions, working_hours) VALUES ($1, $2, $3) RETURNING courier_id`
    queryGetEarnings    = `SELECT cost FROM orders WHERE completed_time >= $1 AND completed_time < $2 AND cour_id = $3`
)

type CourierRepository struct {
    db *database.Database
}

func NewCourierRepository(dbConn *database.Database) *CourierRepository {
    return &CourierRepository{
        db: dbConn,
    }
}

func (r *CourierRepository) GetById(ctx context.Context, id int64) (*model.Courier, error) {
    var courier model.Courier

    //err := r.db.Conn.GetContext(ctx, &courier, query, id)
    err := r.db.Conn.QueryRowContext(ctx, queryGetCourierById, id).
        Scan(&courier.CourierId, &courier.CourierType, pq.Array(&courier.Regions), pq.Array(&courier.WorkingHours))
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        } else {
            return nil, fmt.Errorf("CourierRepository - GetByID: %w", err)
        }
    }

    return &courier, nil
}

func (r *CourierRepository) GetCouriers(ctx context.Context, limit, offset int32) ([]*model.Courier, error) {
    couriers := make([]*model.Courier, 0, limit)

    rows, err := r.db.Conn.QueryContext(ctx, queryGetAllCouriers, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("CourierRepository - GetCouriers: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        temp := model.Courier{}
        err := rows.Scan(&temp.CourierId, &temp.CourierType, pq.Array(&temp.Regions), pq.Array(&temp.WorkingHours))
        if err != nil {
            return nil, fmt.Errorf("CourierRepository - GetCouriers: %w", err)
        }
        couriers = append(couriers, &temp)
    }

    return couriers, nil
}

// todo create all at once or one by one; and if tx rollbacks?

func (r *CourierRepository) CreateCouriers(ctx context.Context, couriers []*model.CreateCourier) ([]*model.Courier, error) {
    resCours := make([]*model.Courier, 0, len(couriers))

    tx, err := r.db.Conn.BeginTxx(ctx, &sql.TxOptions{}) // nil
    if err != nil {
        return nil, fmt.Errorf("CourierRepository - AddCourier: %w", err)
    }
    defer tx.Rollback()

    var courierId int64
    for _, c := range couriers {
        err := tx.QueryRowContext(ctx, queryCreateCourier, c.CourierType, pq.Array(c.Regions), pq.Array(c.WorkingHours)).Scan(&courierId)
        if err != nil {
            return nil, fmt.Errorf("CourierRepository - CreateCouriers: %w", err)
        }
        temp := model.Courier{
            CourierId:    courierId,
            CourierType:  c.CourierType,
            Regions:      c.Regions,
            WorkingHours: c.WorkingHours,
        }
        resCours = append(resCours, &temp)
    }
    if err = tx.Commit(); err != nil {
        return nil, fmt.Errorf("CourierRepository - CreateCouriers: %w", err)
    }
    log.Info(resCours, "courierRepo - Create - result")
    return resCours, nil
}

func (r *CourierRepository) GetEarnings(ctx context.Context, id int64, startDate, endDate time.Time) ([]int32, error) {
    earnings := make([]int32, 0, 0)
    err := r.db.Conn.SelectContext(ctx, &earnings, queryGetEarnings, startDate, endDate, id)
    if err != nil {
        return nil, fmt.Errorf("CourierRepository - GetEarnings: %w", err)
    }
    return earnings, nil
}
