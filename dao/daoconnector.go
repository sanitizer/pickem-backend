package dao

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	mysqlDriver "github.com/go-sql-driver/mysql"
)

const (
	CONFIG_FILE = "dao/config/config.yaml"
)

type CloudConnection struct {
	cfg *mysqlDriver.Config
}

/*
	Make sure that the connection is closed when needed
*/
func (this *CloudConnection) GetNewConnection() (*sql.DB, error) {
	this.checkSetAttr()
	log.Println("Scanned configuration. Connecting... ")
	return mysql.DialCfg(this.cfg)
}

/*
	Set class variables if they are not set
*/
func (this *CloudConnection) checkSetAttr() {
	if this.cfg == nil {
		this.readConfig()
	}
}

/*
	Based on class variable attrs set them from file
*/
func (this *CloudConnection) readConfig() {
	file, err := ioutil.ReadFile(CONFIG_FILE)

	if err != nil {
		log.Panic(fmt.Sprintf("Was not able to read the file %s", CONFIG_FILE))
	}

	// this is needed for handling scope and not expose the attrs on CloudConnection obj
	type Temp struct {
		ConnectionName string `yaml:"CONNECTION_NAME"`
		User           string `yaml:"USER"`
		Password       string `yaml:"PASSWORD"`
		Database       string `yaml:"DATABASE"`
	}

	temp := new(Temp)

	err = yaml.Unmarshal(file, temp)

	if err != nil {
		log.Panic(fmt.Sprintf("Was not able to unmarshal the file %s", CONFIG_FILE))
	}

	this.cfg = mysql.Cfg(temp.ConnectionName, temp.User, temp.Password)
	this.cfg.DBName = temp.Database
	this.cfg.MultiStatements = true
	this.cfg.Collation = "utf8mb4_unicode_ci"
}
