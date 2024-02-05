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
}

func dbCases(ctx context.Context, bs brand.Service, ps product.Service) error {
	var brandID uuid.UUID

	{
		newId := uuid.New()
		entityIn := model.NewBrand()
		entityIn.ID = newId
		entityIn.Payload.Add(`id`, newId)
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
		newId := uuid.New()
		entityIn := model2.NewProduct()
		entityIn.ID = newId
		entityIn.Payload.Add(`id`, newId)
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
