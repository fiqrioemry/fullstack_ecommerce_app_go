import qs from "qs";
import { authInstance } from ".";

// GET /admin/dashboard/summary?gender=
export const getDashboardSummary = async (gender = "") => {
  const res = await authInstance.get("/admin/dashboard/summary", {
    params: { gender },
  });

  return res.data;
};

// GET /admin/dashboard/revenue?range=daily|monthly|yearly
export const getRevenueStats = async (range = "daily") => {
  const res = await authInstance.get("/admin/dashboard/revenue", {
    params: { range },
  });
  return res.data;
};

// GET /admin/dashboard/customers?page=1&limit=10&q=&sort=
export const getAllCustomers = async (param) => {
  const queryString = qs.stringify(param, { skipNulls: true });
  const res = await authInstance.get(
    `/admin/dashboard/customers?${queryString}`
  );
  return res.data;
};

// GET /admin/dashboard/customers/:id
export const getCustomerDetail = async (id) => {
  const res = await authInstance.get(`/admin/dashboard/customers/${id}`);
  return res.data;
};
