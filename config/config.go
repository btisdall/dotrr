package config

import (
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

// Map represents parsed configuration
type Map map[string]string

// Item represents a key/value configuration pair
type Item struct {
	Key   string
	Value string
}

// NewItem returns a new Item from a key and value
func NewItem(k, v string) Item {
	return Item{
		Key:   k,
		Value: v,
	}
}

// Read reads a parsed dotenv file
func Read(file string) (Map, error) {
	return godotenv.Read(file)
}


// Marshal returns a stringified Map lexically sorted by line
func Marshal(m Map) string {
	config := make([]string, 0)
	for k, v := range m {
		line, _ := godotenv.Marshal(Map{k: v})
		config = append(config, line)
	}
	sort.Strings(config)
	return strings.Join(config, "\n")
}
