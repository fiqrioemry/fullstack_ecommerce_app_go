package seeders

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"server/internal/models"
	"server/internal/utils"
)

func SeedUsers(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 10)

	adminUser := models.User{
		ID:       uuid.New(),
		Email:    "admin@shop.com",
		Password: string(password),
		Role:     "admin",
		Profile: models.Profile{
			Fullname: "Happy Shop Admin",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Admin",
		},
	}

	firstNames := []string{"Olivia", "Liam", "Emma", "Noah", "Ava", "Elijah", "Sophia", "James", "Isabella", "William", "Mia", "Benjamin", "Charlotte", "Lucas", "Amelia"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson"}
	genders := []string{"male", "female"}

	var customerUsers []models.User
	for i := 1; i <= 100; i++ {
		first := firstNames[rand.Intn(len(firstNames))]
		last := lastNames[rand.Intn(len(lastNames))]
		fullname := fmt.Sprintf("%s %s", first, last)
		seed := strings.ToLower(strings.ReplaceAll(fullname, " ", "_"))
		email := fmt.Sprintf("customer%02d@shop.com", i)
		gender := genders[rand.Intn(len(genders))]

		customer := models.User{
			ID:       uuid.New(),
			Email:    email,
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: fullname,
				Avatar:   fmt.Sprintf("https://api.dicebear.com/6.x/initials/svg?seed=%s", seed),
				Gender:   gender,
			},
		}
		customerUsers = append(customerUsers, customer)
	}

	if err := db.Create(&adminUser).Error; err != nil {
		log.Println("Failed to seed admin:", err)
	}
	if err := db.Create(&customerUsers).Error; err != nil {
		log.Println("Failed to seed customers:", err)
	}

	allUsers := []models.User{adminUser}
	allUsers = append(allUsers, customerUsers...)
	for _, user := range allUsers {
		generateNotificationSettingsForUser(db, user)
	}

	log.Println("✅ User seeding completed with notification settings!")
}

func seedProvinces(db *gorm.DB) {
	file, err := os.Open("internal/seeders/province.csv")
	if err != nil {
		log.Fatal("Failed to open province.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var provinces []models.Province

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID: %v", err)
		}

		province := models.Province{
			ID:   uint(id),
			Name: record[1],
		}
		provinces = append(provinces, province)
	}

	if err := db.Create(&provinces).Error; err != nil {
		log.Fatal("Failed to seed provinces:", err)
	}
	log.Println("✅ Province seeding completed")
}

func seedCities(db *gorm.DB) {
	file, err := os.Open("internal/seeders/city.csv")
	if err != nil {
		log.Fatal("Failed to open city.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var cities []models.City

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID: %v", err)
		}

		provinceID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID for city: %v", err)
		}

		city := models.City{
			ID:         uint(id),
			ProvinceID: uint(provinceID),
			Name:       record[1],
		}
		cities = append(cities, city)
	}

	if err := db.Create(&cities).Error; err != nil {
		log.Fatal("Failed to seed cities:", err)
	}
	log.Println("✅ City seeding completed")
}

func seedDistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/district.csv")
	if err != nil {
		log.Fatal("Failed to open district.csv:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var districts []models.District

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID: %v", err)
		}

		cityID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID for district: %v", err)
		}

		district := models.District{
			ID:     uint(id),
			CityID: uint(cityID),
			Name:   record[1],
		}
		districts = append(districts, district)
	}

	if err := db.Create(&districts).Error; err != nil {
		log.Fatal("Failed to seed districts:", err)
	}
	log.Println("✅ District seeding completed")
}

func seedSubdistricts(db *gorm.DB) {
	file, err := os.Open("internal/seeders/subdistrict.csv")
	if err != nil {
		log.Fatal("Failed to open subdistrict.csv:", err)
	}
	defer file.Close()

	var count int64
	db.Model(&models.Subdistrict{}).Count(&count)
	if count > 0 {
		log.Println("✅ Subdistricts already seeded, skipping...")
		return
	}

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var subdistricts []models.Subdistrict

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid subdistrict ID: %v", err)
		}

		districtID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID for subdistrict: %v", err)
		}

		subdistrict := models.Subdistrict{
			ID:         uint(id),
			DistrictID: uint(districtID),
			Name:       record[1],
		}
		subdistricts = append(subdistricts, subdistrict)
	}

	if err := db.CreateInBatches(&subdistricts, 500).Error; err != nil {
		log.Fatal("Failed to seed subdistricts:", err)
	}
	log.Println("✅ Subdistrict seeding completed")
}

func seedPostalCodes(db *gorm.DB) {
	file, err := os.Open("internal/seeders/postal_code.csv")
	if err != nil {
		log.Fatal("Failed to open postal_code.csv:", err)
	}
	defer file.Close()

	var count int64
	db.Model(&models.PostalCode{}).Count(&count)
	if count > 0 {
		log.Println("✅ Postal codes already seeded, skipping...")
		return
	}

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var postalCodes []models.PostalCode

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		id, err := strconv.ParseUint(record[0], 10, 64)
		if err != nil {
			log.Fatalf("Invalid postal ID: %v", err)
		}
		provinceID, err := strconv.ParseUint(record[4], 10, 64)
		if err != nil {
			log.Fatalf("Invalid province ID for postal code: %v", err)
		}
		cityID, err := strconv.ParseUint(record[3], 10, 64)
		if err != nil {
			log.Fatalf("Invalid city ID for postal code: %v", err)
		}
		districtID, err := strconv.ParseUint(record[2], 10, 64)
		if err != nil {
			log.Fatalf("Invalid district ID for postal code: %v", err)
		}
		subdistrictID, err := strconv.ParseUint(record[1], 10, 64)
		if err != nil {
			log.Fatalf("Invalid subdistrict ID for postal code: %v", err)
		}

		postalCode := models.PostalCode{
			ID:            uint(id),
			SubdistrictID: uint(subdistrictID),
			DistrictID:    uint(districtID),
			CityID:        uint(cityID),
			ProvinceID:    uint(provinceID),
			PostalCode:    record[5],
		}
		postalCodes = append(postalCodes, postalCode)
	}

	if err := db.CreateInBatches(&postalCodes, 500).Error; err != nil {
		log.Fatal("Failed to seed postal codes:", err)
	}
	log.Println("✅ Postal code seeding completed")
}

