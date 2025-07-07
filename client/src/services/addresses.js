import qs from "qs";
import { authInstance } from ".";

export const getUserAddresses = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await authInstance.get(`/user/addresses?${queryString}`);
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
