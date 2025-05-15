import { publicInstance, authInstance } from ".";

// GET /api/product?q=keyword&category=...&page=1&limit=10 ...
export const searchProducts = async (
  search,
  page,
  limit,
  sort,
  categoryId,
  minPrice,
  maxPrice,
  rating
) => {
  const params = new URLSearchParams();
  if (search) params.append("q", search);
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));
  if (sort) params.append("sort", sort);
  if (categoryId) params.append("category", categoryId);
  if (minPrice) params.append("minPrice", String(minPrice));
  if (maxPrice) params.append("maxPrice", String(maxPrice));
  if (rating) params.append("rating", String(rating));

  console.log(params.toString());
  const res = await publicInstance.get(`/product?${params.toString()}`);
  return res.data;
};

// GET /api/product/:slug
export const getProductBySlug = async (slug) => {
  const res = await publicInstance.get(`/product/${slug}`);
  return res.data;
};

// POST /api/product (admin)
export const createProduct = async (formData) => {
  const res = await authInstance.post("/product", formData);
  return res.data;
};

// PUT /api/product/:id (admin)
export const updateProduct = async ({ id, data }) => {
  const res = await authInstance.put(`/product/${id}`, data);
  return res.data;
};

// DELETE /api/product/:id (admin)
export const deleteProduct = async (id) => {
  const res = await authInstance.delete(`/product/${id}`);
  return res.data;
};
