import { revenueRangeOptions } from "@/lib/constant";
import { usePaymentsQuery } from "@/hooks/usePayment";
import { useAllOrdersQuery } from "@/hooks/useOrder";
import { useRevenueStats } from "@/hooks/useDashboard";
import { Card, CardContent } from "@/components/ui/Card";
import { useProductStore } from "@/store/useProductStore";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useDashboardSummary } from "@/hooks/useDashboard";
import { SelectFilter } from "@/components/ui/SelectFilter";
import { OrderCard } from "@/components/admin/orders/OrderCard";
import { SummaryCard } from "@/components/admin/dashboard/SummaryCard";
import { RevenueChart } from "@/components/admin/dashboard/RevenueChart";
import { DashboardSkeleton } from "@/components/loading/DashboardSkeleton";
import { TransactionCard } from "@/components/admin/transactions/TransactionCard";

const Dashboard = () => {
  const { range, limit, setRange } = useProductStore();
  const { data: transactions } = usePaymentsQuery({
    limit,
  });
  const { data: revenue } = useRevenueStats(range);
  const { data: orders } = useAllOrdersQuery({ limit, status: "pending" });
  const { data: summary, isLoading, isError, refetch } = useDashboardSummary();

  if (isLoading) return <DashboardSkeleton />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  console.log(revenue);
  return (
    <div className="p-6 space-y-6">
      {/* Header */}
      <div>
        <h3 className="mb-4">Statistic Summary</h3>
        <div className="grid grid-cols-2 lg:grid-cols-4 gap-4">
          <SummaryCard
            title="Total Customers"
            value={summary?.data?.totalCustomers}
          />
          <SummaryCard
            title="Total Products"
            value={summary?.data?.totalProducts}
          />
          <SummaryCard title="Total Orders" value={summary?.data.totalOrders} />
          <SummaryCard
            title="Total Revenue"
            value={`Rp${summary?.data.totalRevenue.toLocaleString("id-ID")}`}
          />
        </div>
      </div>

      <div>
        <h3 className="mb-4">Transaction Volume</h3>
        <div className="bg-background rounded-xl p-6">
          <div className="flex items-center justify-between mb-4">
            <h4>Revenue</h4>

            <SelectFilter
              value={range}
              onChange={setRange}
              options={revenueRangeOptions}
            />
          </div>
          <RevenueChart data={revenue?.revenueSeries} range={revenue?.range} />
        </div>
      </div>

      {orders?.data?.length > 0 && (
        <div>
          <h3 className="mb-4">Pending Order</h3>
          <OrderCard orders={orders?.data} />
        </div>
      )}

      <div>
        <h3 className="mb-4">Recent transaction</h3>
        <Card>
          <CardContent className="p-0">
            <TransactionCard transactions={transactions?.data} />
          </CardContent>
        </Card>
      </div>
    </div>
  );
};
export default Dashboard;
