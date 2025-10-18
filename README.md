# chirp-auth-service

Authentication and profile managment for chirp
Built with Go, Gin and PostgreSQL 


## API Endpoints

BASE_URL: `localhost:8001` (if running locally; unless different port used)

### Root

<b><span style="color:green;">GET</span><b> `/`
Authentication: Public

Returns `"Hello from  chirp-auth-service"`

Response 
```
{
    "message": "Hello from chirp-auth-service"
}
```

### User Registration

<b><span style="color:yellow;">POST</span><b> `/auth/register`
Authentication: Public

Takes username, email and password; adds user to database; returns user and JWT

Request 
```
{
    "username": "test",
    "email": "test@gmail.com",
    "password": "testpassword"
}
```

Response 
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTUsImV4cCI6MTc2MTk5NzMwMywiaWF0IjoxNzYwNzg3NzAzfQ.TXspkqpwJUcUzJghjO_DWaEnAFyaKptGY_J5wjTlDgg",
    "user": {
        "id": 15,
        "username": "test",
        "email": "test@gmail.com",
        "bio": "",
        "likes_count": 0,
        "followers_count": 0,
        "following_count": 0,
        "tweets_count": 0,
        "created_at": "2025-10-18T17:11:40.275206Z"
    }
}
```

### User Login

<b><span style="color:yellow">POST</span><b> `/auth/login`
Authentication: Public

Takes email and password; validates credentials; returns user and JWT

Request 
```
{
    "email": "test@gmail.com",
    "password": "testpassword"
}
```

Response 
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTUsImV4cCI6MTc2MTk5NzM1NSwiaWF0IjoxNzYwNzg3NzU1fQ.JR1pEqO5ZNftHCtcmiLUzPZnYAP0hwFaKe5rXGzZ7uw",
    "user": {
        "id": 15,
        "username": "test",
        "email": "test@gmail.com",
        "bio": "",
        "likes_count": 0,
        "followers_count": 0,
        "following_count": 0,
        "tweets_count": 0,
        "created_at": "2025-10-18T17:11:40.275206Z"
    }
}
```

### User Deletion

<b><span style="color:red;">DELETE</span></b> `/auth/delete`
Authentication: JWT to be provided in `x-jwt-token` header

Deletes user associated with provided JWT

Response 
```
{
    "message": "Successfuly deleted user"
}
```

### Get User 

<b><span style="color:green;">GET</span></b> `/me`
Authentication: JWT to be provided in `x-jwt-token` header

Returns user associated with provided JWT

Response 
```
{
    "user": {
        "id": 15,
        "username": "test@gmail.com",
        "email": "test",
        "bio": "",
        "likes_count": 0,
        "followers_count": 0,
        "following_count": 0,
        "tweets_count": 0,
        "created_at": "2025-10-18T17:11:40.275206Z"
    }
}
```

### Update Bio

<b><span style="color:blue;">PUT</span></b> `/update/bio`
Authentication: JWT to be provided in `x-jwt-token` header

Takes bio and updates it in database for user associated with provided JWT; returns updated bio

Request 
```
{
    "bio": "test bio"
}
```
Response 
```
{
    "bio": "test bio",
    "message": "Updated bio successfully"
}
```

### Update Password

<b><span style="color:blue;">PUT</span></b> `/update/password`
Authentication: JWT to be provided in `x-jwt-token` header

Takes old and new password; validates and updates in database 

Request 
```
{
    "old_password": "testpassword",
    "new_password": "testpasswordnew"
}
```

Response 
```
{
    "message": "Sucessfully updated your password"
}
```


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
```