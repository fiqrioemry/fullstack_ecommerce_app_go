import { UpdateProduct } from "./UpdateProduct";
import { DeleteProduct } from "./DeleteProduct";
import { TableRow, TableCell } from "@/components/ui/table";

const ProductCard = ({ product }) => {
  return (
    <TableRow key={product.id}>
      <TableCell className="text-left">
        <div className="flex items-center gap-4">
          <img
            src={product.images?.[0]}
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
      <TableCell>{product.discount} %</TableCell>
      <TableCell>{product.stock}</TableCell>
      <TableCell>
        {product.isFeatured ? (
          <span className="text-gray-600 text-xs bg-gray-100 px-2 py-1 rounded">
            Featured
          </span>
        ) : (
          <span className="text-red-600 text-xs bg-red-100 px-2 py-1 rounded">
            Unfeatured
          </span>
        )}
      </TableCell>

      <TableCell>
        {product.isActive === true ? (
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
        <UpdateProduct product={product} />
        <DeleteProduct product={product} />
      </TableCell>
    </TableRow>
  );
};

export default ProductCard;
