import {
  Select,
  SelectItem,
  SelectValue,
  SelectTrigger,
  SelectContent,
} from "@/components/ui/Select";
import {
  Table,
  TableRow,
  TableCell,
  TableBody,
  TableHead,
  TableHeader,
} from "@/components/ui/Table";
import { useState } from "react";
import { Badge } from "@/components/ui/Badge";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { useRevenueStats } from "@/hooks/useDashboard";
import { Card, CardContent } from "@/components/ui/Card";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useAdminPaymentsQuery } from "@/hooks/usePayment";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useDashboardSummary } from "@/hooks/useDashboard";
import { OrderCard } from "@/components/admin/orders/OrderCard";
import { SummaryCard } from "@/components/admin/dashboard/SummaryCard";
import { RevenueChart } from "@/components/admin/dashboard/RevenueChart";
import { DashboardSkeleton } from "@/components/loading/DashboardSkeleton";

const Dashboard = () => {
  const [range, setRange] = useState("daily");
  const [status, setStatus] = useState("pending");
  const { data: revenue } = useRevenueStats(range);
  const { data: response } = useAdminPaymentsQuery();
  const { data: orders } = useAllOrdersQuery({ status });
  const { data: data, isLoading, isError, refetch } = useDashboardSummary();

  if (isLoading) return <DashboardSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const summary = data.data || [];

  const transactions = response.data || [];

  return (
    <div className="p-6 space-y-6">
      {/* Header */}
      <div>
        <h3 className="mb-4">Statistic Summary</h3>
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <SummaryCard
            title="Total Customers"
            value={summary?.totalCustomers}
          />
          <SummaryCard title="Total Products" value={summary?.totalProducts} />
          <SummaryCard title="Total Orders" value={summary?.totalOrders} />
          <SummaryCard
            title="Total Revenue"
            value={`Rp${summary.totalRevenue.toLocaleString("id-ID")}`}
          />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Transaction Volume</h3>
        <div className="bg-background rounded-xl shadow p-6">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold">Revenue</h2>
            <Select value={range} onValueChange={setRange}>
              <SelectTrigger className="w-[120px]">
                <SelectValue placeholder="Range" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="daily">Daily</SelectItem>
                <SelectItem value="monthly">Monthly</SelectItem>
                <SelectItem value="yearly">Yearly</SelectItem>
              </SelectContent>
            </Select>
          </div>

          <RevenueChart data={revenue?.revenueSeries} range={revenue?.range} />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Pending Order</h3>
        <OrderCard orders={orders?.data} />
      </div>

      <div>
        <h3 className="mb-4">Recent transaction</h3>
        <Card className="border shadow-sm">
          <CardContent className="overflow-x-auto p-0">
            {/* dekstop view */}
            <div className="bg-background hidden md:block w-full">
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
                  {transactions.map((transaction) => (
                    <TableRow key={transaction.id}>
                      <TableCell className="text-left">
                        <div>
                          {transaction.fullname.length > 20
                            ? transaction.fullname.slice(0, 20) + "..."
                            : transaction.fullname}
                        </div>
                      </TableCell>

                      <TableCell className="text-left">
                        {transaction.email}
                      </TableCell>

                      <TableCell className="text-left">
                        {transaction.invoiceNumber || ""}
                      </TableCell>

                      <TableCell className="text-left">
                        {transaction.method || ""}
                      </TableCell>

                      <TableCell className="text-left">
                        {transaction.status === "success" ? (
                          <Badge> success</Badge>
                        ) : "pending" ? (
                          <Badge variant="secondary">pending</Badge>
                        ) : (
                          <Badge variant="destructive">failed</Badge>
                        )}
                      </TableCell>

                      <TableCell className="text-left">
                        {formatRupiah(transaction.total)}
                      </TableCell>
                      <TableCell className="text-left">
                        {transaction.status === "success"
                          ? transaction.paidAt
                          : "-"}
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </div>

            {/* mobile view */}
            <div className="md:hidden space-y-4 p-4  w-full">
              {transactions.map((tx) => (
                <div
                  key={tx.id}
                  className="border rounded-lg p-4 shadow-sm space-y-2"
                >
                  <div>
                    <h3 className="text-base font-semibold">{tx.fullname}</h3>
                    <p className="text-sm text-muted-foreground">
                      {tx.userEmail}
                    </p>
                  </div>
                  <div className="text-sm space-y-4">
                    <div className="grid grid-cols-2 gap-4">
                      <p>
                        <strong>{tx.invoiceNumber}</strong>
                      </p>
                      <p>{formatRupiah(tx.total)}</p>
                      <p>{tx.method?.toUpperCase() || "-"}</p>
                      <p>
                        <Badge
                          variant={
                            tx.status === "success"
                              ? "default"
                              : tx.status === "failed"
                              ? "destructive"
                              : "secondary"
                          }
                        >
                          {tx.status}
                        </Badge>
                      </p>
                    </div>
                    <div className="space-x-4">
                      {" "}
                      <span className="text-muted-foreground">Paid At</span>
                      <span className="text-right whitespace-nowrap">
                        {tx.status === "success"
                          ? formatDateTime(tx.paidAt)
                          : "-"}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Dashboard;
