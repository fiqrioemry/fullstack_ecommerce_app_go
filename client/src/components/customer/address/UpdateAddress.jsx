import { addressSchema } from "@/lib/schema";
import { Button } from "@/components/ui/button";
import { useAddressMutation } from "@/hooks/useAddress";
import LocationSelection from "@/components/input/LocationSelection";
import { FormUpdateDialog } from "@/components/form/FormUpdateDialog";
import { InputTextElement } from "@/components/input/InputTextElement";
import { InputTextareaElement } from "@/components/input/InputTextareaElement";

const UpdateAddress = ({ address }) => {
  const { updateAddress } = useAddressMutation();

  return (
    <FormUpdateDialog
      state={address}
      title="Edit Address"
      schema={addressSchema}
      loading={updateAddress.isPending}
      action={updateAddress.mutateAsync}
      buttonElement={
        <Button variant="outline" size="sm" type="button">
          <span>Update Address</span>
        </Button>
      }
    >
      <InputTextElement
        name="name"
        label="Nama penerima"
        placeholder="Masukkan nama penerima"
      />

      <InputTextElement
        name="phone"
        maxLength={12}
        isNumeric={true}
        label="Nomor Telepon"
        placeholder="Masukkan nomor telepon"
      />
      <InputTextareaElement
        name="address"
        label="Alamat Penerima"
        placeholder="Masukkan Alamat Penerima"
      />
      <LocationSelection />
    </FormUpdateDialog>
  );
};
export { UpdateAddress };
