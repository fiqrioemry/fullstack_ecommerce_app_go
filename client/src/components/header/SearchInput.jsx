import { useRef } from "react";
import { X } from "lucide-react";

export const SearchInput = ({ value, onChange, onKeyDown, onClear }) => {
  const inputRef = useRef(null);

  return (
    <div className="relative w-full">
      <input
        value={value}
        ref={inputRef}
        onKeyDown={onKeyDown}
        placeholder="Search products..."
        onChange={(e) => onChange(e.target.value)}
        className="w-full border px-3 py-2 pr-10 rounded-md text-sm"
      />
      {value && (
        <button
          onClick={() => {
            onClear();
            inputRef.current?.focus();
          }}
          className="absolute right-2 top-1/2 -translate-y-1/2 text-gray-500 hover:text-red-500"
        >
          <X />
        </button>
      )}
    </div>
  );
};
