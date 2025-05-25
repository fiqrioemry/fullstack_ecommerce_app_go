import qs from "qs";
import { publicInstance, authInstance } from ".";
import { buildFormData } from "../lib/utils";

export const searchProducts = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await publicInstance.get(`/product?${queryString}`);
  return res.data;
};

// GET /api/product/:slug
export const getProductBySlug = async (slug) => {
  const res = await publicInstance.get(`/product/${slug}`);
  return res.data;
};

// POST /api/product (admin)
export const createProduct = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/product", formData);
  return res.data;
};

// PUT /api/product/:id (admin)
export const updateProduct = async ({ id, data }) => {
  const formData = buildFormData(data);
  const res = await authInstance.put(`/product/${id}`, formData);
  return res.data;
};

// DELETE /api/product/:id (admin)
export const deleteProduct = async (id) => {
  const res = await authInstance.delete(`/product/${id}`);
  return res.data;
};