func SeedBanner(db *gorm.DB) {
	banners := []models.Banner{
		// Top Banner
		{ID: uuid.New(), Position: "top", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383472/topbanner03_lgpcf5.webp"},
		{ID: uuid.New(), Position: "top", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383471/topbanner02_supj7d.webp"},
		{ID: uuid.New(), Position: "top", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383470/topbanner01_wvpc7l.webp"},

		// Bottom Banner
		{ID: uuid.New(), Position: "bottom", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner02_kh2krk.webp"},
		{ID: uuid.New(), Position: "bottom", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner01_k1lylg.webp"},

		// Side Banner 1
		{ID: uuid.New(), Position: "side-left", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner01_gyfi00.webp"},
		{ID: uuid.New(), Position: "side-left", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner04_bh6d5e.webp"},
		{ID: uuid.New(), Position: "side-left", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747850902/wgkkapox5yeekaqortbv_vi1qte.webp"},

		// Side Banner 2
		{ID: uuid.New(), Position: "side-right", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner02_rdtezb.webp"},
		{ID: uuid.New(), Position: "side-right", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner03_kraq61.webp"},
		{ID: uuid.New(), Position: "side-right", Image: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747850993/i3qgc3ejk6odlnvlvkmm_yhnaod.webp"},
	}

	for _, b := range banners {
		if err := db.FirstOrCreate(&b, "image = ?", b.Image).Error; err != nil {
			log.Printf("failed to seed banner: %v\n", err)
		}
	}
}

func SeedCategories(db *gorm.DB) {
	categories := map[string]string{
		"Fashion and Apparel":     "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747956327/fashion-apparel_sjowyb.webp",
		"Men's & Women's Watches": "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747956327/men-women-watches_hhhqvl.webp",
		"Gadget & Electronics":    "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747956327/gadget-and-electronics_jm74ws.webp",
		"Food & Beverage":         "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747956327/food-beverage_uit5qb.webp",
		"Bags & Wallets":          "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957134/bag-and-wallet_bdljiu.webp",
	}

	for catName, image := range categories {
		cat := models.Category{
			ID:    uuid.New(),
			Name:  catName,
			Slug:  utils.GenerateSlug(catName),
			Image: image,
		}

		err := db.Where("name = ?", cat.Name).FirstOrCreate(&cat).Error
		if err != nil {
			log.Println("❌ Failed to create category:", catName, err)
			continue
		}
	}

	log.Println("✅ SeedCategories completed.")
}

func SeedFashionFirst(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Fashion and Apparel",
			Name:          "Blue Denim Jacket high quality",
			Description:   "Blue Denim Jacket is a classic and timeless piece of clothing that never goes out of style. Made from high-quality denim fabric, this jacket is durable and comfortable to wear. It features a button-up front, two chest pockets, and a relaxed fit that makes it perfect for layering over any outfit. Whether you're dressing up for a night out or keeping it casual for a day out, this jacket is the perfect addition to your wardrobe.",
			IsFeatured:    true,
			Price:         315000,
			AverageRating: 4.6,
			Sold:          5,
			Stock:         15,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429277/erem_shirt_01_shijri.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429275/erem_shirt_02_dusksh.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745429265/erem_shirt_03_ykizqa.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "T-shirt Men's Casual Cotton Tee",
			Description:   "This T-shirt is made from high-quality cotton fabric that is soft and breathable. It features a classic crew neck design and short sleeves, making it perfect for casual wear. The relaxed fit allows for easy movement and comfort, while the variety of colors and sizes make it suitable for everyone.",
			IsFeatured:    false,
			Discount:      5,
			Price:         98500,
			AverageRating: 4.4,
			Sold:          15,
			Stock:         35,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_01_l4sqob.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_02_rzapkt.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/cloth_mens_03_nwcb4c.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509050/cloth_mens_04_xttcat.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Hoodie Addict - Zipper Hoodie Mens Black",
			Description:   "Hoodie Addict is a stylish and comfortable hoodie designed for adults. Made from high-quality fabric, it features a zipper closure and a relaxed fit that makes it perfect for casual wear. The solid black color adds a touch of sophistication, making it suitable for any occasion.",
			IsFeatured:    false,
			Discount:      0.00,
			Price:         138000,
			AverageRating: 4.3,
			Sold:          25,
			Stock:         25,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509457/jaket01_tld8i0.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509457/jaket02_ru71to.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509458/jaket03_ygtnw2.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Hoodie Boxy Oversize Men Decorder Gray",
			Description:   "hoodie boxy is a stylish and comfortable hoodie designed for a relaxed fit. Made from high-quality fabric, it features a boxy silhouette that adds a trendy touch to your outfit. The oversized design provides extra comfort and freedom of movement, making it perfect for casual wear.",
			IsFeatured:    true,
			Discount:      0.00,
			Price:         275000,
			AverageRating: 4.7,
			Sold:          7,
			Stock:         35,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509054/jaket_mens_02_tyjlul.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509053/jaket_mens_01_a21ye5.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Elegant Floral Summer Dress Blossom",
			Description:   "Elegant Floral Summer Dress is a beautiful and stylish dress designed. Made from lightweight and breathable fabric, it features a floral print that adds a touch of femininity. The dress has a flattering silhouette that accentuates the waist and flows gracefully to the knee. Perfect for summer outings, this dress is both comfortable and chic.",
			IsFeatured:    false,
			Discount:      7,
			Price:         215000,
			AverageRating: 4.5,
			Sold:          12,
			Stock:         35,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510300/dress01_w1clnu.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510301/dress02_xnlphu.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510304/dress03_d3y08s.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Chic Long Sleeve Bodycon Dress",
			Description:   "Designer Chic Long Sleeve Bodycon Dress is a stylish and elegant dress designed for special occasions. Made from high-quality fabric, it features a bodycon silhouette that hugs the curves and accentuates the figure. The long sleeves add a touch of sophistication, making it perfect for evening events or formal gatherings.",
			IsFeatured:    false,
			Discount:      12,
			Price:         99000,
			AverageRating: 4.8,
			Sold:          12,
			Stock:         21,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510300/wom_dress03_bqsuif.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510299/wom_dress02_susije.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510299/wom_dress01_zgzscq.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Malvose Celana Pria Formal Bahan Premium Black Slimfit",
			Description:   "Celana Pria Formal Bahan Premium Black Slimfit adalah celana formal dengan potongan slimfit yang terbuat dari bahan premium. Celana ini cocok untuk berbagai acara formal, semi formal, dan bahkan kasual, seperti ke kantor atau kondangan. ",
			IsFeatured:    false,
			Discount:      9,
			Price:         175000,
			AverageRating: 4.4,
			Sold:          12,
			Stock:         10,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510924/pants01_x4memd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510925/pants02_cloota.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510925/pants03_rx1ixk.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Cargo pants mens long pants",
			Description:   "Cargo pants are a versatile and stylish. adaptable for various occasions. Made from durable fabric, these pants feature multiple pockets for added functionality. The relaxed fit and adjustable waistband provide comfort and ease of movement, making them perfect for outdoor activities or casual outings.",
			IsFeatured:    false,
			Discount:      15,
			Price:         155000,
			AverageRating: 4.5,
			Sold:          13,
			Stock:         13,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510916/men_pants02_yjdzug.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510904/men_pants01_tgqmbn.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			Stock:         p.Stock,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}

	}
}

func SeedFoodFirst(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Food & Beverage",
			Name:          "HOTTO PURTO 1 POUCH 16 SACHET | Superfood Multigrain Purple Potato Oat",
			Description:   "Hotto purto is a healthy drink made from purple potato and oat. It is rich in fiber and protein, making it a great choice for a nutritious snack or meal replacement. This product is gluten-free and contains no artificial additives. It is perfect for those who are looking for a healthy and convenient option.",
			Price:         85000,
			AverageRating: 4.7,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424592/hoto_snack_01_lf8uml.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424593/hoto_snack_02_sek5gt.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424599/hoto_snack_03_six5wh.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Covita - Healthy Protein Bar 40 gr Gluten Free - Peanut Choco",
			Description:   "Covita Protein Bar is a healthy snack made from high-quality protein and natural ingredients. It is gluten-free and contains no artificial additives. This protein bar is perfect for those who are looking for a nutritious and convenient snack option.",
			Price:         67000,
			AverageRating: 4.5,
			Stock:         50,
			Sold:          110,
			IsFeatured:    false,
			Discount:      15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424765/bars_snack_01_ghf8uj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424766/bars_snack_02_nsbgth.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424767/bars_snack_03_vcsloc.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Grain Snack - Protein Bar 40 gr Gluten Free - Peanut Choco",
			Description:   "Grain Snack Protein Bar is a healthy snack made from high-quality protein and natural ingredients. It is gluten-free and contains no artificial additives. This protein bar is perfect for those who are looking for a nutritious and convenient snack option.",
			Price:         25000,
			AverageRating: 3.9,
			Stock:         50,
			Sold:          110,
			IsFeatured:    false,
			Discount:      15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425054/grain_snack_01_hurkzb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425057/grain_snack_03_sm9sze.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425055/grain_snack_02_cnqxkk.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Forest Honey Drink 250ml",
			Description:   "Real honey drink with no preservatives. Made from 100% natural honey and pure water. No added sugar or artificial sweeteners. Perfect for a healthy lifestyle.",
			Price:         125000,
			AverageRating: 4.8,
			Stock:         30,
			Sold:          25,
			IsFeatured:    false,
			Discount:      10,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425496/honey_drink_01_qjl69j.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425499/honey_drink_02_dyufai.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Porang noodle - Mie Diet Porang",
			Description:   "porang noodle is a healthy noodle made from porang flour. It is gluten-free and low in calories, making it a great choice for those who are looking for a healthy alternative to regular noodles. This noodle is perfect for those who are on a diet or looking to maintain a healthy lifestyle.",
			Price:         8900,
			AverageRating: 4.3,
			Stock:         100,
			Sold:          1000,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748149921/porang-mie2_yzs1qx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1748149921/porang-mie_szyqfw.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Nestle Pure Life Water 600mL - 1 Pack",
			Description:   "Nestle water is a bottled water product that is sourced from natural springs. It is purified and filtered to ensure the highest quality and safety standards. This product is perfect for those who are looking for a convenient and healthy hydration option.",
			Price:         115000,
			AverageRating: 4.5,
			Stock:         30,
			Sold:          100,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425497/nestle_drink_02_bd5mye.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425501/nestle_drink_01_vgnua8.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "ESSENLI Pure Matcha Powder Japan Bubuk Matcha Drink",
			Description:   "Real matcha powder from Japan. Made from high-quality green tea leaves. No added sugar or artificial flavors. Perfect for making matcha drinks, desserts, and baking.",
			Price:         55500,
			AverageRating: 4.6,
			Stock:         30,
			Sold:          60,
			IsFeatured:    false,
			Discount:      2,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425829/matcha_drink_01_nq1pzd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425832/matcha_drink_02_nviqwj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425827/matcha_drink_03_y1mbxw.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Bihunku All Taste Soto Nyus",
			Description:   "Bihunku is a healthy noodle made from high-quality ingredients. It is gluten-free and low in calories, making it a great choice for those who are looking for a healthy alternative to regular noodles. This noodle is perfect for those who are on a diet or looking to maintain a healthy lifestyle.",
			Price:         11500,
			AverageRating: 4.3,
			Stock:         1500,
			Sold:          1300,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426599/bihun_noodle_02_ibzcpd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426611/bihun_noodle_01_t0egqo.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "ORIMIE noodle for a healthy life",
			Description:   "noodle with a unique taste and texture. Made from high-quality ingredients, this noodle is perfect for those who are looking for a healthy and delicious meal option. It is gluten-free and low in calories, making it a great choice for those who are on a diet or looking to maintain a healthy lifestyle.",
			Price:         23500,
			AverageRating: 4.4,
			Stock:         1500,
			Sold:          1200,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426605/orime_noodle_01_bpuprf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426606/orime_noodle_02_yjx3u0.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426610/orime_noodle_03_k8ljlt.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Stock:         p.Stock,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}
	}
}

func SeedFoodSecond(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Food & Beverage",
			Name:          "Chocolate Creamy Milk Premium",
			Description:   "Chocolate with milk is a delicious and creamy treat that combines the rich flavor of chocolate with the smoothness of milk. Perfect for snacking or baking. Available in various sizes. Indulge in the sweet and creamy taste of chocolate with milk.",
			Price:         50000,
			AverageRating: 4.7,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847154/foodempat1_nzw0sj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847147/foodempat3_i6k9gz.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847136/foodempat2_qnmx20.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Spicy Basreng Crispy Snack",
			Description:   "Spicy Basreng is a crispy snack made from fried fish skin, seasoned with spicy spices. Perfect for snacking or as a topping for your favorite dishes. Available in various flavors. Enjoy the crunchy and spicy taste of Spicy Basreng. great for sharing with friends and family.",
			Price:         25000,
			AverageRating: 4.5,
			Stock:         50,
			Sold:          110,
			IsFeatured:    false,
			Discount:      15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847139/foodsatu1_nubouq.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847142/foodsatu2_qnczej.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847143/foodsatu3_pzbqql.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Gerry Chocolate Creamy Milk Premium",
			Description:   "Gerry Chocolate Creamy Milk Premium is a delicious and creamy treat that combines the rich flavor of chocolate with the smoothness of milk. Perfect for snacking or baking. Available in various sizes. Indulge in the sweet and creamy taste of chocolate with milk.",
			Price:         100000,
			AverageRating: 4.8,
			Stock:         30,
			Sold:          25,
			IsFeatured:    false,
			Discount:      50,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847144/foodtiga1_msrrxa.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847145/foodtiga2_skcmjc.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Sosro tea sachet - 25 sachet",
			Description:   "Sosro tea sachet is a convenient and delicious way to enjoy the refreshing taste of Sosro tea. Each sachet contains high-quality tea leaves, perfect for brewing a cup of tea anytime, anywhere. Enjoy the rich flavor and aroma of Sosro tea with every sip.",
			Price:         25000,
			AverageRating: 4.3,
			Stock:         100,
			Sold:          1000,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847149/fooddua1_egllvu.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847150/fooddua2_ytgjcp.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Stock:         p.Stock,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}
	}
}

func SeedFashionSecond(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Fashion and Apparel",
			Name:          "Casual Sneakers Mens Shoes grey",
			Description:   "casual sneakers mens shoes grey is a stylish and comfortable footwear option for hanging out or casual outings. Made with high-quality materials, these sneakers provide a perfect blend of style and comfort. The grey color adds a modern touch to your outfit, making them versatile for various occasions.",
			Price:         425000,
			AverageRating: 4.5,
			Stock:         100,
			Sold:          50,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/3sneaker_shoes_01_t4lbd5.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536264/3sneaker_shoes_02_atfnsn.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "DES SNEAKERS Mens Shoes Vans Classic",
			Description:   "Des Sneakers Mens Shoes Vans Classic  is a stylish and comfortable footwear option for casual outings. Made with high-quality materials, these sneakers provide a perfect blend of style and comfort. The classic design adds a timeless touch to your outfit, making them versatile for various occasions.",
			Price:         475000,
			AverageRating: 4.6,
			Stock:         100,
			Sold:          80,
			IsFeatured:    false,
			Discount:      15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536262/sneaker_shoes_01_nssqgb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536262/sneaker_shoes_02_mctuky.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536262/sneaker_shoes_03_aiuieg.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Converse Allstar Sepatu Sekolah Sepatu ALL STAR CLASSIC",
			Description:   "Sneakers ikonik Converse All Star dengan konstruksi tahan lama dan gaya klasik yang tetap relevan.",
			Price:         525000,
			AverageRating: 4.7,
			Stock:         100,
			Sold:          90,
			IsFeatured:    false,
			Discount:      15,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker2_shoes_01_rc7i1l.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker2_shoes_02_iluvmx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker3_shoes_02_viyrm9.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Mens Sandal Nike Offcourt Slide Black",
			Description:   "Nike offcourt slide black is a comfortable and stylish sandal designed for casual wear. Made with high-quality materials, these sandals provide a perfect blend of style and comfort. The black color adds a modern touch to your outfit, making them versatile for various occasions.",
			Price:         225000,
			AverageRating: 4.5,
			Stock:         100,
			Sold:          80,
			IsFeatured:    false,
			Discount:      8,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536494/03sandals01_ogodhf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536496/03sandals02_uqknkc.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Air jordan Sandal Mens",
			Description:   "Air Jordan sandal mens is a stylish and comfortable footwear option for casual outings. Made with high-quality materials, these sandals provide a perfect blend of style and comfort. The design adds a modern touch to your outfit, making them versatile for various occasions.",
			Price:         225000,
			AverageRating: 4.5,
			Stock:         100,
			Sold:          80,
			IsFeatured:    false,
			Discount:      8,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536494/02sandals01_otpx9n.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536494/02sandals02_xuz1zl.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536498/02sandals03_f4bphf.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Bata Preseley Feather-Light Synthetic Sandal",
			Description:   "Bata Preseley Feather-light Synthetic sandal is a comfortable and stylish sandal designed for casual wear. Made with high-quality materials, these sandals provide a perfect blend of style and comfort. The synthetic material adds a modern touch to your outfit, making them versatile for various occasions.",
			IsFeatured:    false,
			Discount:      5,
			Price:         135000,
			AverageRating: 4.4,
			Stock:         100,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536493/01sandals01_y4l6vb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536494/01sandals02_euuo47.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536495/01sandals03_dxzjww.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Dockmart Mens Casual Shoes",
			Description:   "Dockmart mens casual shoes is a stylish and comfortable footwear option for casual outings. Made with high-quality materials, these shoes provide a perfect blend of style and comfort. The design adds a modern touch to your outfit, making them versatile for various occasions.",
			Price:         455000,
			AverageRating: 4.3,
			Stock:         100,
			Sold:          70,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/02formal01_nojgda.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/02formal02_ihkwdw.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Kenfa - Mora Black Mens shoes Loafer Formal",
			Description:   "Kenfa - Mora Black mens shoes loafer formal is a stylish and comfortable footwear option for formal occasions. Made with high-quality materials, these shoes provide a perfect blend of style and comfort. The black color adds a modern touch to your outfit, making them versatile for various occasions.",
			IsFeatured:    false,
			Discount:      15,
			Price:         125000,
			AverageRating: 4.6,
			Stock:         100,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536999/01formal03_yzmzs3.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/01formal02_wqdqvd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/01formal01_sxsc4y.webp",
			},
		},
		{
			Category:      "Fashion and Apparel",
			Name:          "Paulmay Formal Work Shoes Venesia",
			Description:   "Paulmay Formal Work Shoes Venesia is a stylish and comfortable footwear option for formal occasions. Made with high-quality materials, these shoes provide a perfect blend of style and comfort. The design adds a modern touch to your outfit, making them versatile for various occasions.",
			IsFeatured:    false,
			Discount:      15,
			Price:         295000,
			AverageRating: 4.3,
			Stock:         100,
			Sold:          40,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536999/03formal02_ysq0pe.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536999/03formal03_kgeocu.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Stock:         p.Stock,
			Weight:        1000.0,
			Width:         40.0,
			Height:        40.0,
			Length:        40.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}
	}
}

func SeedWatchesFirst(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Men's & Women's Watches",
			Name:          "Inshic Women's Watch",
			Description:   "Inshic women's watch is a stylish and elegant time. its design is perfect for any occasion, whether it's a casual outing or a formal event. The watch features a durable strap and a high-quality quartz movement, ensuring accurate timekeeping. With its sleek and modern look, this watch is a must-have accessory for any",
			Price:         325000,
			AverageRating: 4.5,
			Stock:         60,
			Sold:          40,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847119/dua1_ui6aeb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847120/dua2_hqqlag.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847028/dua3_nxjknh.webp",
			},
		},
		{
			Category:      "Men's & Women's Watches",
			Name:          "Men's adventure watch",
			Description:   "Men's adventure watch is a rugged and durable timepiece designed for outdoor enthusiasts. It features a sturdy strap and a high-quality quartz movement, ensuring accurate timekeeping. The watch is water-resistant and has various functions, making it perfect for hiking, camping, and other outdoor activities.",
			Price:         325000,
			AverageRating: 4.5,
			Stock:         60,
			Sold:          40,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847116/tiga2_poorgy.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847117/tiga3_pybsgx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847032/tiga1_cxfmrj.webp",
			},
		},
		{
			Category:      "Men's & Women's Watches",
			Name:          "Women's Beautiful Watch Christian Dior",
			Description:   "Christian Dior watch is a luxurious and elegant timepiece. Its design is perfect for any occasion, whether it's a casual outing or a formal event. The watch features a durable strap and a high-quality quartz movement, ensuring accurate timekeeping. With its sleek and modern look, this watch is a must-have accessory for any fashion-forward individual.",
			Price:         299999,
			AverageRating: 4.4,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847028/empat2_dklkev.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847028/empat1_kpmwni.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847117/empat3_bqahcz.webp",
			},
		},
		{
			Category:      "Men's & Women's Watches",
			Name:          "Black Watch Yolanda Men's",
			Description:   "Yolanda Men's watch is a stylish and elegant timepiece. Its design is perfect for any occasion, whether it's a casual outing or a formal event. The watch features a durable strap and a high-quality quartz movement, ensuring accurate timekeeping. With its sleek and modern look, this watch is a must-have accessory for any fashion-forward individual.",
			IsFeatured:    false,
			Discount:      0.00,
			Price:         325000,
			AverageRating: 4.8,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847028/satu1_yjeolh.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847029/satu2_o1oekd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747847031/satu3_wytplz.webp",
			},
		},
		{
			Category:      "Men's & Women's Watches",
			Name:          "Outdoor Sunlifex men's watch black",
			Description:   "Outdoor Sunlifex men's watch black is a stylish and elegant time. its design is perfect for any occasion, whether it's a casual outing or a formal event. The watch features a durable strap and a high-quality quartz movement, ensuring accurate timekeeping. With its sleek and modern look, this watch is a must-have accessory for any",
			Price:         325000,
			AverageRating: 4.7,
			Stock:         20,
			Sold:          10,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957528/outdoor-sunlifex-1_dzqydm.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957528/outdoor-sunlifex-2_zldmlr.webp",
			},
		},
		{
			Category:      "Men's & Women's Watches",
			Name:          "Elegant Sunlifex women's watch",
			Description:   "Elegant Sunlifex women's watch is a rugged and durable timepiece designed for outdoor enthusiasts. It features a sturdy strap and a high-quality quartz movement, ensuring accurate timekeeping. The watch is water-resistant and has various functions, making it perfect for hiking, camping, and other outdoor activities.",
			Price:         375000,
			AverageRating: 4.5,
			Stock:         60,
			Sold:          40,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957528/sunlifex-watch-1_dcpnqz.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957529/sunlifex-watch-2_gskrd7.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1747957536/sunlifex-watch-3_cuyu6v.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Stock:         p.Stock,
			Weight:        1000.0,
			Width:         10.0,
			Height:        10.0,
			Length:        10.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}
	}
}

