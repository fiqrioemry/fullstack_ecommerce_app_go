import { useDebounce } from "@/hooks/useDebounce";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { paymentStatusOptions } from "@/lib/constant";
import { useOrderStore } from "@/store/useOrderStore";
import { Pagination } from "@/components/ui/pagination";
import { SearchInput } from "@/components/ui/SearchInput";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { NoTransaction } from "@/components/customer/transactions/NoTransaction";
import { TransactionCard } from "@/components/customer/transactions/TransactionCard";

const UserTransactions = () => {
  const { q, status, sort, page, limit, setQ, setPage, setStatus } =
    useOrderStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useAllOrdersQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
    status,
  });

  const transactions = data?.data || [];

  const pagination = data?.pagination || null;

  return (
    <section className="max-w-5xl mx-auto px-4 py-10 space-y-6">
      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search products in your transactions"}
        />

        <SelectFilter
          value={status}
          onChange={setStatus}
          options={paymentStatusOptions}
        />
      </div>

      {/* 🧾 Content */}
      {isLoading ? (
        <LoadingSearch className="mt-10" />
      ) : isError ? (
        <ErrorDialog onRetry={refetch} />
      ) : transactions.length === 0 ? (
        <NoTransaction />
      ) : (
        <TransactionCard transactions={transactions} />
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

export default UserTransactions;
