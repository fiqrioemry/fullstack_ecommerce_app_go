import Barcode from "react-barcode";
import { useParams } from "react-router-dom";
import { formatDateTime } from "@/lib/utils";
import { Loading } from "@/components/ui/Loading";
import { useOrderDetailQuery } from "@/hooks/useOrder";

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
      <div className="mb-4 text-center">
        <Barcode
          value={data.invoiceNumber || data.id}
          height={60}
          width={1.5}
        />
        <p className="text-xs mt-1">Invoice No: {data.invoiceNumber}</p>
      </div>

      {/* Destination Info */}
      <div className="mb-4 space-y-1 text-sm leading-relaxed">
        <h4 className="font-semibold">Ship To:</h4>
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
          Courier: <strong>{data.courierName}</strong>
        </p>
        <p>
          Total Payment: <strong>{formatRupiah(data.amountToPay)}</strong>
        </p>
        <p>
          Status: <strong className="capitalize">{data.status}</strong>
        </p>
      </div>
    </div>
  );
};

export default ShipmentPage;
