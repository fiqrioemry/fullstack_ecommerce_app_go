import qs from "qs";
import { authInstance } from ".";

// POST /api/orders (customer only)
export const checkout = async (data) => {
  const res = await authInstance.post("/orders", data);
  return res.data;
};

// GET /api/orders (admin or customer)
export const getAllOrders = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await authInstance.get(`/orders?${queryString}`);
  return res.data;
};

// GET /api/orders/:orderID (detail for admin/customer)
export const getOrderDetail = async (orderID) => {
  const res = await authInstance.get(`/orders/${orderID}`);
  return res.data;
};

// GET /api/orders/:orderID/shipment (admin/customer)
export const getShipmentByOrderID = async (orderID) => {
  const res = await authInstance.get(`/orders/${orderID}/shipment`);
  return res.data;
};
// POST /api/orders/:orderID/shipment (admin only)
export const createShipment = async ({ orderId, data }) => {
  const res = await authInstance.post(`/orders/${orderId}/shipment`, data);
  return res.data;
};

// POST /api/orders/:orderID/shipment
export const updateShipmentStatus = async (orderID) => {
  const res = await authInstance.put(`/orders/${orderID}/shipment`);
  return res.data;
};

// POST /api/orders/check-shipping (no role restriction)
export const checkShippingCost = async (data) => {
  const res = await authInstance.post("/orders/check-shipping", data);
  return res.data;
};
