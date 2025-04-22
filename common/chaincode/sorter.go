package analysis

import (
	"fmt"
	"os"
)

// WriteExecutionOrder writes final ordered functions to a file.
func WriteExecutionOrder(filePath string, orders []ShardOrder) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, shard := range orders {
		fmt.Fprintf(file, "Shard %d:\n", shard.ShardID)
		for _, fn := range shard.Order {
			fmt.Fprintf(file, "  %s\n", fn)
		}
	}
	return nil
}
