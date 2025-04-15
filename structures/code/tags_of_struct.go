package code

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id" xml:"id" db:"user_id"`
	Name  string `json:"name" xml:"name" db:"user_name"`
	Email string `json:"email" xml:"email" db:"user_email"`
}

/*
Tag Name	Purpose	Example
json	     Controls JSON encoding/decoding (Go standard lib)	json:"user_id,omitempty"
xml	         Controls XML encoding/decoding (Go standard lib)	xml:"UserId,attr"
db	         Used by ORMs or database mappers (e.g., sqlx)	db:"user_id"
gorm	     Used by GORM (popular Go ORM)	gorm:"primaryKey;autoIncrement"
validate	 Used for field validation (e.g., go-playground/validator)	validate:"required,email"
form	     Used for binding HTTP form data (Gin framework)	form:"email"
yaml	     Controls YAML encoding/decoding (gopkg.in/yaml.v2)	yaml:"user_id"
bson	     For MongoDB BSON encoding/decoding (go.mongodb.org/mongo-driver)	bson:"_id,omitempty"
protobuf	 For Protocol Buffers code generation	protobuf:"varint,1,opt,name=id"
*/

/*
The json tag specifies how the struct is encoded or decoded when working with JSON.
In this example, the omitempty tag for Name means that if the Name is empty,
it won't be included in the JSON output.
*/
type Product struct {
	ID    int     `json:"id" xml:"id" db:"product_id" gorm:"primaryKey"`
	Name  string  `json:"name,omitempty" xml:"name" validate:"required"`
	Price float64 `json:"price" xml:"price" validate:"gte=0"`
}

func Json_type() {
	// Example Product
	product := Product{
		ID:    1,
		Price: 99.99,
	}

	// JSON Marshal
	jsonData, _ := json.Marshal(product)
	fmt.Println(string(jsonData)) // Output will exclude "name" since it's empty
}

/*
The gorm tag is used for structuring database fields,
such as indicating the primary key or auto-incrementing fields.
*/
func Gorm_type() {
	// Open a database connection
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	// Auto-migrate the schema
	db.AutoMigrate(&Product{})

	// Create a new product in the database
	product := Product{Name: "Smartphone", Price: 299.99}
	db.Create(&product)

	// Fetch the product
	var fetchedProduct Product
	db.First(&fetchedProduct, 1)
	fmt.Println(fetchedProduct) // Output: {ID 1 Name Smartphone Price 299.99}
}

/*
The validate tag can be used to enforce rules, such as ensuring the Name is required

	or the Price is greater than or equal to 0. We'll need a validation library for this.
	Let's use the go-playground/validator library.
*/
func Validate_type() {
	//validate_rules.txt
}
