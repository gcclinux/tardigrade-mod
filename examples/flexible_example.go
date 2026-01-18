package main

import (
	"encoding/json"
	"fmt"

	"github.com/gcclinux/tardigrade-mod"
)

func main() {
	tar := tardigrade.Tardigrade{}

	// Create database
	tar.CreateDB("flexible.db")

	// Example 1: Using map for multiple fields
	fmt.Println("=== Example 1: Map-based approach ===")
	tar.AddFlexField("user:1", map[string]string{
		"name":     "ricardo wagemaker",
		"status":   "married",
		"location": "london",
	}, "flexible.db")

	// Example 2: Using variadic arguments (easier syntax)
	fmt.Println("\n=== Example 2: Variadic approach ===")
	tar.AddFlexFieldVariadic("app:2", "flexible.db",
		"cost", "299",
		"billing", "monthly",
		"patch", "17",
		"color", "blue",
		"os", "linux",
		"mode", "auto",
	)

	// Example 3: Simple record
	fmt.Println("\n=== Example 3: Simple record ===")
	tar.AddFlexFieldVariadic("item:3", "flexible.db",
		"name", "Widget",
		"price", "49.99",
	)

	// Retrieve and display records
	fmt.Println("\n=== Retrieving Records ===")

	// Get full record as JSON
	result1 := tar.SelectFlexByID(1, "json", "flexible.db")
	fmt.Println("Record 1 (JSON):")
	fmt.Println(result1)

	result2 := tar.SelectFlexByID(2, "json", "flexible.db")
	fmt.Println("\nRecord 2 (JSON):")
	fmt.Println(result2)

	// Get specific field
	fmt.Println("\n=== Getting Specific Fields ===")
	name := tar.GetFlexField(1, "name", "flexible.db")
	fmt.Printf("User 1 name: %s\n", name)

	cost := tar.GetFlexField(2, "cost", "flexible.db")
	fmt.Printf("App 2 cost: %s\n", cost)

	// List all fields in a record
	fmt.Println("\n=== Listing All Fields ===")
	fields := tar.ListFlexFields(2, "flexible.db")
	fmt.Printf("App 2 has fields: %v\n", fields)

	// Search across records
	fmt.Println("\n=== Searching Records ===")
	format, results := tar.SelectFlexSearch("london", "json", "flexible.db")
	fmt.Printf("Search results (%s):\n", format)

	var records []tardigrade.FlexStruct
	json.Unmarshal(results, &records)
	for _, rec := range records {
		fmt.Printf("ID: %d, Key: %s, Fields: %v\n", rec.Id, rec.Key, rec.Fields)
	}

	// Modify a record
	fmt.Println("\n=== Modifying Record ===")
	tar.ModifyFlexField(1, "user:1", map[string]string{
		"name":     "ricardo wagemaker",
		"status":   "married",
		"location": "paris", // Changed from london
		"age":      "35",    // Added new field
	}, "flexible.db")

	updated := tar.SelectFlexByID(1, "json", "flexible.db")
	fmt.Println("Updated Record 1:")
	fmt.Println(updated)

	// Count records
	fmt.Println("\n=== Database Stats ===")
	count := tar.CountSize("flexible.db")
	fmt.Printf("Total records: %d\n", count)

	// Clean up (optional)
	// tar.DeleteDB("flexible.db")
}
