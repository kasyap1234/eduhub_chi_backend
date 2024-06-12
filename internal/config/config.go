package config 
import (
	"github.com/joho/godotenv"
	"log"
	"os"

)
type Config struct {
	Port string 
	Database string 
}
func LoadConfig() Config {
	err :=godotenv.Load()
	if err !=nil {
		log.Fatal(err);
	}
	
}