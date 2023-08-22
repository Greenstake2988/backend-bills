## Bienvenido ala API de Billy Bills

### Diagrama 
![TODO: Image](images/wellcomepipelineservice.png)


## Authentication
#### Login
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/login' \
--header 'Content-Type: text/json' \
--data '{
    "email": "NoelChupaPijas@hotmail.com",
    "password": "muerdeAlmohadas123"
}'
```
#### Response
``` 
{
    "id": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg4MzYzMTksImlhdCI6MTYyODgxMDExOSwiaWQiOjF9.Yb7SR3n_t4TnylPnNvB80UT3Nvf9O1teV7BBHyxdj6o"
}

```

## User Requests

#### Crear User
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users' \
--header 'Content-Type: text/json' \
--data '{
    "email": "NoelChupaPijas@hotmail.com",
    "password": "muerdeAlmohadas123"
}'
```

#### Fetch All Users
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users'
```

#### Fetch User
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID' \
--header 'Content-Type: text/json' \
--data '{
    "ID": 2,
    "CreatedAt": "2023-08-05T15:33:56.610933Z",
    "UpdatedAt": "2023-08-05T15:55:59.774466Z",
    "DeletedAt": null,
    "email": "NoelChupaPijas@hotmail.com",
    "password": "$2a$10$wcDnsqwJ2yNxzRE/40qWOe6HvcOInwdjtlyBtqUFUXgpmQux.jwHS",
    "bills": []
}'
```

#### Update User

``` 
curl --location --request PUT 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID'  \
--header 'Content-Type: text/json' \
--data '{
    "password": "muerdeAlmohadas123"
}'
```

#### Delete User
``` 
curl --location --request DELETE 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID'
```



## Bill Requests

#### Crear Bill
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/bills' \
--header 'Content-Type: text/json' \
--data '{
    "user_id": 1,
    "concept": "oxxo",
    "price":  100.00
}'
```