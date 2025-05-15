import { publicInstance } from ".";

// GET /api/location/provinces
export const getProvinces = async () => {
  const res = await publicInstance.get("/location/provinces");
  return res.data;
};

// GET /api/location/provinces/search?q=Jawa
export const searchProvinces = async (q) => {
  const res = await publicInstance.get(`/location/provinces/search?q=${q}`);
  return res.data;
};

// GET /api/location/provinces/:provinceId/cities
export const getCitiesByProvinceId = async (provinceId) => {
  const res = await publicInstance.get(
    `/location/provinces/${provinceId}/cities`
  );
  return res.data;
};

// GET /api/location/cities/search?q=Bandung
export const searchCities = async (q) => {
  const res = await publicInstance.get(`/location/cities/search?q=${q}`);
  return res.data;
};

// GET /api/location/cities/:cityId/districts
export const getDistrictsByCityId = async (cityId) => {
  const res = await publicInstance.get(`/location/cities/${cityId}/districts`);
  return res.data;
};

// GET /api/location/districts/:districtId/subdistricts
export const getSubdistrictsByDistrictId = async (districtId) => {
  const res = await publicInstance.get(
    `/location/districts/${districtId}/subdistricts`
  );
  return res.data;
};

// GET /api/location/subdistricts/:subdistrictId/postalcodes
export const getPostalcodesBySubdistrictId = async (subdistrictId) => {
  const res = await publicInstance.get(
    `/location/subdistricts/${subdistrictId}/postalcodes`
  );
  return res.data;
};
