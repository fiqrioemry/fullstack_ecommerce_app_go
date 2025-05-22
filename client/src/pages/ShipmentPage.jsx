import Barcode from "react-barcode";
import { useParams } from "react-router-dom";
import { Loading } from "@/components/ui/Loading";
import { useOrderDetailQuery } from "@/hooks/useOrder";
import { formatDateTime, formatRupiah } from "@/lib/utils";

const ShipmentPage = () => {
  const { orderId } = useParams();

  const { data, isLoading } = useOrderDetailQuery(orderId);

  if (isLoading || !data) return <Loading />;

  return (
    <div className="max-w-xl mx-auto p-6 print:p-2 print:max-w-full bg-white border shadow-md print:shadow-none">
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

      {/* Destination Info */}
      <div className="mb-4 space-y-1 text-sm leading-relaxed">
        <h4 className="font-semibold">Shipping To:</h4>
        <p>{data.address}</p>
        <p>
          {data.subdistrict}, {data.district}, {data.city}
        </p>
        <p>
          {data.province}, {data.postalCode}
        </p>
      </div>

      {/* Courier Info */}
      <div className="text-sm space-y-1">
        <p>
          <span className="inline-block w-48">Courier </span> :
          <strong> {data.courierName}</strong>
        </p>
        <p>
          <span className="inline-block w-48">Total Payment </span> :
          <strong> {formatRupiah(data.amountToPay)}</strong>
        </p>
        <p>
          <span className="inline-block w-48">Tracking No.</span> :
        </p>
      </div>
    </div>
  );
};

export default ShipmentPage;
