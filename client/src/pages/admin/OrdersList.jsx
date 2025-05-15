import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { OrderCard } from "@/components/admin/orders/OrderCard";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { NoOrderResult } from "@/components/admin/orders/NoOrderResult";
import { Pagination } from "@/components/ui/pagination";

const OrdersList = () => {
  const {
    search,
    status,
    sort,
    page,
    limit,
    setSearch,
    setSort,
    setPage,
    setStatus,
  } = useQueryParamsStore();

  const { data, isLoading, isError, error } = useAllOrdersQuery(
    search,
    page,
    limit,
    sort,
    status
  );

  const orders = data?.data || [];
  const pagination = data?.pagination;
  console.log(data);

  return (
    <section className="section px-4 py-8 space-y-6">
      <div className="text-center space-y-1 mb-6">
        <h2 className="text-2xl font-bold">Orders List</h2>
        <p className="text-sm text-muted-foreground">See all user orders</p>
      </div>

      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2"
          placeholder="Search something in orders"
        />

        <Select
          value={sort}
          onValueChange={(val) => {
            setPage(1);
            setStatus(val);
          }}
        >
          <SelectTrigger className="w-60 bg-background">
            <SelectValue placeholder="Select status" />
          </SelectTrigger>
          <SelectContent>
            <SelectGroup>
              <SelectLabel>Select status</SelectLabel>
              <SelectItem>all status</SelectItem>
              <SelectItem value="success">success</SelectItem>
              <SelectItem value="pending">pending</SelectItem>
              <SelectItem value="failed">failed</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      {/* 🧾 Content */}
      {isLoading ? (
        <LoadingSearch />
      ) : isError ? (
        <ErrorDialog onRetry={refetch} />
      ) : orders.length === 0 ? (
        <NoOrderResult />
      ) : (
        <>
          <OrderCard orders={orders} />
          {pagination && (
            <Pagination
              page={pagination.page}
              limit={pagination.limit}
              total={pagination.totalRows}
              onPageChange={(p) => setPage(p)}
            />
          )}
        </>
      )}
    </section>
  );
};

export default OrdersList;
