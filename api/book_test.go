package api

import (
	"fmt"
	"testing"
)

func TestToJSON(t *testing.T) {
	fmt.Println("TestToJSON")
	b := Book{Title: "Test Title", Author: "Test Author", ISBN: "100001"}
	bJSON, err := ToJSON(b)
	if err != nil {
		fmt.Printf("ToJSON failed with err: %v\n", err)
		t.FailNow()
	}
	expectedJSON := "{\"title\":\"Test Title\",\"author\":\"Test Author\",\"isbn\":\"100001\"}"
	if string(bJSON) != expectedJSON {
		fmt.Printf("TestToJSON: Expected: %v | Actual: %v", expectedJSON, string(bJSON))
		t.FailNow()
	}
	fmt.Println("Success")
}

func TestFromJSON(t *testing.T) {
	fmt.Println("TestFromJSON")
	bJSON := "{\"title\":\"Test Title\",\"author\":\"Test Author\",\"isbn\":\"100001\"}"
	b, err := FromJSON([]byte(bJSON))
	if err != nil {
		fmt.Printf("FromJSON failed with err: %v\n", err)
		t.FailNow()
	}
	expectedBook := Book{Title: "Test Title", Author: "Test Author", ISBN: "100001"}
	if *b != expectedBook {
		fmt.Printf("TestToJSON: Expected: %v | Actual: %v", expectedBook, *b)
		t.FailNow()
	}
	fmt.Println("Success")
}
