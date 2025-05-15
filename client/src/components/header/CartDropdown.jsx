import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
} from "@/components/ui/dropdown-menu";
import { formatRupiah } from "@/lib/utils";
import { ShoppingCart } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useCartQuery } from "@/hooks/useCart";

const CartDropdown = () => {
  const navigate = useNavigate();
  const { data, isLoading } = useCartQuery();

  const cartItems = data?.items || [];
  const totalItemsPrice = data?.total || 0;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button className="relative text-muted-foreground hover:text-primary transition-colors">
          <ShoppingCart className="w-6 h-6" />
          <span className="absolute -top-2 -right-2 bg-destructive text-white text-[10px] w-5 h-5 rounded-full flex items-center justify-center">
            {cartItems.length || 0}
          </span>
        </button>
      </DropdownMenuTrigger>

      <DropdownMenuContent
        align="end"
        className="w-[320px] bg-card border border-border shadow-xl rounded-lg overflow-hidden"
      >
        {isLoading ? (
          <div className="p-4 text-center text-sm text-muted-foreground">
            Loading cart...
          </div>
        ) : cartItems.length > 0 ? (
          <>
            <div className="max-h-64 overflow-y-auto divide-y">
              {cartItems.map((item) => (
                <div key={item.productId} className="flex gap-4 p-4">
                  <img
                    src={item.imageUrl}
                    alt={item.name}
                    className="w-16 h-16 object-cover rounded-md border"
                  />
                  <div className="flex-1 text-sm">
                    <p className="font-medium line-clamp-2 text-foreground">
                      {item.name}
                    </p>

                    {item.discount > 0 ? (
                      <div className="text-xs mt-1">
                        <span className="text-muted-foreground line-through mr-2">
                          {formatRupiah(item.price)}
                        </span>
                        <span className="text-primary font-medium">
                          {formatRupiah(item.discountedPrice)} × {item.quantity}
                        </span>
                      </div>
                    ) : (
                      <p className="text-xs text-muted-foreground mt-1">
                        {item.quantity} × {formatRupiah(item.price)}
                      </p>
                    )}
                  </div>
                </div>
              ))}
            </div>

            <div className="p-4 border-t border-border">
              <div className="flex items-center justify-between mb-3 text-sm text-muted-foreground">
                <span>Total:</span>
                <span className="font-semibold text-foreground">
                  {formatRupiah(totalItemsPrice)}
                </span>
              </div>
              <Button
                onClick={() => navigate("/cart")}
                className="w-full"
                variant="default"
              >
                View Cart
              </Button>
            </div>
          </>
        ) : (
          <div className="p-4 flex flex-col items-center justify-center text-center h-60 text-muted-foreground">
            <img
              src="/empty-cart.webp"
              alt="Empty Cart"
              className="w-24 h-24 opacity-70 mb-3"
            />
            <p className="text-sm">Your cart is currently empty.</p>
          </div>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export { CartDropdown };
