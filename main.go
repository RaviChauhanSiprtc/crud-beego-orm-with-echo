package main

import (
	"net/http"
	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id    int    `orm:"column(id)"`
	Name  string `orm:"column(name)"`
	Email string `orm:"column(email)"`
}

func CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}

	o := orm.NewOrm()
	_, err := o.Insert(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	o := orm.NewOrm()
	user := User{Id: id}
	if err := c.Bind(&user); err != nil {
		return err
	}

	_, err := o.Update(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	o := orm.NewOrm()
	user := User{Id: id}
	_, err := o.Delete(&user)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func main() {

	e := echo.New()

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/demo?charset=utf8")
	orm.RegisterModel(new(User))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", CreateUser)
	e.GET("/users/:id", GetUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
