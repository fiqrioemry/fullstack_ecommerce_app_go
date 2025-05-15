import { publicInstance, authInstance } from ".";

// GET /api/payments?q=&page=&limit=
export const getAllUserPayments = async (search, page, limit, sort, status) => {
  const params = new URLSearchParams();
  if (search) params.append("q", search);
  if (sort) params.append("sort", sort);
  if (status) params.append("status", status);
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));

  const res = await authInstance.get(`/payments?${params.toString()}`);
  return res.data;
};

// POST /api/payments/notification (webhook - public)
export const handlePaymentNotification = async (data) => {
  const res = await publicInstance.post("/payments/notification", data);
  return res.data;
};
