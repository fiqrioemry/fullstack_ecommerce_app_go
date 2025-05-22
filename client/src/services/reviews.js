import { authInstance, publicInstance } from ".";

// GET /api/reviews/:productID
export const getProductReviews = async (productID) => {
  const res = await publicInstance.get(`/reviews/${productID}`);
  return res.data;
};

// POST /api/reviews/order/:productID
export const createReview = async ({ productID, data }) => {
  const res = await authInstance.post(`/reviews/order/${productID}`, data);
  return res.data;
};