func SeedGadgetElectronic(db *gorm.DB) {
	products := []struct {
		Category      string
		Name          string
		Description   string
		Price         float64
		AverageRating float64
		Stock         int
		Sold          int
		IsFeatured    bool
		Discount      float64
		Images        []string
	}{
		{
			Category:      "Gadget & Electronics",
			Name:          "Motorola G45 Snapdragon 6s Gen 3",
			Description:   "Motorola G45 is a smartphone with Snapdragon 6s Gen 3, 6.5-inch FHD+ display, 50MP camera, and 5000mAh battery. It features 120Hz refresh rate, 8GB RAM, and 128GB storage. The phone runs on Android 14 and has a sleek design.",
			Price:         1450000,
			AverageRating: 4.5,
			Stock:         60,
			Sold:          40,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_01_hpmjaf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421820/motorola_phone_02_wqlrdz.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_03_pbvpd1.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "HUAWEI WATCH FIT Special Edition Smartwatch",
			Description:   "HUAWEI WATCH FIT Special Edition Smartwatch | 1.64 HD AMOLED | 24/7 Active Health Management | Built-in GPS | Fast Charging. Notifications | Music Control | 96 Workout Modes | 5 ATM Water Resistant | 10 Days Battery Life | 30+ Watch Faces. It's a smartwatch with a 1.64-inch AMOLED display, 24/7 health monitoring, built-in GPS, and fast charging. It has 96 workout modes, 5 ATM water resistance, and a battery life of up to 10 days.",
			IsFeatured:    false,
			Discount:      5,
			Price:         625000,
			AverageRating: 4.2,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_02_ihjja7.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421186/huawei_smartwatch_04_r8ftp5.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_03_wswy7h.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_01_iwdoic.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421187/huawei_smartwatch_05_qbvhc7.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Samsung Galaxy A16 - Guarantee 2 Years",
			Description:   "Galaxy A16 is a smartphone with a 6.5-inch HD+ display, powered by Exynos 1330 processor, 4GB RAM, and 128GB storage. It features a 50MP main camera, 8MP front camera, and a 5000mAh battery. The phone runs on Android 14 and has a sleek design.",
			Price:         2799999,
			AverageRating: 4.4,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422220/samsung_phone_01_icqmh8.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422219/samsung_phone_02_juvdm6.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422219/samsung_phone_03_own5bj.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Samsung Galaxy Watch 4 Classic 42mm",
			Description:   "Samsung Watch 4 is a smartwatch with a 1.4-inch AMOLED display, powered by Exynos W920 processor, 16GB storage, and 1.5GB RAM. It features a heart rate monitor, ECG, and SpO2 sensor. The watch runs on Wear OS 3.5 and has a battery life of up to 40 hours.",
			IsFeatured:    false,
			Discount:      0.00,
			Price:         1225000,
			AverageRating: 4.8,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_02_szbzqg.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_04_uh1fjs.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420674/samsung_watch_01_qpf1dz.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Asus Zenfone 11 Ultra 12 5G",
			Description:   "Zenfone 11 Ultra uses a 6.8-inch AMOLED display with a 120Hz refresh rate, powered by Snapdragon 8 Gen 2 processor, 12GB RAM, and 256GB storage. It features a 200MP main camera, 12MP ultra-wide camera, and a 5000mAh battery. The phone runs on Android 14 and has a sleek design.",
			Price:         8899000,
			AverageRating: 4.8,
			Stock:         60,
			Sold:          90,
			IsFeatured:    false,
			Discount:      8,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423036/asus_phone_05_bgoxso.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423035/asus_phone_04_qe1lqw.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_01_wyvgsx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_03_ptjmet.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745422573/asus_phone_02_mbvwyi.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Xiaomi Mi band 4 Smartwatch",
			Description:   "Xiaomi Mi Band 4 is a fitness tracker with a 0.95-inch AMOLED display, heart rate monitor, sleep tracking, and 20-day battery life. It features 5 ATM water resistance, 6-axis accelerometer, and supports notifications from apps. The band is compatible with Android and iOS devices.",
			Price:         750000,
			AverageRating: 4.6,
			Stock:         50,
			Sold:          120,
			IsFeatured:    true,
			Discount:      5,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_n3vcip.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_2_mjutcx.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Infinix XPad 11 Tablet 5G Premium",
			Description:   "Infinix XPad 11 is a tablet with a 11-inch FHD+ display, powered by MediaTek Helio G90T processor, 4GB RAM, and 64GB storage. It features a 13MP main camera, 8MP front camera, and a 7000mAh battery. The tablet runs on Android 14 and has a sleek design.",
			IsFeatured:    true,
			Discount:      5,
			Price:         1950000,
			AverageRating: 4.2,
			Stock:         30,
			Sold:          50,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423645/infinix_tablet_01_mh0wgd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423643/infinix_tablet_02_fptycg.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Huawei MatePad 11 Snapdragon 865",
			Description:   "Huawei MatePad 11 is a tablet with a 11-inch 2K display, powered by Snapdragon 865 processor, 6GB RAM, and 128GB storage. It features a 13MP main camera, 8MP front camera, and a 7250mAh battery. The tablet runs on HarmonyOS and has a sleek design.",
			IsFeatured:    false,
			Discount:      0.0,
			Price:         3550000,
			AverageRating: 4.4,
			Stock:         30,
			Sold:          50,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423869/huawei_tablet_01_qz7bbi.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423858/huawei_tablet_02_twk4ey.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423859/huawei_tablet_03_qbokzz.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Xiaomi Pad SE NEW Guearantee 1 Year",
			Description:   "Xiaomi Redmi Pad SE is a tablet with a 10.1-inch FHD+ display, powered by Snapdragon 680 processor, 4GB RAM, and 128GB storage. It features a 8MP main camera, 5MP front camera, and a 8000mAh battery. The tablet runs on MIUI and has a sleek design.",
			IsFeatured:    false,
			Discount:      0.0,
			Price:         1975000,
			AverageRating: 4.6,
			Stock:         20,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424296/xiaomi_tablet_02_oxh1ad.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424295/xiaomi_tablet_01_wkjuec.webp",
			},
		},
	}

	for _, p := range products {
		var cat models.Category
		db.Where("name = ?", p.Category).First(&cat)

		product := models.Product{
			ID:            uuid.New(),
			CategoryID:    cat.ID,
			Name:          p.Name,
			Description:   p.Description,
			Price:         p.Price,
			Sold:          p.Sold,
			Stock:         p.Stock,
			Weight:        1000.0,
			Width:         15.0,
			Height:        15.0,
			Length:        15.0,
			Slug:          utils.GenerateSlug(p.Name),
			IsFeatured:    p.IsFeatured,
			IsActive:      true,
			Discount:      &p.Discount,
			AverageRating: p.AverageRating,
		}
		db.Create(&product)

		for _, img := range p.Images {
			db.Create(&models.ProductGallery{
				ID:        uuid.New(),
				ProductID: product.ID,
				Image:     img,
			})
		}
	}
}

