import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import TransactionCard from "@/components/admin/transactions/TransactionCard";
import NoTransactionResult from "@/components/admin/transactions/NoTransactionResult";

const TransactionsList = () => {
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

  const { data, isLoading, isError, error } = useAdminPaymentsQuery(
    search,
    page,
    limit,
    sort,
    status
  );

  const payments = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <div className="text-center space-y-1 mb-6">
        <h2 className="text-2xl font-bold">Payment Transactions</h2>
        <p className="text-sm text-muted-foreground">
          All user payment history
        </p>
      </div>

      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row justify-between gap-4">
        <Input
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          placeholder="Search by email or user ID"
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
          </select>
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
            <option value="total desc">Total High → Low</option>
            <option value="total asc">Total Low → High</option>
            <option value="status asc">Status A-Z</option>
            <option value="status desc">Status Z-A</option>
          </select>
        </div>
      </div>

      {/* 🔄 Content */}
      {isLoading ? (
        <LoadingSearch className="mt-10" />
      ) : isError ? (
        <ErrorDialog
          message={error?.message || "Failed to load transactions"}
        />
      ) : payments.length === 0 ? (
        <NoTransactionResult search={search} />
      ) : (
        <>
          <div className="space-y-4">
            {payments.map((transaction) => (
              <TransactionCard key={transaction.id} transaction={transaction} />
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

export default TransactionsList;
