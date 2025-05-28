import qs from "qs";
import { buildFormData } from "@/lib/utils";
import { authInstance, publicInstance } from ".";

// GET /api/reviews/:productID
export const getProductReviews = async (productID, param) => {
  console.log(productID);
  console.log(param);
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await publicInstance.get(`/reviews/${productID}?${queryString}`);
  return res.data;
};

// POST /api/reviews/order/:productID
export const createReview = async ({ itemId, data }) => {
  const formData = buildFormData(data);
  const res = await authInstance.post(`/reviews/order/${itemId}`, formData);
  return res.data;
};
