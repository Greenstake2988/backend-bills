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
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID'
```

### Update User

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
