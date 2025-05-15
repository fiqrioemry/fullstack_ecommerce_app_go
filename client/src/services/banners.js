import { authInstance, publicInstance } from ".";

// GET /api/banners
export const getAllBanners = async () => {
  const res = await publicInstance.get("/banners");
  return res.data;
};

// GET /api/banners/:position
export const getBannerByPosition = async (position) => {
  const res = await publicInstance.get(`/banners/${position}`);
  return res.data;
};

// POST /api/banners
export const createBanner = async (formData) => {
  const res = await authInstance.post("/banners", formData);
  return res.data;
};

// PUT /api/banners/:id
export const updateBanner = async ({ id, data }) => {
  const res = await authInstance.put(`/banners/${id}`, data);
  return res.data;
};

// DELETE /api/banners/:id
export const deleteBanner = async (id) => {
  const res = await authInstance.delete(`/banners/${id}`);
  return res.data;
};
