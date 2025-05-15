package seeders

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
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
		Email:    "admin@example.com",
		Password: string(password),
		Role:     "admin",
		Profile: models.Profile{
			Fullname: "Admin User",
			Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Admin",
		},
	}

	customerUsers := []models.User{
		{
			ID:       uuid.New(),
			Email:    "customer01@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Customer User 01",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
				Gender:   "female",
			},
		},
		{
			ID:       uuid.New(),
			Email:    "customer02@example.com",
			Password: string(password),
			Role:     "customer",
			Profile: models.Profile{
				Fullname: "Customer User 02",
				Avatar:   "https://api.dicebear.com/6.x/initials/svg?seed=Customer",
				Gender:   "female",
			},
		},
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

	log.Println("User seeding completed with notification settings!")
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
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383472/topbanner03_lgpcf5.webp"},
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383471/topbanner02_supj7d.webp"},
		{ID: uuid.New(), Position: "top", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383470/topbanner01_wvpc7l.webp"},

		// Bottom Banner
		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner02_kh2krk.webp"},
		{ID: uuid.New(), Position: "bottom", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383469/bottombanner01_k1lylg.webp"},

		// Side Banner 1
		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner01_gyfi00.webp"},
		{ID: uuid.New(), Position: "side1", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner04_bh6d5e.webp"},

		// Side Banner 2
		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner02_rdtezb.webp"},
		{ID: uuid.New(), Position: "side2", ImageURL: "https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745383406/sidebanner03_kraq61.webp"},
	}

	for _, b := range banners {
		if err := db.FirstOrCreate(&b, "image_url = ?", b.ImageURL).Error; err != nil {
			log.Printf("failed to seed banner: %v\n", err)
		}
	}
}

func SeedCategories(db *gorm.DB) {
	placeholder := "https://placehold.co/400x400"
	categories := map[string][]string{
		"Fashion & Apparel":    {"Men's Clothing", "Women's Skirt", "Men's Pants", "Women's Dress"},
		"Men's Shoes":          {"Sneakers", "Sandals", "Formal Shoes"},
		"Gadget & Electronics": {"Phone & Tablet", "Electronic Devices", "Weareable Devices"},
		"Food & Beverage":      {"Health Drink", "Noodle & Pasta", "Snack food"},
	}

	for catName, _ := range categories {
		cat := models.Category{
			ID:       uuid.New(),
			Name:     catName,
			Slug:     utils.GenerateSlug(catName),
			ImageURL: placeholder,
		}

		err := db.Where("name = ?", cat.Name).FirstOrCreate(&cat).Error
		if err != nil {
			log.Println("failed to create category:", catName, err)
			continue
		}

	}
}

func SeedFashionAndApparel(db *gorm.DB) {
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
			Category:      "Fashion & Apparel",
			Name:          "Jacket Denim Warna Biru Bahan Ekslusif",
			Description:   "Jaket denim warna biru dongker adalah jaket yang terbuat dari bahan denim yang memiliki warna biru tua...",
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
			Category:      "Fashion & Apparel",
			Name:          "Kaos Distro Pria Lengan Pendek NY Kaos Oblong Cowok",
			Description:   "Kaos Distro Pria Lengan Pendek NY Kaos Oblong Cowok adalah jenis kaos yang diproduksi dengan jumlah terbatas...",
			IsFeatured:    false,
			Discount:      0.05,
			Price:         98500,
			AverageRating: 4.4,
			Sold:          15,
			Stock:         35,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_01_l4sqob.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745509051/cloth_mens_02_rzapkt.webp",
			},
		},
		{
			Category:      "Fashion & Apparel",
			Name:          "Hoodie Addict - Zipper Hoodie Dewasa Polos Hitam Pria",
			Description:   "Hoodie Addict Zipper adalah jaket hoodie dengan ritsleting (zipper) yang populer...",
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
			Category:      "Fashion & Apparel",
			Name:          "Hoodie Boxy Oversize Men Decorder Gray",
			Description:   "Hoodie boxy oversize adalah hoodie dengan siluet yang lebih lebar dan berbentuk kotak (boxy)...",
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
			Category:      "Fashion & Apparel",
			Name:          "Elegant Floral Summer Dress Blossom",
			Description:   "Dress ini dirancang untuk memberikan kesan anggun dan modern bagi setiap wanita. Menggunakan bahan berkualitas tinggi yang ringan dan nyaman dipakai sepanjang hari. Potongannya mengikuti lekuk tubuh dengan elegan namun tetap memberikan kenyamanan.",
			IsFeatured:    false,
			Discount:      0.07,
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
			Category:      "Fashion & Apparel",
			Name:          "Chic Long Sleeve Bodycon Dress",
			Description:   "Didesain dengan gaya timeless yang tak lekang oleh tren. Panjang rok yang midi membuatnya tetap sopan namun tetap stylish. Dress ini dirancang untuk memberikan kesan anggun dan modern bagi setiap wanita. Bagian pinggang dibuat elastis untuk fleksibilitas ukuran dan kenyamanan ekstra.",
			IsFeatured:    false,
			Discount:      0.12,
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
			Category:      "Fashion & Apparel",
			Name:          "Malvose Celana Pria Formal Bahan Premium Black Slimfit",
			Description:   "Celana Pria Formal Bahan Premium Black Slimfit adalah celana formal dengan potongan slimfit yang terbuat dari bahan premium. Celana ini cocok untuk berbagai acara formal, semi formal, dan bahkan kasual, seperti ke kantor atau kondangan. ",
			IsFeatured:    false,
			Discount:      0.09,
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
			Category:      "Fashion & Apparel",
			Name:          "celana cargo panjang pria celana outdoor pria longgar kasual korduroi kulot",
			Description:   "Celana cargo panjang pria ini adalah pilihan ideal untuk kegiatan outdoor, dikarenakan desainnya yang longgar dan kasual, serta dilengkapi dengan saku-saku besar di samping (cargo pockets). Bahan korduroi memberikan kesan unik dan nyaman, cocok untuk berbagai aktivitas, termasuk kulot.",
			IsFeatured:    false,
			Discount:      0.15,
			Price:         155000,
			AverageRating: 4.5,
			Sold:          13,
			Stock:         13,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510904/men_pants01_tgqmbn.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745510916/men_pants02_yjdzug.webp",
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
				ImageURL:  img,
			})
		}

	}
}

