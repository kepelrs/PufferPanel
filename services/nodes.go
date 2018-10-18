/*
 Copyright 2018 Padduck, LLC
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 	http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package services

import (
	"github.com/jinzhu/gorm"
	"github.com/pufferpanel/pufferpanel/database"
	"github.com/pufferpanel/pufferpanel/models"
)

type NodeService struct {
	db *gorm.DB
}

func GetNodeService() (*NodeService, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	service := &NodeService{
		db: db,
	}

	return service, nil
}

func (ns *NodeService) GetAll() (models.Node, error) {
	nodes := models.Node{}

	res := ns.db.Find(&nodes)

	return nodes, res.Error
}

func (ns *NodeService) Get(id int) (models.Node, bool, error) {
	model := models.Node{}

	res := ns.db.First(&model, id)

	return model, model.Id != 0, res.Error
}

func (ns *NodeService) Update(node models.Node) error {
	res := ns.db.Update(&node)
	return res.Error
}

func (ns *NodeService) Delete(id int) error {
	model := models.Node{
		Id: id,
	}

	res := ns.db.Delete(&model)
	return res.Error
}