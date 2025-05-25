import qs from "qs";
import { publicInstance, authInstance } from ".";

// GET /api/payments?q=&page=&limit=
export const getAllUserPayments = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await authInstance.get(`/payments?${queryString}`);
  return res.data;
};

// POST /api/payments/notification (webhook - public)
export const handlePaymentNotification = async (data) => {
  const res = await publicInstance.post("/payments/notification", data);
  return res.data;
};
