package mongodb

import (
	"reflect"
	"testing"
)

func TestGetActions(t *testing.T) {
	privilege := PrivilegeDto{Db: "test", Collection: "test", Actions: []string{"find", "update", "remove", "insert"}}

	expectedActions := []string{"find", "insert", "remove", "update"}
	var actions []string = getActions(privilege)

	if reflect.DeepEqual(actions, expectedActions) == false {
		t.Errorf("Obtained actions = %v; want %v", actions, expectedActions)
	}
}

func TestGetPrivilegesFromDto(t *testing.T) {
	privilegesDto := PrivilegeDto{Db: "test", Collection: "test", Actions: []string{"remove", "update", "insert", "find"}}
	privilegesDto2 := PrivilegeDto{Db: "test", Collection: "test2", Actions: []string{"remove", "update", "find"}}

	expectedPrivilege1 := Privilege{Resource{Db: privilegesDto.Db, Collection: privilegesDto.Collection}, getActions(privilegesDto)}
	expectedPrivilege2 := Privilege{Resource{Db: privilegesDto2.Db, Collection: privilegesDto2.Collection}, getActions(privilegesDto2)}

	expectedPrivileges := []Privilege{expectedPrivilege1, expectedPrivilege2}
	privileges := getPrivilegesFromDto([]PrivilegeDto{privilegesDto, privilegesDto2})

	if reflect.DeepEqual(expectedPrivileges, privileges) == false {
		t.Errorf("Obtained privileges = %v; want %v", privileges, expectedPrivileges)
	}

}