func SeedFoodBeverage(db *gorm.DB) {
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
			Description:   "Hotto Purto merupakan minuman tinggi serat dan rendah kalori dengan 15 multigrain seperti oat Swedia dan ubi ungu.",
			Price:         135000,
			AverageRating: 4.7,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424592/hoto_snack_01_lf8uml.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424593/hoto_snack_02_sek5gt.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Covita - Healthy Protein Bar 40 gr Gluten Free - Peanut Choco",
			Description:   "Cemilan sehat berprotein dari tanaman alami. Cocok untuk olahraga, mengandung 15 multigrain, vitamin dan serat tinggi.",
			Price:         67000,
			AverageRating: 4.5,
			Stock:         50,
			Sold:          110,
			IsFeatured:    false,
			Discount:      0.14,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424765/bars_snack_01_ghf8uj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424766/bars_snack_02_nsbgth.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424767/bars_snack_03_vcsloc.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Madu Asli Hutan Honey Life Gold 650ml",
			Description:   "Madu hutan asli, tanpa pengawet, alami dan segar. Cocok untuk meningkatkan daya tahan tubuh dan kesehatan harian.",
			Price:         168000,
			AverageRating: 4.8,
			Stock:         30,
			Sold:          25,
			IsFeatured:    false,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425496/honey_drink_01_qjl69j.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425499/honey_drink_02_dyufai.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Mie Porang dietmeal GORENG rendah kalori",
			Description:   "Mie diet porang yang rendah kalori, bebas gluten, cocok untuk program diet dan tinggi serat.",
			Price:         8900,
			AverageRating: 4.3,
			Stock:         100,
			Sold:          1000,
			IsFeatured:    false,
			Discount:      0.04,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426605/indomie_noodle_02_leaptj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426601/indomie_noodle_01_wztuyg.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Nestle Pure Life Air Minum Ukuran 600mL - 1 Pack",
			Description:   "Air minum Nestle Pure Life 600mL adalah air mineral yang diproduksi dengan Standar Internasional oleh Nestle Global Waters.",
			Price:         115000,
			AverageRating: 4.5,
			Stock:         30,
			Sold:          100,
			IsFeatured:    false,
			Discount:      0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425497/nestle_drink_02_bd5mye.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425501/nestle_drink_01_vgnua8.webp",
			},
		},
		{
			Category: "Food & Beverage",

			Name:          "ESSENLI Pure Matcha Powder Japan Bubuk Matcha Murni Drink",
			Description:   "Matcha Jepang asli, kaya antioksidan & vitamin, cocok untuk minuman dan makanan sehat.",
			Price:         75500,
			AverageRating: 4.6,
			Stock:         30,
			Sold:          60,
			IsFeatured:    false,
			Discount:      0.02,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425829/matcha_drink_01_nq1pzd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425832/matcha_drink_02_nviqwj.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745425827/matcha_drink_03_y1mbxw.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "Bihunku All Rasa Soto Nyus",
			Description:   "Bihunku All Rasa adalah bihun instan rendah lemak dan kolesterol, cocok untuk santapan harian.",
			Price:         11600,
			AverageRating: 4.3,
			Stock:         1500,
			Sold:          1300,
			IsFeatured:    false,
			Discount:      0.05,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426599/bihun_noodle_02_ibzcpd.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745426611/bihun_noodle_01_t0egqo.webp",
			},
		},
		{
			Category:      "Food & Beverage",
			Name:          "ORIMIE Goreng dari Orimen Kids",
			Description:   "Mie sehat untuk anak-anak tanpa MSG & bahan kimia berbahaya. Bumbu alami & aman dikonsumsi.",
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
				ImageURL:  img,
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
			Description:   "Moto G45 5G pakai prosesor Snapdragon 6s Gen 3. Didukung RAM 8GB + 8GB virtual, storage 256GB, multitasking lancar.",
			Price:         1450000,
			AverageRating: 4.5,
			Stock:         60,
			Sold:          40,
			IsFeatured:    true,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421821/motorola_phone_01_hpmjaf.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421820/motorola_phone_02_wqlrdz.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Samsung Galaxy A16 - Garansi Resmi Sein Tam",
			Description:   "Galaxy A16 hadir dengan layar Super AMOLED 6.7 inci, baterai 5000mAh, kamera 50MP, dan desain tipis 7.9mm.",
			Price:         2799999,
			AverageRating: 4.4,
			Stock:         50,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.04,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Asus Zenfone 11 Ultra 12 5G",
			Description:   "Zenfone 11 Ultra pakai Snapdragon 8 Gen 3, layar 6.78 inci AMOLED, kamera gimbal 50MP, RAM 12GB, storage 256GB.",
			Price:         8899000,
			AverageRating: 4.8,
			Stock:         60,
			Sold:          90,
			IsFeatured:    false,
			Discount:      0.07,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423035/asus_phone_04_qe1lqw.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423036/asus_phone_05_bgoxso.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Xiaomi Mi band 4 Smartwatch",
			Description:   "Mi Band 4 hadir dengan kapasitas baterai lebih besar dan konektivitas Bluetooth 4.2, tahan air hingga 50 meter.",
			Price:         750000,
			AverageRating: 4.6,
			Stock:         50,
			Sold:          120,
			IsFeatured:    true,
			Discount:      0.045,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424296/xiaomi_tablet_02_oxh1ad.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745424295/xiaomi_tablet_01_wkjuec.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Infinix XPad 11 Tablet 5G Premium",
			Description:   "Infinix XPad 11 adalah tablet Android dengan layar 11 inci dan refresh rate 90Hz, ditenagai oleh chipset MediaTek Helio G99. 7000mAh, RAM hingga 8GB, dan Android 14. Ia juga dilengkapi dengan fitur-fitur seperti Folax Voice Assistant, Multi-Device Collaboration, dan pengisian cepat.",
			IsFeatured:    true,
			Discount:      0.02,
			Price:         2250000,
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
			Description:   "Huawei MatePad 11 adalah tablet dengan layar 11 inci, ditenagai oleh chipset Snapdragon 865, RAM 6GB, dan memori internal 128GB yang dapat diperluas. Tablet ini juga dilengkapi dengan sistem operasi Harmony OS 3.1. Secara keseluruhan, Huawei MatePad 11 adalah tablet yang menawarkan performa baik, layar yang bagus, dan berbagai fitur tambahan, menjadikannya pilihan yang menarik untuk berbagai kebutuhan, mulai dari produktivitas hingga hiburan.",
			IsFeatured:    false,
			Discount:      0.0,
			Price:         3550000,
			AverageRating: 4.4,
			Stock:         30,
			Sold:          50,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423869/huawei_tablet_01_qz7bbi.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423859/huawei_tablet_03_qbokzz.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745423858/huawei_tablet_02_twk4ey.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Xiaomi Pad SE NEW Garansi",
			Description:   "Xiaomi Redmi Pad SE adalah tablet Android yang memiliki layar FHD+ 11 inci dengan refresh rate 90 Hz, ditenagai oleh prosesor Snapdragon 680, RAM 4GB, dan penyimpanan internal 128GB, serta baterai 8000mAh. Tablet ini dilengkapi dengan empat speaker dengan Dolby Atmos, dan kamera depan 5MP dan kamera belakang 8MP. Redmi Pad SE hadir dengan layar IPS LCD berukuran 10,1 inci, memberikan tampilan yang luas dan jelas. Resolusi layar sebesar 1200 x 2000 piksel, dengan tingkat kecerahan hingga 340 nits dan rasio kontras 1500:1, cocok untuk berbagai kebutuhan mulai dari streaming video, browsing, hingga bermain game.",
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
		{
			Category:      "Gadget & Electronics",
			Name:          "Xiaomi Mi band 4 Smartwatch",
			Description:   "Miliki smartband pintar xiaomi Mi Band 4 Generasi terbaru, hadir dengan beragam fitur canggih dengan peningkatan yang lebih baik dari generasi sebelumnya. Kapasitas baterai Xiaomi Mi Band 4 50 % lebih besar dari xiaomi mi band 2 yang mampu bertahan hingga lebih dari 20 hari penggunaan. XIaomi Mi Band 4 dilengkapi dengan bluetooth 4.2 untuk konektivitasnya dan untuk ketahanan airnya pun turut ditingkatkan yang kini mampu bertahan hingga kedalaman 50 meter.",
			IsFeatured:    true,
			Discount:      0.045,
			Price:         775000,
			AverageRating: 4.8,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_2_mjutcx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420230/smart_watch_mi_band_4_n3vcip.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "Samsung Galaxy Watch 4 Classic 42mm",
			Description:   "Samsung Watch 4 hadir dengan display Sapphire Crystal, GPS, sleep tracker dan body composition. Smartwatch yang menawarkan berbagai fitur kesehatan dan kebugaran, serta integrasi yang mulus dengan perangkat Galaxy lainnya. Smartwatch ini dilengkapi dengan sensor BioActive yang mampu memantau detak jantung, tekanan darah, kadar oksigen dalam darah, dan kualitas tidur. Selain itu, Galaxy Watch juga mendukung fitur-fitur lain seperti menerima panggilan dan pesan, mengontrol musik, dan memberikan notifikasi.",
			IsFeatured:    false,
			Discount:      0.00,
			Price:         1225000,
			AverageRating: 4.8,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745420675/samsung_watch_03_bmlayk.webp",
			},
		},
		{
			Category:      "Gadget & Electronics",
			Name:          "HUAWEI WATCH FIT Special Edition Smartwatch",
			Description:   "HUAWEI WATCH FIT Special Edition Smartwatch | 1.64 HD AMOLED | 24/7 Active Health Management | Built-in GPS | Fast Charging. Notifikasi panggilan Bluetooth & balas pesan cepat Kompatibel dengan luas, bisa digunakan bersama semua OS Tersedia dalam 3 varian warna: Nebula Pink, Forest Green, Starry Black.",
			IsFeatured:    false,
			Discount:      0.03,
			Price:         625000,
			AverageRating: 4.2,
			Stock:         10,
			Sold:          20,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421186/huawei_smartwatch_04_r8ftp5.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745421185/huawei_smartwatch_02_ihjja7.webp",
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
				ImageURL:  img,
			})
		}
	}
}

