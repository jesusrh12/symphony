// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by entc, DO NOT EDIT.

package checklistitem

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/facebook/ent"
	"github.com/facebookincubator/symphony/pkg/ent/schema/enum"
)

const (
	// Label holds the string label denoting the checklistitem type in the database.
	Label = "check_list_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldIsMandatory holds the string denoting the is_mandatory field in the database.
	FieldIsMandatory = "is_mandatory"
	// FieldChecked holds the string denoting the checked field in the database.
	FieldChecked = "checked"
	// FieldStringVal holds the string denoting the string_val field in the database.
	FieldStringVal = "string_val"
	// FieldEnumValues holds the string denoting the enum_values field in the database.
	FieldEnumValues = "enum_values"
	// FieldEnumSelectionModeValue holds the string denoting the enum_selection_mode_value field in the database.
	FieldEnumSelectionModeValue = "enum_selection_mode_value"
	// FieldSelectedEnumValues holds the string denoting the selected_enum_values field in the database.
	FieldSelectedEnumValues = "selected_enum_values"
	// FieldYesNoVal holds the string denoting the yes_no_val field in the database.
	FieldYesNoVal = "yes_no_val"
	// FieldHelpText holds the string denoting the help_text field in the database.
	FieldHelpText = "help_text"

	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeWifiScan holds the string denoting the wifi_scan edge name in mutations.
	EdgeWifiScan = "wifi_scan"
	// EdgeCellScan holds the string denoting the cell_scan edge name in mutations.
	EdgeCellScan = "cell_scan"
	// EdgeCheckListCategory holds the string denoting the check_list_category edge name in mutations.
	EdgeCheckListCategory = "check_list_category"

	// Table holds the table name of the checklistitem in the database.
	Table = "check_list_items"
	// FilesTable is the table the holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "check_list_item_files"
	// WifiScanTable is the table the holds the wifi_scan relation/edge.
	WifiScanTable = "survey_wi_fi_scans"
	// WifiScanInverseTable is the table name for the SurveyWiFiScan entity.
	// It exists in this package in order to avoid circular dependency with the "surveywifiscan" package.
	WifiScanInverseTable = "survey_wi_fi_scans"
	// WifiScanColumn is the table column denoting the wifi_scan relation/edge.
	WifiScanColumn = "survey_wi_fi_scan_checklist_item"
	// CellScanTable is the table the holds the cell_scan relation/edge.
	CellScanTable = "survey_cell_scans"
	// CellScanInverseTable is the table name for the SurveyCellScan entity.
	// It exists in this package in order to avoid circular dependency with the "surveycellscan" package.
	CellScanInverseTable = "survey_cell_scans"
	// CellScanColumn is the table column denoting the cell_scan relation/edge.
	CellScanColumn = "survey_cell_scan_checklist_item"
	// CheckListCategoryTable is the table the holds the check_list_category relation/edge.
	CheckListCategoryTable = "check_list_items"
	// CheckListCategoryInverseTable is the table name for the CheckListCategory entity.
	// It exists in this package in order to avoid circular dependency with the "checklistcategory" package.
	CheckListCategoryInverseTable = "check_list_categories"
	// CheckListCategoryColumn is the table column denoting the check_list_category relation/edge.
	CheckListCategoryColumn = "check_list_category_check_list_items"
)

// Columns holds all SQL columns for checklistitem fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldType,
	FieldIndex,
	FieldIsMandatory,
	FieldChecked,
	FieldStringVal,
	FieldEnumValues,
	FieldEnumSelectionModeValue,
	FieldSelectedEnumValues,
	FieldYesNoVal,
	FieldHelpText,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the CheckListItem type.
var ForeignKeys = []string{
	"check_list_category_check_list_items",
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
	Hooks  [1]ent.Hook
	Policy ent.Policy
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enum.CheckListItemType) error {
	switch _type {
	case "simple", "string", "enum", "files", "yes_no", "cell_scan", "wifi_scan":
		return nil
	default:
		return fmt.Errorf("checklistitem: invalid enum value for type field: %q", _type)
	}
}

// EnumSelectionModeValueValidator is a validator for the "enum_selection_mode_value" field enum values. It is called by the builders before save.
func EnumSelectionModeValueValidator(esmv enum.CheckListItemEnumSelectionMode) error {
	switch esmv {
	case "multiple", "single":
		return nil
	default:
		return fmt.Errorf("checklistitem: invalid enum value for enum_selection_mode_value field: %q", esmv)
	}
}

// YesNoVal defines the type for the yes_no_val enum field.
type YesNoVal string

// YesNoVal values.
const (
	YesNoValYes YesNoVal = "YES"
	YesNoValNo  YesNoVal = "NO"
)

func (ynv YesNoVal) String() string {
	return string(ynv)
}

// YesNoValValidator is a validator for the "yes_no_val" field enum values. It is called by the builders before save.
func YesNoValValidator(ynv YesNoVal) error {
	switch ynv {
	case YesNoValYes, YesNoValNo:
		return nil
	default:
		return fmt.Errorf("checklistitem: invalid enum value for yes_no_val field: %q", ynv)
	}
}

var (
	// enum.CheckListItemType must implement graphql.Marshaler.
	_ graphql.Marshaler = enum.CheckListItemType("")
	// enum.CheckListItemType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.CheckListItemType)(nil)
)

var (
	// enum.CheckListItemEnumSelectionMode must implement graphql.Marshaler.
	_ graphql.Marshaler = enum.CheckListItemEnumSelectionMode("")
	// enum.CheckListItemEnumSelectionMode must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enum.CheckListItemEnumSelectionMode)(nil)
)

// MarshalGQL implements graphql.Marshaler interface.
func (ynv YesNoVal) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(ynv.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (ynv *YesNoVal) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*ynv = YesNoVal(str)
	if err := YesNoValValidator(*ynv); err != nil {
		return fmt.Errorf("%s is not a valid YesNoVal", str)
	}
	return nil
}
