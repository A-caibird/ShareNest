package db

import (
	"database/sql"
	"fmt"
)

func Query[T any](db *sql.DB, query string, args ...interface{}) ([]T, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare error: %w", err)
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	var results []T

	for rows.Next() {
		var result T
		// 假设 T 是可以通过 rows.Scan 直接映射的类型
		err := rows.Scan(&result)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return results, nil
}
