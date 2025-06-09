package config

import (
    
    "fmt"
    "log"
    "os"

   
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

//load env file
//connecting to database
//load the .env file variables
//var DB *sql.DB

func ConnectDB()(*gorm.DB){

    //loading env file 
    errEnv := godotenv.Load()
    if errEnv!=nil {
        log.Fatalf("Unable to load the env file %v",errEnv)
    }
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost :=os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbNAME := os.Getenv("DB_NAME")
    dsn :=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser,dbPassword,dbHost,dbPort,dbNAME)
    //Open the connection to the database
    //DB,err := sql.Open("mysql",dsn)
     DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err!=nil{
        log.Fatalf("Error in connecting to database: %v",err)
    }
    log.Println("Successfully Connected to Database")
    return DB;
}

