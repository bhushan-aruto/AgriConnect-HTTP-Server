package postgres

import (
	"context"
	"log"

	"github.com/bhushn-aruto/krushi-sayak-http-server/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(pool *pgxpool.Pool) *postgresRepo {
	return &postgresRepo{
		pool: pool,
	}
}

func (repo *postgresRepo) Init() {
	conn, err := repo.pool.Acquire(context.Background())

	if err != nil {
		log.Fatalln("error occurred while aquring the connection ", err)
	}

	defer conn.Release()

	quries := []string{
		`CREATE TABLE IF NOT EXISTS farmers (
			farmer_id VARCHAR(255) PRIMARY KEY,
			full_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			phone_number VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS buyers (
			buyer_id VARCHAR(255) PRIMARY KEY,
			full_name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			phone_number VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS food_variants (
			variant_id VARCHAR(255) PRIMARY KEY,
			farmer_id VARCHAR(255) NOT NULL,
			variant_name VARCHAR(255) NOT NULL,
			banner_image_url VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(farmer_id) REFERENCES farmers(farmer_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS foods (
			food_id VARCHAR(255) PRIMARY KEY,
			food_variant_id VARCHAR(255) NOT NULL,
			food_name VARCHAR(255) NOT NULL,
			food_unit VARCHAR(255) NOT NULL,
			food_qty VARCHAR(255) NOT NULL,
			food_price VARCHAR(255) NOT NULL,
			food_image_url VARCHAR(255) NOT NULL,
			ratings VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(food_variant_id) REFERENCES food_variants(variant_id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS orders(
			order_id VARCHAR(255) PRIMARY KEY,
			food_id VARCHAR(255) NOT NULL, 
			buyer_name VARCHAR(255) NOT NULL,
			buyer_phone VARCHAR(255) NOT NULL,
			buyer_email VARCHAR(255) NOT NULL,
			buyer_address VARCHAR(255) NOT NULL,
			qty VARCHAR(255) NOT NULL,
			FOREIGN KEY(food_id) REFERENCES foods(food_id) ON DELETE CASCADE
		)`,
	}

	for _, query := range quries {
		if _, err := conn.Exec(context.Background(), query); err != nil {
			log.Fatalln("error occurred while initialzing database, Err: ", err.Error())
		}
	}

	log.Println("database initialized successfully")
}

func (repo *postgresRepo) CheckFamerEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM farmers WHERE email=$1)`
	var exists bool
	err := repo.pool.QueryRow(context.Background(), query, email).Scan(&exists)
	return exists, err
}

func (repo *postgresRepo) CheckFarmerPhoneNumberExists(phNum string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM farmers WHERE phone_number=$1)`
	var exists bool
	err := repo.pool.QueryRow(context.Background(), query, phNum).Scan(&exists)
	return exists, err
}

func (repo *postgresRepo) CreateFarmer(farmer *entity.Farmer) error {
	query := `INSERT INTO  farmers (farmer_id,full_name,email,phone_number,password) VALUES ($1,$2,$3,$4,$5)`

	_, err := repo.pool.Exec(context.Background(), query, farmer.FarmerId, farmer.FullName, farmer.Email, farmer.PhoneNumber, farmer.Password)

	return err
}

func (repo *postgresRepo) GetFarmerForLogin(email string) (*entity.Farmer, error) {
	query := `SELECT farmer_id,full_name,email,phone_number,password FROM farmers WHERE email=$1`
	farmer := new(entity.Farmer)
	err := repo.pool.QueryRow(context.Background(), query, email).Scan(&farmer.FarmerId, &farmer.FullName, &farmer.Email, &farmer.PhoneNumber, &farmer.Password)
	return farmer, err

}

func (repo *postgresRepo) GetFarmerPhoneNumberByFoodId(foodId string) (string, error) {
	query := `SELECT f.phone_number
			 FROM foods AS fo
			 JOIN food_variants AS fv ON fo.food_variant_id = fv.variant_id
			 JOIN farmers AS f ON fv.farmer_id = f.farmer_id
             WHERE fo.food_id = $1`
	var phoneNumber string
	err := repo.pool.QueryRow(context.Background(), query, foodId).Scan(&phoneNumber)
	return phoneNumber, err
}

