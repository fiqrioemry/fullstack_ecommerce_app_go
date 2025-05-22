import { toast } from "sonner";
import { useOrderMutation } from "./useOrder";
import { loadStripe } from "@stripe/stripe-js";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PUBLISHABLE_KEY);

export function useStripePayment(getPayload) {
  const { checkout } = useOrderMutation();

  const triggerPayment = () => {
    const payload = getPayload();
    if (!payload) return;

    checkout.mutateAsync(payload, {
      onSuccess: async (res) => {
        console.log(res);
        if (!res.sessionId) {
          toast.error("No session ID received.");
          return;
        }

        const stripe = await stripePromise;
        if (!stripe) {
          toast.error("Stripe SDK not loaded.");
          return;
        }

        const result = await stripe.redirectToCheckout({
          sessionId: res.sessionId,
        });

        if (result.error) {
          toast.error(result.error.message || "Failed to redirect to Stripe.");
        }
      },
      onError: () => {
        toast.error("Failed to create Stripe session.");
      },
    });
  };

  return {
    triggerPayment,
    isPending: checkout.isPending,
  };
}
