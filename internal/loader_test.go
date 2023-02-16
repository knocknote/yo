package internal

import (
	"fmt"
	"testing"

	"github.com/knocknote/yo/models"
)

func Test_setIndexesToTables(t *testing.T) {
	tests := []struct {
		table  map[string]*Type
		ix     map[string]*Index
		result map[string]int
	}{
		{
			table: map[string]*Type{
				"TableA": &Type{
					Indexes: []*Index{},
				},
			},
			ix: map[string]*Index{
				"TableA_Index1": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
				"TableA_Index2": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
			},
			result: map[string]int{
				"TableA": 2,
			},
		},
		{
			table: map[string]*Type{
				"TableA": &Type{
					Indexes: []*Index{},
				},
				"TableB": &Type{
					Indexes: []*Index{},
				},
			},
			ix: map[string]*Index{
				"TableA_Index1": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
				"TableA_Index2": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
			},
			result: map[string]int{
				"TableA": 2,
				"TableB": 0,
			},
		},
		{
			table: map[string]*Type{
				"TableA": &Type{
					Indexes: []*Index{},
				},
				"TableB": &Type{
					Indexes: []*Index{},
				},
			},
			ix: map[string]*Index{
				"TableA_Index1": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
				"TableA_Index2": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableA",
						},
					},
				},
				"TableB_Index1": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableB",
						},
					},
				},
				"TableB_Index2forTableA_Hoge": &Index{
					Type: &Type{
						Table: &models.Table{
							TableName: "TableB",
						},
					},
				},
			},
			result: map[string]int{
				"TableA": 2,
				"TableB": 2,
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			setIndexesToTables(tt.table, tt.ix)
			for k, v := range tt.table {
				if len(v.Indexes) != tt.result[k] {
					t.Errorf("error. want:%d got:%d", tt.result[k], len(v.Indexes))
				}
			}
		})
	}
}
