import {
  Table,
  TableRow,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { useAllCustomers } from "@/hooks/useDashboard";
import { Pagination } from "@/components/ui/pagination";
import { Card, CardContent } from "@/components/ui/card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { LoadingSearch } from "@/components/ui/LoadingSearch";
import { SectionTitle } from "@/components/header/SectionTitle";
import { useQueryParamsStore } from "@/store/useQueryParamsStore";
import { CustomerCard } from "@/components/admin/users/CustomerCard";
import { NoCustomerResult } from "@/components/admin/users/NoCustomerResult";

const CustomersList = () => {
  const { search, sort, page, limit, setSearch, setPage } =
    useQueryParamsStore();

  const { data, isLoading, isError, refetch } = useAllCustomers({
    search,
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
        <Input
          type="text"
          value={search}
          onChange={(e) => {
            setPage(1);
            setSearch(e.target.value);
          }}
          className="md:w-1/2 input"
          placeholder="Search for customer name or email"
        />
      </div>

      <Card className="border shadow-sm">
        <CardContent className="overflow-x-auto p-0">
          {isLoading ? (
            <LoadingSearch />
          ) : isError ? (
            <ErrorDialog onRetry={refetch} />
          ) : customers.length === 0 ? (
            <NoCustomerResult search={search} />
          ) : (
            <div className="hidden md:block w-full">
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead className="text-left">Avatar</TableHead>
                    <TableHead className="text-left">Fullname</TableHead>
                    <TableHead className="text-left">Email</TableHead>
                    <TableHead className="text-left">JoinedAt</TableHead>
                    <TableHead className="text-center">Detail</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody className="h-12">
                  {customers.map((customer) => (
                    <CustomerCard customer={customer} />
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

export default CustomersList;
