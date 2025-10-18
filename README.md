# chirp-auth-service

Authentication and profile managment for chirp
Built with Go, Gin and Postgres 


## API Endpoints

BASE_URL: `localhost:8001` (if running locally; unless different port used)

### User Registration

![POST](https://img.shields.io/badge/POST-%23FF5733?style=flat&logo=postman) `/auth/register`
Authentication: public

### User Login

![POST](https://img.shields.io/badge/POST-%23FF5733?style=flat&logo=postman) `/auth/login`
![POST](https://img.shields.io/badge/POST-%23FF5733?style=for-the-badge&logo=postman&logoColor=white) `auth/login`
Authentication: public

### User Deletion

DELETE `/auth/delete`
Authentication: JWT to be provided in `x-jwt-token` header

### Get User 

GET `/me`
Authentication: JWT to be provided in `x-jwt-token` header

### Update Bio

PUT `/update/bio`
Authentication: JWT to be provided in `x-jwt-token` header

### Update Password

PUT `/update/password`
Authentication: JWT to be provided in `x-jwt-token` header


## Project Structure

```
.
├── cmd/auth/           # main function
├── internal/
│   ├── controllers/    # request handlers
│   ├── db/             # database connection and table creation
│   ├── middleware/     # JWT authentication middleware
│   ├── models/         # database model structs 
│   └── utils/          # JWT and password hash creation and validation
├── .env         
├── go.mod              # go module declaration and dependencies
├── go.sum              
├── Makefile            # build and run commands
└── README.md