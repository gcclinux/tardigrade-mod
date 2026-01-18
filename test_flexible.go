package main

import (
	"fmt"
	"os"

	"github.com/gcclinux/tardigrade-mod"
)

func main() {
	tar := tardigrade.Tardigrade{}

	// Clean up any existing test database
	os.Remove("test_flex.db")

	// Test 1: Create database and add flexible record
	fmt.Println("=== Test 1: AddFlexFieldVariadic ===")
	tar.CreateDB("test_flex.db")
	success := tar.AddFlexFieldVariadic("user:1", "test_flex.db",
		"name", "ricardo wagemaker",
		"status", "married",
		"location", "london")
	fmt.Printf("Add record: %v\n", success)

	// Test 2: Add another record with different fields
	fmt.Println("\n=== Test 2: Add record with many fields ===")
	success = tar.AddFlexFieldVariadic("app:2", "test_flex.db",
		"cost", "299",
		"billing", "monthly",
		"patch", "17",
		"color", "blue",
		"os", "linux",
		"mode", "auto")
	fmt.Printf("Add record: %v\n", success)

	// Test 3: Retrieve by ID
	fmt.Println("\n=== Test 3: SelectFlexByID ===")
	result := tar.SelectFlexByID(1, "json", "test_flex.db")
	fmt.Println(result)

	// Test 4: Get specific field
	fmt.Println("\n=== Test 4: GetFlexField ===")
	name := tar.GetFlexField(1, "name", "test_flex.db")
	fmt.Printf("Name: %s\n", name)

	location := tar.GetFlexField(1, "location", "test_flex.db")
	fmt.Printf("Location: %s\n", location)

	// Test 5: List all fields
	fmt.Println("\n=== Test 5: ListFlexFields ===")
	fields := tar.ListFlexFields(2, "test_flex.db")
	fmt.Printf("App fields: %v\n", fields)

	// Test 6: Search
	fmt.Println("\n=== Test 6: SelectFlexSearch ===")
	format, results := tar.SelectFlexSearch("london", "json", "test_flex.db")
	fmt.Printf("Search format: %s\n", format)
	fmt.Printf("Results: %s\n", string(results))

	// Test 7: Modify record
	fmt.Println("\n=== Test 7: ModifyFlexField ===")
	msg, status := tar.ModifyFlexField(1, "user:1", map[string]string{
		"name":     "ricardo wagemaker",
		"status":   "married",
		"location": "paris",
		"age":      "35",
	}, "test_flex.db")
	fmt.Printf("Modified: %v, Status: %v\n", msg, status)

	// Verify modification
	updated := tar.SelectFlexByID(1, "json", "test_flex.db")
	fmt.Println("Updated record:")
	fmt.Println(updated)

	// Test 8: Count records
	fmt.Println("\n=== Test 8: CountSize ===")
	count := tar.CountSize("test_flex.db")
	fmt.Printf("Total records: %d\n", count)

	// Clean up
	fmt.Println("\n=== Cleanup ===")
	os.Remove("test_flex.db")
	fmt.Println("Test database removed")

	fmt.Println("\n✅ All flexible field tests completed successfully!")
}
