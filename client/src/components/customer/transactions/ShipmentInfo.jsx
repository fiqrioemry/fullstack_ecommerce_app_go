import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogTrigger,
} from "@/components/ui/Dialog";
import { formatDate } from "@/lib/utils";
import { Button } from "@/components/ui/Button";
import { useShipmentQuery } from "@/hooks/useOrder";
import { ShipmentInfoSkeleton } from "@/components/loading/ShipmentInfoSkeleton";

export const ShipmentInfo = ({ orderId }) => {
  const { data, isLoading } = useShipmentQuery(orderId);

  const statusColor = {
    returned: "bg-red-100 text-red-700",
    delivered: "bg-green-100 text-green-700",
    shipped: "bg-yellow-100 text-yellow-700",
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" className="w-32" size="sm">
          View Detail
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl p-6 space-y-4">
        {isLoading || !data ? (
          <ShipmentInfoSkeleton />
        ) : (
          <>
            <DialogTitle className="text-xl font-semibold">
              Shipment Detail
            </DialogTitle>

            <div className="space-y-3">
              <div className="flex items-center gap-2">
                <span className="text-sm font-medium">Status:</span>
                <span
                  className={`px-3 py-1 rounded-full text-sm font-semibold capitalize ${
                    statusColor[data.status]
                  }`}
                >
                  {data.status}
                </span>
              </div>

              <div className="text-sm space-y-1">
                <p>
                  <span className="font-medium">Tracking Code:</span>{" "}
                  {data.trackingCode}
                </p>
                <p>
                  <span className="font-medium">Shipped At:</span>{" "}
                  {formatDate(data.shippedAt)}
                </p>
                {data.deliveredAt && (
                  <p>
                    <span className="font-medium">Delivered At:</span>{" "}
                    {formatDate(data.deliveredAt)}
                  </p>
                )}
                {data.notes && (
                  <p>
                    <span className="font-medium">Notes:</span> {data.notes}
                  </p>
                )}
              </div>

              <div className="mt-4">
                <div className="relative h-2 bg-gray-200 rounded-full">
                  <div
                    className={`absolute top-0 left-0 h-2 rounded-full transition-all duration-300 ${
                      data.status === "shipped"
                        ? "w-1/2 bg-yellow-400"
                        : data.status === "delivered"
                        ? "w-full bg-green-500"
                        : data.status === "returned"
                        ? "w-full bg-red-500"
                        : "w-1/3 bg-gray-400"
                    }`}
                  />
                </div>
                <div className="flex justify-between text-xs text-gray-500 mt-1">
                  <span>Shipped</span>
                  <span>Delivered</span>
                </div>
              </div>
            </div>
          </>
        )}
      </DialogContent>
    </Dialog>
  );
};
