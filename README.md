# Binder + Validator

A package that simplify binding and validating process for Gin framework using validator/v10.
```bash
go get github.com/fanesz/bindator
```

Functions:
- bindator.BindBody()
- bindator.BindBodies()
- bindator.BindParam()
- bindator.BindParams()
- bindator.BindUri()
- bindator.BindUris()

Example:
```go
package main

import (
	"net/http"

	"github.com/fanesz/bindator"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func main() {
	router := gin.Default()

	router.POST("/users", func(c *gin.Context) {
		var user User

		res := bindator.BindBody(c, &user)
		if !res.Ok {
			c.JSON(http.StatusBadRequest, res)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":    user.Email,
			"password": user.Password,
		})
	})

	router.Run("127.0.0.1:8080")
}
```
### Fetch Result
#### When no body json provided:
```json
{
    "message": "Invalid JSON data: unexpected end of JSON input"
}
```

#### When invalid body json:
```json
{
    "message": "Invalid body type",
    "errors": [
        {
            "field": "email",
            "message": "Field is required"
        },
        {
            "field": "password",
            "message": "Field is required"
        }
    ]
}
```

#### When invalid specific field (email, gte, lte, min, max, len):
```json
{
    "message": "Invalid body type",
    "errors": [
        {
            "field": "email",
            "message": "Email is not valid"
        }
    ]
}
```

## Embedded Binding and Validating
Used when you need to bind an embedded struct, for example:
```go
type Company struct {
	User     User     `json:"user"`
	Contract Contract `json:"contract"`
}

ok, res := bindator.BindBodies(c, &company)
```
note: you can't mix normal field with embedded field.