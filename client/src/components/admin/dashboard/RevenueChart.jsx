import {
  Line,
  XAxis,
  YAxis,
  Tooltip,
  LineChart,
  CartesianGrid,
  ResponsiveContainer,
} from "recharts";
import { formatRupiah, formatDate } from "@/lib/utils";

const RevenueChart = ({ data, range }) => {
  const formattedData = data?.map((item) => {
    const date = new Date(item.date);
    let label = item.date;

    if (range === "daily") {
      label = date.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "short",
      });
    } else if (range === "monthly") {
      label = date.toLocaleDateString("id-ID", {
        month: "short",
        year: "numeric",
      });
    } else if (range === "yearly") {
      label = date.getFullYear().toString();
    }

    return { ...item, label };
  });

  const formatAxisNumber = (value) => {
    if (value >= 1_000_000) return `${value / 1_000_000}M`;
    if (value >= 1_000) return `${value / 1_000}k`;
    return value;
  };

  return (
    <ResponsiveContainer width="100%" height={300}>
      <LineChart data={formattedData}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="label" />
        <YAxis tickFormatter={formatAxisNumber} />
        <Tooltip
          formatter={(value) => `${formatRupiah(value)}`}
          labelFormatter={(label) => `Tanggal: ${formatDate(label)}`}
        />
        <Line
          type="monotone"
          dataKey="total"
          stroke="#2563eb"
          strokeWidth={2}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export { RevenueChart };
