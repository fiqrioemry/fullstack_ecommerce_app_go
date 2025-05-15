import { publicInstance, authInstance } from ".";

// GET /api/categories
export const getAllCategories = async (search, page, limit, sort) => {
  const params = new URLSearchParams();
  if (search) params.append("q", search);
  if (sort) params.append("sort", sort);
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));

  const res = await publicInstance.get(`/categories?${params.toString()}`);
  return res.data;
};

// GET /api/categories/:id
export const getCategoryById = async (id) => {
  const res = await publicInstance.get(`/categories/${id}`);
  return res.data;
};

// POST /api/categories
export const createCategory = async (data) => {
  const res = await authInstance.post("/categories", data);
  return res.data;
};

// PUT /api/categories/:id
export const updateCategory = async (id, data) => {
  const res = await authInstance.put(`/categories/${id}`, data);
  return res.data;
};

// DELETE /api/categories/:id
export const deleteCategory = async (id) => {
  const res = await authInstance.delete(`/categories/${id}`);
  return res.data;
};
