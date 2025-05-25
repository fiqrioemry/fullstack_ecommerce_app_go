import { toast } from "sonner";
import * as cart from "@/services/cart";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCartQuery = () =>
  useQuery({
    queryKey: ["cart"],
    queryFn: cart.getCart,
  });

export const useCartMutation = () => {
  const queryClient = useQueryClient();

  const invalidateCart = () =>
    queryClient.invalidateQueries({ queryKey: ["cart"] });

  const baseMutation = (fn, msg, showToast = true) => ({
    mutationFn: fn,
    onSuccess: (res) => {
      if (showToast) {
        toast.success(res?.message || msg);
      }
      invalidateCart();
    },
    onError: (err) => {
      toast.error("Please login to start shopping");
    },
  });

  return {
    addToCart: useMutation(baseMutation(cart.addToCart, "Item added to cart")),
    updateQuantity: useMutation(
      baseMutation(cart.updateCartQuantity, "", false)
    ),
    toggleCheck: useMutation(
      baseMutation(cart.toggleCartItemChecked, "", false)
    ),
    removeItem: useMutation(
      baseMutation(cart.removeCartItem, "Item removed from cart")
    ),
    clearCart: useMutation(baseMutation(cart.clearCart, "Cart cleared")),
  };
};
