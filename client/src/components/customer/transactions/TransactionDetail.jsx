import {
  Dialog,
  DialogContent,
  DialogTrigger,
  DialogTitle,
} from "@/components/ui/dialog";
import { Loader2 } from "lucide-react";
import { Link } from "react-router-dom";
import { CreateReview } from "./CreateReview";
import { ShipmentInfo } from "./ShipmentInfo";
import { Button } from "@/components/ui/button";
import { useOrderDetailQuery } from "@/hooks/useOrder";
import { formatDateTime, formatRupiah } from "@/lib/utils";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { TransactionDetailSkeleton } from "@/components/loading/TransactionDetailSkeleton";

export const TransactionDetail = ({ transaction }) => {
  const { data, isLoading } = useOrderDetailQuery(transaction.id);

  console.log("Transaction Detail Data:", data);
  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" className="w-32" size="sm">
          View Detail
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl p-6 space-y-2">
        {isLoading || !data ? (
          <TransactionDetailSkeleton />
        ) : (
          <>
            <DialogTitle className="text-xl font-semibold">
              Transaction Detail
            </DialogTitle>
            <div className="border flex justify-between p-2 rounded-md bg-muted">
              <div className="space-y-2">
                <p className="font-medium capitalize">
                  Order{" "}
                  {data.status === "pending"
                    ? data.status
                    : data.status === "process"
                    ? data.status
                    : "success"}
                </p>
                <p className="text-sm">
                  <span className="font-medium">Order No:</span>{" "}
                  <span className="text-primary font-medium">
                    {data.invoiceNumber || ""}
                  </span>
                </p>
                <p className="text-sm">
                  <span className="font-medium">Order Date:</span>{" "}
                  {formatDateTime(data.createdAt)}
                </p>
              </div>

              <Link to={`/invoice/${data.id}`} target="_blank">
                <Button className="w-28" size="sm">
                  Print Invoice
                </Button>
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
                      {item.quantity} x {formatRupiah(item.price)}
                    </p>
                  </div>

                  {data.status === "success" &&
                    (item.isReviewed ? (
                      <Link to={`/products/${item.slug}`}>
                        <Button variant="outline" size="sm">
                          Buy Again
                        </Button>
                      </Link>
                    ) : (
                      <CreateReview productId={item.id} />
                    ))}
                </div>
              ))}
            </div>
            {/* Shipping information  */}
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
                    <ShipmentInfo orderId={data.id} />
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
                <p>
                  <span className="inline-block w-48">Tax (10%)</span>:{" "}
                  {formatRupiah(data.tax)}
                </p>
                <p className="text-base font-semibold text-foreground">
                  <span className="inline-block font-bold w-48">
                    Grand total
                  </span>
                  : {formatRupiah(data.amountToPay)}
                </p>
              </div>

              <p className="text-xs text-muted-foreground">
                * see invoice for details.
              </p>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
