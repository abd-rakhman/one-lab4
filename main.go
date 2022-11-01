package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Email    string
	Password string
}

func main() {
	accounts := make(map[string]User)
	// simple user-register and login crud operations
	e := echo.New()

	//login user with email and password
	e.GET("/login/:email", func(c echo.Context) error {
		email := c.Param("email")
		password := c.QueryParam("password")

		if _, ok := accounts[email]; !ok {
			return c.String(http.StatusBadRequest, "User does not exist")
		}
		if accounts[email].Password == password {
			return c.String(http.StatusOK, "OK")
		} else {
			return c.String(http.StatusBadRequest, "Invalid password")
		}
	})

	// updating user password
	e.PATCH("/update/:email", func(c echo.Context) error {
		email := c.Param("email")
		oldPassword := c.QueryParam("password")
		newPassword := c.FormValue("password")

		if _, ok := accounts[email]; !ok {
			return c.String(http.StatusBadRequest, "User does not exist")
		}
		user := accounts[email]
		if user.Password == oldPassword && len(newPassword) != 0 {
			user.Password = newPassword

			accounts[user.Email] = user
			return c.String(http.StatusOK, "OK")
		} else {
			return c.String(http.StatusBadRequest, "Invalid old password")
		}
	})

	//creating user
	e.POST("/create-user", func(c echo.Context) error {
		newUser := new(User)

		bindingErr := c.Bind(newUser)

		if bindingErr != nil {
			return bindingErr
		}

		if len(newUser.Email) == 0 || len(newUser.Password) == 0 {
			return c.String(http.StatusBadRequest, "Invalid body")
		}

		if _, ok := accounts[newUser.Email]; ok {
			return c.String(http.StatusBadRequest, "Error: User already exists")
		}

		accounts[newUser.Email] = *newUser
		return c.String(http.StatusOK, "OK")
	})

	// deleting user
	e.DELETE("/delete/:email", func(c echo.Context) error {
		email := c.Param("email")
		password := c.QueryParam("password")

		if _, ok := accounts[email]; !ok {
			return c.String(http.StatusBadRequest, "User does not exist")
		}
		user := accounts[email]
		if user.Password == password {
			delete(accounts, user.Email)
			return c.String(http.StatusOK, "OK")
		} else {
			return c.String(http.StatusBadRequest, "Invalid old password")
		}
	})

	e.Logger.Fatal(e.Start(":3000"))
}