func (repo *postgresRepo) CheckBuyerEmailExists(email string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM buyers WHERE email=$1)`
	var exists bool
	err := repo.pool.QueryRow(context.Background(), query, email).Scan(&exists)
	return exists, err
}

func (repo *postgresRepo) CheckBuyerPhoneNumberExists(phNum string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM buyers WHERE phone_number=$1)`
	var exists bool
	err := repo.pool.QueryRow(context.Background(), query, phNum).Scan(&exists)
	return exists, err
}

func (repo *postgresRepo) CreateBuyer(buyer *entity.Buyer) error {
	query := `INSERT INTO buyers (buyer_id,full_name,email,phone_number,password) VALUES ($1,$2,$3,$4,$5)`
	_, err := repo.pool.Exec(context.Background(), query, buyer.BuyerId, buyer.FullName, buyer.Email, buyer.PhoneNumber, buyer.Password)
	return err
}

func (repo *postgresRepo) GetBuyerForLogin(email string) (*entity.Buyer, error) {
	query := `SELECT buyer_id,full_name,email,phone_number,password FROM buyers WHERE email=$1`
	buyer := new(entity.Buyer)
	err := repo.pool.QueryRow(context.Background(), query, email).Scan(&buyer.BuyerId, &buyer.FullName, &buyer.Email, &buyer.PhoneNumber, &buyer.Password)
	return buyer, err

}

func (repo *postgresRepo) CheckFoodVariantExists(farmerId string, name string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM food_variants WHERE variant_name=$1 AND farmer_id=$2)`
	var exists bool
	err := repo.pool.QueryRow(context.Background(), query, name, farmerId).Scan(&exists)
	return exists, err
}

func (repo *postgresRepo) CreateFoodVariant(variant *entity.FoodVariant) error {
	query := `INSERT INTO food_variants (variant_id,farmer_id,variant_name,banner_image_url) VALUES ($1,$2,$3,$4)`
	_, err := repo.pool.Exec(
		context.Background(),
		query,
		variant.Id,
		variant.FarmerId,
		variant.Name,
		variant.BannerImageUrl,
	)
	return err
}

func (repo *postgresRepo) GetFoodVariantsByFormerId(farmerId string) ([]*entity.FoodVariant, error) {
	query := `SELECT variant_id,farmer_id,variant_name,banner_image_url FROM food_variants WHERE farmer_id=$1`

	var foodVariants []*entity.FoodVariant

	rows, err := repo.pool.Query(context.Background(), query, farmerId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		variant := new(entity.FoodVariant)

		err := rows.Scan(
			&variant.Id,
			&variant.FarmerId,
			&variant.Name,
			&variant.BannerImageUrl,
		)

		if err != nil {
			return nil, err
		}

		foodVariants = append(foodVariants, variant)

	}

	return foodVariants, nil
}

func (repo *postgresRepo) GetFoodVariantImageUrl(id string) (string, error) {
	query := `SELECT banner_image_url FROM food_variants WHERE variant_id=$1`

	var url string

	err := repo.pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&url)

	return url, err
}

func (repo *postgresRepo) DeleteFoodVariant(id string) error {
	query := `DELETE FROM food_variants WHERE variant_id=$1`

	_, err := repo.pool.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}

func (repo *postgresRepo) CheckFoodExists(variantId string, foodName string) (bool, error) {
	query := `SELECT EXISTS ( SELECT 1 FROM foods WHERE food_name=$1 AND food_variant_id=$2)`

	var exists bool

	err := repo.pool.QueryRow(
		context.Background(),
		query,
		foodName,
		variantId,
	).Scan(&exists)

	return exists, err
}

func (repo *postgresRepo) CreateFood(food *entity.Food) error {
	query := `INSERT INTO foods (food_id,food_variant_id,food_name,food_unit,food_qty,food_price,food_image_url,ratings) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := repo.pool.Exec(
		context.Background(),
		query,
		food.Id,
		food.VariantId,
		food.Name,
		food.Unit,
		food.Qty,
		food.Price,
		food.ImageUrl,
		food.Ratings,
	)

	return err
}

