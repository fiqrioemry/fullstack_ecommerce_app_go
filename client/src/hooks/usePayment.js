// src/hooks/usePayment.js
import { toast } from "sonner";
import * as paymentService from "@/services/payment";
import { useMutation, useQuery } from "@tanstack/react-query";

// GET /api/payments?q=&page=&limit= (admin only)
export const usePaymentsQuery = (param) =>
  useQuery({
    queryKey: ["payments", param],
    queryFn: () => paymentService.getAllUserPayments(param),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 15,
  });

// POST /api/payments/notification (webhook - public)
export const useHandlePaymentNotification = () =>
  useMutation({
    mutationFn: paymentService.handlePaymentNotification,
    onSuccess: () => {
      toast.success("Payment notification handled");
    },
    onError: (err) => {
      toast.error(
        err?.response?.data?.message || "Failed to handle notification"
      );
    },
  });
