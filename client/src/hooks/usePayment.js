// src/hooks/usePayment.js
import { toast } from "sonner";
import * as paymentService from "@/services/payment";
import { useMutation, useQuery } from "@tanstack/react-query";

// GET /api/payments?q=&page=&limit= (admin only)
export const useAdminPaymentsQuery = (search, page, limit, sort, status) =>
  useQuery({
    queryKey: ["admin-payments", search, page, limit, sort, status],
    queryFn: () =>
      paymentService.getAllUserPayments(search, page, limit, sort, status),
    keepPreviousData: true,
    staleTime: 0,
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
