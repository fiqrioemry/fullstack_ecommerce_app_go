import { useDebounce } from "@/hooks/useDebounce";
import { usePaymentsQuery } from "@/hooks/usePayment";
import { paymentStatusOptions } from "@/lib/constant";
import { Pagination } from "@/components/ui/Pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { SearchInput } from "@/components/ui/SearchInput";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { RecordNotFound } from "@/components/ui/RecordNotFound";
import { useTransactionStore } from "@/store/useTransactionStore";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const TransactionsList = () => {
  const { page, limit, q, sort, setPage, status, setQ, setSort, setStatus } =
    useTransactionStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = usePaymentsQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
    status,
  });

  const transactions = data?.data || [];
  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Transactions List"
        description="See all user transactions"
      />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by customer name or email"}
        />
        <SelectFilter
          value={status}
          onChange={setStatus}
          options={paymentStatusOptions}
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : transactions.length === 0 ? (
            <RecordNotFound title="No transactions found" q={q} />
          ) : (
            <TransactionCard
              sort={sort}
              setSort={setSort}
              transactions={transactions}
            />
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

export default TransactionsList;
