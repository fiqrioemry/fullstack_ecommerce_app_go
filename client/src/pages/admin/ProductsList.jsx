import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
import {
  Select,
  SelectItem,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { SortableHeader } from "@/components/ui/SortableHeader";
import ProductCard from "@/components/admin/products/ProductCard";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import NoProductResult from "@/components/admin/products/NoProductResult";

const ProductsList = () => {
  const {
    search,
    status,
    setStatus,
    sort,
    page,
    limit,
    setSearch,
    setSort,
    setPage,
  } = useQueryParamsStore();

  const { data, isLoading, isError } = useSearchProductsQuery({
    search,
    page,
    limit,
    status,
    sort,
  });

  const products = data?.data || [];

  const pagination = data?.pagination || null;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Products List"
        description="Manage all user transactions and monitor payment activities."
      />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          placeholder="Search product name or description"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2 input"
        />
        <Select
          value={status}
          onValueChange={(val) => {
            setPage(1);
            setStatus(val);
          }}
        >
          <SelectTrigger className="w-48 bg-background">
            <SelectValue placeholder="Filter Status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All</SelectItem>
            <SelectItem value="active">Active</SelectItem>
            <SelectItem value="inactive">Inactive</SelectItem>
            <SelectItem value="featured">Featured</SelectItem>
            <SelectItem value="unfeatured">Unfeatured</SelectItem>
          </SelectContent>
        </Select>
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
                    <TableHead className="text-center">Stock</TableHead>
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
