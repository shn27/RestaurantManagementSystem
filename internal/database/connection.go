package database

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var Connection = &cobra.Command{
	Use:   "db",
	Short: "Database connection",
	Long:  "Database connection",
	Run: func(cmd *cobra.Command, args []string) {
		ConnectDB()
		deleteTables()
		migrateDatabase()
	},
}

var DB *gorm.DB
var RedisClient *redis.Client

func ConnectDB() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	for {
		var err error
		DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database: ", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Successfully connected to database")
			break
		}
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println("Error retrieve database: ", err)
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error closing database: ", err)
	}
	fmt.Println("Successfully closed database")
}

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "mypassword",
		DB:       0,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successfully connected to redis")
	}
}

func CloseRedis() {
	err := RedisClient.Close()
	if err != nil {
		fmt.Println("Error closing redis: ", err)
		return
	}
	fmt.Println("Successfully closed redis")
}

var EsClient *elasticsearch.Client

func ConnectElasticsearch() {
	fmt.Println("Connecting to Elasticsearch...")
	var err error
	for {
		EsClient, err = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{
				"http://elasticsearch:9200", // e.g., "http://localhost:9200"
			},
		})
		if err != nil {
			log.Println("Error creating the client: ", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Connected to Elasticsearch Client")
			break
		}
	}
	res, err := EsClient.Info()
	if err != nil {
		log.Println("Error getting Elasticsearch info: ", err)
	}
	if err != nil {
		for {
			res, err = EsClient.Info()
			if err != nil {
				log.Println("Error getting Elasticsearch info: ", err)
				log.Println("Retrying in 5 seconds...")
				time.Sleep(5 * time.Second)
			} else {
				break
			}
		}
	}
	defer res.Body.Close()
	log.Println("Successfully connected to Elasticsearch")
}

func CloseEsClient() {
	err := EsClient.ClosePointInTime
	if err != nil {
		fmt.Println("Error closing redis: ", err)
		return
	}
	fmt.Println("Successfully closed redis")
}
