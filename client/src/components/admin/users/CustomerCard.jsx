import { CustomerDetail } from "./CustomerDetail";
import { TableRow, TableCell } from "@/components/ui/table";

export const CustomerCard = ({ customer }) => {
  return (
    <TableRow key={customer.id}>
      <TableCell className="text-left">
        <img
          src={customer.avatar}
          alt={customer.fullname}
          className="h-12 w-12 object-cover rounded-full"
        />
      </TableCell>
      <TableCell className="text-left">
        <div>
          {customer.fullname.length > 20
            ? customer.fullname.slice(0, 20) + "..."
            : customer.fullname}
        </div>
      </TableCell>
      <TableCell className="text-left">{customer.email}</TableCell>
      <TableCell className="text-left">{customer.createdAt}</TableCell>
      <TableCell clasName="text-center">
        <CustomerDetail userId={customer.id} />
      </TableCell>
    </TableRow>
  );
};
