package controller

import (
	"NotesAPI/lib/database"
	"NotesAPI/model"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPicturesController(c echo.Context) error {
	pictures := database.GetPictures()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all pictures",
		"data":    pictures,
	})
}

func GetPictureByIDController(c echo.Context) error {
	id := c.Param("id")
	picture := database.GetPictureByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get picture By ID",
		"data":    picture.ID,
	})
}

func DeletePictureByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeletePictureByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Picture Successfully Deleted",
	})
}

func UpdatePictureByIDController(c echo.Context) error {
	id := c.Param("id")

	var picture model.Picture
	if err := c.Bind(&picture); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update picture",
			"error":   err.Error(),
		})
	}
	database.UpdatePictureByID(id, picture)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Picture Successfully Updated",
	})
}

func CreatePictureController(c echo.Context) error {
	var newPicture model.Picture
	file, err := c.FormFile("picture")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//destination
	// dst, err := os.Create(file.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// this is path which  we want to store the file
	f, err := os.OpenFile("pictures/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer f.Close()

	if _, err = io.Copy(f, src); err != nil {
		return err
	}

	newPicture.Picture_name = file.Filename

	noteid, err := strconv.Atoi(c.FormValue("noteid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}
	newPicture.NoteID = uint(noteid)

	newPicture = database.CreatePicture(newPicture)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Picture Successfully Uploaded",
		"data":    newPicture,
	})
}
