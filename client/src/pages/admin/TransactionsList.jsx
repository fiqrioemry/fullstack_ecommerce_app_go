import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
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
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { Card, CardContent } from "@/components/ui/card";
import { Pagination } from "@/components/ui/pagination";
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
    <section className="section space-y-6">
      <div className="space-y-1 text-center">
        <h2 className="text-2xl font-bold">Transactions List</h2>
        <p className="text-muted-foreground text-sm">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi,
          blanditiis.
        </p>
      </div>

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
