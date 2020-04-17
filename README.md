# Golang storage server
## API Endpoints
### Upload 
Store file into service.
* URL `http://hostname:8080/files/upload`
* Content type `multipart/form-data`
* Method `POST`
* Data Params `files[]`
* Success Response `application/json`
    * Code: 201
    * Content:
    ````
    {
        "files": [
            {
                "fileName": "photo-1541727687969-ce40493cd847.jpeg",
                "uuid": "e9509777-3811-442c-9c7a-51c9f04f63eb"
            }
        ],
        "status": "ok"
    }
  ````
---
### Delete
Delete stored file
* URL `http://hostname:8080/files/:uuid`
* Method `DELETE`
* URL Params
  
  Required: 
 
        uuid=[string]

* Success Response `application/json`
    * Code: 200
    * Content:
    ````
    {
      "status": "ok",
      "data": "e9509777-3811-442c-9c7a-51c9f04f63eb"
    }
    ````
---
### Get stored file
* URL `http://hostname:8080/files/:uuid`
* Method `GET`
* URL Params
  
  Required: 
 
        uuid=[string]
        
* Success Response: `binary-data`
---
### Get resized Image
* URL `http://hostname:8080/files/:uuid/:profile`
* Method `GET`
* URL Params
  
  Required: 
 
        uuid=[string]
        profile=[string]

* Available options of `profile`
        
        - small [500x500 px]
        - medium [1024x768 px]
        - thumbnail [164x164 px]
        
* Success Response: `binary-data`
___
### Service health-check
* URL `http://hostname:8080/__healthcheck`
* Method `GET`
* Success Response `application/json`
    * Code: 200
    * Content:
    ````
    {
      "service": "sm-drive",
    }
    ````
---