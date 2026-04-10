package database

import (
	"database/sql/driver"
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

type Participants []uint32

func (p Participants) Value() (driver.Value, error) {
	return p, nil
}

func (l *Participants) Scan(value any) error {
	if value == nil {
		return nil
	}

	// * Ensure the value is in the right format
	var pgArray string
	switch v := value.(type) {
	case []byte:
		pgArray = string(v)
	case string:
		pgArray = v
	default:
		return fmt.Errorf("unsupported Scan type for List: %T", value)
	}

	// * Postgres formats arrays in curly braces,
	// * such as `{"string1"}`
	if len(pgArray) < 2 || pgArray[0] != '{' || pgArray[len(pgArray)-1] != '}' {
		return fmt.Errorf("invalid PostgreSQL array format: %s", pgArray)
	}

	pgArray = strings.TrimSuffix(pgArray, "}")
	pgArray = strings.TrimPrefix(pgArray, "{")

	// * Bail if array is empty, who cares
	if pgArray == "" {
		return nil
	}

	var elements []string
	var err error

	reader := csv.NewReader(strings.NewReader(pgArray))
	reader.Comma = ','

	elements, err = reader.Read()

	if err != nil {
		return nil
	}

	result := make(Participants, 0, len(elements))

	for _, element := range elements {
		i, err := strconv.ParseUint(element, 10, 32)
		if err != nil {
			return err
		}

		result = append(result, uint32(i))
	}

	*l = result
	return nil
}
