package godror

import (
  "database/sql"
  log "wrk-connector/src/shared/infrastructure/config"
)

func ExecuteQuery(conn *sql.DB, query string) ([][]interface{}, error) {
	var result [][]interface{}
	rows, err := conn.Query(query)
	if err != nil {
		log.WithError(err).Fatal("error executing query")
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.WithError(err).Error("error closing DB")
		}
	}(rows)

	cols, _ := rows.Columns()

	for rows.Next() {
		values := make([]interface{}, len(cols))
		scanArgs := make([]interface{}, len(cols))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		_ = rows.Scan(scanArgs...)
		result = append(result, values)
	}

	return result, nil
}
