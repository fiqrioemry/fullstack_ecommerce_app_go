import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { NoTransaction } from "@/components/customer/transactions/NoTransaction";
import { TransactionCard } from "@/components/customer/transactions/TransactionCard";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";

const UserTransactions = () => {
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

  const transactions = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="max-w-5xl mx-auto px-4 py-10 space-y-6">
      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          placeholder="Search product in transaction"
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
            value={status}
            onChange={(e) => {
              setPage(1);
              setStatus(e.target.value);
            }}
          >
            <option value="">All status</option>
            <option value="success">Success</option>
            <option value="pending">Pending</option>
            <option value="failed">Failed</option>
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

      {/* 🧾 Content */}
      {isLoading ? (
        <LoadingSearch className="mt-10" />
      ) : isError ? (
        <ErrorDialog message={error?.message || "Something went wrong"} />
      ) : transactions.length === 0 ? (
        <NoTransaction />
      ) : (
        <>
          <TransactionCard transactions={transactions} />

          {/* 📄 Pagination */}
          {pagination && (
            <div className="flex items-center justify-between pt-6">
              <Button
                variant="outline"
                onClick={() => setPage((p) => Math.max(p - 1, 1))}
                disabled={page === 1}
              >
                Sebelumnya
              </Button>
              <p className="text-sm text-muted-foreground">
                Page {pagination.page} of {pagination.totalPages}
              </p>
              <Button
                variant="outline"
                onClick={() => setPage((p) => p + 1)}
                disabled={page >= pagination.totalPages}
              >
                Selanjutnya
              </Button>
            </div>
          )}
        </>
      )}
    </section>
  );
};

export default UserTransactions;
