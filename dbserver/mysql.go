package dbserver

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/KasissLiu/go-tools/loadConfig"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	Host    string
	Port    int
	User    string
	Passwd  string
	Dbname  string
	Charset string
}

var MysqlConfigs map[string]MysqlConfig

func init() {
	localdb := loadConfig.New("mysql", "./config/db.ini")
	configs := localdb.GetAll()

	MysqlConfigs = make(map[string]MysqlConfig, 0)

	if m, ok := configs.(map[string]interface{}); ok {
		for connName, config := range m {
			if c, ok := config.(map[string]interface{}); ok {
				mysql := MysqlConfig{}
				for k, v := range c {
					if k == "type" {
						val, _ := v.(string)
						if val != "mysql" {
							break
						}
						continue
					}

					switch k {
					case "user":
						val, _ := v.(string)
						mysql.User = val
					case "port":
						val, _ := v.(int)
						mysql.Port = val
					case "dbname":
						val, _ := v.(string)
						mysql.Dbname = val
					case "host":
						val, _ := v.(string)
						mysql.Host = val
					case "passwd":
						val, _ := v.(string)
						mysql.Passwd = val
					case "charset":
						val, _ := v.(string)
						mysql.Charset = val
					default:
					}

				}

				MysqlConfigs[connName] = mysql

			}
		}
	}
}

func GetMysql(connName string) (*sql.DB, error) {

	if c, ok := MysqlConfigs[connName]; ok {

		conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", c.User, c.Passwd, c.Host, c.Port, c.Dbname, c.Charset))
		if err != nil {
			return nil, errors.New("Connection Cannot Create")
		}

		err = conn.Ping()
		if err != nil {
			return nil, errors.New("Connection Lost")
		}
		return conn, nil
	}

	return nil, errors.New("Cannot load Config")
}