func SeedMenShoes(db *gorm.DB) {
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
			Category:      "Men's Shoes",
			Name:          "Sepatu Sneakers Olahraga Pria Casual",
			Description:   "Sepatu sneakers gaya sporty & nyaman untuk kegiatan harian, cocok untuk nongkrong dan jalan santai.",
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
			Category:      "Men's Shoes",
			Name:          "DES SNEAKERS Sepatu Pria Vans Classic",
			Description:   "Sepatu Vans dengan desain klasik, toe cap kuat, collar empuk, outsole waffle karet khas Vans.",
			Price:         475000,
			AverageRating: 4.6,
			Stock:         100,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.12,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536262/sneaker_shoes_01_nssqgb.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536262/sneaker_shoes_02_mctuky.webp",
			},
		},
		{
			Category:      "Men's Shoes",
			Name:          "Converse Allstar Sepatu Sekolah Sepatu ALL STAR CLASSIC",
			Description:   "Sneakers ikonik Converse All Star dengan konstruksi tahan lama dan gaya klasik yang tetap relevan.",
			Price:         525000,
			AverageRating: 4.7,
			Stock:         100,
			Sold:          90,
			IsFeatured:    false,
			Discount:      0.12,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker2_shoes_01_rc7i1l.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker2_shoes_02_iluvmx.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536263/sneaker3_shoes_02_viyrm9.webp",
			},
		},
		{
			Category:      "Men's Shoes",
			Name:          "Sandal Pria Nike Offcourt Slide Black",
			Description:   "Sandal ringan, empuk, dan sporty dengan busa lembut di tali dan midsole untuk kenyamanan ekstra.",
			Price:         225000,
			AverageRating: 4.5,
			Stock:         100,
			Sold:          80,
			IsFeatured:    false,
			Discount:      0.07,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536494/03sandals01_ogodhf.webp",
			},
		},
		{
			Category:      "Men's Shoes",
			Name:          "Sepatu Dokmart pria terlaris xaxinara footwear",
			Description:   "Sepatu boot kokoh dengan kualitas kulit premium, cocok untuk tampilan punk dan kasual.",
			Price:         465000,
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
			Category:      "Men's Shoes",
			Name:          "Bata Preseley Feather-Light Sendal Sintetis Kulit",
			Description:   "Sandal Bata adalah merek alas kaki yang populer di Indonesia, dikenal dengan kualitas dan keawetannya. Bata menawarkan berbagai jenis sandal, mulai dari model flat hingga sandal dengan hak, dengan desain yang beragam dan cocok untuk berbagai kegiatan, baik santai sehari-hari maupun untuk acara khusus. Sandal Bata seringkali terbuat dari bahan seperti PU (Polyurethane), kulit asli, dan karet, yang memberikan kenyamanan dan daya tahan.",
			IsFeatured:    false,
			Discount:      0.05,
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
			Category:      "Men's Shoes",
			Name:          "Sepatu Dokmart pria terlaris xaxinara footwear",
			Description:   "Sepatu Docmart pria terlaris Xaxinara Footwear adalah sepatu boot dengan desain ikonik yang kokoh dan tahan lama, dikenal karena kualitas kulitnya yang premium dan jahitan yang kuat. Sepatu ini sering dipilih untuk tampilan kasual atau punk, serta cocok untuk berbagai aktivitas karena sol karetnya yang tahan slip dan nyaman.",
			IsFeatured:    false,
			Price:         225000,
			AverageRating: 4.2,
			Stock:         100,
			Sold:          70,
			Discount:      0.00,
			Images: []string{
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/02formal01_nojgda.webp",
				"https://res.cloudinary.com/dp1xbgxdn/image/upload/v1745536998/02formal02_ihkwdw.webp",
			},
		},
		{
			Category:      "Men's Shoes",
			Name:          "Kenfa - Mora Black Sepatu Pria Loafer Formal Kerja Kantor Kuliah Slip On Basic Hitam",
			Description:   "Sepatu Kenfa Mora Basic Hitam adalah sepatu formal pria dengan model slip-on yang elegan dan cocok untuk berbagai acara, baik formal maupun kasual. Sepatu ini dibuat dengan material berkualitas tinggi dari pengrajin berpengalaman, memberikan tampilan yang berkelas dan nyaman untuk dipakai sehari-hari, misalnya di kantor atau kuliah",
			IsFeatured:    false,
			Discount:      0.12,
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
			Category:      "Men's Shoes",
			Name:          "Paulmay Sepatu Formal Kerja Venesia",
			Description:   "Paulmay Sepatu Formal Kerja Venesia adalah sepatu kulit formal yang cocok untuk berbagai acara, termasuk kerja dan kegiatan formal lainnya. Sepatu ini dikenal sebagai produk dari merek Paulmay, sebuah brand fashion lokal Indonesia yang awalnya fokus pada sepatu kulit. Venesia kemungkinan adalah nama model spesifik dari sepatu formal ini.",
			IsFeatured:    false,
			Discount:      0.12,
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
				ImageURL:  img,
			})
		}
	}
}

