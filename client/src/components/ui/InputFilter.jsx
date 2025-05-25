import { Input } from "@/components/ui/input";

export const InputFilter = ({
  label = "",
  value,
  onChange,
  placeholder = "",
  inputMode = "numeric",
  type = "text",
}) => {
  return (
    <div className="space-y-1">
      {label && <label className="text-sm font-medium block">{label}</label>}
      <Input
        type={type}
        inputMode={inputMode}
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder={placeholder}
        className="text-sm"
      />
    </div>
  );
};
