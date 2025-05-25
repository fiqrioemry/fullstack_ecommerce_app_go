import qs from "qs";
import { buildFormData } from "../lib/utils";
import { publicInstance, authInstance } from ".";

// GET /api/categories
export const getAllCategories = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await publicInstance.get(`/categories?${queryString}`);
  return res.data;
};

// GET /api/categories/:id
export const getCategoryById = async (id) => {
  const res = await publicInstance.get(`/categories/${id}`);
  return res.data;
};

// POST /api/categories
export const createCategory = async (data) => {
  const formData = buildFormData(data);
  const res = await authInstance.post("/categories", formData);
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
