import { Button } from "@/components/ui/Button";
import { ChevronLeft, ChevronRight } from "lucide-react";

const Pagination = ({ page, limit, total, onPageChange }) => {
  const totalPages = Math.ceil(total / limit);
  const start = (page - 1) * limit + 1;
  const end = Math.min(start + limit - 1, total);

  return (
    <div className="flex items-center justify-between p-4 text-sm w-full">
      <div className="text-muted-foreground space-x-2">
        <span> Showing</span>
        <span className="font-medium text-primary">
          {start}-{end}
        </span>
        <span>of</span>
        <span className="font-medium text-primary">{total}</span>
      </div>
      <div className="flex items-center gap-2">
        <Button
          size="icon"
          variant="outline"
          disabled={page === 1}
          className="disabled:opacity-50"
          onClick={() => onPageChange(page - 1)}
        >
          <ChevronLeft className="w-4 h-4" />
        </Button>
        <span className="font-medium text-sm w-8 text-center">{page}</span>
        <Button
          size="icon"
          variant="outline"
          disabled={page >= totalPages}
          className="disabled:opacity-50"
          onClick={() => onPageChange(page + 1)}
        >
          <ChevronRight className="w-4 h-4" />
        </Button>
      </div>
    </div>
  );
};

export { Pagination };
