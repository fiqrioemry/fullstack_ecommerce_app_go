// Checkout.jsx (refactored with Zustand)
import { toast } from "sonner";
import { useEffect } from "react";
import { formatRupiah } from "@/lib/utils";
import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router-dom";
import { useCartQuery } from "@/hooks/useCart";
import { Button } from "@/components/ui/button";
import { Loading } from "@/components/ui/Loading";
import { useOrderMutation } from "@/hooks/useOrder";
import { useVoucherMutation } from "@/hooks/useVouchers";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { useUserAddressesQuery } from "@/hooks/useAddress";
import { useCheckoutStore } from "@/store/useCheckoutStore";
import AddAddress from "@/components/customer/address/AddAddress";
import { useMidtransPayment } from "@/hooks/useMidtransPayment";
import { MidtransScriptLoader } from "@/components/midtrans/MidtransScriptLoader";

const Checkout = () => {
  const navigate = useNavigate();
  const { data: addressesRes = { data: [], pagination: {} } } =
    useUserAddressesQuery();
  const { checkShippingCost } = useOrderMutation();
  const { data: carts, isLoading, isError, refetch } = useCartQuery();
  const { applyVoucher } = useVoucherMutation();

  const {
    note,
    courier,
    voucherCode,
    voucherInfo,
    selectedShipping,
    shippingOptions,
    setNote,
    setCourier,
    setVoucherCode,
    setVoucherInfo,
    setSelectedShipping,
    setShippingOptions,
  } = useCheckoutStore();

  const mainAddress = addressesRes.data.find((a) => a.isMain);
  const checkedItems = carts?.items?.filter((item) => item.isChecked) || [];

  useEffect(() => {
    if (!carts || checkedItems.length === 0) {
      navigate("/cart");
    }
  }, [carts, checkedItems, navigate]);

  const totalWeight = checkedItems.reduce(
    (acc, item) => acc + item.weight * item.quantity,
    0
  );
  const subtotal = checkedItems.reduce((acc, item) => acc + item.subtotal, 0);
  const voucherDiscount = voucherInfo?.discountValue || 0;
  const priceAfterVoucher = subtotal - voucherDiscount;
  const tax = priceAfterVoucher * 0.1;
  const total = priceAfterVoucher + tax + (selectedShipping?.cost || 0);

  const handleCheckShipping = () => {
    if (!mainAddress) return;
    checkShippingCost.mutate(
      {
        provinceId: mainAddress.provinceId,
        cityId: mainAddress.cityId,
        weight: totalWeight,
        courier,
      },
      {
        onSuccess: (res) => {
          const costs = res?.costs || [];
          setShippingOptions(costs);
          setSelectedShipping(costs[0]);
        },
      }
    );
  };

  const handleApplyVoucher = () => {
    if (!voucherCode.trim()) return;
    applyVoucher.mutate(
      { code: voucherCode, total: priceAfterVoucher },
      {
        onSuccess: (res) => {
          setVoucherInfo(res);
          toast.success(`Voucher "${res.code}" applied`);
        },
        onError: () => {
          setVoucherInfo(null);
        },
      }
    );
  };

  const { triggerPayment, isPending } = useMidtransPayment(() => {
    if (!mainAddress || !selectedShipping) return null;
    return {
      courier,
      shippingCost: selectedShipping.cost,
      voucherCode: voucherInfo?.code || null,
      note,
    };
  });

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  return (
    <>
      <MidtransScriptLoader />
      <section className="section py-10 md:py-16 space-y-6">
        <h2 className="text-2xl font-bold">Checkout</h2>
        <div className="grid md:grid-cols-3 gap-8">
          {/* LEFT SIDE */}
          <div className="md:col-span-2 space-y-6">
            {/* ADDRESS */}
            <div className="border rounded-lg p-4 space-y-2 bg-card">
              <h3 className="font-semibold">Shipping Address</h3>
              {!addressesRes.data?.length ? (
                <AddAddress />
              ) : mainAddress ? (
                <div className="text-sm text-foreground">
                  <p className="font-medium">{mainAddress.name}</p>
                  <p>{mainAddress.address}</p>
                  <p>
                    {mainAddress.province}, {mainAddress.city},{" "}
                    {mainAddress.subdistrict}, {mainAddress.district},{" "}
                    {mainAddress.postalCode}
                  </p>
                </div>
              ) : (
                <p className="text-muted-foreground">
                  No main address selected.
                </p>
              )}
            </div>

            {/* CART ITEMS */}
            <div className="space-y-4">
              <h3 className="font-semibold">Order Items</h3>
              {checkedItems.map((item) => (
                <div
                  key={item.productId}
                  className="flex items-center gap-4 border p-4 rounded-lg"
                >
                  <img
                    src={item.image}
                    alt={item.name}
                    className="w-16 h-16 object-cover rounded-md border"
                  />
                  <div className="flex-1 text-sm">
                    <p className="font-medium">{item.name}</p>
                    <p className="text-muted-foreground">
                      Qty: {item.quantity}
                    </p>
                    <p className="text-muted-foreground">
                      Subtotal: {formatRupiah(item.subtotal)}
                    </p>
                  </div>
                </div>
              ))}
            </div>

            {/* SHIPPING OPTION */}
            {mainAddress && (
              <div className="space-y-4 border p-4 rounded-lg mt-4">
                <h3 className="font-semibold">Shipping Option</h3>
                <select
                  value={courier}
                  onChange={(e) => setCourier(e.target.value)}
                  className="w-full px-3 py-2 border rounded-md text-sm"
                >
                  <option value="jne">JNE</option>
                  <option value="sicepat">SiCepat</option>
                </select>
                <Button size="sm" onClick={handleCheckShipping}>
                  Check Shipping Cost
                </Button>

                {shippingOptions.length > 0 && (
                  <div className="space-y-2">
                    {shippingOptions.map((opt) => (
                      <div
                        key={opt.service}
                        className={`p-3 border rounded-md cursor-pointer ${
                          selectedShipping?.service === opt.service
                            ? "bg-primary/10 border-primary"
                            : "hover:bg-muted"
                        }`}
                        onClick={() => setSelectedShipping(opt)}
                      >
                        <p className="font-medium text-sm">
                          {opt.service} - {formatRupiah(opt.cost)}
                        </p>
                        <p className="text-xs text-muted-foreground">
                          {opt.description} ({opt.etd})
                        </p>
                      </div>
                    ))}
                  </div>
                )}
              </div>
            )}
          </div>

          {/* RIGHT SIDE: SUMMARY */}
          <div className="border p-6 rounded-lg shadow-sm space-y-4 bg-card">
            <h3 className="text-lg font-semibold mb-2">Order Summary</h3>

            <div className="flex justify-between text-sm">
              <span>Subtotal</span>
              <span>{formatRupiah(subtotal)}</span>
            </div>
            <div className="flex justify-between text-sm">
              <span>Tax (10%)</span>
              <span>{formatRupiah(tax)}</span>
            </div>
            {selectedShipping && (
              <div className="flex justify-between text-sm">
                <span>Shipping ({selectedShipping.service})</span>
                <span>{formatRupiah(selectedShipping.cost)}</span>
              </div>
            )}

            <div className="space-y-2">
              <label className="text-sm font-medium">Promo Code</label>
              <div className="flex items-center gap-2">
                <Input
                  value={voucherCode}
                  onChange={(e) => setVoucherCode(e.target.value)}
                  placeholder="Enter code"
                  className="flex-1"
                />
                <Button
                  size="sm"
                  onClick={handleApplyVoucher}
                  disabled={applyVoucher.isPending}
                >
                  Apply
                </Button>
              </div>
              {voucherInfo && (
                <p className="text-xs text-green-600">
                  Applied: {voucherInfo.code} (Save Rp{" "}
                  {voucherDiscount.toLocaleString("id-ID")})
                </p>
              )}
            </div>

            <div className="flex justify-between font-semibold text-base pt-2">
              <span>Total</span>
              <span>{formatRupiah(total)}</span>
            </div>

            <div className="space-y-2">
              <label className="text-sm font-medium">Order Note</label>
              <Input
                value={note}
                onChange={(e) => setNote(e.target.value)}
                placeholder="e.g. please deliver after 3 PM"
              />
            </div>
            <Button
              size="lg"
              className="w-full mt-2"
              onClick={triggerPayment}
              disabled={!mainAddress || !selectedShipping || isPending}
            >
              {isPending ? "Processing..." : "Proceed to Payment"}
            </Button>
          </div>
        </div>
      </section>
    </>
  );
};

export default Checkout;
