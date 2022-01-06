package main

import (
	"database/sql"

	"github.com/jackc/pgtype"
	"github.com/ramzes13/structs"
)

func UuidFromString(strUuid string) pgtype.UUID {
	var response pgtype.UUID
	response.Scan(strUuid)
	return response
}

type TestStruct struct {
	Id    pgtype.UUID    `json:"id" db:"id"`
	Name  string         `json:"name" db:"name"`
	Name1 sql.NullString `json:"name1" db:"name_1"`
	Name2 sql.NullString `json:"name2" db:"name_2"`
	Name3 pgtype.UUID    `json:"name3" db:"name_3"`
	Name4 int            `json:"name4" db:"name_4"`
	Name5 sql.NullString `json:"name5" db:"name_5"`
	Name6 sql.NullString `json:"name6" db:"name_6"`
	Name7 bool           `json:"name7" db:"name_7"`
}

func getDifference(is1 interface{}, is2 interface{}) (map[string]interface{}, error) {
	response := map[string]interface{}{}
	s1 := is1.(TestStruct)
	s2 := is2.(TestStruct)

	if s1.Name != s2.Name {
		response["name"] = s1.Name
	}

	if s1.Name1 != s2.Name1 {
		response["name_1"] = s1.Name1
	}

	if s1.Name2 != s2.Name2 {
		response["name_2"] = s1.Name2
	}

	if s1.Name3 != s2.Name3 {
		response["name_3"] = s1.Name3
	}

	if s1.Name4 != s2.Name4 {
		response["name_4"] = s1.Name4
	}

	if s1.Name5 != s2.Name5 {
		response["name_5"] = s1.Name5
	}

	if s1.Name6 != s2.Name6 {
		response["name_6"] = s1.Name6
	}
	if s1.Name7 != s2.Name7 {
		response["name_7"] = s1.Name7
	}

	return response, nil
}

func getDifferenceAuto(is1 interface{}, is2 interface{}) (map[string]interface{}, error) {
	return structs.GenerateDiff(is1, is2), nil
}

func main() {}
