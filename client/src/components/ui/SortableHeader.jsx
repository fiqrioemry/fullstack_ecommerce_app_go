// components/ui/SortableHeader.tsx
import { ArrowDown, ArrowUp } from "lucide-react";
import { TableHead } from "@/components/ui/Table";

export const SortableHeader = ({
  label,
  sortKey,
  currentSort,
  onSortChange,
}) => {
  const [key, direction] = currentSort.split("_");
  const isActive = key === sortKey;
  const nextSort = isActive && direction === "asc" ? "desc" : "asc";
  const finalSort = `${sortKey}_${nextSort}`;

  return (
    <TableHead
      onClick={() => onSortChange(finalSort)}
      className="cursor-pointer select-none text-left"
    >
      <div className="flex items-center text-muted-foreground gap-4">
        {label}
        {isActive ? (
          direction === "asc" ? (
            <ArrowUp size={14} />
          ) : (
            <ArrowDown size={14} />
          )
        ) : (
          <span className="text-xs text-muted-foreground"></span>
        )}
      </div>
    </TableHead>
  );
};
