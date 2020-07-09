// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package adapter

import (
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/jinzhu/gorm"
)

type TCasbinRule struct {
	PType string `gorm:"size:100"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

type TCasbinPolicy struct {
	PType string `gorm:"size:100"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

func (c *TCasbinRule) TableName() string {
	return "casbin_rule" //as Gorm keeps table names are plural, and we love consistency
}

// TAdapter represents the Gorm adapter for policy storage.
// TODO : refactor db propeties using IOC
type TAdapter struct {
	driverName     string
	dataSourceName string
	dbSpecified    bool
	db             *gorm.DB
}

func NewAdapter(db *gorm.DB) *TAdapter {
	a := &TAdapter{}
	a.open(db)
	return a
}

func (a *TAdapter) open(db *gorm.DB) {
	a.db = db
	a.createTable()
}

func (a *TAdapter) close() {
	a.db.Close()
	a.db = nil
}
func (a *TAdapter) createTable() {
	if a.db.HasTable(&TCasbinRule{}) {
		return
	}

	err := a.db.CreateTable(&TCasbinRule{}).Error
	if err != nil {
		Log.Error(err)
	}
}

func (a *TAdapter) dropTable() {
	err := a.db.DropTable(&TCasbinRule{}).Error
	if err != nil {
		Log.Error(err)
	}
}

func loadPolicyLine(line TCasbinRule, model model.Model) {
	lineText := line.PType
	if line.V0 != "" {
		lineText += ", " + line.V0
	}
	if line.V1 != "" {
		lineText += ", " + line.V1
	}
	if line.V2 != "" {
		lineText += ", " + line.V2
	}
	if line.V3 != "" {
		lineText += ", " + line.V3
	}
	if line.V4 != "" {
		lineText += ", " + line.V4
	}
	if line.V5 != "" {
		lineText += ", " + line.V5
	}

	persist.LoadPolicyLine(lineText, model)
}

// LoadPolicy loads policy from database.
func (a *TAdapter) LoadPolicy(model model.Model) error {
	var lines []TCasbinRule
	err := a.db.Find(&lines).Error
	if err != nil {
		return err
	}

	for _, line := range lines {
		loadPolicyLine(line, model)
	}

	return nil
}

func savePolicyLine(ptype string, rule []string) TCasbinRule {
	line := TCasbinRule{}

	line.PType = ptype
	if len(rule) > 0 {
		line.V0 = rule[0]
	}
	if len(rule) > 1 {
		line.V1 = rule[1]
	}
	if len(rule) > 2 {
		line.V2 = rule[2]
	}
	if len(rule) > 3 {
		line.V3 = rule[3]
	}
	if len(rule) > 4 {
		line.V4 = rule[4]
	}
	if len(rule) > 5 {
		line.V5 = rule[5]
	}

	return line
}

// SavePolicy saves policy to database.
func (a *TAdapter) SavePolicy(model model.Model) error {
	a.dropTable()
	a.createTable()

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err := a.db.Create(&line).Error
			if err != nil {
				return err
			}
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			err := a.db.Create(&line).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *TAdapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := a.db.Create(&line).Error
	if err != nil {
		Log.Error(err)
	}
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (a *TAdapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	err := rawDelete(a.db, line) //can't use db.Delete as we're not using primary key http://jinzhu.me/gorm/crud.html#delete
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *TAdapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	line := TCasbinRule{}

	line.PType = ptype
	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		line.V0 = fieldValues[0-fieldIndex]
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		line.V1 = fieldValues[1-fieldIndex]
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		line.V2 = fieldValues[2-fieldIndex]
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		line.V3 = fieldValues[3-fieldIndex]
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		line.V4 = fieldValues[4-fieldIndex]
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		line.V5 = fieldValues[5-fieldIndex]
	}
	err := rawDeleteAll(a.db, line)
	return err
}

func rawDelete(db *gorm.DB, line TCasbinRule) error {
	err := db.Delete(TCasbinRule{}, "p_type = ? and v0 = ?"+
		" and v1 = ? and v2 = ? and v3 = ? and v4 = ? and v5 = ?",
		line.PType, line.V0, line.V1, line.V2, line.V3, line.V4, line.V5).Error
	return err
}

func rawDeleteAll(db *gorm.DB, line TCasbinRule) error {
	queryArgs := []interface{}{line.PType}

	queryStr := "p_type = ?"
	if line.V0 != "" {
		queryStr += " and v0 = ?"
		queryArgs = append(queryArgs, line.V0)
	}
	if line.V1 != "" {
		queryStr += " and v1 = ?"
		queryArgs = append(queryArgs, line.V1)
	}
	if line.V2 != "" {
		queryStr += " and v2 = ?"
		queryArgs = append(queryArgs, line.V2)
	}
	if line.V3 != "" {
		queryStr += " and v3 = ?"
		queryArgs = append(queryArgs, line.V3)
	}
	if line.V4 != "" {
		queryStr += " and v4 = ?"
		queryArgs = append(queryArgs, line.V4)
	}
	if line.V5 != "" {
		queryStr += " and v5 = ?"
		queryArgs = append(queryArgs, line.V5)
	}
	args := append([]interface{}{queryStr}, queryArgs...)
	err := db.Delete(TCasbinRule{}, args...).Error
	return err
}
