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

package models

import (
	"github.com/go-xorm/xorm"
	"github.com/google/uuid"
	"outofmemory/errors"
	"time"
)

type Tag struct {
	Id        uint32 `json:"-"`
	TagId     uint32 `json:"tag_id"`
	Name      string `json:"name"`
	State     uint8  `json:"-"` // 0: disable 1: enable
	CreatedAt time.Time   `xorm:"created" json:"-"`
	UpdatedAt time.Time   `xorm:"updated" json:"-"`
	DeletedAt time.Time   `xorm:"deleted" json:"-"`
}

var TagStateToUint = map[string]uint8{
	"disable": 0,
	"enable":  1,
	"deleted": 2,
}

func AddTag(data map[string]interface{}) (interface{}, error) {
	var (
		tagData = parseTagData(data)
		tag     Tag
	)
	_, err := existTagByName(tagData.Name)
	// make sure tag name is not exit
	if err == nil {
		return nil, errors.ErrTagIsExist
	}
	// generate a short unique tag id
	tagData.TagId = uuid.New().ID()
	if _, err := engine.Insert(&tagData); err != nil {
		return nil, errors.ErrCreateTagFailed
	}
	return tag, nil
}

func GetTagById(tagID uint32) (interface{}, error) {
	return existTagByID(tagID)
}

func GetTags(pageNum, perPage int) (interface{}, error) {
	var tags []*Tag
	err := engine.Select("tag_id, name").Where("state = ?", TagStateToUint["enable"]).Limit((pageNum-1)*perPage, perPage).Find(&tags)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
			return nil, nil
		default:
			return nil, errors.ErrGetArticleFailed

		}
	}
	result := map[string]interface{}{
		"total":    MustCountTags(),
		"per_page": perPage,
		"page":     pageNum,
		"tags":     tags,
	}
	return result, nil
}

func DeleteTag(tagIds interface{}) error {
	var err error
	switch inst := tagIds.(type) {
	case uint32:
		err = deleteTagByID(inst)
		break
	case []uint32:
		for _, tagID := range inst {
			err = deleteTagByID(tagID)
		}
		break
	}
	return err
}

func UpdateTagByID(data map[string]interface{}) (interface{}, error) {
	tag := parseTagData(data)
	_, err := existTagByName(tag.Name)
	// make sure tag name is not exist
	if err == nil || err != errors.ErrTagNotExist {
		return nil, errors.ErrTagIsExist
	}
	// make sure tag id is exist
	_, err = existTagByID(tag.TagId)
	if err != nil {
		return nil, err
	}
	_, err = engine.Table("tag").Select("name, state").Where("tag_id = ?", tag.TagId).Update(&tag)
	if err != nil {
		return nil, errors.ErrUpdateTagFailed
	}
	return tag, nil
}

func deleteTagByID(tagID uint32) error {
	var tag Tag
	err := engine.Select("tag_id, name, state").Where("tag_id = ?", tagID).Find(&tag)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
			return errors.ErrTagNotExist
		default:
			return errors.ErrDeleteTagFailed
		}
	}
	_, err = engine.Table(&tag).Select("state").Where("tag_id = ?", tagID).Update("state", TagStateToUint["deleted"])
	if err != nil {
		return errors.ErrDeleteTagFailed
	}
	return nil
}

func existTagByID(tagID uint32) (*Tag, error) {
	var tag Tag
	err := engine.Select("tag_id, name, state").Where("tag_id = ? and state < ?", tagID, TagStateToUint["deleted"]).Find(&tag)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
			return nil, errors.ErrTagNotExist
		default:
			return nil, errors.ErrGetTagFailed
		}
	}
	return &tag, nil
}

func existTagByName(name string) (*Tag, error) {
	var tag Tag
	err := engine.Select("tag_id, name, state").Where("name = ? and state < ?", name, TagStateToUint["deleted"]).Find(&tag)
	if err != nil {
		switch err {
		case xorm.ErrNotExist:
			return nil, errors.ErrTagNotExist
		default:
			return nil, errors.ErrGetTagFailed
		}
	}
	return &tag, nil
}

func MustCountTags() uint32 {
	var count uint32
	_, _ = engine.Table("tag").Where("state = ?", TagStateToUint["enable"]).Count(&count)
	return count
}

func parseTagData(data map[string]interface{}) Tag {
	tag := Tag{
		TagId: data["tag_id"].(uint32),
		State: data["state"].(uint8),
		Name:  data["name"].(string),
	}
	return tag
}
