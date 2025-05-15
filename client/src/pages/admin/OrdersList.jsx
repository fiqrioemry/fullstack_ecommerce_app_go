import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import OrderCard from "@/components/admin/orders/OrderCard";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import NoOrderResult from "@/components/admin/orders/NoOrderResult";

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

  return (
    <section className="section px-4 py-8 space-y-6">
      <div className="text-center space-y-1 mb-6">
        <h2 className="text-2xl font-bold">Orders List</h2>
        <p className="text-sm text-muted-foreground">See all user orders</p>
      </div>

      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row justify-between gap-4">
        <Input
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          placeholder="Search by invoice, user name or email"
          className="md:w-1/2"
        />
        <div className="flex gap-3">
          <select
            className="border rounded px-3 py-2 text-sm"
            value={status}
            onChange={(e) => {
              setPage(1);
              setStatus(e.target.value);
            }}
          >
            <option value="">All Status</option>
            <option value="pending">Pending</option>
            <option value="success">Success</option>
            <option value="failed">Failed</option>
            <option value="canceled">Canceled</option>
          </select>
          <select
            className="border rounded px-3 py-2 text-sm"
            value={sort}
            onChange={(e) => {
              setPage(1);
              setSort(e.target.value);
            }}
          >
            <option value="created_at desc">Newest</option>
            <option value="created_at asc">Oldest</option>
            <option value="product_name asc">Product Name A-Z</option>
            <option value="product_name desc">Product Name Z-A</option>
          </select>
        </div>
      </div>

      {/* 🔄 Content */}
      {isLoading ? (
        <LoadingSearch className="mt-10" />
      ) : isError ? (
        <ErrorDialog message={error?.message || "Failed to load orders"} />
      ) : orders.length === 0 ? (
        <NoOrderResult search={search} />
      ) : (
        <>
          <div className="space-y-4">
            {orders.map((order) => (
              <OrderCard key={order.id} order={order} />
            ))}
          </div>

          {/* Pagination */}
          {pagination && (
            <div className="flex items-center justify-between pt-6">
              <Button
                variant="outline"
                onClick={() => setPage((p) => Math.max(p - 1, 1))}
                disabled={page === 1}
              >
                Previous
              </Button>
              <p className="text-sm text-muted-foreground">
                Page {pagination.page} of {pagination.totalPages}
              </p>
              <Button
                variant="outline"
                onClick={() => setPage((p) => p + 1)}
                disabled={page >= pagination.totalPages}
              >
                Next
              </Button>
            </div>
          )}
        </>
      )}
    </section>
  );
};

export default OrdersList;
