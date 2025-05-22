import {
  Select,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/select";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { NoTransaction } from "@/components/customer/transactions/NoTransaction";
import { TransactionCard } from "@/components/customer/transactions/TransactionCard";

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

  const { data, isLoading, isError, error } = useAllOrdersQuery({
    search,
    page,
    limit,
    sort,
    status,
  });

  const transactions = data?.data || [];

  console.log(data);

  const pagination = data?.pagination || null;

  return (
    <section className="max-w-5xl mx-auto px-4 py-10 space-y-6">
      {/* 🔍 Filter Bar */}
      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          placeholder="Search in transactions"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2"
        />

        <div className="flex gap-3">
          <Select
            value={status}
            onValueChange={(val) => {
              setPage(1);
              setStatus(val);
            }}
          >
            <SelectTrigger className="w-48 h-11 bg-background">
              <SelectValue placeholder="Filter by Status" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Status</SelectLabel>
                <SelectItem value="all">All Status</SelectItem>
                <SelectItem value="success">Success</SelectItem>
                <SelectItem value="pending">Pending</SelectItem>
                <SelectItem value="failed">Failed</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>

          <Select
            value={sort}
            onValueChange={(val) => {
              setPage(1);
              setSort(val);
            }}
          >
            <SelectTrigger className="w-60 h-11 bg-background">
              <SelectValue placeholder="Sort By" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Sort By</SelectLabel>
                <SelectItem value="created_at desc">Newest</SelectItem>
                <SelectItem value="created_at asc">Oldest</SelectItem>
                <SelectItem value="product_name asc">
                  Product Name A-Z
                </SelectItem>
                <SelectItem value="product_name desc">
                  Product Name Z-A
                </SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
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
