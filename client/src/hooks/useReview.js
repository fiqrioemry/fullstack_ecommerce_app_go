import { toast } from "sonner";
import * as review from "@/services/reviews";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useProductReviewsQuery = (productID, param) =>
  useQuery({
    queryKey: ["reviews", productID, param],
    queryFn: () => review.getProductReviews(productID, param),
    enabled: !!productID,
  });

export const useReviewMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ itemId, data, orderId }) =>
      review.createReview({ itemId, data, orderId }),
    onSuccess: (res, variables) => {
      toast.success(res?.message || "Review submitted");
      queryClient.invalidateQueries({
        queryKey: ["orderDetail", variables.orderId],
      });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
