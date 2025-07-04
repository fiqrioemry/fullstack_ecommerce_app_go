import { toast } from "sonner";
import * as category from "@/services/categories";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useCategoriesQuery = (param) =>
  useQuery({
    queryKey: ["categories", param],
    queryFn: () => category.getAllCategories(param),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 30,
  });

export const useCategoryDetailQuery = (id) =>
  useQuery({
    queryKey: ["category", id],
    queryFn: () => category.getCategoryById(id),
    enabled: !!id,
  });

export const useCategoryMutation = () => {
  const queryClient = useQueryClient();

  const mutationBase = (fn, msg, invalidate = true) => ({
    mutationFn: fn,
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      if (invalidate) {
        queryClient.invalidateQueries({ queryKey: ["categories"] });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createCategory: useMutation(
      mutationBase(category.createCategory, "Category created successfully")
    ),
    updateCategory: useMutation(
      mutationBase(category.updateCategory, "Category updated successfully")
    ),
    deleteCategory: useMutation(
      mutationBase(category.deleteCategory, "Category deleted successfully")
    ),
  };
};
