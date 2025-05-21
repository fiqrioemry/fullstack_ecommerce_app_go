import { authInstance } from ".";

export const getUserAddresses = async (search, page, limit, sort) => {
  const params = new URLSearchParams();

  if (sort) params.append("sort", sort);
  if (search) params.append("q", search);
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));

  const res = await authInstance.get(`/user/addresses?${params.toString()}`);
  return res.data;
};

export const createUserAddress = async (data) => {
  const res = await authInstance.post("/user/addresses", data);
  return res.data;
};

// PUT /api/user/addresses/:id
export const updateUserAddress = async (id, data) => {
  const res = await authInstance.put(`/user/addresses/${id}`, data);
  return res.data;
};

// DELETE /api/user/addresses/:id
export const deleteUserAddress = async (id) => {
  const res = await authInstance.delete(`/user/addresses/${id}`);
  return res.data;
};

// PATCH /api/user/addresses/:id/main
export const setMainAddress = async (id) => {
  const res = await authInstance.patch(`/user/addresses/${id}/main`);
  return res.data;
};
