import { toast } from "sonner";
import * as review from "@/services/reviews";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useProductReviewsQuery = (productID) =>
  useQuery({
    queryKey: ["reviews", productID],
    queryFn: () => review.getProductReviews(productID),
    enabled: !!productID,
  });

export const useReviewMutation = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: review.createReview,
    onSuccess: (res) => {
      toast.success(res?.message || "Review submitted");
      queryClient.invalidateQueries({ queryKey: "orders" });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Failed to submit review");
    },
  });
};
