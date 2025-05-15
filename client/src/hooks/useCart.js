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

  const baseMutation = (fn, msg) => ({
    mutationFn: fn,
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      invalidateCart();
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    addToCart: useMutation(baseMutation(cart.addToCart, "Item added to cart")),
    updateQuantity: useMutation(
      baseMutation(cart.updateCartQuantity, "Quantity updated")
    ),
    toggleCheck: useMutation(baseMutation(cart.toggleCartItemChecked)),
    removeItem: useMutation(
      baseMutation(cart.removeCartItem, "Item removed from cart")
    ),
    clearCart: useMutation(baseMutation(cart.clearCart, "Cart cleared")),
  };
};
