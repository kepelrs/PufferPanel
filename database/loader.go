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

package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pufferpanel/pufferpanel/config"
	"github.com/pufferpanel/pufferpanel/models"
)

var dbConn *gorm.DB

func Load() error {
	err := openConnection()
	if err != nil {
		return err
	}

	dbObjects := []interface{} {
		&models.Node{},
		&models.Server{},
		&models.User{},
	}

	for _, v := range dbObjects {
		dbConn.AutoMigrate(v)
	}

	return nil
}

func openConnection() (error) {
	dialect := config.Get().Database.Dialect
	if dialect == "" {
		dialect = "mysql"
	}
	connString := config.Get().Database.Url

	//attempt to open database connection to validate
	var err error
	dbConn, err = gorm.Open(dialect, connString)
	return err
}

func GetConnection() (*gorm.DB, error) {
	var err error
	if dbConn == nil {
		err = openConnection()
	}

	return dbConn, err
}

func Close() {
	dbConn.Close()
}