func SeedReviews(db *gorm.DB) {
	var products []models.Product
	var customers []models.User

	// Ambil semua produk
	if err := db.Find(&products).Error; err != nil {
		log.Println("❌ Failed to fetch products:", err)
		return
	}

	// Ambil 5 customer pertama berdasarkan email
	customerEmails := []string{
		"customer01@shop.com",
		"customer02@shop.com",
		"customer03@shop.com",
		"customer04@shop.com",
		"customer05@shop.com",
	}
	if err := db.Where("email IN ?", customerEmails).Find(&customers).Error; err != nil {
		log.Println("❌ Failed to fetch customers for reviews:", err)
		return
	}
	if len(customers) < 5 {
		log.Println("❌ Not enough seeded customers to create reviews")
		return
	}

	sampleComments := []string{
		"The product is great and matches the description!",
		"Fast delivery and good quality.",
		"Affordable price and high-quality item.",
		"Highly recommended, will definitely buy again.",
		"Neatly and safely packed, awesome!",
		"Perfect for daily use.",
		"The product has a really cool features",
		"The Package delivery was so awesome, it's incredibly fast",
		"Superb packaging and the product exceeded my expectations.",
		"Great value for money. Will order again.",
	}

	// Loop produk dan tambahkan review dari 5 customer
	for _, product := range products {
		var reviews []models.Review
		for _, customer := range customers {
			review := models.Review{
				ID:        uuid.New(),
				ProductID: product.ID,
				UserID:    customer.ID,
				Rating:    rand.Intn(2) + 4, // rating antara 4 atau 5
				Comment:   sampleComments[rand.Intn(len(sampleComments))],
			}
			reviews = append(reviews, review)
		}
		if err := db.Create(&reviews).Error; err != nil {
			log.Printf("❌ Failed to save review for product %s: %v", product.Name, err)
		}
	}

	log.Println("✅ Review seeding from 5 customers completed")
}

