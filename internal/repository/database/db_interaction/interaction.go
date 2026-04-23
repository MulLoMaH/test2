package dbinteraction

import (
	"context"
	"errors"
	"records/internal/employee"
	"records/internal/repository/database/model"
	"records/internal/repository/interface_employee"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{
		pool: pool,
	}
}

var _ interface_employee.Employee_Repository = (*PostgresRepo)(nil)

func (p *PostgresRepo) Add_employee(ctx context.Context, fullname string, position_company string) error {
	sqlQuery := `
	INSERT INTO Employee (fullname, position_company, reception_at)
	VALUES ($1, $2, $3);
	`

	_, err := p.pool.Exec(ctx, sqlQuery, fullname, position_company, time.Now())

	return err
}

func (p *PostgresRepo) GetAll(ctx context.Context) (map[string]*employee.Employee, error) {
	sqlQuery := `
	SELECT id, fullname, position_company, reception_at
	FROM Employee;
	`

	rows, err := p.pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]*employee.Employee)
	for rows.Next() {
		var dbModel model.DbEmployee
		if err := rows.Scan(
			&dbModel.ID,
			&dbModel.Fullname,
			&dbModel.Position,
			&dbModel.Reception_at,
		); err != nil {
			return nil, err
		}

		empl := employee.NewEmployeeFromDB(
			dbModel.ID,
			dbModel.Fullname,
			dbModel.Position,
			dbModel.Reception_at,
		)

		result[dbModel.Fullname] = empl
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("no employees")
	}

	return result, nil
}

func (p *PostgresRepo) Delete_Employee(ctx context.Context, fullname string) error {
	sqlQuery := `
	DELETE 
	FROM Employee
	WHERE fullname=$1
	`

	cmdTag, err := p.pool.Exec(ctx, sqlQuery, fullname)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("employee not found")
	}

	return nil
}
