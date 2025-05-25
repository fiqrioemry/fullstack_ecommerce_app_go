import { Checkbox } from "@/components/ui/checkbox";

export const CheckboxFilter = ({
  options = [],
  selectedValue,
  onChange,
  title = "Filter",
}) => {
  return (
    <div className="space-y-2">
      <h4>{title}</h4>
      <div className="flex flex-col gap-2">
        {options.map((opt) => (
          <label
            key={opt.value}
            className="flex items-center gap-2 cursor-pointer text-sm"
          >
            <Checkbox
              checked={selectedValue === opt.value}
              onCheckedChange={(checked) => onChange(checked ? opt.value : "")}
            />
            {opt.label}
          </label>
        ))}
      </div>
    </div>
  );
};
