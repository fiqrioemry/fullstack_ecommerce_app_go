import { useDebounce } from "@/hooks/useDebounce";
import { useUserStore } from "@/store/useUserStore";
import { useCustomersQuery } from "@/hooks/useDashboard";
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { SearchInput } from "@/components/ui/SearchInput";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { RecordNotFound } from "@/components/ui/RecordNotFound";
import { SectionTitle } from "@/components/header/SectionTitle";
import { CustomerCard } from "@/components/admin/users/CustomerCard";

const CustomersList = () => {
  const { page, limit, q, sort, setPage, setQ, setSort } = useUserStore();

  const debouncedQ = useDebounce(q, 500);
  const { data, isLoading, isError, refetch } = useCustomersQuery({
    q: debouncedQ,
    page,
    limit,
    sort,
  });

  const customers = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Customers List"
        description="See all active customers and their activities."
      />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <SearchInput
          q={q}
          setQ={setQ}
          setPage={setPage}
          placeholder={"search by product name or description"}
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : customers.length === 0 ? (
            <RecordNotFound title="No Record found for user" q={q} />
          ) : (
            <CustomerCard users={customers} sort={sort} setSort={setSort} />
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

export default CustomersList;
