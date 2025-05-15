/* eslint-disable react/prop-types */
import { Button } from "@/components/ui/button";
import { useAddressMutation } from "@/hooks/useAddress";
import { FormDelete } from "@/components/form/FormDelete";

const DeleteAddress = ({ address }) => {
  const { deleteAddress } = useAddressMutation();

  const handleDeleteAddress = () => {
    deleteAddress.mutate(address.id);
  };
  return (
    <FormDelete
      title="Delete Address"
      loading={deleteAddress.isPending}
      onDelete={handleDeleteAddress}
      buttonElement={
        <Button size="sm" type="button" variant="destructive">
          <span>Delete</span>
        </Button>
      }
      description="Are you sure want to delete this address ?"
    />
  );
};

export { DeleteAddress };
