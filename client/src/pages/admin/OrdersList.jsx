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
import { SectionTitle } from "@/components/header/SectionTitle";

const OrdersList = () => {
  const { search, status, sort, page, limit, setSearch, setPage, setStatus } =
    useQueryParamsStore();

  const { data, isLoading, isError } = useAllOrdersQuery(
    search,
    page,
    limit,
    sort,
    status
  );

  const orders = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle title="Orders List" description="See all user orders" />

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
          value={status}
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
              <SelectItem value="all">all status</SelectItem>
              <SelectItem value="success">success</SelectItem>
              <SelectItem value="pending">pending</SelectItem>
              <SelectItem value="waiting_payment">waiting_payment</SelectItem>
              <SelectItem value="failed">canceled</SelectItem>
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
