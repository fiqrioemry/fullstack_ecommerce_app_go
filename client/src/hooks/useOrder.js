import { toast } from "sonner";
import * as order from "@/services/orders";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// 🔄 Queries
export const useAllOrdersQuery = ({ search, page, limit, sort, status }) =>
  useQuery({
    queryKey: ["orders", search, page, limit, sort, status],
    queryFn: () => order.getAllOrders(search, page, limit, sort, status),
    keepPreviousData: true,
  });
export const useOrderDetailQuery = (id) =>
  useQuery({
    queryKey: ["orderDetail", id],
    queryFn: () => order.getOrderDetail(id),
    enabled: !!id,
  });

export const useShipmentQuery = (orderId) =>
  useQuery({
    queryKey: ["shipment", orderId],
    queryFn: () => order.getShipmentByOrderID(orderId),
    enabled: !!orderId,
  });

export const useOrderMutation = () => {
  const queryClient = useQueryClient();

  const baseMutation = (fn, msg, invalidateKey) => ({
    mutationFn: fn,
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      if (invalidateKey) {
        queryClient.invalidateQueries({ queryKey: invalidateKey });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    checkout: useMutation(
      baseMutation(order.checkout, "Order created successfully", ["orders"])
    ),

    createShipment: useMutation(
      baseMutation(order.createShipment, "Order Shipped Succcessfully", [
        "orders",
      ])
    ),

    updateShipment: useMutation(
      baseMutation(order.updateShipmentStatus, "Order confirmed as delivered", [
        "orders",
      ])
    ),

    checkShippingCost: useMutation({
      mutationFn: order.checkShippingCost,
      onError: (err) => {
        toast.error(err?.response?.data?.message || "Failed to check shipping");
      },
    }),
  };
};
