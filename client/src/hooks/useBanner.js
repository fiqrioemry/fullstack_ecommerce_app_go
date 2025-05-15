import { toast } from "sonner";
import * as banner from "@/services/banners";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

// 🔄 QUERY
export const useBannersQuery = () =>
  useQuery({
    queryKey: ["banners"],
    queryFn: banner.getAllBanners,
  });

export const useBannerByPositionQuery = (position) =>
  useQuery({
    queryKey: ["banner", position],
    queryFn: () => banner.getBannerByPosition(position),
    enabled: !!position,
  });

// ✏️ MUTATIONS
export const useBannerMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (successMsg, invalidate = true) => ({
    onSuccess: (res) => {
      toast.success(res?.message || successMsg);
      if (invalidate) queryClient.invalidateQueries({ queryKey: ["banners"] });
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createBanner: useMutation({
      mutationFn: banner.createBanner,
      ...mutationOptions("Banner created successfully"),
    }),

    updateBanner: useMutation({
      mutationFn: banner.updateBanner,
      ...mutationOptions("Banner updated successfully"),
    }),

    deleteBanner: useMutation({
      mutationFn: banner.deleteBanner,
      ...mutationOptions("Banner deleted successfully"),
    }),
  };
};
