import { Button } from "@/components/ui/Button";
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
      onDelete={handleDeleteAddress}
      loading={deleteAddress.isPending}
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
