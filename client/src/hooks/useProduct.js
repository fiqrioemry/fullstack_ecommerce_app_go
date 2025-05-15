import { toast } from "sonner";
import * as product from "@/services/product";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// 🔄 Queries

export const useSearchProductsQuery = (params = {}) =>
  useQuery({
    queryKey: ["products", params],
    queryFn: () =>
      product.searchProducts(
        params.search,
        params.page,
        params.limit,
        params.sort,
        params.categoryId,
        params.minPrice,
        params.maxPrice,
        params.rating
      ),
    keepPreviousData: true,
  });

export const useProductDetailQuery = (slug) =>
  useQuery({
    queryKey: ["product", slug],
    queryFn: () => product.getProductBySlug(slug),
    enabled: !!slug,
  });

export const useProductMutation = () => {
  const queryClient = useQueryClient();

  const mutationBase = (fn, msg, invalidate = true) => ({
    mutationFn: fn,
    onSuccess: (res) => {
      toast.success(res?.message || msg);
      if (invalidate) {
        queryClient.invalidateQueries({ queryKey: ["products"] });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createProduct: useMutation(
      mutationBase(product.createProduct, "Product created successfully")
    ),
    updateProduct: useMutation(
      mutationBase(product.updateProduct, "Product updated successfully")
    ),
    deleteProduct: useMutation(
      mutationBase(product.deleteProduct, "Product deleted successfully")
    ),
  };
};
