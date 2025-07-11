import { Badge } from "@/components/ui/Badge";
import { ProceedOrder } from "./ProceedOrder";
import { Button } from "@/components/ui/Button";
import { Card, CardContent } from "@/components/ui/Card";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { useLocation, useNavigate } from "react-router-dom";

export const OrderCard = ({ orders }) => {
  const navigate = useNavigate();
  const location = useLocation();
  const openModal = (id) => {
    navigate(`/admin/orders/${id}`, {
      state: { backgroundLocation: location },
    });
  };

  return (
    <div className="space-y-6">
      {orders.map((order) => (
        <Card key={order.id}>
          <CardContent className="p-5 space-y-4">
            {/* Header */}
            <div className="flex flex-col md:flex-row md:items-center md:justify-between w-full">
              <div className="space-y-1 text-start">
                <div className="text-xs text-muted-foreground">
                  Ordered on : {formatDateTime(order.createdAt)}
                </div>
                <div className="text-xs text-muted-foreground">
                  Invoice No : INV/{order.invoiceNumber}
                </div>
              </div>
              <Badge
                variant={
                  order.status === "success"
                    ? "success"
                    : order.status === "pending"
                    ? "outline"
                    : order.status === "waiting_payment"
                    ? "destructive"
                    : "ghost"
                }
                className="capitalize w-fit text-xs mt-2 md:mt-0"
              >
                {order.status}
              </Badge>
            </div>

            {/* Items */}
            <div className="border-t pt-4 flex gap-4 items-center w-full">
              <img
                src={order.items[0]?.image}
                alt={order.items[0]?.name}
                className="w-20 h-20 object-cover rounded border"
              />
              <div className="flex-1 text-start">
                <p className="font-medium text-sm text-foreground line-clamp-1">
                  {order.items[0]?.name}
                </p>
                <p className="text-muted-foreground text-sm">
                  {order.items.length} items x{" "}
                  {formatRupiah(order.total / order.items[0]?.quantity || 1)}
                </p>
              </div>
              <div className="text-right hidden md:block">
                <p className="text-sm text-muted-foreground">Total Amount</p>
                <p className="text-lg font-bold text-foreground">
                  {formatRupiah(order.total)}
                </p>
              </div>
            </div>

            {/* Mobile total */}
            <div className="md:hidden text-right w-full">
              <p className="text-sm text-muted-foreground">Total Amount</p>
              <p className="text-lg font-bold text-foreground">
                {formatRupiah(order.total)}
              </p>
            </div>

            {/* Actions */}
            <div className="pt-2 flex justify-end gap-3 w-full">
              {order.status === "pending" && <ProceedOrder order={order} />}
              {(order.status === "success" || order.status === "process") && (
                <Button
                  size="sm"
                  className="w-32"
                  variant="secondary"
                  onClick={() => openModal(order.id)}
                >
                  View Detail
                </Button>
              )}
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};
