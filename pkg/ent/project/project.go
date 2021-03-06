// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package project

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"

	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// EdgeLocation holds the string denoting the location edge name in mutations.
	EdgeLocation = "location"
	// EdgeComments holds the string denoting the comments edge name in mutations.
	EdgeComments = "comments"
	// EdgeWorkOrders holds the string denoting the work_orders edge name in mutations.
	EdgeWorkOrders = "work_orders"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"

	// Table holds the table name of the project in the database.
	Table = "projects"
	// TypeTable is the table the holds the type relation/edge.
	TypeTable = "projects"
	// TypeInverseTable is the table name for the ProjectType entity.
	// It exists in this package in order to avoid circular dependency with the "projecttype" package.
	TypeInverseTable = "project_types"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "project_type_projects"
	// TemplateTable is the table the holds the template relation/edge.
	TemplateTable = "projects"
	// TemplateInverseTable is the table name for the ProjectTemplate entity.
	// It exists in this package in order to avoid circular dependency with the "projecttemplate" package.
	TemplateInverseTable = "project_templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "project_template"
	// LocationTable is the table the holds the location relation/edge.
	LocationTable = "projects"
	// LocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationInverseTable = "locations"
	// LocationColumn is the table column denoting the location relation/edge.
	LocationColumn = "project_location"
	// CommentsTable is the table the holds the comments relation/edge.
	CommentsTable = "comments"
	// CommentsInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentsInverseTable = "comments"
	// CommentsColumn is the table column denoting the comments relation/edge.
	CommentsColumn = "project_comments"
	// WorkOrdersTable is the table the holds the work_orders relation/edge.
	WorkOrdersTable = "work_orders"
	// WorkOrdersInverseTable is the table name for the WorkOrder entity.
	// It exists in this package in order to avoid circular dependency with the "workorder" package.
	WorkOrdersInverseTable = "work_orders"
	// WorkOrdersColumn is the table column denoting the work_orders relation/edge.
	WorkOrdersColumn = "project_work_orders"
	// PropertiesTable is the table the holds the properties relation/edge.
	PropertiesTable = "properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// PropertiesColumn is the table column denoting the properties relation/edge.
	PropertiesColumn = "project_properties"
	// CreatorTable is the table the holds the creator relation/edge.
	CreatorTable = "projects"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "project_creator"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldDescription,
	FieldPriority,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Project type.
var ForeignKeys = []string{
	"project_template",
	"project_location",
	"project_creator",
	"project_type_projects",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/facebookincubator/symphony/pkg/ent/runtime"
//
var (
	Hooks  [2]ent.Hook
	Policy ent.Policy
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Priority defines the type for the priority enum field.
type Priority string

// PriorityNone is the default Priority.
const DefaultPriority = PriorityNone

// Priority values.
const (
	PriorityUrgent Priority = "URGENT"
	PriorityHigh   Priority = "HIGH"
	PriorityMedium Priority = "MEDIUM"
	PriorityLow    Priority = "LOW"
	PriorityNone   Priority = "NONE"
)

func (pr Priority) String() string {
	return string(pr)
}

// PriorityValidator is a validator for the "priority" field enum values. It is called by the builders before save.
func PriorityValidator(pr Priority) error {
	switch pr {
	case PriorityUrgent, PriorityHigh, PriorityMedium, PriorityLow, PriorityNone:
		return nil
	default:
		return fmt.Errorf("project: invalid enum value for priority field: %q", pr)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (pr Priority) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(pr.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (pr *Priority) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*pr = Priority(str)
	if err := PriorityValidator(*pr); err != nil {
		return fmt.Errorf("%s is not a valid Priority", str)
	}
	return nil
}
