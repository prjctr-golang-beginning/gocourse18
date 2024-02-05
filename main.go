package main

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	mysql2 "gocourse18/internal/core/db/sql/mysql"
	"gocourse18/internal/domains/brand"
	"gocourse18/internal/domains/brand/adapter"
	"gocourse18/internal/domains/brand/model"
	"gocourse18/internal/domains/product"
	adapter2 "gocourse18/internal/domains/product/adapter"
	model2 "gocourse18/internal/domains/product/model"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cnf := mysql.Config{
		User:   `user`,
		Passwd: `password`,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}
	conn := mysql2.NewConnectionPool(&cnf)
	bs := adapter.NewBrandService(brand.NewBrandRepository(conn.ReadPool()))
	ps := adapter2.NewProductService(product.NewProductRepository(conn.ReadPool()))

	if err := dbCases(ctx, bs, ps); err != nil {
		log.Fatalln(err)
	}

	//var err error
	//db, err = sql.Open("mysql", cnf.FormatDSN())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//pingErr := db.Ping()
	//if pingErr != nil {
	//	log.Fatal(pingErr)
	//}
	//fmt.Println("Connected!")
	//
	//// Select
	//rows, err := db.Query(`SELECT id FROM users`)
	//if pingErr != nil {
	//	log.Fatal(pingErr)
	//}
	//
	//var us []User
	//
	//for rows.Next() {
	//	u := User{}
	//	err = rows.Scan(&u.ID)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	us = append(us, u)
	//}
	//rows.Close()
	//fmt.Println(us)
	//
	//// Insert
	//res, err := db.Exec(`INSERT INTO courses (price, name, description) VALUES (127, 'Some new', 'Details about some new')`)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res.RowsAffected())
	//
	//rdb := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "mycomplicatedpassword",
	//	DB:       0, // use default DB
	//})
	//
	//err := rdb.Set(ctx, "user:name", "Maks Morozov", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "user:name").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("user:name", val)
	//
	//val2, err := rdb.Get(ctx, "user:phone-number").Result()
	//if err == redis.Nil {
	//	fmt.Println("user:phone-number does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("user:phone-number", val2)
	//}

	// exercise:
	// Реалізувати функціонал, який нагадує месенджер:
	// - Треба зберігати повідомлення
	// - Треба редагувати повідомлення
	// - Треба видаляти повідомлення
	// - Треба повертату результат операції
	//
	// Використовуємо пакет flag
	// Формат повідомлення: <операція> <Імʼя користувача>: <повідомлення>
	// Наприклад:
	// create: Max: Hello to all!
	// > created id: 1
	// create: Alex: Hi! All is good.
	// > created id: 2
	// update 2: Hi! All os awesome.
	// > message 2 updated
}

func dbCases(ctx context.Context, bs brand.Service, ps product.Service) error {
	var brandID uuid.UUID

	{
		entityIn := model.NewBrand()
		entityIn.Payload.Add(`id`, uuid.New())
		entityIn.Payload.Add(`name`, `Adidas`)
		pk, err := bs.Create(ctx, entityIn)
		if err != nil {
			return err
		}

		entityOut, err := bs.GetOne(ctx, pk)
		if err != nil {
			return err
		}
		if entityIn.ID == entityOut.ID {
			log.Println(`All is OK with Brand`)
		}

		brandID = entityOut.ID
	}

	{
		entityIn := model2.NewProduct()
		entityIn.Payload.Add(`id`, uuid.New())
		entityIn.Payload.Add(`brand_id`, brandID)
		pk, err := ps.Create(ctx, entityIn)
		if err != nil {
			return err
		}

		entityOut, err := ps.GetOne(ctx, pk)
		if entityIn.ID == entityOut.ID {
			log.Println(`All is OK with Product`)
		}
	}

	return nil
}
