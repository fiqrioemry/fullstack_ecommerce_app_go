import { toast } from "sonner";
import * as address from "@/services/addresses";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";

export const useUserAddressesQuery = (param) =>
  useQuery({
    queryKey: ["addresses", param],
    queryFn: () => address.getUserAddresses(param),
    keepPreviousData: true,
  });

export const useAddressMutation = () => {
  const queryClient = useQueryClient();

  const mutationOptions = (successMsg, invalidate = true) => ({
    onSuccess: (res) => {
      toast.success(res?.message || successMsg);
      if (invalidate) {
        queryClient.invalidateQueries({ queryKey: ["addresses"] });
      }
    },
    onError: (err) => {
      toast.error(err?.response?.data?.message || "Something went wrong");
    },
  });

  return {
    createAddress: useMutation({
      mutationFn: address.createUserAddress,
      ...mutationOptions("Address created successfully"),
    }),

    updateAddress: useMutation({
      mutationFn: ({ id, data }) => address.updateUserAddress(id, data),
      ...mutationOptions("Address updated successfully"),
    }),

    deleteAddress: useMutation({
      mutationFn: address.deleteUserAddress,
      ...mutationOptions("Address deleted successfully"),
    }),

    setMainAddress: useMutation({
      mutationFn: address.setMainAddress,
      ...mutationOptions("Main address set successfully"),
    }),
  };
};
