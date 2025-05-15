import { Pencil, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { TableRow, TableCell } from "@/components/ui/table";

const ProductCard = ({ product }) => {
  return (
    <TableRow key={product.id}>
      <TableCell className="text-left">
        <div className="flex items-center gap-4">
          <img
            src={product.imageUrl?.[0]}
            alt={product.name}
            className="h-12 w-12 object-cover rounded"
          />
          <div>
            {product.name.length > 20
              ? product.name.slice(0, 20) + "..."
              : product.name}
          </div>
        </div>
      </TableCell>
      <TableCell className="text-left">{product.category}</TableCell>
      <TableCell className="text-left">
        Rp {product.price.toLocaleString("id-ID")}
      </TableCell>
      <TableCell>
        {product.discount ? `${product.discount.toLocaleString()}` : "-"}
      </TableCell>
      <TableCell>
        {product.isFeatured ? (
          <span className="text-gray-600 text-xs bg-gray-100 px-2 py-1 rounded">
            featured
          </span>
        ) : (
          <span>-</span>
        )}
      </TableCell>
      <TableCell>
        {product.isActive ? (
          <span className="text-green-600 text-xs bg-green-100 px-2 py-1 rounded">
            Active
          </span>
        ) : (
          <span className="text-gray-600 text-xs bg-gray-100 px-2 py-1 rounded">
            Inactive
          </span>
        )}
      </TableCell>
      <TableCell className="text-center space-x-4">
        <Button variant="outline" size="icon" className="gap-1">
          <Pencil size={16} />
        </Button>
        <Button variant="destructive" size="icon" className="gap-1">
          <Trash2 size={16} />
        </Button>
      </TableCell>
    </TableRow>
  );
};

export default ProductCard;
