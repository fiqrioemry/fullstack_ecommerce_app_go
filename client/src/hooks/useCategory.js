import { toast } from "sonner";
import * as category from "@/services/categories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCategoriesQuery = (search, page, limit, sort) =>
  useQuery({
    queryKey: ["categories", search, page, limit, sort],
    queryFn: () => category.getAllCategories(search, page, limit, sort),
    keepPreviousData: true,
  });

export const useCategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["category", id],
    queryFn: () => category.getCategoryById(id),
    enabled: !!id,
  });

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (successMsg, refetchFn) => ({
    onSuccess: (res, vars) => {
      toast.success(res?.message || successMsg);
      if (typeof refetchFn === "function") {
        refetchFn(vars);
      } else {
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createOptions: useMutation({
      mutationFn: category.createCategory,
      ...mutationOptions("Category created successfully"),
    }),

    updateOptions: useMutation({
      mutationFn: ({ id, data }) => category.updateCategory(id, data),
      ...mutationOptions("Category updated successfully", ({ id }) => {
        queryClient.invalidateQueries({ queryKey: ["category", id] });
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      }),
    }),

    deleteOptions: useMutation({
      mutationFn: category.deleteCategory,
      ...mutationOptions("Category deleted successfully"),
    }),
  };
};