func SeedNotificationTypes(db *gorm.DB) {
	types := []models.NotificationType{
		// Transaksi Pembelian
		{ID: uuid.New(), Code: "pending_payment", Title: "Waiting for Payment", Category: "transaction", DefaultEnabled: true},
		{ID: uuid.New(), Code: "order_processed", Title: "Order is Being Processed", Category: "transaction", DefaultEnabled: true},
		{ID: uuid.New(), Code: "order_shipped", Title: "Order Shipped", Category: "transaction", DefaultEnabled: true},
		{ID: uuid.New(), Code: "order_completed", Title: "Order Completed", Category: "transaction", DefaultEnabled: true},
		{ID: uuid.New(), Code: "promo_offer", Title: "Promo & Discount", Category: "promotion", DefaultEnabled: true},
		{ID: uuid.New(), Code: "system_message", Title: "System Announcement", Category: "announcement", DefaultEnabled: false},
	}

	for _, t := range types {
		db.FirstOrCreate(&t, "code = ?", t.Code)
	}

	log.Println("✅ Notification types seeded successfully")
}

func SeedVouchers(db *gorm.DB) {
	var count int64
	db.Model(&models.Voucher{}).Count(&count)
	if count > 0 {
		log.Println("Vouchers already seeded, skipping...")
		return
	}

	now := time.Now()
	expired := now.AddDate(0, 1, 0)

	max1 := 30000.0
	max2 := 50000.0

	voucher1 := models.Voucher{
		ID:           uuid.New(),
		Code:         "SALE50",
		Description:  "Dapatkan diskon 50% hingga 30.000",
		DiscountType: "percentage",
		Discount:     50,
		MaxDiscount:  &max1,
		Quota:        10,
		IsReusable:   false,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	voucher2 := models.Voucher{
		ID:           uuid.New(),
		Code:         "EUFORIA100",
		Description:  "Diskon langsung 100.000",
		DiscountType: "fixed",
		Discount:     100000,
		MaxDiscount:  &max2,
		Quota:        10,
		IsReusable:   true,
		ExpiredAt:    expired,
		CreatedAt:    now,
	}

	if err := db.Create([]models.Voucher{voucher1, voucher2}).Error; err != nil {
		log.Printf("Failed to seed vouchers: %v", err)
		return
	}

	log.Println("Vouchers seeding completed!")

}

func SeedCustomerTransactions(db *gorm.DB) {

	var customers []models.User
	if err := db.Preload("Profile").Where("role = ?", "customer").Find(&customers).Error; err != nil || len(customers) == 0 {
		log.Println("❌ Tidak ditemukan user dengan role customer")
		return
	}

	var products []models.Product
	if err := db.Preload("ProductGallery").Where("is_active = ?", true).Find(&products).Error; err != nil || len(products) == 0 {
		log.Println("❌ Tidak ditemukan produk aktif")
		return
	}

	for _, user := range customers {
		var addrCount int64
		db.Model(&models.Address{}).Where("user_id = ?", user.ID).Count(&addrCount)
		if addrCount == 0 {
			addr := models.Address{
				ID:         uuid.New(),
				UserID:     user.ID,
				Name:       "Alamat Utama",
				IsMain:     true,
				Address:    "Jl. Raya No.123",
				ProvinceID: 11, CityID: 156, DistrictID: 1928, SubdistrictID: 25033, PostalCodeID: 25043,
				Province: "DKI Jakarta", City: "Jakarta Barat", District: "Grogol", Subdistrict: "Tomang", PostalCode: "11440",
				Phone:     "081234567890",
				CreatedAt: time.Now(), UpdatedAt: time.Now(),
			}
			db.Create(&addr)
		}
	}

	dayToCount := map[int]int{
		120: 1, 119: 2, 118: 0, 117: 3, 116: 1, 115: 2, 114: 1, 113: 0, 112: 2, 111: 1,
		110: 2, 109: 3, 108: 1, 107: 0, 106: 2, 105: 4, 104: 3, 103: 1, 102: 2, 101: 2,
		100: 3, 99: 2, 98: 1, 97: 3, 96: 2, 95: 3, 94: 1, 93: 4, 92: 2, 91: 4,
		90: 3, 89: 2, 88: 3, 87: 3, 86: 4, 85: 2, 84: 4, 83: 5, 82: 4, 81: 3,
		80: 4, 79: 5, 78: 6, 77: 3, 76: 1, 75: 2, 74: 1, 73: 4, 72: 3, 71: 1,
		70: 5, 69: 3, 68: 6, 67: 4, 66: 7, 65: 3, 64: 5, 63: 3, 62: 2, 61: 3,
		60: 1, 59: 0, 58: 2, 57: 1, 56: 0, 55: 2, 54: 1, 53: 0, 52: 2, 51: 1,
		50: 2, 49: 1, 48: 2, 47: 2, 46: 1, 45: 2, 44: 3, 43: 1, 42: 2, 41: 2,
		40: 3, 39: 2, 38: 1, 37: 3, 36: 2, 35: 3, 34: 1, 33: 4, 32: 2, 31: 4,
		30: 3, 29: 2, 28: 3, 27: 3, 26: 4, 25: 2, 24: 4, 23: 5, 22: 4, 21: 3,
		20: 4, 19: 5, 18: 6, 17: 3, 16: 1, 15: 2, 14: 1, 13: 4, 12: 3, 11: 1,
		10: 5, 9: 3, 8: 6, 7: 4, 6: 7, 5: 3, 4: 5, 3: 3, 2: 2, 1: 3,
		0: 4,
	}

	index := 0
	for day, count := range dayToCount {
		for i := range count {
			customer := customers[(index+i)%len(customers)]
			address := models.Address{}
			db.Where("user_id = ?", customer.ID).First(&address)
			product := products[(index+i)%len(products)]

			orderID := uuid.New()
			qty := rand.Intn(2) + 1
			total := float64(qty) * product.Price
			shipping := 15000.0
			tax := 10000.0
			voucher := 10000.0
			amount := total + shipping + tax - voucher

			statuses := []string{"waiting_payment", "pending", "success", "process"}
			status := statuses[rand.Intn(len(statuses))]
			paymentStatus := "pending"
			if status == "success" || status == "process" {
				paymentStatus = "success"
			}

			order := models.Order{
				ID:              orderID,
				UserID:          customer.ID,
				InvoiceNumber:   fmt.Sprintf("INV/SEED/%d", time.Now().UnixNano()),
				Phone:           address.Phone,
				RecipientName:   customer.Profile.Fullname,
				ShippingAddress: fmt.Sprintf("%s, %s, %s %s", address.Address, address.Subdistrict, address.City, address.PostalCode),
				Courier:         "JNE",
				Status:          status,
				Total:           total,
				ShippingCost:    shipping,
				Tax:             tax,
				VoucherDiscount: voucher,
				AmountToPay:     amount,
				PaymentLink:     "https://www.ahmadfiqrioemry.com",
				CreatedAt:       time.Now().AddDate(0, 0, -day),
				UpdatedAt:       time.Now().AddDate(0, 0, -day),
				Items: []models.OrderItem{
					{
						ID:          uuid.New(),
						ProductID:   product.ID,
						IsReviewed:  false,
						ProductName: product.Name,
						ProductSlug: product.Slug,
						Image:       product.ProductGallery[0].Image,
						Price:       product.Price,
						Quantity:    qty,
						Subtotal:    total,
					},
				},
			}
			db.Create(&order)

			payment := models.Payment{
				ID:       uuid.New(),
				UserID:   customer.ID,
				Fullname: customer.Profile.Fullname,
				Email:    customer.Email,
				OrderID:  orderID,
				Method:   "BANK_TRANSFER",
				Status:   paymentStatus,
				PaidAt:   time.Now().AddDate(0, 0, -day),
				Total:    amount,
			}
			db.Create(&payment)

			if status == "process" || status == "success" {
				shipmentStatus := "shipped"
				if status == "success" {
					shipmentStatus = "delivered"
				}
				shipment := models.Shipment{
					ID:           uuid.New(),
					OrderID:      orderID,
					TrackingCode: fmt.Sprintf("JNE-SEED-%d", rand.Intn(99999)),
					Status:       shipmentStatus,
					Notes:        utils.ToPtr("Segera dikirim"),
					ShippedAt:    utils.ToPtr(time.Now().AddDate(0, 0, -day)),
				}
				db.Create(&shipment)
			}
		}
		index++
	}

	log.Println("✅ Seed multiple days of customer transactions completed.")
}

func SeedCustomerNotifications(db *gorm.DB) {
	var customers []models.User
	if err := db.Where("role = ?", "customer").Find(&customers).Error; err != nil {
		log.Println("❌ Gagal mengambil daftar customer:", err)
		return
	}

	if len(customers) == 0 {
		log.Println("❌ Tidak ada user dengan role 'customer' ditemukan.")
		return
	}

	for _, user := range customers {
		notifications := []models.Notification{
			{
				UserID:   user.ID,
				TypeCode: "order_processed",
				Title:    "Order is being processed",
				Message:  "Your order has been received and is currently being processed by the seller.",
				Channel:  "browser",
				IsRead:   false,
			},
			{
				UserID:   user.ID,
				TypeCode: "promo_offer",
				Title:    "Special Promotion Just for You!",
				Message:  "Don't miss out! Get 20% off on selected items today only.",
				Channel:  "browser",
				IsRead:   false,
			},
		}

		if err := db.Create(&notifications).Error; err != nil {
			log.Printf("❌ Gagal membuat notifikasi untuk user %s: %v", user.Email, err)
		} else {
			log.Printf("✅ Notifikasi disimpan untuk %s", user.Email)
		}
	}

	log.Println("✅ Notifikasi untuk semua customer berhasil disimpan.")
}

func generateNotificationSettingsForUser(db *gorm.DB, user models.User) {
	var notifTypes []models.NotificationType
	if err := db.Find(&notifTypes).Error; err != nil {
		log.Printf("Failed to get notification types for user %s: %v", user.Email, err)
		return
	}

	for _, nt := range notifTypes {
		for _, channel := range []string{"browser"} {
			setting := models.NotificationSetting{
				ID:                 uuid.New(),
				UserID:             user.ID,
				NotificationTypeID: nt.ID,
				Channel:            channel,
				Enabled:            nt.DefaultEnabled,
			}
			if err := db.Create(&setting).Error; err != nil {
				log.Printf("Failed to create notification setting for user %s: %v", user.Email, err)
			}
		}
	}
}
