import { useRef } from "react";
import html2pdf from "html2pdf.js";
import { useParams } from "react-router-dom";
import { useOrderDetailQuery } from "@/hooks/useOrder";
import { formatRupiah } from "@/lib/utils";
import { Loading } from "@/components/ui/Loading";
import { DollarSign, Pencil } from "lucide-react";

const formatDate = (iso) => {
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

const InvoicePage = () => {
  const { orderId } = useParams();
  const { data, isLoading } = useOrderDetailQuery(orderId);
  const invoiceRef = useRef();

  const handleDownload = () => {
    const element = invoiceRef.current;
    const opt = {
      margin: 0.3,
      filename: `invoice-${orderId}.pdf`,
      image: { type: "jpeg", quality: 0.98 },
      html2canvas: { scale: 2 },
      jsPDF: { unit: "in", format: "a4", orientation: "portrait" },
    };
    html2pdf().from(element).set(opt).save();
  };

  if (isLoading || !data) return <Loading />;

  return (
    <div className="max-w-3xl mx-auto px-8 py-10 print:p-0 print:max-w-full">
      {/* Download Button */}
      <div className="flex justify-end mb-4 print:hidden">
        <button
          onClick={handleDownload}
          className="bg-primary text-white px-4 py-2 rounded text-sm"
        >
          Download Invoice PDF
        </button>
      </div>

      {/* Invoice Content */}
      <div
        ref={invoiceRef}
        className="relative bg-white p-10 shadow-md print:shadow-none"
      >
        {/* Watermark */}
        {data.status === "success" && (
          <div className="absolute flex items-center gap-4 border border-red-500 border-10 p-10 text-[72px] font-bold text-red-500 opacity-20 rotate-[-30deg] top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 pointer-events-none select-none">
            <DollarSign className="h-20 w-20" /> <span>PAID</span>
          </div>
        )}

        <h2 className="text-2xl font-bold text-center mb-6">INVOICE</h2>

        {/* Header Info */}
        <div className="mb-6 text-sm space-y-1">
          <p>
            Invoice No.:{" "}
            <span className="font-medium">
              {data.invoiceNumber || data.id.slice(0, 8).toUpperCase()}
            </span>
          </p>
          <p>Transaction Date: {formatDate(data.createdAt)}</p>
          <p>Status: {data.status}</p>
        </div>

        {/* Product Table */}
        <div className="border rounded mb-6">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b bg-muted">
                <th className="p-2 text-left">Product</th>
                <th className="p-2 text-center">Qty</th>
                <th className="p-2 text-right">Price</th>
                <th className="p-2 text-right">Subtotal</th>
              </tr>
            </thead>
            <tbody>
              {data.items.map((item) => (
                <tr key={item.id} className="border-b last:border-b-0">
                  <td className="p-2">{item.name}</td>
                  <td className="p-2 text-center">{item.quantity}</td>
                  <td className="p-2 text-right">{formatRupiah(item.price)}</td>
                  <td className="p-2 text-right">
                    {formatRupiah(item.subtotal)}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* Payment Summary - Structured Layout */}
        <div className="text-sm mb-6">
          <div className="grid gap-2 text-right w-full md:w-[60%] ml-auto">
            <div className="flex justify-between items-center">
              <span className="text-muted-foreground">Subtotal</span>
              <span>{formatRupiah(data.total)}</span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-muted-foreground">Shipping Cost</span>
              <span>{formatRupiah(data.shippingCost)}</span>
            </div>

            <hr className="my-1 border-muted" />

            <div className="flex justify-between items-center font-medium">
              <span>Total Shopping</span>
              <span>{formatRupiah(data.total + data.shippingCost)}</span>
            </div>

            <div className="flex justify-between items-center">
              <span className="text-muted-foreground">App Fee</span>
              <span>{formatRupiah(2000)}</span>
            </div>

            <hr className="my-1 border-muted" />

            <div className="flex justify-between items-center font-bold text-lg">
              <span>Total Payment</span>
              <span>{formatRupiah(data.amountToPay + 2000)}</span>
            </div>
          </div>
        </div>

        {/* Shipping Info */}
        <div className="text-sm">
          <h4 className="font-semibold mb-1">Shipping To:</h4>
          <p>{data.address}</p>
          <p>Courier: {data.courierName}</p>
        </div>
      </div>
    </div>
  );
};

export default InvoicePage;