func SeedReviews(db *gorm.DB) {
	var products []models.Product
	var customer1 models.User
	var customer2 models.User

	// Ambil semua produk
	if err := db.Find(&products).Error; err != nil {
		log.Println("Failed to fetch products:", err)
		return
	}

	// Ambil user berdasarkan email
	db.Where("email = ?", "customer01@example.com").First(&customer1)
	db.Where("email = ?", "customer02@example.com").First(&customer2)

	if customer1.ID == uuid.Nil || customer2.ID == uuid.Nil {
		log.Println("Customer seeding belum dilakukan.")
		return
	}

	sampleComments := []string{
		"Produk sangat bagus dan sesuai deskripsi!",
		"Pengiriman cepat dan kualitas oke.",
		"Harga terjangkau dan barang berkualitas.",
		"Sangat direkomendasikan, pasti beli lagi.",
		"Packing rapi dan aman, mantap!",
		"Cocok banget dipakai harian.",
	}

	for _, product := range products {
		reviews := []models.Review{
			{
				ID:        uuid.New(),
				ProductID: product.ID,
				UserID:    customer1.ID,
				Rating:    rand.Intn(2) + 4, // 4 atau 5
				Comment:   sampleComments[rand.Intn(len(sampleComments))],
			},
			{
				ID:        uuid.New(),
				ProductID: product.ID,
				UserID:    customer2.ID,
				Rating:    rand.Intn(2) + 4,
				Comment:   sampleComments[rand.Intn(len(sampleComments))],
			},
		}
		if err := db.Create(&reviews).Error; err != nil {
			log.Printf("Gagal menyimpan review untuk produk %s: %v", product.Name, err)
		}
	}

	log.Println("✅ Review seeding completed")
}

