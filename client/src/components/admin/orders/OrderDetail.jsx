import {
  Dialog,
  DialogTitle,
  DialogHeader,
  DialogDescription,
  DialogContent,
} from "@/components/ui/Dialog";
import { Button } from "@/components/ui/Button";
import { useOrderDetailQuery } from "@/hooks/useOrder";
import { formatRupiah, formatDateTime } from "@/lib/utils";
import { ShipmentConfirmation } from "./ShipmentConfirmation";
import { Link, useNavigate, useParams } from "react-router-dom";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/Alert";
import { TransactionDetailSkeleton } from "@/components/loading/TransactionDetailSkeleton";
import { Loader2 } from "lucide-react";

export const OrderDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { data, isLoading } = useOrderDetailQuery(id);

  return (
    <Dialog open={true} onOpenChange={() => navigate(-1)}>
      <DialogContent className="max-w-2xl p-6 space-y-6">
        <DialogHeader>
          <DialogTitle className="text-xl font-semibold">
            Order Detail
          </DialogTitle>
          <DialogDescription>
            Detailed information about order
          </DialogDescription>
        </DialogHeader>
        {isLoading || !data ? (
          <TransactionDetailSkeleton />
        ) : (
          <>
            {/* Main Info */}
            <div className="border flex justify-between p-4 rounded-md bg-muted">
              <div className="space-y-2">
                <p className="font-medium">
                  Order Status:{" "}
                  {data.status === "pending"
                    ? "Pending"
                    : data.status === "process"
                    ? "Processing"
                    : "Completed"}
                </p>
                <p className="text-sm">
                  <span className="font-medium">Order No:</span>{" "}
                  <span className="text-primary font-medium">
                    {data.invoiceNumber || data.id.slice(0, 8).toUpperCase()}
                  </span>
                </p>
                <p className="text-sm">
                  <span className="font-medium">Order Date:</span>{" "}
                  {formatDateTime(data.createdAt)}
                </p>
              </div>

              <Link to={`/invoice/${data.id}`} target="_blank">
                <Button size="sm">Print Invoice</Button>
              </Link>
            </div>

            {/* Product Items */}
            <div className="border rounded-md">
              {data.items.map((item) => (
                <div
                  key={item.id}
                  className="flex gap-4 items-center p-4 border-b last:border-b-0"
                >
                  <img
                    src={item.image}
                    alt={item.name}
                    className="w-16 h-16 object-cover border rounded"
                  />
                  <div className="flex-1">
                    <p className="font-semibold">{item.name}</p>
                    <p className="text-sm text-muted-foreground">
                      Quantity : {item.quantity} x {formatRupiah(item.price)}
                    </p>
                  </div>
                </div>
              ))}
            </div>

            {/* Shipping Info */}
            <div className="border p-4 rounded-md space-y-2 bg-muted/50">
              <h4 className="font-semibold text-lg">Shipping Info</h4>

              <div className="text-sm space-y-1">
                <p>
                  <span className="font-medium">Courier:</span>{" "}
                  {data.courierName}
                </p>
                <p>
                  <span className="font-medium">Address:</span>{" "}
                  {data.shippingAddress}
                </p>
              </div>

              {data.status === "pending" ? (
                <div className="mt-3 flex items-center gap-2 text-yellow-600">
                  <Loader2 className="w-4 h-4 animate-spin" />
                  <span className="text-sm">
                    Your order is being processed for shipment...
                  </span>
                </div>
              ) : (
                <div className="mt-3">
                  <Alert variant="success">
                    <AlertTitle className="text-green-600">
                      Shipment Created
                    </AlertTitle>
                    <AlertDescription className="text-sm">
                      Your package is on the way. You can view the tracking
                      information below.
                    </AlertDescription>
                  </Alert>

                  <div className="mt-2">
                    <ShipmentConfirmation data={data} />
                  </div>
                </div>
              )}
            </div>

            {/* Payment Summary */}
            <div className="border p-4 rounded-md space-y-2">
              <h4 className="font-medium">Payment Summary</h4>
              <div className="text-sm text-muted-foreground space-y-1">
                <p>
                  <span className="inline-block w-48">Subtotal</span>:{" "}
                  {formatRupiah(data.total)}
                </p>
                <p>
                  <span className="inline-block w-48">Shipping Cost</span>:{" "}
                  {formatRupiah(data.shippingCost)}
                </p>
                {data.voucherDiscount > 0 && (
                  <p>
                    <span className="inline-block w-48">Voucher</span>: -{" "}
                    {formatRupiah(data.voucherDiscount)}
                  </p>
                )}
                <p className="text-base font-semibold text-foreground">
                  <span className="inline-block font-bold w-48">
                    Grand total
                  </span>
                  :{" "}
                  {formatRupiah(
                    data.total + data.shippingCost - data.voucherDiscount
                  )}
                </p>
              </div>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
