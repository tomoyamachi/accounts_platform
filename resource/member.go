package resource

import (
	"net/http"
	"strconv"

	"platform_accounts/model"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

func CreateMember(c echo.Context) error {
	var j model.Member
	c.Bind(&j)

	tx := c.Get("Tx").(*dbr.Tx)

	member := model.NewMember(j.Number, j.Name)

	if err := member.SaveMember(tx); err != nil {
		c.Error(err)
		return err
	} else {
		return c.JSON(http.StatusCreated, member)
	}
}

func GetMember(c echo.Context) error {

	number, _ := strconv.ParseInt(c.Param("id"), 0, 64)

	tx := c.Get("Tx").(*dbr.Tx)

	var member model.Member
	if err := member.LoadMember(tx, number); err != nil {
		c.Error(err)
		return err
	} else {
		return c.JSON(http.StatusOK, member)
	}
}

func GetMembers(c echo.Context) error {

	tx := c.Get("Tx").(*dbr.Tx)
	var members model.Members

	if err := members.LoadMembers(tx); err != nil {
		c.Error(err)
		return err
	} else {
		return c.JSON(http.StatusOK, members)
	}
}
