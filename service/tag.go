// Copyright 2019 sekfung
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package service

import "outofmemory/models"

type Tag struct {
	TagID   uint32
	Name    string
	State   uint8
	Page    int
	PerPage int
}

func (t *Tag) GetTags() (interface{}, error) {
	return models.GetTags(t.Page, t.PerPage)
}

func (t *Tag) GetTagByID() (interface{}, error) {
	return models.GetTagById(t.TagID)
}

func (t *Tag) DeleteTagById() error {
	return models.DeleteTag(t.TagID)
}

func (t *Tag) UpdateTag() (interface{}, error) {
	data := makeTagData(t)
	return models.UpdateTagByID(data)
}

func (t *Tag) AddTag() (interface{}, error) {
	data := makeTagData(t)
	return models.AddTag(data)
}

func makeTagData(t *Tag) map[string]interface{} {
	data := map[string]interface{}{
		"tag_id": t.TagID,
		"state":  t.State,
		"name":   t.Name,
	}
	return data
}
