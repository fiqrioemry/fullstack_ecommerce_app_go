import { useDebounce } from "@/hooks/useDebounce";
import { productStatusOptions } from "@/lib/constant";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { useSearchProductsQuery } from "@/hooks/useProduct";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { RecordNotFound } from "@/components/ui/RecordNotFound";
import { SectionTitle } from "@/components/header/SectionTitle";
import { ProductCard } from "@/components/admin/products/ProductCard";
import { useProductStore } from "@/store/useProductStore";

const ProductsList = () => {
  const { page, limit, q, sort, setPage, status, setQ, setSort, setStatus } =
    useProductStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError } = useSearchProductsQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
    status,
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
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by product name or description"}
        />

        <SelectFilter
          value={status}
          onChange={setStatus}
          options={productStatusOptions}
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch className="mt-10" />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : products.length === 0 ? (
            <RecordNotFound title={"No Product record found"} q={q} />
          ) : (
            <ProductCard sort={sort} setSort={setSort} products={products} />
          )}

          {pagination && (
            <Pagination
              page={pagination.page}
              onPageChange={setPage}
              limit={pagination.limit}
              total={pagination.totalRows}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default ProductsList;
