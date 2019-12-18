### ABOUT

ESPGAME written in Golang as a Web app 

### SETUP

- Modify DB mysql user in init.sh and src/interfaces/config/config.go
- run script init.sh
- init.sh will import db install dependencies

To run the server use command
- go run src/app.go

- open http://localhost:8080/login.html


login as admin by credentials as to insert questions or view stats
- Username:admin
- Password:admin

Create a new user to play the game

- To insert new Questions login as admin, u'll be redirected to add question page
- To insert multiple questions at once with random answer options(for testing) use url localhost:8080/insert_question_random
- To clear all past answers delete all rows from table ESPGAME.user_questions_answers

### ABOUT

The application is written in GoLang, it uses The Clean Architecture
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

- DB files  =  files/dbinit.sql
- DB dump   =  files/espdump.sql
- APIs      =  src/interfaces/api/routes.go
- DB queries=  src/interfaces/db/mysql/queries.go
- Config file= src/entities/constant/constants.go


CheckList
https://docs.google.com/document/d/1wkcQ9KL9zewMTE2Q40cm31vv0a1Auv-xO8iCcmSPxFg/edit?usp=sharing
