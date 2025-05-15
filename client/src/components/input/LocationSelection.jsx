import {
  useProvincesQuery,
  useCitiesByProvinceQuery,
  useDistrictsByCityQuery,
  useSubdistrictsByDistrictQuery,
  usePostalcodesBySubdistrictQuery,
} from "@/hooks/useLocation";
import { useFormContext, Controller } from "react-hook-form";
const SelectField = ({
  name,
  label,
  options,
  optionLabelKey = "name",
  optionValueKey = "id",
}) => {
  const { control } = useFormContext();

  return (
    <Controller
      control={control}
      name={name}
      rules={{ required: true }}
      render={({ field }) => (
        <div>
          <label className="block mb-1 font-medium text-sm">{label}</label>
          <select
            {...field}
            onChange={(e) => field.onChange(Number(e.target.value))}
            value={field.value || ""}
            className="w-full border px-3 py-2 rounded text-sm"
          >
            <option value="">Select {label}</option>
            {options.map((option) => (
              <option
                key={option[optionValueKey]}
                value={option[optionValueKey]}
              >
                {option[optionLabelKey]}
              </option>
            ))}
          </select>
        </div>
      )}
    />
  );
};

const LocationSelection = () => {
  const { watch } = useFormContext();

  const cityId = watch("cityId");
  const provinceId = watch("provinceId");
  const districtId = watch("districtId");
  const subdistrictId = watch("subdistrictId");

  const { data: provinces = [] } = useProvincesQuery();
  const { data: cities = [] } = useCitiesByProvinceQuery(provinceId);
  const { data: districts = [] } = useDistrictsByCityQuery(cityId);
  const { data: subdistricts = [] } =
    useSubdistrictsByDistrictQuery(districtId);
  const { data: postalCodes = [] } =
    usePostalcodesBySubdistrictQuery(subdistrictId);

  return (
    <div className="space-y-4">
      <SelectField name="provinceId" label="Province" options={provinces} />
      {provinceId ? (
        <SelectField name="cityId" label="City" options={cities} />
      ) : null}
      {cityId ? (
        <SelectField name="districtId" label="District" options={districts} />
      ) : null}
      {districtId ? (
        <SelectField
          name="subdistrictId"
          label="Subdistrict"
          options={subdistricts}
        />
      ) : null}
      {subdistrictId ? (
        <SelectField
          name="postalCodeId"
          label="Postal Code"
          options={postalCodes}
          optionLabelKey="postalCode"
        />
      ) : null}
    </div>
  );
};

export default LocationSelection;
