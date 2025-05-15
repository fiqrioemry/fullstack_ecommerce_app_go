import { toast } from "sonner";
import { useOrderMutation } from "./useOrder";
import { useNavigate } from "react-router-dom";

export function useMidtransPayment(getPayload, successRedirect = "/cart") {
  const navigate = useNavigate();
  const { checkout } = useOrderMutation();

  const triggerPayment = () => {
    const payload = getPayload();
    if (!payload) return;

    checkout.mutateAsync(payload, {
      onSuccess: (res) => {
        if (res.snapToken && window.snap) {
          window.snap.pay(res.snapToken, {
            onSuccess: () => {
              toast.success("Payment successful!");
            },
            onPending: () => {
              toast("Waiting for payment confirmation...");
            },
            onError: () => {
              toast.error("Payment failed.");
            },
            onClose: () => {
              toast.info("You closed the payment popup.");
              navigate(successRedirect);
            },
          });
        } else {
          toast.error("Failed to load Midtrans Snap.");
        }
      },
      onError: () => {
        toast.error("Failed to create transaction.");
      },
    });
  };

  return {
    triggerPayment,
    isPending: checkout.isPending,
  };
}
