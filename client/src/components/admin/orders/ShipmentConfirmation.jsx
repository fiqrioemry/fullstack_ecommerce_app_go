import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogTrigger,
} from "@/components/ui/dialog";
import { useRef } from "react";
import html2pdf from "html2pdf.js";
import Barcode from "react-barcode";
import { Button } from "@/components/ui/button";
import { useOrderMutation } from "@/hooks/useOrder";
import { formatRupiah, formatDateTime } from "@/lib/utils";

export const ShipmentConfirmation = ({ data }) => {
  const labelRef = useRef();
  const handleDownload = () => {
    const element = labelRef.current;
    const opt = {
      margin: 0.3,
      filename: `invoice-${data.invoiceNumber}.pdf`,
      image: { type: "jpeg", quality: 0.98 },
      html2canvas: { scale: 2 },
      jsPDF: { unit: "in", format: "a4", orientation: "portrait" },
    };
    html2pdf().from(element).set(opt).save();
  };
  const { updateShipment } = useOrderMutation();

  const handleCreateShipment = () => {
    updateShipment.mutateAsync(data.id);
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="outline" className="w-32" size="sm">
          Get Shipment Label
        </Button>
      </DialogTrigger>
      <DialogContent className="max-w-2xl p-6">
        <>
          <DialogTitle className="text-xl font-semibold">
            Shipment Detail
          </DialogTitle>

          <div className="flex gap-2 justify-end">
            <div className="flex justify-end print:hidden">
              <button
                onClick={handleCreateShipment}
                className="bg-primary text-white px-4 py-2 rounded text-sm"
              >
                Confirm Delivery
              </button>
            </div>

            <div className="flex justify-end print:hidden">
              <button
                onClick={handleDownload}
                className="bg-primary text-white px-4 py-2 rounded text-sm"
              >
                Download Label PDF
              </button>
            </div>
          </div>
          <div
            ref={labelRef}
            className="p-4 print:p-2 print:max-w-full bg-white border shadow-md print:shadow-none"
          >
            {/* Label Title */}
            <div className="text-center mb-4">
              <h2 className="text-lg font-bold">Shipping Label</h2>
              <p className="text-sm text-muted-foreground">
                {formatDateTime(data.createdAt)}
              </p>
            </div>

            {/* Barcode & Invoice */}
            <div className="mb-4 flex flex-col items-center justify-center">
              <Barcode
                value={data.invoiceNumber || data.id}
                height={60}
                width={1.5}
              />
              <p className="text-xs mt-1">Invoice No: {data.invoiceNumber}</p>
            </div>

            <div className="text-sm space-y-2">
              <div className="flex items-center gap-2">
                <span className="font-medium inline-block w-48">Courier </span>{" "}
                :<p> {data.courierName}</p>
              </div>

              <div className="flex items-center gap-2">
                <span className="font-medium inline-block w-48">
                  Total Payment{" "}
                </span>{" "}
                :<p> {formatRupiah(data.amountToPay)}</p>
              </div>

              <div className="flex items-center gap-2">
                <span className="font-medium inline-block w-48">
                  Recipient{" "}
                </span>{" "}
                :<p> {data.recipientName}</p>
              </div>

              <div>
                <div className="flex items-center gap-2">
                  <span className="font-medium inline-block w-48">
                    Shipping To
                  </span>{" "}
                  :
                </div>
                <p>{data.shippingAddress}</p>
              </div>

              <div className="flex items-center gap-2">
                <span className="font-medium inline-block w-48">
                  Tracking No.{" "}
                </span>{" "}
                :<p> {data.recipientName}</p>
              </div>
            </div>
          </div>
        </>
      </DialogContent>
    </Dialog>
  );
};
