import { Input } from "@/components/ui/input";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import ProductCard from "@/components/admin/products/ProductCard";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import NoProductResult from "@/components/admin/products/NoProductResult";
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/card";
import { SortableHeader } from "@/components/ui/SortableHeader";

import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";

const ProductsList = () => {
  const { search, sort, page, limit, setSearch, setSort, setPage } =
    useQueryParamsStore();

  const { data, isLoading, isError } = useSearchProductsQuery({
    search,
    page,
    limit,
    sort,
  });

  const products = data?.data || [];

  const pagination = data?.pagination;
  console.log(products);
  console.log(pagination);
  return (
    <section className="section space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Products List</h2>
        <p className="text-muted-foreground text-sm">
          Manage all user transactions and monitor payment activities.
        </p>
      </div>

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          placeholder="Search product name or description"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2"
        />
        <div className="flex gap-3">
          <select
            className="border rounded px-3 py-2 text-sm"
            value={sort}
            onChange={(e) => {
              setPage(1);
              setSort(e.target.value);
            }}
          >
            <option value="paid_at desc">Newest</option>
            <option value="paid_at asc">Oldest</option>
          </select>
        </div>
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch className="mt-10" />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : products.length === 0 ? (
            <NoProductResult search={search} />
          ) : (
            <div className="hidden md:block w-full">
              <Table>
                <TableHeader>
                  <TableRow>
                    <SortableHeader
                      label="Name"
                      sortKey="name"
                      currentSort={sort}
                      onSortChange={(val) => {
                        setPage(1);
                        setSort(val);
                      }}
                    />
                    <TableHead className="text-left">Category</TableHead>
                    <SortableHeader
                      label="Price"
                      sortKey="price"
                      currentSort={sort}
                      onSortChange={(val) => {
                        setPage(1);
                        setSort(val);
                      }}
                    />
                    <TableHead className="text-center">Discount</TableHead>
                    <TableHead className="text-center">Featured</TableHead>
                    <TableHead className="text-center">Status</TableHead>
                    <TableHead className="text-center">Action</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody className="h-12">
                  {products.map((product) => (
                    <ProductCard key={product.id} product={product} />
                  ))}
                </TableBody>
              </Table>
            </div>
          )}

          {pagination && (
            <Pagination
              page={pagination.page}
              limit={pagination.limit}
              total={pagination.totalRows}
              onPageChange={(p) => setPage(p)}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default ProductsList;
