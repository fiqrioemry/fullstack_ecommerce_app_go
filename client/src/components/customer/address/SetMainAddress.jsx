/* eslint-disable react/prop-types */
import { useAddressMutation } from "@/hooks/useAddress";
import { FormToggle } from "@/components/form/FormToggle";

const SetMainAddress = ({ address }) => {
  const { setMainAddress } = useAddressMutation();

  const handleSetMainAddress = () => {
    setMainAddress.mutate(address.id);
  };
  return (
    <FormToggle
      title="Set as main address"
      onToggle={handleSetMainAddress}
      loading={setMainAddress.isPending}
      description="Are you sure want to set this as main address ?"
    />
  );
};

export { SetMainAddress };
