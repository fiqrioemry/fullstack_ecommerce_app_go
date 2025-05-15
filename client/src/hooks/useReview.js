import { toast } from "sonner";
import * as review from "@/services/reviews";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// 🔄 GET reviews by product
export const useProductReviewsQuery = (productID) =>
  useQuery({
    queryKey: ["productReviews", productID],
    queryFn: () => review.getProductReviews(productID),
    enabled: !!productID,
  });

export const useReviewMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: review.createReview,
    onSuccess: (res, { productID }) => {
      toast.success(res?.message || "Review submitted");
      queryClient.invalidateQueries({
        queryKey: ["productReviews", productID],
      });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
