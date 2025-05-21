import { useQuery } from "@tanstack/react-query";
import * as adminService from "@/services/dashboard";

// GET all customers
export const useAllCustomers = (params = {}) =>
  useQuery({
    queryKey: ["customers", params],
    queryFn: () =>
      adminService.getAllCustomers(
        params.search,
        params.page,
        params.limit,
        params.sort
      ),
    keepPreviousData: true,
  });

// GET customer detail by ID
export const useCustomerDetail = (id) =>
  useQuery({
    queryKey: ["customers", id],
    queryFn: () => adminService.getCustomerDetail(id),
    enabled: !!id,
  });

// GET summary stats (with optional gender filter)
export const useDashboardSummary = (gender = "") =>
  useQuery({
    queryKey: ["dashboard", "summary", gender],
    queryFn: () => adminService.getDashboardSummary(gender),
  });

// GET revenue stats by range (daily, monthly, yearly)
export const useRevenueStats = (range = "daily") =>
  useQuery({
    queryKey: ["dashboard", "revenue", range],
    queryFn: () => adminService.getRevenueStats(range),
  });
