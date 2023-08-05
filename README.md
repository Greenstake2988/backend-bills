## Bienvenido ala API de Billy Bills

### Sequence Diagram
![TODO: Image](images/wellcomepipelineservice.png)

## User Requests

#### Create User
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users' \
--header 'Content-Type: text/json' \
--data '{
    "email": "NoelChupaPijas@hotmail.com",
    "password": "muerdeAlmohadas123"
}'
```

#### GET User
``` 
curl --location 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID'
```

#### DELETE User
``` 
curl --location --request DELETE 'https://backend-bills-v7nohlccfa-uc.a.run.app/users/REPLACE_WITH_USER_ID'
```
