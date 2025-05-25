package services

import (
	"errors"

	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type AddressService interface {
	GetAddresses(userID string, param dto.AddressQueryParam) ([]dto.AddressResponse, *dto.PaginationResponse, error)
	AddAddressWithLocation(userID string, req dto.CreateAddressRequest) error
	UpdateAddressWithLocation(userID string, addressID string, req dto.UpdateAddressRequest) error
	DeleteAddress(userID string, addressID string) error
	SetMainAddress(userID string, addressID string) error
	GetMainAddress(userID string) (*models.Address, error)
}

type addressService struct {
	AddressRepo  repositories.AddressRepository
	LocationRepo repositories.LocationRepository
}

func NewAddressService(AddressRepo repositories.AddressRepository, locationRepo repositories.LocationRepository) AddressService {
	return &addressService{AddressRepo: AddressRepo, LocationRepo: locationRepo}
}

func (s *addressService) GetMainAddress(userID string) (*models.Address, error) {
	return s.AddressRepo.GetMainAddress(userID)
}

func (s *addressService) GetAddresses(userID string, param dto.AddressQueryParam) ([]dto.AddressResponse, *dto.PaginationResponse, error) {

	addresses, total, err := s.AddressRepo.GetAddressesByUserID(userID, param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.AddressResponse
	for _, a := range addresses {
		result = append(result, dto.AddressResponse{
			ID:            a.ID.String(),
			Name:          a.Name,
			Address:       a.Address,
			ProvinceID:    int(a.ProvinceID),
			Province:      a.Province,
			CityID:        int(a.CityID),
			City:          a.City,
			DistrictID:    int(a.DistrictID),
			District:      a.District,
			SubdistrictID: int(a.SubdistrictID),
			Subdistrict:   a.Subdistrict,
			PostalCode:    a.PostalCode,
			PostalCodeID:  int(a.PostalCodeID),
			Phone:         a.Phone,
			IsMain:        a.IsMain,
			CreatedAt:     a.CreatedAt.Format("2006-01-02"),
		})
	}

	totalPages := int((total + int64(param.Limit) - 1) / int64(param.Limit))
	pagination := &dto.PaginationResponse{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *addressService) AddAddressWithLocation(userID string, req dto.CreateAddressRequest) error {
	// 1. Ambil Province
	provinces, err := s.LocationRepo.GetAllProvinces()
	if err != nil {
		return errors.New("failed to fetch provinces")
	}
	var province models.Province
	for _, p := range provinces {
		if p.ID == req.ProvinceID {
			province = p
			break
		}
	}
	if province.ID == 0 {
		return errors.New("invalid province ID")
	}

	// 2. Ambil City
	cities, err := s.LocationRepo.GetCitiesByProvinceID(req.ProvinceID)
	if err != nil {
		return errors.New("failed to fetch cities")
	}
	var city models.City
	for _, c := range cities {
		if c.ID == req.CityID {
			city = c
			break
		}
	}
	if city.ID == 0 {
		return errors.New("invalid city ID")
	}

	// 3. Ambil District
	districts, err := s.LocationRepo.GetDistrictsByCityID(req.CityID)
	if err != nil {
		return errors.New("failed to fetch districts")
	}
	var district models.District
	for _, d := range districts {
		if d.ID == req.DistrictID {
			district = d
			break
		}
	}
	if district.ID == 0 {
		return errors.New("invalid district ID")
	}

	// 4. Ambil Subdistrict
	subdistricts, err := s.LocationRepo.GetSubdistrictsByDistrictID(req.DistrictID)
	if err != nil {
		return errors.New("failed to fetch subdistricts")
	}
	var subdistrict models.Subdistrict
	for _, sd := range subdistricts {
		if sd.ID == req.SubdistrictID {
			subdistrict = sd
			break
		}
	}
	if subdistrict.ID == 0 {
		return errors.New("invalid subdistrict ID")
	}

	// 5. Ambil Postal Code
	postalCodes, err := s.LocationRepo.GetPostalCodesBySubdistrictID(req.SubdistrictID)
	if err != nil {
		return errors.New("failed to fetch postal codes")
	}
	var postalCode models.PostalCode
	for _, pc := range postalCodes {
		if pc.ID == req.PostalCodeID {
			postalCode = pc
			break
		}
	}
	if postalCode.ID == 0 {
		return errors.New("invalid postal code ID")
	}

	// 6. Buat Address
	addr := models.Address{
		ID:            uuid.New(),
		UserID:        uuid.MustParse(userID),
		Name:          req.Name,
		Address:       req.Address,
		ProvinceID:    req.ProvinceID,
		CityID:        req.CityID,
		DistrictID:    req.DistrictID,
		SubdistrictID: req.SubdistrictID,
		PostalCodeID:  req.PostalCodeID,
		Province:      province.Name,
		City:          city.Name,
		District:      district.Name,
		Subdistrict:   subdistrict.Name,
		PostalCode:    postalCode.PostalCode,
		Phone:         req.Phone,
		IsMain:        req.IsMain,
	}

	if req.IsMain {
		_ = s.AddressRepo.UnsetAllMain(userID)
	}

	return s.AddressRepo.CreateAddress(&addr)
}

func (s *addressService) UpdateAddressWithLocation(userID string, addressID string, req dto.UpdateAddressRequest) error {
	addr, err := s.AddressRepo.GetAddressByID(addressID)
	if err != nil {
		return err
	}

	// 1. Cek Province
	provinces, err := s.LocationRepo.GetAllProvinces()
	if err != nil {
		return errors.New("failed to fetch provinces")
	}
	var province models.Province
	for _, p := range provinces {
		if p.ID == req.ProvinceID {
			province = p
			break
		}
	}
	if province.ID == 0 {
		return errors.New("invalid province ID")
	}

	// 2. Cek City
	cities, err := s.LocationRepo.GetCitiesByProvinceID(req.ProvinceID)
	if err != nil {
		return errors.New("failed to fetch cities")
	}
	var city models.City
	for _, c := range cities {
		if c.ID == req.CityID {
			city = c
			break
		}
	}
	if city.ID == 0 {
		return errors.New("invalid city ID")
	}

	// 3. Cek District
	districts, err := s.LocationRepo.GetDistrictsByCityID(req.CityID)
	if err != nil {
		return errors.New("failed to fetch districts")
	}
	var district models.District
	for _, d := range districts {
		if d.ID == req.DistrictID {
			district = d
			break
		}
	}
	if district.ID == 0 {
		return errors.New("invalid district ID")
	}

	// 4. Cek Subdistrict
	subdistricts, err := s.LocationRepo.GetSubdistrictsByDistrictID(req.DistrictID)
	if err != nil {
		return errors.New("failed to fetch subdistricts")
	}
	var subdistrict models.Subdistrict
	for _, sd := range subdistricts {
		if sd.ID == req.SubdistrictID {
			subdistrict = sd
			break
		}
	}
	if subdistrict.ID == 0 {
		return errors.New("invalid subdistrict ID")
	}

	// 5. Cek PostalCode
	postalCodes, err := s.LocationRepo.GetPostalCodesBySubdistrictID(req.SubdistrictID)
	if err != nil {
		return errors.New("failed to fetch postal codes")
	}
	var postalCode models.PostalCode
	for _, pc := range postalCodes {
		if pc.ID == req.PostalCodeID {
			postalCode = pc
			break
		}
	}
	if postalCode.ID == 0 {
		return errors.New("invalid postal code ID")
	}

	addr.Name = req.Name
	addr.Address = req.Address
	addr.ProvinceID = req.ProvinceID
	addr.CityID = req.CityID
	addr.DistrictID = req.DistrictID
	addr.SubdistrictID = req.SubdistrictID
	addr.PostalCodeID = req.PostalCodeID
	addr.Province = province.Name
	addr.City = city.Name
	addr.District = district.Name
	addr.Subdistrict = subdistrict.Name
	addr.PostalCode = postalCode.PostalCode
	addr.Phone = req.Phone

	return s.AddressRepo.UpdateAddress(addr)
}

func (s *addressService) DeleteAddress(userID string, addressID string) error {
	add, err := s.AddressRepo.GetAddressByID(addressID)
	if err != nil {
		return err
	}

	if add.IsMain {
		return errors.New("cannot delete main address")
	}

	return s.AddressRepo.DeleteAddress(addressID, userID)
}

func (s *addressService) SetMainAddress(userID string, addressID string) error {
	if err := s.AddressRepo.UnsetAllMain(userID); err != nil {
		return err
	}
	return s.AddressRepo.SetMainAddress(addressID, userID)
}
