import { publicInstance, authInstance } from ".";
import { buildFormData } from "../lib/utils";

export const searchProducts = async (
  search,
  status,
  page,
  limit,
  sort,
  category,
  minPrice,
  maxPrice,
  rating
) => {
  const params = new URLSearchParams();
  if (search) params.append("q", search);
  if (status) params.append("status", String(status));
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));
  if (sort) params.append("sort", sort);
  if (category) params.append("category", category);
  if (minPrice) params.append("minPrice", String(minPrice));
  if (maxPrice) params.append("maxPrice", String(maxPrice));
  if (rating) params.append("rating", String(rating));
  const res = await publicInstance.get(`/product?${params.toString()}`);
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
