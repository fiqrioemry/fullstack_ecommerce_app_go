import { buildFormData } from "@/lib/utils";
import { authInstance, publicInstance } from ".";

// GET /api/reviews/:productID
export const getProductReviews = async (productID) => {
  const res = await publicInstance.get(`/reviews/${productID}`);
  return res.data;
};

// POST /api/reviews/order/:productID
export const createReview = async ({ productId, data }) => {
  console.log(data);
  const formData = buildFormData(data);
  const res = await authInstance.post(`/reviews/order/${productId}`, formData);
  return res.data;
};
