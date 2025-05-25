import { useQuery } from "@tanstack/react-query";
import * as admin from "@/services/dashboard";

// GET all customers
export const useCustomersQuery = (param) =>
  useQuery({
    queryKey: ["customers", param],
    queryFn: () => admin.getAllCustomers(param),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 15,
  });

// GET customer detail by ID
export const useCustomerDetailQuery = (id) =>
  useQuery({
    queryKey: ["customers", id],
    queryFn: () => admin.getCustomerDetail(id),
    enabled: !!id,
  });

// GET summary stats (with optional gender filter)
export const useDashboardSummary = (gender = "") =>
  useQuery({
    queryKey: ["dashboard", "summary", gender],
    queryFn: () => admin.getDashboardSummary(gender),
  });

// GET revenue stats by range (daily, monthly, yearly)
export const useRevenueStats = (range = "daily") =>
  useQuery({
    queryKey: ["dashboard", "revenue", range],
    queryFn: () => admin.getRevenueStats(range),
  });
