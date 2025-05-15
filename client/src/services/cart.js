import { authInstance } from ".";

// GET /api/cart
export const getCart = async () => {
  const res = await authInstance.get("/cart");
  return res.data;
};

// POST /api/cart
export const addToCart = async (data) => {
  const res = await authInstance.post("/cart", data);
  return res.data;
};

// DELETE /api/cart
export const clearCart = async () => {
  const res = await authInstance.delete("/cart");
  return res.data;
};

// DELETE /api/cart/:productId
export const removeCartItem = async (productId) => {
  const res = await authInstance.delete(`/cart/${productId}`);
  return res.data;
};

// PUT /api/cart/:productId
export const updateCartQuantity = async ({ productId, quantity }) => {
  const res = await authInstance.put(`/cart/${productId}`, { quantity });
  return res.data;
};

export const toggleCartItemChecked = async ({ productId, isChecked }) => {
  console.log(isChecked);
  const res = await authInstance.patch(`/cart/${productId}/checked`, {
    isChecked,
  });
  return res.data;
};
