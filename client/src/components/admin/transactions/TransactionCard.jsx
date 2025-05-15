import { Badge } from "@/components/ui/badge";
import { formatDate, formatRupiah } from "@/lib/utils";
import { TableRow, TableCell } from "@/components/ui/table";

const TransactionCard = ({ transaction }) => {
  return (
    <TableRow key={transaction.id}>
      <TableCell className="text-left">
        <div>
          {transaction.fullname.length > 20
            ? transaction.fullname.slice(0, 20) + "..."
            : transaction.fullname}
        </div>
      </TableCell>

      <TableCell className="text-left">{transaction.email}</TableCell>

      <TableCell className="text-left">
        {transaction.invoiceNumber || ""}
      </TableCell>

      <TableCell className="text-left">{transaction.method || ""}</TableCell>

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
          ? formatDate(transaction.paidAt)
          : "-"}
      </TableCell>
    </TableRow>
  );
};

export default TransactionCard;
