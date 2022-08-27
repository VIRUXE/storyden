package seed

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/Southclaws/storyden/pkg/resources/category"
)

var (
	SeedCategory_01_General = category.Category{
		ID:          category.CategoryID(uuid.MustParse("00000000-0000-0000-0000-000000000000")),
		Name:        "General",
		Description: "General stuff",
		Colour:      "#ffffff",
		Sort:        0,
	}

	SeedCategory_02_Photos = category.Category{
		ID:          category.CategoryID(uuid.MustParse("00000000-0000-0000-0000-000000000001")),
		Name:        "Media",
		Description: "Movies and tv shows",
		Colour:      "#ffffff",
		Sort:        1,
	}

	SeedCategory_03_Movies = category.Category{
		ID:          category.CategoryID(uuid.MustParse("00000000-0000-0000-0000-000000000002")),
		Name:        "Movies",
		Description: "Movies discussion",
		Colour:      "#ffffff",
		Sort:        2,
	}

	SeedCategory_04_Music = category.Category{
		ID:          category.CategoryID(uuid.MustParse("00000000-0000-0000-0000-000000000003")),
		Name:        "Music",
		Description: "Music discussion",
		Colour:      "#ffffff",
		Sort:        3,
	}

	SeedCategory_05_Admin = category.Category{
		ID:          category.CategoryID(uuid.MustParse("00000000-0000-0000-0000-000000000004")),
		Name:        "Admin",
		Description: "Admin area",
		Colour:      "#ffffff",
		Sort:        4,
		Admin:       true,
	}
)

func categories(r category.Repository) {
	ctx := context.Background()

	create := func(c *category.Category) category.CategoryID {
		c, err := r.CreateCategory(ctx, c.Name, c.Description, c.Colour, c.Sort, c.Admin)
		if err != nil {
			panic(err)
		}

		return c.ID
	}

	SeedCategory_01_General.ID = create(&SeedCategory_01_General)
	SeedCategory_02_Photos.ID = create(&SeedCategory_02_Photos)
	SeedCategory_03_Movies.ID = create(&SeedCategory_03_Movies)
	SeedCategory_04_Music.ID = create(&SeedCategory_04_Music)
	SeedCategory_05_Admin.ID = create(&SeedCategory_05_Admin)

	fmt.Println("created seed categories")
}