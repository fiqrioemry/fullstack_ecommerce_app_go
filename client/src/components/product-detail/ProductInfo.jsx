import { useState } from "react";
import { formatRupiah } from "@/lib/utils";
import { Minus, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useCartMutation } from "@/hooks/useCart";
import { useAuthStore } from "@/store/useAuthStore";
import { useNavigate } from "react-router-dom";

const ProductInfo = ({ product }) => {
  const navigate = useNavigate();
  const { user } = useAuthStore();
  const [quantity, setQuantity] = useState(1);
  const { addToCart } = useCartMutation();

  const handleAddToCart = async () => {
    if (!user) navigate("/signin");
    if (quantity > product.stock) return;
    await addToCart.mutateAsync({
      productId: product.id,
      quantity,
    });
  };

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
          {formatRupiah(product.price)}
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
