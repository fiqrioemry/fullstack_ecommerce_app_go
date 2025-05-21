import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
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
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";
import { NoTransactionResult } from "@/components/admin/transactions/NoTransactionResult";

const TransactionsList = () => {
  const { search, status, sort, page, limit, setSearch, setPage, setStatus } =
    useQueryParamsStore();

  const { data, isLoading, isError } = useAdminPaymentsQuery(
    search,
    page,
    limit,
    sort,
    status
  );

  const transactions = data?.data || [];

  const pagination = data?.pagination;

  return (
    <section className="section px-4 py-8 space-y-6">
      <SectionTitle
        title="Transactions List"
        description="See all user transactions"
      />

      <div className="flex flex-col md:flex-row items-start md:items-center justify-between gap-4">
        <Input
          type="text"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2 input"
          placeholder="Search product name or description"
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
              <SelectItem value="failed">failed</SelectItem>
            </SelectGroup>
          </SelectContent>
        </Select>
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : transactions.length === 0 ? (
            <NoTransactionResult search={search} />
          ) : (
            <div className="hidden md:block w-full">
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead className="text-left">Name</TableHead>
                    <TableHead className="text-left">Email</TableHead>
                    <TableHead className="text-left">Invoice No.</TableHead>
                    <TableHead className="text-left">Method</TableHead>
                    <TableHead className="text-left">Status</TableHead>
                    <TableHead className="text-left">Total</TableHead>
                    <TableHead className="text-left">Paid At</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody className="h-12">
                  {transactions.map((tx) => (
                    <TransactionCard key={tx.id} transaction={tx} />
                  ))}
                </TableBody>
              </Table>
            </div>
          )}

          {pagination && (
            <Pagination
              page={pagination.page}
              limit={pagination.limit}
              total={pagination.totalRows}
              onPageChange={(p) => setPage(p)}
            />
          )}
        </CardContent>
      </Card>
    </section>
  );
};

export default TransactionsList;
