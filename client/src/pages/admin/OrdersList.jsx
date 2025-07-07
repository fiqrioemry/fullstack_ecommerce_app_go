import { useDebounce } from "@/hooks/useDebounce";
import { orderStatusOptions } from "@/lib/constant";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { useOrderStore } from "@/store/useOrderStore";
import { Pagination } from "@/components/ui/Pagination";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { OrderCard } from "@/components/admin/orders/OrderCard";
import { RecordNotFound } from "@/components/ui/RecordNotFound";

const OrdersList = () => {
  const { q, status, sort, page, limit, setQ, setPage, setSort, setStatus } =
    useOrderStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError } = useAllOrdersQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
    status,
  });

  const orders = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle title="Orders List" description="See all user orders" />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by product"}
        />

        <SelectFilter
          value={status}
          onChange={setStatus}
          options={orderStatusOptions}
        />
      </div>

      {/* 🧾 Content */}
      {isLoading ? (
        <LoadingSearch />
      ) : isError ? (
        <ErrorDialog onRetry={refetch} />
      ) : orders.length === 0 ? (
        <RecordNotFound title="No transactions found" q={q} />
      ) : (
        <OrderCard orders={orders} sort={sort} setSort={setSort} />
      )}
      {pagination && (
        <Pagination
          page={pagination.page}
          onPageChange={setPage}
          limit={pagination.limit}
          total={pagination.totalRows}
        />
      )}
    </section>
  );
};

export default OrdersList;
