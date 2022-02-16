package handler

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	//get images
	files := form.File["image"]
	filename := ""

	for _, file := range files {
		filename = file.Filename

		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:3000/api/uploads/" + filename,
	})
}
