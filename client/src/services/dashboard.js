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
export const getAllCustomers = async (search, page, limit, sort) => {
  const params = new URLSearchParams();
  if (search) params.append("q", search);
  if (page) params.append("page", String(page));
  if (limit) params.append("limit", String(limit));
  if (sort) params.append("sort", sort);

  const res = await authInstance.get(
    `/admin/dashboard/customers?${params.toString()}`
  );
  return res.data;
};

// GET /admin/dashboard/customers/:id
export const getCustomerDetail = async (id) => {
  const res = await authInstance.get(`/admin/dashboard/customers/${id}`);
  return res.data;
};
