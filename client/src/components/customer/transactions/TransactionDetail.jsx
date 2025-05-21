import {
  Dialog,
  DialogContent,
  DialogTrigger,
  DialogTitle,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { formatRupiah } from "@/lib/utils";
import { useOrderDetailQuery } from "@/hooks/useOrder";
import { Loading } from "@/components/ui/Loading";
import { Link } from "react-router-dom";

const formatDateTime = (iso) => {
  const d = new Date(iso);
  return (
    d.toLocaleString("en-GB", {
      day: "2-digit",
      month: "long",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      timeZone: "Asia/Jakarta",
    }) + " WIB"
  );
};

const TransactionDetail = ({ transaction }) => {
  const { data, isLoading } = useOrderDetailQuery(transaction.id);

  console.log(data);

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" className="w-32" size="sm">
          View Detail
        </Button>
      </DialogTrigger>

      <DialogContent className="max-w-2xl p-6 space-y-6">
        {isLoading || !data ? (
          <Loading />
        ) : (
          <>
            <DialogTitle className="text-xl font-semibold">
              Transaction Detail
            </DialogTitle>

            {/* Main Info */}
            <div className="border flex justify-between p-4 rounded-md bg-muted">
              <div className="space-y-2">
                <p className="font-medium capitalize">
                  Order {data.status === "success" ? "Completed" : data.status}
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
                      {item.quantity} x {formatRupiah(item.price)}
                    </p>
                  </div>
                  <Link to={`/products/${item.slug}`}>
                    <Button variant="outline" size="sm">
                      Buy Again
                    </Button>
                  </Link>
                </div>
              ))}
            </div>

            {/* Shipping Info */}
            <div className="border p-4 rounded-md space-y-2">
              <h4 className="font-medium">Shipping Info</h4>
              <p className="text-sm">
                <span className="font-medium">Courier:</span> {data.courierName}
              </p>
              <p className="text-sm">
                <span className="font-medium">Address:</span> {data.address}
              </p>
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

              <p className="text-xs text-muted-foreground">
                * Transaction fees not included, see invoice for details.
              </p>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};

export { TransactionDetail };
