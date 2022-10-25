# go_rest_gin

### Use Compile Daemon 
Watches your go files in a directory and invokes go build if file changed.  
Github location:  
https://github.com/githubnemo/CompileDaemon
  
From your project folder where the go.mod is located:  
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon 
Dependency will be added to your go.mod. 

The executable will be put in your $HOME/go/bin 
  
How to use it  
CompileDaemon -command=“./<Your_executable_filename>”  
Every time you change your code, the CompileDaemon will compile and run your executable.  
  
  
  
### Go dot env  
Github location:  
https://github.com/joho/godotenv
  
  
From your project folder where the go.mod is located:  
go get github.com/joho/godotenv 
  
Dependency will be added to your go.mod. 
  
  
  
### Web Framework - Gin
GitHub location:  
https://github.com/gin-gonic/gin
  
  
Load it from your project folder where the go.mod is located:  
go get -u github.com/gin-gonic/gin
  
Dependency will be added to your go.mod. 


### Go ORM
https://gorm.io/
  
Look at the web site doc on how to load it for the latest instruction. 
  
go get -u gorm.io/gorm. 
  
go get -u gorm.io/driver/sqlite -> for sqlite. 
  
go get -u gorm.io/driver/postgres -> for Postgres. 
  
