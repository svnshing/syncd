// Copyright 2019 syncd Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import(
    "time"
)

type DeployApply struct {
    ID              int     `gorm:"primary_key"`
    SpaceId         int     `gorm:"type:int(11);not null;default:0"`
    Project         int     `gorm:"type:int(11);not null;default:0"`
    Name            string  `gorm:"type:varchar(100);not null;default:''"`
    Description     string  `gorm:"type:varchar(500);not null;default:''"`
    BranchName      string  `gorm:"type:varchar(100);not null;default:''"`
    TagName         string  `gorm:"type:varchar(100);not null;default:''"`
    Status          int     `gorm:"type:int(11);not null;default:0"`
}

func (m *Project) TableName() string {
    return "syd_project"
}

func (m *Project) Create() bool {
    m.Ctime = int(time.Now().Unix())
    return Create(m)
}

func (m *Project) Update() bool {
    return UpdateByPk(m)
}

func (m *Project) UpdateByFields(data map[string]interface{}, query QueryParam) bool {
    return Update(m, data, query)
}

func (m *Project) List(query QueryParam) ([]Project, bool) {
    var data []Project
    ok := GetMulti(&data, query)
    return data, ok
}

func (m *Project) Count(query QueryParam) (int, bool) {
    var count int
    ok := Count(m, &count, query)
    return count, ok
}

func (m *Project) Delete() bool {
    return DeleteByPk(m)
}

func (m *Project) Get(id int) bool {
    return GetByPk(m, id)
}