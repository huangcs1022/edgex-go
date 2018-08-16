/*******************************************************************************
 * Copyright 2018
 * Dell Inc.
 * Cavium
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package memory

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"github.com/edgexfoundry/edgex-go/internal/export"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
)

type ExportMemoryDB struct {
	regs []export.Registration
}

func NewExportMemoryClient() *ExportMemoryDB {
	return &ExportMemoryDB{
		regs: make([]export.Registration, 0),
	}
}

func (mc *ExportMemoryDB) Registrations() ([]export.Registration, error) {
	return mc.regs, nil
}

func (mc *ExportMemoryDB) CloseSession() {
}

func (mc *ExportMemoryDB) AddRegistration(reg *export.Registration) (bson.ObjectId, error) {
	ticks := time.Now().Unix()
	reg.Created = ticks
	reg.Modified = ticks
	reg.ID = bson.NewObjectId()

	mc.regs = append(mc.regs, *reg)

	return reg.ID, nil
}

func (mc *ExportMemoryDB) UpdateRegistration(reg export.Registration) error {
	for i, r := range mc.regs {
		if r.ID == reg.ID {
			mc.regs[i] = reg
			return nil
		}
	}
	return db.ErrNotFound
}

func (mc *ExportMemoryDB) RegistrationById(id string) (export.Registration, error) {
	for _, reg := range mc.regs {
		if reg.ID.Hex() == id {
			return reg, nil
		}
	}

	return export.Registration{}, db.ErrNotFound
}

func (mc *ExportMemoryDB) RegistrationByName(name string) (export.Registration, error) {
	for _, reg := range mc.regs {
		if reg.Name == name {
			return reg, nil
		}
	}

	return export.Registration{}, db.ErrNotFound
}

func (mc *ExportMemoryDB) DeleteRegistrationById(id string) error {
	for i, reg := range mc.regs {
		if reg.ID.Hex() == id {
			mc.regs = append(mc.regs[:i], mc.regs[i+1:]...)
			return nil
		}
	}
	return db.ErrNotFound
}

func (mc *ExportMemoryDB) DeleteRegistrationByName(name string) error {
	for i, reg := range mc.regs {
		if reg.Name == name {
			mc.regs = append(mc.regs[:i], mc.regs[i+1:]...)
			return nil
		}
	}
	return db.ErrNotFound
}
