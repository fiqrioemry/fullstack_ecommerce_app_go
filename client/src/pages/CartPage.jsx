import { formatRupiah } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Minus, Plus, Trash2 } from "lucide-react";
import { useCartQuery } from "@/hooks/useCart";
import { Loading } from "@/components/ui/Loading";
import { useCartMutation } from "@/hooks/useCart";
import { ErrorDialog } from "@/components/ui/ErrorDialog";

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

  if (isLoading) return <Loading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;

  const cartItems = carts.items || [];
  const totalCheckedItems = cartItems.filter((c) => c.isChecked);
  const totalPriceItems = carts.total || 0;
  const isEmpty = cartItems.length === 0;

  console.log(carts.items);

  return (
    <section className="section py-16 md:py-20 space-y-8">
      <h2 className="text-2xl font-bold">Shopping Cart</h2>

      {isEmpty ? (
        <div className="flex flex-col items-center justify-center mt-16">
          <img
            src="/empty-cart.webp"
            alt="Empty Cart"
            className=" h-96 opacity-70"
          />
          <p className="text-muted-foreground mt-4">
            Your cart is currently empty.
          </p>
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Cart Items */}
          <div className="lg:col-span-2 space-y-6">
            {cartItems.map((item) => (
              <div
                key={item.productId}
                className="flex items-center gap-4 border border-border p-4 rounded-lg shadow-sm bg-card"
              >
                <input
                  type="checkbox"
                  checked={item.isChecked}
                  onChange={() =>
                    toggleCheck.mutate({
                      productId: item.productId,
                      isChecked: !item.isChecked,
                    })
                  }
                  className="w-5 h-5 accent-primary"
                />
                <img
                  src={item.imageUrl}
                  alt={item.name}
                  className="w-20 h-20 rounded object-cover border"
                />
                <div className="flex-1 text-sm">
                  <h5 className="font-semibold text-foreground">{item.name}</h5>

                  {item.discount > 0 ? (
                    <div className="text-xs mt-1">
                      <span className="line-through text-muted-foreground mr-2">
                        {formatRupiah(item.price)}
                      </span>
                      <span className="text-primary font-medium">
                        {formatRupiah(item.discountedPrice)} × {item.quantity}
                      </span>
                    </div>
                  ) : (
                    <p className="text-muted-foreground mt-1">
                      {item.quantity} × {formatRupiah(item.price)}
                    </p>
                  )}

                  <p className="text-muted-foreground mt-1">
                    Subtotal:{" "}
                    <span className="font-medium text-foreground">
                      {formatRupiah(item.subtotal)}
                    </span>
                  </p>

                  <div className="flex items-center gap-2 mt-2">
                    <button
                      onClick={() => handleQuantityChange(item, -1)}
                      className="px-2 py-1 border border-border rounded"
                    >
                      <Minus className="w-4 h-4" />
                    </button>
                    <span className="w-6 text-center">{item.quantity}</span>
                    <button
                      onClick={() => handleQuantityChange(item, 1)}
                      className="px-2 py-1 border border-border rounded"
                    >
                      <Plus className="w-4 h-4" />
                    </button>
                  </div>
                </div>
                <button
                  onClick={() => removeItem.mutate(item.productId)}
                  className="text-destructive hover:text-red-700 ml-2"
                  title="Remove item"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>
            ))}
          </div>

          {/* Summary */}
          <div className="border border-border p-6 rounded-lg shadow-sm bg-card">
            <h2 className="text-lg font-semibold mb-4 text-foreground">
              Order Summary
            </h2>
            <div className="flex justify-between mb-2 text-muted-foreground text-sm">
              <span>Total Items</span>
              <span>{totalCheckedItems.length}</span>
            </div>
            <div className="flex justify-between font-semibold text-base text-foreground mb-4">
              <span>Total</span>
              <span>{formatRupiah(totalPriceItems)}</span>
            </div>
            <Button
              className="w-full"
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
