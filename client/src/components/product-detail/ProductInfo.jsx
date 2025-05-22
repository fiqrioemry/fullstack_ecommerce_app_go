import { useState } from "react";
import { formatRupiah } from "@/lib/utils";
import { Minus, Plus } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useCartMutation } from "@/hooks/useCart";
import { useAuthStore } from "@/store/useAuthStore";

const ProductInfo = ({ product }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const { addToCart } = useCartMutation();
  const [quantity, setQuantity] = useState(1);

  const handleAddToCart = async () => {
    if (!user) navigate("/signin");
    if (quantity > product.stock) return;
    await addToCart.mutateAsync({
      productId: product.id,
      quantity,
    });
  };

  const hasDiscount = product.discount > 0;
  const finalPrice = hasDiscount
    ? product.price * (1 - product.discount / 100)
    : product.price;

  return (
    <div>
      <div className="space-y-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight text-foreground">
            {product.name}
          </h1>
          <p className="text-sm text-muted-foreground">
            {product.category} &raquo; {product.subcategory?.name}
          </p>
        </div>

        <div className="text-primary text-2xl font-semibold">
          {hasDiscount ? (
            <div className="flex items-center gap-2">
              <span className="line-through text-gray-400 text-base">
                {formatRupiah(product.price)}
              </span>
              <span className="text-red-600 font-semibold">
                {formatRupiah(finalPrice)}
              </span>
            </div>
          ) : (
            <span className="text-gray-900 font-semibold">
              {formatRupiah(product.price)}
            </span>
          )}
        </div>

        <p className="text-sm text-muted-foreground leading-relaxed">
          {product.description}
        </p>
      </div>

      <div>
        <div className="flex items-center gap-3 mt-6">
          <p className="text-sm font-medium text-muted-foreground">
            Quantity :
          </p>
          <div className="flex h-10 items-center border border-border rounded-md px-2 bg-secondary">
            <button
              onClick={() => setQuantity(Math.max(1, quantity - 1))}
              className="p-1 text-muted-foreground hover:text-primary"
            >
              <Minus className="w-4 h-4" />
            </button>
            <span className="px-3 text-sm font-medium text-foreground">
              {quantity}
            </span>
            <button
              onClick={() => setQuantity(Math.min(quantity + 1, product.stock))}
              className="p-1 text-muted-foreground hover:text-primary"
            >
              <Plus className="w-4 h-4" />
            </button>
          </div>
          <span className="text-xs text-muted-foreground">
            In stock: {product.stock}
          </span>
        </div>

        <div className="mt-4">
          <Button
            disabled={product.stock === 0 || quantity > product.stock}
            onClick={handleAddToCart}
            className="w-full"
          >
            Add to Cart
          </Button>
        </div>
      </div>
    </div>
  );
};

export { ProductInfo };
