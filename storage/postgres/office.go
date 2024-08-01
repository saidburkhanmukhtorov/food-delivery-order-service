package postgres

import (
	"context"
	"fmt"

	"github.com/food-delivery/food-delivery-order-service/genproto/order"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type OfficeRepo struct {
	db *pgx.Conn
}

func NewOfficeRepo(db *pgx.Conn) *OfficeRepo {
	return &OfficeRepo{
		db: db,
	}
}

func (r *OfficeRepo) CreateOffice(ctx context.Context, req *order.CreateOfficeRequest) (*order.Office, error) {
	if req.Office.Id == "" {
		req.Office.Id = uuid.NewString()
	}
	query := `
		INSERT INTO offices (
			id,
			name,
			address,
			latitude,
			longitude
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id
	`

	err := r.db.QueryRow(ctx, query,
		req.Office.Id,
		req.Office.Name,
		req.Office.Address,
		req.Office.Latitude,
		req.Office.Longitude,
	).Scan(&req.Office.Id)

	if err != nil {
		return nil, err
	}

	return req.Office, nil
}

func (r *OfficeRepo) GetOfficeByID(ctx context.Context, id string) (*order.Office, error) {
	var (
		officeModel order.Office
	)
	query := `
		SELECT 
			id,
			name,
			address,
			latitude,
			longitude
		FROM offices
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&officeModel.Id,
		&officeModel.Name,
		&officeModel.Address,
		&officeModel.Latitude,
		&officeModel.Longitude,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &officeModel, nil
}

func (r *OfficeRepo) GetAllOffices(ctx context.Context, req *order.GetOfficesRequest) (*order.GetOfficesResponse, error) {
	var args []interface{}
	count := 1
	query := `
		SELECT 
			id,
			name,
			address,
			latitude,
			longitude
		FROM 
			offices
		WHERE 1=1
	`

	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" AND name ILIKE $%d", count)
		args = append(args, "%"+req.Name+"%")
		count++
	}

	if req.Address != "" {
		filter += fmt.Sprintf(" AND address ILIKE $%d", count)
		args = append(args, "%"+req.Address+"%")
		count++
	}

	query += filter
	// Handle invalid page or limit values
	if req.Page <= 0 {
		req.Page = 1 // Default to page 1
	}
	if req.Limit <= 0 {
		req.Limit = 10 // Default to a limit of 10
	}

	// Add LIMIT and OFFSET for pagination using the proto fields
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
	args = append(args, req.Limit, (req.Page-1)*req.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var officeList []*order.Office

	for rows.Next() {
		var (
			officeModel order.Office
		)
		err = rows.Scan(
			&officeModel.Id,
			&officeModel.Name,
			&officeModel.Address,
			&officeModel.Latitude,
			&officeModel.Longitude,
		)
		if err != nil {
			return nil, err
		}
		officeList = append(officeList, &officeModel)
	}

	return &order.GetOfficesResponse{
		Offices: officeList,
	}, nil
}

func (r *OfficeRepo) UpdateOffice(ctx context.Context, req *order.UpdateOfficeRequest) (*order.Office, error) {
	query := `
		UPDATE offices
		SET 
			name = $1,
			address = $2,
			latitude = $3,
			longitude = $4
		WHERE id = $5
		RETURNING id, name, address, latitude, longitude
	`

	err := r.db.QueryRow(ctx, query,
		req.Office.Name,
		req.Office.Address,
		req.Office.Latitude,
		req.Office.Longitude,
		req.Office.Id,
	).Scan(
		&req.Office.Id,
		&req.Office.Name,
		&req.Office.Address,
		&req.Office.Latitude,
		&req.Office.Longitude,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return req.Office, nil
}

func (r *OfficeRepo) PatchOffice(ctx context.Context, req *order.PatchOfficeRequest) (*order.Office, error) {
	var args []interface{}
	count := 1
	query := `
		UPDATE offices
		SET 
	`

	filter := ""

	if req.Name != "" {
		filter += fmt.Sprintf(" name = $%d, ", count)
		args = append(args, req.Name)
		count++
	}

	if req.Address != "" {
		filter += fmt.Sprintf(" address = $%d, ", count)
		args = append(args, req.Address)
		count++
	}

	if req.Latitude != 0 {
		filter += fmt.Sprintf(" latitude = $%d, ", count)
		args = append(args, req.Latitude)
		count++
	}

	if req.Longitude != 0 {
		filter += fmt.Sprintf(" longitude = $%d, ", count)
		args = append(args, req.Longitude)
		count++
	}

	if filter == "" {
		return nil, fmt.Errorf("at least one field to update is required")
	}

	filter = filter[:len(filter)-2] // Remove the trailing comma and space
	query += filter + fmt.Sprintf(" WHERE id = $%d RETURNING id, name, address, latitude, longitude", count)
	args = append(args, req.Id)

	var response order.Office
	err := r.db.QueryRow(ctx, query, args...).Scan(
		&response.Id,
		&response.Name,
		&response.Address,
		&response.Latitude,
		&response.Longitude,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, pgx.ErrNoRows
		}
		return nil, err
	}

	return &response, nil
}

func (r *OfficeRepo) DeleteOffice(ctx context.Context, id string) error {
	query := `
		DELETE FROM offices
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}