func (repo *postgresRepo) GetFoodsByVariantId(variantId string) ([]*entity.Food, error) {
	query := `SELECT food_id,food_variant_id,food_name,food_unit,food_qty,food_price,food_image_url,ratings FROM foods WHERE food_variant_id = $1`

	rows, err := repo.pool.Query(
		context.Background(),
		query,
		variantId,
	)

	if err != nil {
		return nil, err
	}

	var foods []*entity.Food

	for rows.Next() {
		food := new(entity.Food)

		if err := rows.Scan(
			&food.Id,
			&food.VariantId,
			&food.Name,
			&food.Unit,
			&food.Qty,
			&food.Price,
			&food.ImageUrl,
			&food.Ratings,
		); err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (repo *postgresRepo) GetFoodImageUrl(id string) (string, error) {
	query := `SELECT food_image_url FROM foods WHERE food_id=$1`

	var url string

	err := repo.pool.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(&url)

	return url, err
}

func (repo *postgresRepo) DeleteFood(foodId string) error {
	query := `DELETE FROM foods WHERE food_id=$1`

	_, err := repo.pool.Exec(
		context.Background(),
		query,
		foodId,
	)

	return err
}

func (repo *postgresRepo) GetAllFoods() ([]*entity.Food, error) {
	query := `SELECT food_id,food_variant_id,food_name,food_unit,food_qty,food_price,food_image_url FROM foods`

	rows, err := repo.pool.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	var foods []*entity.Food

	for rows.Next() {
		food := new(entity.Food)

		if err := rows.Scan(
			&food.Id,
			&food.VariantId,
			&food.Name,
			&food.Unit,
			&food.Qty,
			&food.Price,
			&food.ImageUrl,
		); err != nil {
			return nil, err
		}

		foods = append(foods, food)
	}

	return foods, nil
}

func (repo *postgresRepo) GetBuyerDetails(buyerId string) (*entity.Buyer, error) {
	query := `SELECT full_name,email,phone_number,password FROM buyers WHERE buyer_id=$1`

	buyer := new(entity.Buyer)

	err := repo.pool.QueryRow(
		context.Background(),
		query,
		buyerId,
	).Scan(&buyer.FullName, &buyer.Email, &buyer.PhoneNumber, &buyer.Password)

	return buyer, err
}

func (repo *postgresRepo) GetFoodQty(itemId string) (string, error) {
	query := `SELECT food_qty FROM foods WHERE food_id=$1`

	var qty string

	err := repo.pool.QueryRow(
		context.Background(),
		query,
		itemId,
	).Scan(
		&qty,
	)

	return qty, err
}

func (repo *postgresRepo) CreateBuyerOrder(order *entity.Order, qty string) error {

	c, err := repo.pool.Acquire(context.Background())

	if err != nil {
		return err
	}

	defer c.Release()

	tx, err := c.Begin(context.Background())

	if err != nil {
		return err
	}

	query1 := `UPDATE foods SET food_qty=$2 WHERE food_id=$1`

	query2 := `INSERT INTO orders  (
				order_id,
				food_id,
				buyer_name,
				buyer_phone,
				buyer_email,
				buyer_address,
				qty
			 ) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	if _, err := tx.Exec(
		context.Background(),
		query1,
		order.FoodId,
		qty,
	); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	if _, err := tx.Exec(
		context.Background(),
		query2,
		order.Id,
		order.FoodId,
		order.BuyerName,
		order.BuyerPhone,
		order.BuyerEmail,
		order.BuyerAddress,
		order.Qty,
	); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (repo *postgresRepo) GetOrdersByFarmerId(farmerId string) ([]*entity.OrderResponse, error) {
	query := `SELECT 
					o.order_id,
					o.buyer_name,
					o.buyer_phone,
					o.buyer_email,
					o.buyer_address,
					f.food_name,
					f.food_unit,
					f.food_price,
					o.qty
			  FROM orders o
			  JOIN foods f ON o.food_id=f.food_id
			  JOIN food_variants fv ON f.food_variant_id=fv.variant_id
			  WHERE fv.farmer_id=$1`

	rows, err := repo.pool.Query(
		context.Background(),
		query,
		farmerId,
	)

	if err != nil {
		return nil, err
	}

	var orders []*entity.OrderResponse

	for rows.Next() {
		order := new(entity.OrderResponse)

		err := rows.Scan(
			&order.OrderId,
			&order.BuyerName,
			&order.BuyerPhone,
			&order.BuyerEmail,
			&order.BuyerAddress,
			&order.ItemName,
			&order.ItemUnit,
			&order.ItemPrice,
			&order.TotalQty,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *postgresRepo) DeleteOrder(orderId string) error {
	query := `DELETE FROM orders WHERE order_id=$1`
	_, err := repo.pool.Exec(
		context.Background(),
		query,
		orderId,
	)
	return err
}
