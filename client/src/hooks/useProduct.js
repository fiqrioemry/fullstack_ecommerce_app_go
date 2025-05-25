import { toast } from "sonner";
import * as product from "@/services/product";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useSearchProductsQuery = (param) =>
  useQuery({
    queryKey: ["products", param],
    queryFn: () => product.searchProducts(param),
    keepPreviousData: true,
    staleTime: 1000 * 60 * 5,
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
