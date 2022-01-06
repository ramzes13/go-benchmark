package main

import (
	"database/sql"
	"testing"
)

var result = map[string]interface{}{}

func BenchmarkGetDifferenceSimpleEmpty(b *testing.B) {
	s1 := TestStruct{}
	s2 := TestStruct{}
	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifference(s1, s2)
	}

	result = r
}

func BenchmarkGetDifferenceSimpleLeft(b *testing.B) {
	s1 := TestStruct{
		Id:    UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa3"),
		Name:  "Name",
		Name1: sql.NullString{Valid: true, String: "Name1"},
		Name2: sql.NullString{Valid: true, String: "Name2"},
		Name3: UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa1"),
		Name4: 4,
		Name5: sql.NullString{Valid: true, String: "Name5"},
		Name6: sql.NullString{Valid: true, String: "Name6"},
		Name7: true,
	}
	s2 := TestStruct{}
	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifference(s1, s2)
	}

	result = r
}

func BenchmarkGetDifferenceSimpleRight(b *testing.B) {
	s1 := TestStruct{}
	s2 := TestStruct{
		Id:    UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa3"),
		Name:  "Name",
		Name1: sql.NullString{Valid: true, String: "Name1"},
		Name2: sql.NullString{Valid: true, String: "Name2"},
		Name3: UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa1"),
		Name4: 4,
		Name5: sql.NullString{Valid: true, String: "Name5"},
		Name6: sql.NullString{Valid: true, String: "Name6"},
		Name7: true,
	}

	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifference(s1, s2)
	}

	result = r
}

func BenchmarkGetDifferenceAutoEmpty(b *testing.B) {
	s1 := TestStruct{}
	s2 := TestStruct{}
	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifferenceAuto(s1, s2)
	}

	result = r
}

func BenchmarkGetDifferenceAutoLeft(b *testing.B) {
	s1 := TestStruct{
		Id:    UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa3"),
		Name:  "Name",
		Name1: sql.NullString{Valid: true, String: "Name1"},
		Name2: sql.NullString{Valid: true, String: "Name2"},
		Name3: UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa1"),
		Name4: 4,
		Name5: sql.NullString{Valid: true, String: "Name5"},
		Name6: sql.NullString{Valid: true, String: "Name6"},
		Name7: true,
	}
	s2 := TestStruct{}
	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifferenceAuto(s1, s2)
	}

	result = r
}

func BenchmarkGetDifferenceAutoRight(b *testing.B) {
	s1 := TestStruct{}
	s2 := TestStruct{
		Id:    UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa3"),
		Name:  "Name",
		Name1: sql.NullString{Valid: true, String: "Name1"},
		Name2: sql.NullString{Valid: true, String: "Name2"},
		Name3: UuidFromString("a229c54d-327f-455a-9d61-03083e2f6aa1"),
		Name4: 4,
		Name5: sql.NullString{Valid: true, String: "Name5"},
		Name6: sql.NullString{Valid: true, String: "Name6"},
		Name7: true,
	}
	var r = map[string]interface{}{}

	for i := 0; i < b.N; i++ {
		r, _ = getDifferenceAuto(s1, s2)
	}

	result = r
}
