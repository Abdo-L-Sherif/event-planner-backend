To run this backend you should first install mysql to you machine create a user that has a password and a database associated with that user

After creating the user and the database you'll have to create a file with the name of .env in the same level of main.go and have inside it be as follows 
DB_USER="your usernmame"
DB_PASSWORD="your password"
DB_NAME="your databaseName (preferrably event-planner)"
DB_HOST=127.0.0.1
DB_PORT=3306

After that, assuming you have go and it's necessary dependancies installed you'll simply open up a terminal page on the directory of the aplication and type the command "go run main.go"

If you have problems with dependancies here are all the dependancies that I installed while making this app so far (I am using Ubuntu 24.04)
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get gorm.io/driver/mysql

