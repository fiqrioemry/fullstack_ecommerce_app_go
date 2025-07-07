import { formatRupiah } from "@/lib/utils";
import { useNavigate } from "react-router-dom";
import { useCartQuery } from "@/hooks/useCart";
import { Button } from "@/components/ui/Button";
import { Loading } from "@/components/ui/Loading";
import { useCartMutation } from "@/hooks/useCart";
import { ErrorDialog } from "@/components/ui/ErrorDialog";
import { Minus, Plus, Trash2, XCircle } from "lucide-react";

const CartPage = () => {
  const navigate = useNavigate();
  const { toggleCheck, updateQuantity, removeItem, clearCart } =
    useCartMutation();
  const { data: carts, isLoading, isError, refetch } = useCartQuery();

  const handleQuantityChange = (item, delta) => {
    const newQuantity = item.quantity + delta;
    if (newQuantity >= 1) {
      updateQuantity.mutate({
        productId: item.productId,
        quantity: newQuantity,
      });
    }
  };

  const handleCheckedItem = (item) => {
    toggleCheck.mutate({ productId: item.productId });
  };

  if (isLoading) return <Loading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  const cartItems = carts.items || [];
  const totalCheckedItems = cartItems.filter((c) => c.isChecked);
  const totalPriceItems = totalCheckedItems.reduce(
    (acc, item) => acc + item.subtotal,
    0
  );

  const isEmpty = cartItems.length === 0;

  return (
    <section className="section py-16 md:py-20 space-y-8">
      <div className="flex items-center justify-between">
        <h2 className="text-2xl font-bold">🛒 Shopping Cart</h2>
      </div>

      {isEmpty ? (
        <div className="flex flex-col items-center justify-center mt-16">
          <img
            src="/empty-cart.webp"
            alt="Empty Cart"
            className="h-96 opacity-70"
          />
          <p className="text-muted-foreground mt-4">
            Your cart is currently empty.
          </p>
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Cart Items */}
          <div className="lg:col-span-2 space-y-6">
            <div className="flex items-center justify-end mb-2">
              <button
                onClick={() => clearCart.mutate()}
                className="text-sm text-red-500 hover:underline flex items-center gap-1"
              >
                <XCircle className="w-4 h-4" />
                Clear All
              </button>
            </div>
            {cartItems.map((item) => (
              <div
                key={item.productId}
                className="flex flex-col md:flex-row items-start md:items-center gap-4 border border-border p-4 rounded-2xl shadow bg-white dark:bg-card"
              >
                <input
                  type="checkbox"
                  checked={item.isChecked}
                  className="w-5 h-5 accent-primary mt-1"
                  onChange={() => handleCheckedItem(item)}
                />
                <img
                  src={item.image}
                  alt={item.name}
                  className="w-24 h-24 rounded-xl object-cover border"
                />
                <div className="flex-1 text-sm w-full">
                  <div className="flex justify-between items-center mb-1">
                    <h5 className="font-semibold text-lg text-foreground">
                      {item.name}
                    </h5>
                    <button
                      onClick={() => removeItem.mutate(item.productId)}
                      className="text-destructive hover:text-red-700"
                      title="Remove item"
                    >
                      <Trash2 className="w-5 h-5" />
                    </button>
                  </div>

                  {item.discount > 0 ? (
                    <p className="text-sm">
                      <span className="line-through text-muted-foreground mr-2">
                        {formatRupiah(item.price)}
                      </span>
                      <span className="text-primary font-medium">
                        {formatRupiah(item.discountedPrice)} × {item.quantity}
                      </span>
                    </p>
                  ) : (
                    <p className="text-muted-foreground">
                      {item.quantity} × {formatRupiah(item.price)}
                    </p>
                  )}

                  <div className="flex justify-between items-center mt-3">
                    <div className="flex items-center gap-2">
                      <button
                        onClick={() => handleQuantityChange(item, -1)}
                        className="px-2 py-1 border rounded-lg"
                      >
                        <Minus className="w-4 h-4" />
                      </button>
                      <span className="w-8 text-center font-semibold">
                        {item.quantity}
                      </span>
                      <button
                        onClick={() => handleQuantityChange(item, 1)}
                        className="px-2 py-1 border rounded-lg"
                      >
                        <Plus className="w-4 h-4" />
                      </button>
                    </div>
                    <p className="font-medium text-foreground">
                      {formatRupiah(item.subtotal)}
                    </p>
                  </div>
                </div>
              </div>
            ))}
          </div>

          {/* Summary */}
          <div className="border border-border p-6 rounded-2xl shadow bg-white dark:bg-card space-y-4">
            <h2 className="text-xl font-semibold text-foreground mb-2">
              Order Summary
            </h2>
            <div className="flex justify-between text-sm text-muted-foreground">
              <span>Selected Items</span>
              <span>{totalCheckedItems.length}</span>
            </div>
            <div className="flex justify-between text-base font-bold text-foreground">
              <span>Total</span>
              <span>{formatRupiah(totalPriceItems)}</span>
            </div>
            <Button
              className="w-full mt-2"
              disabled={totalCheckedItems.length === 0}
              onClick={() => navigate("/cart/checkout")}
            >
              Proceed to Checkout
            </Button>
          </div>
        </div>
      )}
    </section>
  );
};

export default CartPage;