func SeedNotificationTypes(db *gorm.DB) {
	types := []models.NotificationType{
		// Transaksi Pembelian
		{ID: uuid.New(), Code: "pending_payment", Title: "Waiting for Payment", Category: "transaction", DefaultEnabled: true},
		{ID: uuid.New(), Code: "waiting_confirmation", Title: "Waiting for Confirmation", Category: "transaction", DefaultEnabled: true},
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
	var customer models.User
	if err := db.Where("email = ?", "customer01@example.com").First(&customer).Error; err != nil {
		log.Println("User customer01@example.com tidak ditemukan")
		return
	}

	// 🔹 1. Seed Addresses
	addresses := []models.Address{
		{
			ID:         uuid.New(),
			UserID:     customer.ID,
			Name:       "Alamat Rumah",
			IsMain:     true,
			Address:    "Jl. Merdeka No.123",
			ProvinceID: 11, CityID: 156, DistrictID: 1928, SubdistrictID: 25033, PostalCodeID: 25043,
			Province: "DKI Jakarta", City: "Jakarta Barat", District: "Grogol Petamburan", Subdistrict: "Tomang", PostalCode: "11440",
			Phone:     "081234567890",
			CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
		{
			ID:     uuid.New(),
			UserID: customer.ID,

			Name:       "Alamat Kantor",
			IsMain:     false,
			Address:    "Jl. Sudirman No.99",
			ProvinceID: 11, CityID: 156, DistrictID: 1904, SubdistrictID: 24978, PostalCodeID: 24988,
			Province: "DKI Jakarta", City: "Jakarta Pusat", District: "Tanah Abang", Subdistrict: "Bendungan Hilir", PostalCode: "10210",
			Phone:     "089876543210",
			CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&addresses).Error; err != nil {
		log.Println("Gagal seed address:", err)
	}

	for i := 1; i <= 2; i++ {
		orderID := uuid.New()
		address := addresses[i-1]
		status := []string{"success", "pending"}[i-1]
		paymentStatus := []string{"success", "pending"}[i-1]

		shipmentID := uuid.New()
		shipment := models.Shipment{
			ID:           shipmentID,
			OrderID:      orderID,
			TrackingCode: fmt.Sprintf("JNE00%d", i),
			Status:       "pending",
			Notes:        utils.ToPtr("Segera dikirim"),
		}

		order := models.Order{
			ID:              orderID,
			ShipmentID:      shipmentID, // ✅ Disisipkan di sini
			UserID:          customer.ID,
			InvoiceNumber:   fmt.Sprintf("INV/SEED/%d", time.Now().UnixNano()),
			AddressID:       address.ID,
			Courier:         "JNE",
			Status:          status,
			Total:           200000,
			ShippingCost:    15000,
			Tax:             10000,
			PaymentLink:     "https://www.ahmadfiqrioemry.com",
			AmountToPay:     205000,
			VoucherDiscount: 20000,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			Items: []models.OrderItem{
				{
					ID:          uuid.New(),
					ProductID:   uuid.New(),
					ProductName: "Produk Dummy " + strconv.Itoa(i),
					ProductSlug: utils.GenerateSlug("Produk Dummy " + strconv.Itoa(i)),
					ImageURL:    "https://placehold.co/300x300",
					Price:       100000,
					Quantity:    2,
					Subtotal:    200000,
				},
			},
		}

		payment := models.Payment{
			ID:      uuid.New(),
			UserID:  customer.ID,
			OrderID: orderID,
			Method:  "bank_transfer",
			Status:  paymentStatus,
			PaidAt:  time.Now(),
			Total:   205000,
		}

		// ✅ Urutan create: Order -> Payment -> Shipment
		if err := db.Create(&order).Error; err != nil {
			log.Println("Gagal seed order:", err)
		}
		if err := db.Create(&payment).Error; err != nil {
			log.Println("Gagal seed payment:", err)
		}
		if err := db.Create(&shipment).Error; err != nil {
			log.Println("Gagal seed shipment:", err)
		}
	}

	log.Println("✅ Seed orders, payments, shipments, addresses untuk customer01@example.com selesai")
}

func SeedCustomerNotifications(db *gorm.DB) {
	var user models.User
	if err := db.Where("email = ?", "customer01@example.com").First(&user).Error; err != nil {
		log.Println("❌ Gagal menemukan user: customer01@example.com")
		return
	}

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

	for _, n := range notifications {
		if err := db.Create(&n).Error; err != nil {
			log.Printf("❌ Gagal membuat notifikasi: %s\n", n.Title)
		}
	}

	log.Println("✅ Notifikasi untuk customer01@example.com berhasil disimpan.")
}

func generateNotificationSettingsForUser(db *gorm.DB, user models.User) {
	var notifTypes []models.NotificationType
	if err := db.Find(&notifTypes).Error; err != nil {
		log.Printf("Failed to get notification types for user %s: %v", user.Email, err)
		return
	}

	for _, nt := range notifTypes {
		for _, channel := range []string{"email", "browser"} {
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
