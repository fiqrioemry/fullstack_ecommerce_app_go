import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { Badge } from "@/components/ui/badge";
import { formatDate, formatRupiah } from "@/lib/utils";
import { ChevronDown, ChevronUp } from "lucide-react";
import { UpdateProduct } from "./UpdateProduct";
import { DeleteProduct } from "./DeleteProduct";

export const ProductCard = ({ products, sort, setSort }) => {
  const renderSortIcon = (field) => {
    if (sort === `${field}_asc`)
      return <ChevronUp className="w-4 h-4 inline" />;
    if (sort === `${field}_desc`)
      return <ChevronDown className="w-4 h-4 inline" />;
    return null;
  };

  return (
    <>
      <div className="hidden md:block w-full">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("name")}
              >
                product name {renderSortIcon("name")}
              </TableHead>
              <TableHead>Category</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("price")}
              >
                Price
                {renderSortIcon("price")}
              </TableHead>
              <TableHead>Discount</TableHead>
              <TableHead
                className="cursor-pointer"
                onClick={() => setSort("stock")}
              >
                Stock {renderSortIcon("stock")}
              </TableHead>
              <TableHead>Featured </TableHead>
              <TableHead>Status</TableHead>
              <TableHead>Action</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody className="h-12">
            {products.map((product) => (
              <TableRow key={product.id}>
                <TableCell>
                  <div className="flex items-center justify-center gap-4">
                    <img
                      alt={product.name}
                      src={product.images?.[0]}
                      className="h-12 w-12 object-cover rounded"
                    />
                    <div>
                      {product.name.length > 15
                        ? product.name.slice(0, 15) + "..."
                        : product.name}
                    </div>
                  </div>
                </TableCell>
                <TableCell>{product.category}</TableCell>
                <TableCell>{formatRupiah(product.price)}</TableCell>
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
            ))}
          </TableBody>
        </Table>
      </div>

      {/* Mobile view */}
      {/* Mobile view */}
      <div className="md:hidden space-y-4 p-4 w-full">
        {products.map((product) => (
          <div
            key={product.id}
            className="border rounded-xl p-4 shadow-sm space-y-3"
          >
            {/* Header */}
            <div className="flex items-center gap-4">
              <img
                src={product.images?.[0]}
                alt={product.name}
                className="h-16 w-16 object-cover rounded"
              />
              <div className="flex-1">
                <h3 className="text-base font-semibold truncate">
                  {product.name}
                </h3>
                <p className="text-sm text-muted-foreground">
                  {product.category}
                </p>
              </div>
            </div>

            {/* Info Grid */}
            <div className="grid grid-cols-2 gap-x-4 gap-y-2 text-sm text-muted-foreground">
              <div>
                <span className="block font-medium text-foreground">Price</span>
                {formatRupiah(product.price)}
              </div>
              <div>
                <span className="block font-medium text-foreground">
                  Discount
                </span>
                {product.discount}%
              </div>
              <div>
                <span className="block font-medium text-foreground">Stock</span>
                {product.stock}
              </div>
              <div>
                <span className="block font-medium text-foreground">
                  Rating
                </span>
                {product.averageRating ?? "-"}
              </div>
            </div>

            {/* Status Badges */}
            <div className="flex gap-2 flex-wrap">
              {product.isFeatured ? (
                <Badge
                  className="bg-gray-100 text-gray-800"
                  variant="secondary"
                >
                  Featured
                </Badge>
              ) : (
                <Badge className="bg-red-100 text-red-700" variant="outline">
                  Unfeatured
                </Badge>
              )}

              {product.isActive ? (
                <Badge
                  className="bg-green-100 text-green-700"
                  variant="secondary"
                >
                  Active
                </Badge>
              ) : (
                <Badge className="bg-gray-200 text-gray-700" variant="outline">
                  Inactive
                </Badge>
              )}
            </div>

            {/* Actions */}
            <div className="flex justify-end gap-2 pt-2">
              <UpdateProduct product={product} />
              <DeleteProduct product={product} />
            </div>
          </div>
        ))}
      </div>
    </>
  );
};
