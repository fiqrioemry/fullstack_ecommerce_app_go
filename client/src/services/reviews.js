import { authInstance, publicInstance } from ".";

// GET /api/reviews/:productID
export const getProductReviews = async (productID) => {
  const res = await publicInstance.get(`/reviews/${productID}`);
  return res.data;
};

// POST /api/reviews/order/:orderID/product/:productID
export const createReview = async ({ orderID, productID, data }) => {
  const res = await authInstance.post(
    `/reviews/order/${orderID}/product/${productID}`,
    data
  );
  return res.data;
};
