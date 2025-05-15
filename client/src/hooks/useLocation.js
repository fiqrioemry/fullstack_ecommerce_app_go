import { useQuery } from "@tanstack/react-query";
import * as location from "@/services/location";

// PROVINCES
export const useProvincesQuery = () =>
  useQuery({ queryKey: ["provinces"], queryFn: location.getProvinces });

export const useSearchProvincesQuery = (q) =>
  useQuery({
    queryKey: ["searchProvinces", q],
    queryFn: () => location.searchProvinces(q),
    enabled: !!q,
  });

// CITIES
export const useSearchCitiesQuery = (q) =>
  useQuery({
    queryKey: ["searchCities", q],
    queryFn: () => location.searchCities(q),
    enabled: !!q,
  });

export const useCitiesByProvinceQuery = (provinceId) =>
  useQuery({
    queryKey: ["cities", provinceId],
    queryFn: () => location.getCitiesByProvinceId(provinceId),
    enabled: !!provinceId,
  });

// DISTRICTS
export const useDistrictsByCityQuery = (cityId) =>
  useQuery({
    queryKey: ["districts", cityId],
    queryFn: () => location.getDistrictsByCityId(cityId),
    enabled: !!cityId,
  });

// SUBDISTRICTS
export const useSubdistrictsByDistrictQuery = (districtId) =>
  useQuery({
    queryKey: ["subdistricts", districtId],
    queryFn: () => location.getSubdistrictsByDistrictId(districtId),
    enabled: !!districtId,
  });

// POSTAL CODES
export const usePostalcodesBySubdistrictQuery = (subdistrictId) =>
  useQuery({
    queryKey: ["postalcodes", subdistrictId],
    queryFn: () => location.getPostalcodesBySubdistrictId(subdistrictId),
    enabled: !!subdistrictId,
  });
